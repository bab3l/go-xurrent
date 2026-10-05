#!/usr/bin/env python3
"""
List OpenAPI operations that have no matching 200/201 JSON skeleton row in an NDJSON capture.

Uses the same (method, path) keys as utils/merge_skeleton_into_openapi.py (path must match
OpenAPI path templates, e.g. /v1/people/{id}).

Usage (repo root):
  python utils/report_ndjson_openapi_gaps.py
  python utils/report_ndjson_openapi_gaps.py --ndjson test/integration/_live_shapes/full_run.jsonl
  python utils/report_ndjson_openapi_gaps.py --write
  python utils/report_ndjson_openapi_gaps.py --write path/to/custom_gap_report.md

`--write` with no path writes to docs/internal/generated/NDJSON_OPENAPI_GAP.md (see utils/doc_paths.py).

Requires: pip install pyyaml
"""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

_UTILS_DIR = Path(__file__).resolve().parent
if str(_UTILS_DIR) not in sys.path:
    sys.path.insert(0, str(_UTILS_DIR))
from doc_paths import NDJSON_GAP_REPORT_MD  # noqa: E402

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    sys.exit(1)

_HTTP = frozenset({"get", "post", "put", "patch", "delete", "head", "options"})


def load_openapi_operations(spec_path: Path) -> list[tuple[str, str, str]]:
    """(METHOD, path_template, operationId) for each HTTP operation."""
    doc = yaml.safe_load(spec_path.read_text(encoding="utf-8"))
    paths = doc.get("paths") or {}
    out: list[tuple[str, str, str]] = []
    for path_key, path_item in paths.items():
        if not isinstance(path_item, dict):
            continue
        for method, op in path_item.items():
            m = method.lower()
            if m not in _HTTP:
                continue
            if not isinstance(op, dict):
                continue
            oid = str(op.get("operationId") or "")
            out.append((method.upper(), path_key, oid))
    return sorted(out, key=lambda x: (x[1], x[0], x[2]))


def captured_keys_from_ndjson(ndjson_path: Path) -> set[tuple[str, str]]:
    """
    (METHOD, path) for each response row suitable for merge (200/201 + json_skeleton).
    """
    keys: set[tuple[str, str]] = set()
    if not ndjson_path.is_file():
        return keys
    for raw in ndjson_path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line:
            continue
        try:
            row = json.loads(line)
        except json.JSONDecodeError:
            continue
        if row.get("kind") != "response":
            continue
        try:
            status = int(row.get("status", 0))
        except (TypeError, ValueError):
            continue
        if status not in (200, 201):
            continue
        if row.get("json_skeleton") is None:
            continue
        method = str(row.get("method", "GET")).upper()
        path = str(row.get("path", ""))
        if path:
            keys.add((method, path))
    return keys


def openapi_keys_with_json_success(spec_path: Path) -> set[tuple[str, str]]:
    """
    (METHOD, path) where the operation declares 200 or 201 with application/json content
    (schema key or inline empty object). Matches what merge_skeleton_into_openapi can patch.
    """
    doc = yaml.safe_load(spec_path.read_text(encoding="utf-8"))
    paths = doc.get("paths") or {}
    found: set[tuple[str, str]] = set()

    for path_key, path_item in paths.items():
        if not isinstance(path_item, dict):
            continue
        for method, op in path_item.items():
            m = method.lower()
            if m not in _HTTP:
                continue
            if not isinstance(op, dict):
                continue
            resp = op.get("responses")
            if not isinstance(resp, dict):
                continue
            for code in ("200", "201"):
                block = resp.get(code)
                if not isinstance(block, dict):
                    continue
                content = block.get("content")
                if not isinstance(content, dict):
                    continue
                aj = content.get("application/json")
                if not isinstance(aj, dict):
                    continue
                if "schema" in aj or aj == {}:
                    found.add((method.upper(), path_key))
                    break
    return found


