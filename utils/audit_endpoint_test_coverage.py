#!/usr/bin/env python3
"""
Cross-check OpenAPI operationIds against references in test/**/*.go and test/integration/**/*.go.

Usage (repo root):
  python utils/audit_endpoint_test_coverage.py
  python utils/audit_endpoint_test_coverage.py --write
  python utils/audit_endpoint_test_coverage.py --write path/to/custom_audit.md

`--write` with no path writes to docs/internal/generated/ENDPOINT_TEST_AUDIT.md (see utils/doc_paths.py).

"Tested" = operationId string appears somewhere in those Go files (heuristic: method names, docs).
"""
from __future__ import annotations

import argparse
import re
import sys
from pathlib import Path

_UTILS_DIR = Path(__file__).resolve().parent
if str(_UTILS_DIR) not in sys.path:
    sys.path.insert(0, str(_UTILS_DIR))
from doc_paths import ENDPOINT_TEST_AUDIT_MD, NDJSON_GAP_REPORT_MD  # noqa: E402

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    sys.exit(1)

_HTTP = frozenset({"get", "post", "put", "patch", "delete", "head", "options"})


def operation_has_live_structure_marker(op: dict) -> bool:
    """True if any 200/201 application/json block has x-structure-from-redacted-live-run."""
    resp = op.get("responses") or {}
    for code in ("200", "201"):
        block = resp.get(code)
        if not isinstance(block, dict):
            continue
        content = block.get("content") or {}
        aj = content.get("application/json")
        if isinstance(aj, dict) and aj.get("x-structure-from-redacted-live-run"):
            return True
    return False


def live_shape_operation_counts(spec_path: Path) -> tuple[int, int, int]:
    """
    Returns (raw_marker_occurrences_in_file, operations_with_at_least_one_marker, total_operations).
    Raw count can exceed unique operations when multiple response blocks are marked.
    """
    doc = yaml.safe_load(spec_path.read_text(encoding="utf-8"))
    paths = doc.get("paths") or {}
    total = 0
    with_marker = 0
    for path_key, path_item in paths.items():
        if not isinstance(path_item, dict):
            continue
        for method, op in path_item.items():
            if method.lower() not in _HTTP:
                continue
            if not isinstance(op, dict) or not op.get("operationId"):
                continue
            total += 1
            if operation_has_live_structure_marker(op):
                with_marker += 1
    raw = spec_path.read_text(encoding="utf-8").count("x-structure-from-redacted-live-run:")
    return raw, with_marker, total


def load_operation_ids(spec_path: Path) -> list[tuple[str, str, str]]:
    doc = yaml.safe_load(spec_path.read_text(encoding="utf-8"))
    paths = doc.get("paths") or {}
    out: list[tuple[str, str, str]] = []
    for path_key, path_item in paths.items():
        if not isinstance(path_item, dict):
            continue
        for method, op in path_item.items():
            if method.lower() not in ("get", "post", "put", "patch", "delete", "head", "options"):
                continue
            if not isinstance(op, dict):
                continue
            oid = op.get("operationId") or ""
            out.append((method.upper(), path_key, oid))
    return out


def scan_test_files(root: Path) -> str:
    chunks: list[str] = []
    for pat in ("test/**/*.go", "test/integration/**/*.go"):
        for p in root.glob(pat):
            chunks.append(p.read_text(encoding="utf-8", errors="replace"))
    return "\n".join(chunks)


def operation_id_to_go_identifier(operation_id: str) -> str:
    """openapi operationId get_foo_bar -> Go method name GetFooBar (heuristic)."""
    if not operation_id:
        return ""
    parts = operation_id.split("_")
    return "".join(p[:1].upper() + p[1:] if p else "" for p in parts)


def main() -> None:
    root = Path(__file__).resolve().parents[1]
    _gap_rel = NDJSON_GAP_REPORT_MD.relative_to(root).as_posix()
    _audit_rel = ENDPOINT_TEST_AUDIT_MD.relative_to(root).as_posix()
    _closure_rel = "docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md"
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument(
        "--write",
        type=Path,
        nargs="?",
        const=ENDPOINT_TEST_AUDIT_MD,
        default=None,
        metavar="PATH",
        help=f"Write markdown report (default path if flag is used with no value: {_audit_rel})",
    )
    args = ap.parse_args()
    spec_path = root / "openapi" / "openapi.yaml"
    if not spec_path.is_file():
        print(f"Missing {spec_path}", file=sys.stderr)
        sys.exit(1)

    ops = load_operation_ids(spec_path)
    n_ops = len(ops)
    n_live_shape_raw, n_ops_with_live_shape, n_ops_check = live_shape_operation_counts(spec_path)
    if n_ops_check != n_ops:
        print(
            f"warning: operation count mismatch (load_operation_ids={n_ops} vs walk={n_ops_check})",
            file=sys.stderr,
        )
    n_live_gap_ops = n_ops - n_ops_with_live_shape
    haystack = scan_test_files(root)

    tested: list[tuple[str, str, str]] = []
    untested: list[tuple[str, str, str]] = []
    for method, path, oid in sorted(ops, key=lambda x: (x[2], x[1], x[0])):
        if not oid:
            untested.append((method, path, oid))
            continue
        go_name = operation_id_to_go_identifier(oid)
        if oid in haystack or (go_name and go_name in haystack):
            tested.append((method, path, oid))
        else:
            untested.append((method, path, oid))

    lines: list[str] = [
        "# Endpoint test coverage audit",
        "",
        "Generated from `openapi/openapi.yaml`. **Heuristic:** *tested* if `operationId` or its PascalCase Go-style name "
        "(e.g. `get_products` → `GetProducts`) appears under `test/` (including skipped generated stubs).",
        "",
        f"- **Total operations:** {len(ops)}",
        f"- **With operationId:** {sum(1 for _, _, o in ops if o)}",
        f"- **Matched in test files:** {len(tested)}",
        f"- **Not matched:** {len(untested)}",
        "",
        "## OpenAPI, client, and live-response shapes",
        "",
        f"- **Operations in the published spec** (`paths` × HTTP methods): **{len(ops)}**. The Go client in this repo is generated from the same document, so **implemented API operations match that count** (one method per `operationId`).",
        f"- **Live-captured anonymized JSON** (`x-structure-from-redacted-live-run: true` under `application/json` on a **200**/**201** response, from NDJSON merge): **{n_ops_with_live_shape}** / **{n_ops}** operations have at least one such block; **{n_live_shape_raw}** total marker occurrences in the file (some operations have more than one marked block).",
        f"- **Operations without any such marker yet:** **{n_live_gap_ops}** (see *Gaps to full live-shape coverage*). "
        f"For **mergeable** `200`/`201` JSON operations vs NDJSON keys, see `{_gap_rel}`.",
        "",
        "### Current status (short)",
        "",
        f"- Spec + generated client surface: **{n_ops}** operations.",
        f"- Test heuristic (string match in `test/**/*.go`): **{len(tested)}** referenced, **{len(untested)}** missing references.",
        f"- Marker occurrences in YAML (includes multi-block operations): **{n_live_shape_raw}**.",
        "",
        "### Gaps to full live-shape coverage",
        "",
        "“Full” here means every operation’s successful **`application/json`** **`200` or `201`** response has been logged, merged, and marked with `x-structure-from-redacted-live-run`. Typical reasons the remaining "
        f"**{n_live_gap_ops}** operations are not marked:",
        "",
        "- Operation never called during integration + sweep (or only with **4xx** / **401** in your tenant).",
        "- GET sweep **skip** (no path match / not in sweep list) or **fail** (network, rate limit, permissions).",
        "- Response is **not JSON** or success code is not 200/201 (e.g. **204**, **302**, binary, HTML error pages).",
        "- **Writes** not executed: opt-in `XURRENT_ALLOW_MUTATIONS=1` (and related flags) adds POST/PATCH bodies to NDJSON.",
        "- Merge only applies when OpenAPI already has `content` → `application/json` with `schema:` or `application/json: {}` under that status.",
        "",
        "Doc parity, **`components/schemas`**, and strict typing are separate from this counter; see `docs/ROADMAP.md`.",
        "",
        "### Suggested next steps",
        "",
        "1. Re-run integration tests with `XURRENT_STRUCTURE_LOG_FILE` and env flags for mutations/coverage as needed.",
        "2. Run `python utils/collect_live_response_shapes.py` (see script help); append to NDJSON, then `python utils/merge_skeleton_into_openapi.py …`.",
        "3. Fix tenant/token gaps that cause sweep **fail** (e.g. service create 401); set optional test IDs from integration test headers.",
        "4. After merge: `python -c \"import yaml; yaml.safe_load(open('openapi/openapi.yaml', encoding='utf-8'))\"` to confirm YAML is valid.",
        "5. List exact operations still missing from the capture: `python utils/report_ndjson_openapi_gaps.py --ndjson test/integration/_live_shapes/combined_for_gap_report.jsonl --write --only-mergeable` (build `combined_for_gap_report.jsonl` by concatenating NDJSON logs such as `full_run.jsonl` + `entity_get_by_id.jsonl`).",
        f"6. Service spine check (permissions): `XURRENT_PHASE0_POST_SERVICES=1 go test -tags=integration ./test/integration/ -run TestPhase0_PostServicesSpine -v` (see `{_closure_rel}`).",
        f"7. Full live capture: `python utils/live_capture_pipeline.py --fresh -v` (loads `.env`; merges into OpenAPI; refreshes `{_gap_rel}`).",
        "",
        "Latest local capture artifacts (paths under `test/integration/_live_shapes/`, gitignored):",
        "",
        "- `test/integration/_live_shapes/integration_test_run.log` — integration tests with structure logging.",
        "- `test/integration/_live_shapes/collect_sweep.log` — `utils/collect_live_response_shapes.py` GET sweep.",
        "- `test/integration/_live_shapes/full_run.jsonl` — combined NDJSON input to `utils/merge_skeleton_into_openapi.py`.",
        "- `test/integration/_live_shapes/entity_get_by_id.jsonl` — optional focused run for `GetTeamsId` / `GetSitesId` / `GetSlasId` / `GetTasksId` list→id capture.",
        "",
        "## Untested operationIds",
        "",
        "| Method | Path | operationId |",
        "|--------|------|-------------|",
    ]
    for method, path, oid in untested:
        if not oid:
            oid = "*(missing)*"
        lines.append(f"| {method} | `{path}` | `{oid}` |")

    lines.extend(
        [
            "",
            "## Integration coverage",
            "",
            "- `test/integration/production_smoke_test.go` — subset of authenticated GETs + rate limit; `GET /v1/problems/{id}` when a problem id is available (env, `parent_fixture.env`, or problem list discovery).",
            "- `test/integration/mutations_test.go` — opt-in `XURRENT_ALLOW_MUTATIONS=1` for POST/PATCH (person + product).",
            "- `test/integration/high_value_mutations_test.go` — opt-in `XURRENT_ALLOW_MUTATIONS=1` for teams, sites, services, SLAs (POST/PATCH + disable cleanup); chained calendar → service → service offering → request template (`TestMutation_Chained_CalendarServiceOfferingRequestTemplate`); optional workflow create/patch/trash when `XURRENT_TEST_WORKFLOW_TEMPLATE_ID` is set (`TestMutation_Chained_WorkflowCreateTrash`).",
            "- `test/integration/lifecycle_mutations_test.go` — opt-in `XURRENT_ALLOW_MUTATIONS=1` and `XURRENT_ALLOW_LIFECYCLE_MUTATIONS=1` for archive/trash/restore on workflows (incl. `PostWorkflowsIdTasks`), problems, requests, people; POST→GET fixtures with cleanup.",
            "- `test/integration/coverage_get_test.go` — opt-in `XURRENT_ENDPOINT_COVERAGE=1` for broad GET smoke (`TestCoverage_HighValueGETs`: workflows with list/env/fixture id → `GetWorkflowsId`, calendars, teams, sites, services, offerings, request templates, SLA lists; `TestCoverage_EntityGETByID`: `GetTeamsId`, `GetSitesId`, `GetSlasId`, `GetTasksId` via list discovery; filters for problems/requests/people; usage statements and product CIs where permitted).",
            "- `utils/collect_live_response_shapes.py` — optional authenticated GET sweep → NDJSON; merge with `utils/merge_skeleton_into_openapi.py` (fills `application/json: {}` and existing `schema:` blocks with anonymized shapes). **Merge only within the same OpenAPI operation** (verb block); otherwise a `POST` without `200` could incorrectly match a sibling `GET` `200` and corrupt YAML.",
            "",
            "Regenerate this file:",
            "",
            "```bash",
            "python utils/audit_endpoint_test_coverage.py --write",
            "```",
            "",
        ]
    )

    text = "\n".join(lines) + "\n"
    if args.write:
        out = args.write if args.write.is_absolute() else root / args.write
        out.parent.mkdir(parents=True, exist_ok=True)
        out.write_text(text, encoding="utf-8")
        print(f"Wrote {out}")
    else:
        print(text)


if __name__ == "__main__":
    main()