def main() -> int:
    root = Path(__file__).resolve().parents[1]
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument(
        "--ndjson",
        type=Path,
        default=Path("test/integration/_live_shapes/full_run.jsonl"),
        help="NDJSON capture (default: test/integration/_live_shapes/full_run.jsonl)",
    )
    ap.add_argument(
        "--spec",
        type=Path,
        default=Path("openapi/openapi.yaml"),
        help="OpenAPI 3 YAML",
    )
    _gap_rel = NDJSON_GAP_REPORT_MD.relative_to(root).as_posix()
    ap.add_argument(
        "--write",
        type=Path,
        nargs="?",
        const=NDJSON_GAP_REPORT_MD,
        default=None,
        metavar="PATH",
        help=f"Write Markdown report (default path if flag is used with no value: {_gap_rel})",
    )
    ap.add_argument(
        "--only-mergeable",
        action="store_true",
        help="Restrict gap list to operations that declare 200/201 application/json in the spec",
    )
    args = ap.parse_args()
    spec_path = args.spec if args.spec.is_absolute() else root / args.spec
    ndjson_path = args.ndjson if args.ndjson.is_absolute() else root / args.ndjson

    def rel(p: Path) -> str:
        try:
            return p.resolve().relative_to(root.resolve()).as_posix()
        except ValueError:
            return p.as_posix()

    if not spec_path.is_file():
        print(f"Missing spec: {spec_path}", file=sys.stderr)
        return 1

    ops = load_openapi_operations(spec_path)
    captured = captured_keys_from_ndjson(ndjson_path)
    mergeable = openapi_keys_with_json_success(spec_path)

    if args.only_mergeable:
        candidates = [(m, p, o) for m, p, o in ops if (m, p) in mergeable]
    else:
        candidates = ops

    missing = [(m, p, o) for m, p, o in candidates if (m, p) not in captured]
    captured_in_spec = [(m, p, o) for m, p, o in candidates if (m, p) in captured]

    extra = sorted(captured - {(m, p) for m, p, _ in ops})

    lines: list[str] = [
        "# NDJSON vs OpenAPI live-shape gaps",
        "",
        "Generated by `python utils/report_ndjson_openapi_gaps.py`.",
        "",
        f"- **Spec:** `{rel(spec_path)}`",
        f"- **NDJSON:** `{rel(ndjson_path)}`"
        + (" *(file missing: no captures)*" if not ndjson_path.is_file() else ""),
        f"- **OpenAPI HTTP operations:** {len(ops)}",
        f"- **NDJSON (200/201 + json_skeleton) keys:** {len(captured)}",
        "",
    ]
    if args.only_mergeable:
        lines.append(
            f"- **Scope:** operations with **200/201** and **`application/json`** (`schema` or empty mapping) only: "
            f"**{len(mergeable)}** `(method, path)` keys; **{len(missing)}** without capture."
        )
    else:
        lines.append(
            f"- **Gap (no matching capture row):** **{len(missing)}** operations "
            f"(of {len(candidates)} in scope)."
        )
    lines.extend(
        [
            "",
            "## Missing capture (method + path + operationId)",
            "",
            "| Method | Path | operationId |",
            "|--------|------|-------------|",
        ]
    )
    for m, p, oid in missing:
        oid_disp = oid or "*(missing)*"
        lines.append(f"| {m} | `{p}` | `{oid_disp}` |")

    if extra:
        lines.extend(
            [
                "",
                "## Captured keys not present in current OpenAPI",
                "",
                "*(Often harmless: old paths, typos, or manual NDJSON lines.)*",
                "",
                "| Method | Path |",
                "|--------|------|",
            ]
        )
        for m, p in extra:
            lines.append(f"| {m} | `{p}` |")

    lines.extend(
        [
            "",
            "## Summary counts",
            "",
            f"| Metric | Count |",
            f"|--------|------:|",
            f"| Operations in scope | {len(candidates)} |",
            f"| With NDJSON capture | {len(captured_in_spec)} |",
            f"| Missing capture | {len(missing)} |",
        ]
    )
    if extra:
        lines.append(f"| NDJSON keys not in spec | {len(extra)} |")
    lines.append("")

    text = "\n".join(lines)

    if args.write:
        out = args.write if args.write.is_absolute() else root / args.write
        out.parent.mkdir(parents=True, exist_ok=True)
        out.write_text(text, encoding="utf-8")
        print(f"Wrote {out}", file=sys.stderr)

    print(text)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
