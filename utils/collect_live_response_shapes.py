#!/usr/bin/env python3
"""
Walk OpenAPI GET operations and call the live API, appending NDJSON lines compatible with
utils/merge_skeleton_into_openapi.py.

**Anonymization** matches test/integration/http_structure_log.go: object keys preserved,
all scalar leaves replaced with "", 0, or false; arrays collapse to one skeleton element.

Requires:
  XURRENT_TOKEN, XURRENT_ACCOUNT

Optional:
  XURRENT_API_BASE   (default https://api.xurrent.com)
  XURRENT_SWEEP_DELAY  seconds between requests (default 0.12)
  XURRENT_SWEEP_WORKFLOW_TEMPLATE_ID  numeric template id — seed POST /v1/workflows when list is empty
  XURRENT_SWEEP_WORKFLOW_TYPE         default application_change
  XURRENT_SWEEP_PRODUCT_ID            numeric product id — seed POST /v1/cis when list is empty (needs team + service)

  XURRENT_GAPTEST_<RESOURCE>_ID       optional overrides for GET paths with {{id}} — used before list discovery
                                      (see docs/internal/GAP_GET_CAPTURE_IDS.md). Example: XURRENT_GAPTEST_REQUEST_ID=12345
  XURRENT_GAPTEST_IMPORT_TOKEN        optional string for GET /v1/import/{{token}} (else discovered from GET /v1/import)

By default the sweep **seeds** minimal records (POST) when a collection is empty so GET /v1/foo/{{id}}/… can resolve
an id. Disable with **--no-seed**.

**Gap-only mode:** only the GET paths in ``GAP_GET_ONLY_PATHS`` (see gap report under docs/internal/generated/).

  python utils/collect_live_response_shapes.py --gap-gets-only --no-seed --out test/integration/_live_shapes/full_run.jsonl

Uses the same **429 backoff** as the full sweep (`http_get_json`). Prefer a conservative **XURRENT_SWEEP_DELAY** (e.g. 0.35).

Usage (repo root, after loading .env):
  python utils/collect_live_response_shapes.py --out test/integration/_live_shapes/sweep.jsonl
  python utils/merge_skeleton_into_openapi.py test/integration/_live_shapes/sweep.jsonl openapi/openapi.yaml

Use --truncate to overwrite --out; default is append.
"""

from __future__ import annotations

import argparse
import json
import os
import re
import sys
import time
import urllib.error
import urllib.request
from pathlib import Path
from typing import Any

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    sys.exit(1)

_DIGITS = re.compile(r"^\d+$")

# GET sweep: optional explicit ids (readonly production) before discover_resource_id(list).
# Keys: collection prefix up to first path parameter (same as collection_prefix()).
GAP_ID_ENV_BY_PREFIX: dict[str, str] = {
    "/v1/automation_rules": "XURRENT_GAPTEST_AUTOMATION_RULE_ID",
    "/v1/cis": "XURRENT_GAPTEST_CI_ID",
    "/v1/contracts": "XURRENT_GAPTEST_CONTRACT_ID",
    "/v1/custom_collections": "XURRENT_GAPTEST_CUSTOM_COLLECTION_ID",
    "/v1/custom_collection_elements": "XURRENT_GAPTEST_CUSTOM_COLLECTION_ELEMENT_ID",
    "/v1/flsas": "XURRENT_GAPTEST_FLSA_ID",
    "/v1/invoices": "XURRENT_GAPTEST_INVOICE_ID",
    "/v1/people": "XURRENT_GAPTEST_PERSON_ID",
    "/v1/problems": "XURRENT_GAPTEST_PROBLEM_ID",
    "/v1/projects": "XURRENT_GAPTEST_PROJECT_ID",
    "/v1/request_templates": "XURRENT_GAPTEST_REQUEST_TEMPLATE_ID",
    "/v1/requests": "XURRENT_GAPTEST_REQUEST_ID",
    "/v1/reservations": "XURRENT_GAPTEST_RESERVATION_ID",
    "/v1/service_offerings": "XURRENT_GAPTEST_SERVICE_OFFERING_ID",
    "/v1/service_instances": "XURRENT_GAPTEST_SERVICE_INSTANCE_ID",
    "/v1/services": "XURRENT_GAPTEST_SERVICE_ID",
    "/v1/slas": "XURRENT_GAPTEST_SLA_ID",
    "/v1/surveys": "XURRENT_GAPTEST_SURVEY_ID",
    "/v1/ui_extensions": "XURRENT_GAPTEST_UI_EXTENSION_ID",
    "/v1/workflows": "XURRENT_GAPTEST_WORKFLOW_ID",
}


def gap_id_from_env(prefix: str) -> str | None:
    env_key = GAP_ID_ENV_BY_PREFIX.get(prefix)
    if not env_key:
        return None
    raw = os.environ.get(env_key, "").strip()
    return raw or None


def normalize_openapi_path(path: str) -> str:
    segs = [s for s in path.strip("/").split("/") if s]
    out: list[str] = []
    for i, s in enumerate(segs):
        if _DIGITS.match(s):
            out.append("{id}")
        elif (
            len(segs) >= 3
            and segs[0] == "v1"
            and segs[1] == "import"
            and i == 2
        ):
            # OpenAPI template: /v1/import/{token}
            out.append("{token}")
        else:
            out.append(s)
    return "/" + "/".join(out)


def json_skeleton(v: Any) -> Any:
    if isinstance(v, dict):
        keys = sorted(v.keys())
        return {k: json_skeleton(v[k]) for k in keys}
    if isinstance(v, list):
        if len(v) == 0:
            return []
        return [json_skeleton(v[0])]
    if isinstance(v, str):
        return ""
    if isinstance(v, bool):
        return False
    if isinstance(v, (int, float)):
        return 0
    if v is None:
        return None
    return {"_type": type(v).__name__}


def collection_prefix(path: str) -> str | None:
    """For /v1/foo/{id}/bar return /v1/foo; no brace -> None."""
    brace = path.find("{")
    if brace < 0:
        return None
    return path[:brace].rstrip("/")


# When GET /v1/<collection> is missing or empty, try these list URLs (first 200 with ids wins).
LIST_FALLBACKS: dict[str, list[str]] = {
    "/v1/requests": [
        "/v1/requests/open",
        "/v1/requests/waiting_for_me",
        "/v1/requests/assigned_to_me",
        "/v1/requests/completed",
        "/v1/requests/requested_by_or_for_me",
    ],
    "/v1/problems": [
        "/v1/problems/active",
        "/v1/problems/assigned_to_me",
        "/v1/problems/managed_by_me",
        "/v1/problems/assigned_to_my_team",
        "/v1/problems/solved",
        "/v1/problems/known_errors",
        "/v1/problems/progress_halted",
    ],
    "/v1/tasks": [
        "/v1/tasks/open",
        "/v1/tasks/assigned_to_me",
        "/v1/tasks/managed_by_me",
        "/v1/tasks/assigned_to_my_team",
        "/v1/tasks/assigned_by_me",
        "/v1/tasks/finished",
        "/v1/tasks/approval_by_me",
    ],
}


def pick_id(row: dict[str, Any]) -> Any | None:
    if "id" in row:
        v = row["id"]
        if v is not None:
            return v
    for k, v in row.items():
        if k.endswith("_id") and v is not None and isinstance(v, (int, float, str)):
            return v
    return None


def _sleep_after_rate_limit(retry_after: str | None, attempt: int) -> None:
    """Backoff for 429: max(Retry-After seconds, exponential), capped at 120s (probe_id_paths_with_real_ids)."""
    ra = 0.0
    if retry_after:
        try:
            ra = float(max(int(str(retry_after).strip()), 0))
        except ValueError:
            ra = 0.0
    exp = min(120.0, 3.0 * (1.5 ** min(attempt, 8)))
    wait_s = min(120.0, max(ra, exp, 0.5))
    time.sleep(wait_s)


def extract_first_ids(data: Any) -> list[Any]:
    if isinstance(data, list):
        if not data or not isinstance(data[0], dict):
            return []
        pid = pick_id(data[0])
        return [pid] if pid is not None else []
    if isinstance(data, dict):
        top = pick_id(data)
        if top is not None:
            return [top]
        for k in ("data", "items", "results", "rows"):
            if k in data:
                return extract_first_ids(data[k])
    return []


def http_get_json(url: str, headers: dict[str, str], timeout: float = 60.0) -> tuple[int, str, Any | None]:
    """GET JSON; on HTTP 429 uses Retry-After + exponential backoff (max 12 attempts)."""
    max_429_attempts = 12
    for attempt in range(max_429_attempts + 1):
        req = urllib.request.Request(url, headers=headers, method="GET")
        try:
            with urllib.request.urlopen(req, timeout=timeout) as resp:
                status = resp.getcode()
                raw = resp.read(2 << 20).decode("utf-8", errors="replace")
                ct = (resp.headers.get("Content-Type") or "").lower()
                if status != 200 or "application/json" not in ct:
                    return status, ct, None
                try:
                    return status, ct, json.loads(raw)
                except json.JSONDecodeError:
                    return status, ct, None
        except urllib.error.HTTPError as e:
            if e.code == 429 and attempt < max_429_attempts:
                print(
                    f"GET 429 — backoff ({attempt + 1}/{max_429_attempts}) {url[:80]}…",
                    file=sys.stderr,
                )
                try:
                    e.read()
                except Exception:
                    pass
                _sleep_after_rate_limit(e.headers.get("Retry-After") if e.headers else None, attempt)
                continue
            try:
                raw = e.read(2 << 20).decode("utf-8", errors="replace")
            except Exception:
                raw = ""
            try:
                body = json.loads(raw) if raw else None
            except json.JSONDecodeError:
                body = None
            return e.code, "", body
        except Exception:
            return 0, "", None
    return 0, "", None


def http_post_json(
    url: str, headers: dict[str, str], body: dict[str, Any], timeout: float = 60.0
) -> tuple[int, Any | None]:
    """POST JSON; on 429 backs off and retries the same body once (seed POSTs are idempotent enough for sweep)."""
    max_429_attempts = 4
    payload = json.dumps(body).encode("utf-8")
    h = {**headers, "Content-Type": "application/json"}
    for attempt in range(max_429_attempts + 1):
        req = urllib.request.Request(url, data=payload, headers=h, method="POST")
        try:
            with urllib.request.urlopen(req, timeout=timeout) as resp:
                status = resp.getcode()
                raw = resp.read(2 << 20).decode("utf-8", errors="replace")
                ct = (resp.headers.get("Content-Type") or "").lower()
                if status not in (200, 201) or "application/json" not in ct:
                    return status, None
                try:
                    return status, json.loads(raw)
                except json.JSONDecodeError:
                    return status, None
        except urllib.error.HTTPError as e:
            if e.code == 429 and attempt < max_429_attempts:
                print(
                    f"POST 429 — backoff ({attempt + 1}/{max_429_attempts}) {url[:80]}…",
                    file=sys.stderr,
                )
                try:
                    e.read()
                except Exception:
                    pass
                _sleep_after_rate_limit(e.headers.get("Retry-After") if e.headers else None, attempt)
                continue
            try:
                raw = e.read(2 << 20).decode("utf-8", errors="replace")
            except Exception:
                raw = ""
            try:
                parsed = json.loads(raw) if raw else None
            except json.JSONDecodeError:
                parsed = None
            return e.code, parsed
        except Exception:
            return 0, None
    return 0, None


def get_me_person_id(api_base: str, headers: dict[str, str], delay: float) -> Any | None:
    time.sleep(delay)
    st, _, data = http_get_json(f"{api_base}/v1/me", headers)
    if st != 200 or not isinstance(data, dict):
        return None
    return data.get("id")


def post_create(
    api_base: str,
    path: str,
    headers: dict[str, str],
    body: dict[str, Any],
    delay: float,
    label: str,
) -> Any | None:
    """POST JSON; return created resource id or None."""
    time.sleep(delay)
    url = f"{api_base}{path}"
    st, data = http_post_json(url, headers, body)
    if st not in (200, 201) or not isinstance(data, dict):
        print(f"seed: POST {path} -> HTTP {st} ({label})", file=sys.stderr)
        return None
    rid = data.get("id")
    if rid is not None:
        print(f"seed: POST {path} -> id={rid} ({label})", file=sys.stderr)
    return rid


def run_bootstrap_seeds(
    api_base: str,
    headers: dict[str, str],
    delay: float,
) -> None:
    """
    When key collections are empty, POST minimal rows so discover_resource_id succeeds later.
    Order matches dependencies (organization from GET; service before offering/template; team for problems).
    """
    org_id = discover_resource_id("/v1/organizations", api_base, headers, delay)
    if org_id is None:
        print("seed: skip — no organization visible (GET /v1/organizations)", file=sys.stderr)
        return

    tag = str(int(time.time()))

    # Team (for POST /v1/problems)
    if discover_resource_id("/v1/teams", api_base, headers, delay) is None:
        post_create(
            api_base,
            "/v1/teams",
            headers,
            {"name": f"sweep-seed-team-{tag}"},
            delay,
            "team",
        )

    # Service (spine for offerings, templates, problems, optional workflow/CI)
    if discover_resource_id("/v1/services", api_base, headers, delay) is None:
        post_create(
            api_base,
            "/v1/services",
            headers,
            {"name": f"sweep-seed-svc-{tag}", "provider": {"id": org_id}},
            delay,
            "service",
        )

    service_id = discover_resource_id("/v1/services", api_base, headers, delay)

    if service_id is not None and discover_resource_id(
        "/v1/service_offerings", api_base, headers, delay
    ) is None:
        post_create(
            api_base,
            "/v1/service_offerings",
            headers,
            {"name": f"sweep-seed-offering-{tag}", "service": {"id": service_id}},
            delay,
            "service_offering",
        )

    if service_id is not None and discover_resource_id(
        "/v1/request_templates", api_base, headers, delay
    ) is None:
        post_create(
            api_base,
            "/v1/request_templates",
            headers,
            {
                "subject": f"sweep-seed-rt-{tag}",
                "service": {"id": service_id},
                "category": "rfi",
            },
            delay,
            "request_template",
        )

    if discover_resource_id("/v1/slas", api_base, headers, delay) is None:
        post_create(
            api_base,
            "/v1/slas",
            headers,
            {"name": f"sweep-seed-sla-{tag}", "customer": {"id": org_id}},
            delay,
            "sla",
        )

    team_id = discover_resource_id("/v1/teams", api_base, headers, delay)
    me_id = get_me_person_id(api_base, headers, delay)

    if (
        discover_resource_id("/v1/problems", api_base, headers, delay) is None
        and service_id is not None
        and team_id is not None
        and me_id is not None
    ):
        post_create(
            api_base,
            "/v1/problems",
            headers,
            {
                "service_id": service_id,
                "requested_by_id": me_id,
                "manager": me_id,
                "impact": "medium",
                "subject": f"sweep-seed-problem-{tag}",
                "team": str(team_id),
                "category": "proactive",
                "note": "sweep seed",
            },
            delay,
            "problem",
        )

    tpl = os.environ.get("XURRENT_SWEEP_WORKFLOW_TEMPLATE_ID", "").strip()
    wf_type = os.environ.get("XURRENT_SWEEP_WORKFLOW_TYPE", "application_change").strip() or "application_change"
    if (
        tpl
        and service_id is not None
        and me_id is not None
        and discover_resource_id("/v1/workflows", api_base, headers, delay) is None
    ):
        try:
            tid = int(tpl, 10)
        except ValueError:
            print("seed: invalid XURRENT_SWEEP_WORKFLOW_TEMPLATE_ID", file=sys.stderr)
            tid = 0
        if tid:
            post_create(
                api_base,
                "/v1/workflows",
                headers,
                {
                    "subject": f"sweep-seed-wf-{tag}",
                    "service": {"id": service_id},
                    "category": "rfc",
                    "manager": {"id": me_id},
                    "workflow_type": wf_type,
                    "workflow_template": {"id": tid},
                },
                delay,
                "workflow",
            )

    prod_raw = os.environ.get("XURRENT_SWEEP_PRODUCT_ID", "").strip()
    if prod_raw and service_id is not None and team_id is not None:
        if discover_resource_id("/v1/cis", api_base, headers, delay) is None:
            try:
                product_id: int | str = int(prod_raw, 10)
            except ValueError:
                product_id = prod_raw
            post_create(
                api_base,
                "/v1/cis",
                headers,
                {
                    "source": "4me",
                    "systemID": f"sweep-seed-ci-{tag}",
                    "label": f"sweep-ci-{tag}",
                    "name": f"Sweep seed CI {tag}",
                    "status": "in_service",
                    "product_id": product_id,
                    "service": {"id": service_id},
                    "support_team_id": team_id,
                },
                delay,
                "ci",
            )


# Curated GET paths for --gap-gets-only (align with docs/internal/generated/NDJSON_OPENAPI_GAP.md); order stable for diffs.
GAP_GET_ONLY_PATHS: tuple[str, ...] = (
    "/v1/attachments/upload",
    "/v1/automation_rules/{id}",
    "/v1/cis/{id}",
    "/v1/cis/{id}/ci_relations",
    "/v1/cis/{id}/users",
    "/v1/custom_collection_elements/{id}",
    "/v1/custom_collections/{id}",
    "/v1/import",
    "/v1/import/{token}",
    "/v1/people/{id}/ci",
    "/v1/problems",
    "/v1/requests/requests_of_my_organization",
    "/v1/reservations/{id}",
    "/v1/tasks/assigned_by_me",
)


def load_repo_dotenv() -> None:
    """Load repo-root `.env` into ``os.environ`` (overrides existing keys so token/account switches apply)."""
    root = Path(__file__).resolve().parents[1]
    path = root / ".env"
    if not path.is_file():
        return
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        key, _, val = line.partition("=")
        key, val = key.strip(), val.strip().strip('"').strip("'")
        if key:
            os.environ[key] = val


def discover_import_job_token(api_base: str, headers: dict[str, str], delay: float) -> str | None:
    """First token/id from GET /v1/import list for GET /v1/import/{token}."""
    time.sleep(delay)
    st, _, data = http_get_json(f"{api_base}/v1/import?per_page=5", headers)
    if st != 200 or data is None:
        return None
    rows: Any = data
    if isinstance(data, dict):
        for k in ("data", "items", "results"):
            if k in data:
                rows = data[k]
                break
    if not isinstance(rows, list) or not rows:
        return None
    first = rows[0]
    if not isinstance(first, dict):
        if isinstance(first, str) and first.strip():
            return first.strip()
        return None
    for k in ("token", "id", "job_token", "import_token"):
        v = first.get(k)
        if v is not None and str(v).strip() != "":
            return str(v).strip()
    return None


def discover_resource_id(
    base: str,
    api_base: str,
    headers: dict[str, str],
    delay: float,
) -> Any | None:
    candidates = [base] + LIST_FALLBACKS.get(base, [])
    for path in candidates:
        url = f"{api_base}{path}?per_page=1&fields=id"
        time.sleep(delay)
        status, _, data = http_get_json(url, headers)
        if status != 200 or data is None:
            continue
        ids = extract_first_ids(data)
        if ids:
            return ids[0]
    return None


def load_get_operations(spec_path: Path) -> list[tuple[str, str]]:
    doc = yaml.safe_load(spec_path.read_text(encoding="utf-8"))
    paths = doc.get("paths") or {}
    out: list[tuple[str, str]] = []
    for pkey, pitem in sorted(paths.items()):
        if not isinstance(pitem, dict):
            continue
        get_op = pitem.get("get")
        if not isinstance(get_op, dict):
            continue
        out.append(("GET", pkey))
    return out


def main() -> int:
    load_repo_dotenv()
    ap = argparse.ArgumentParser(description="Collect redacted JSON skeletons for GET /v1/* operations.")
    ap.add_argument(
        "--openapi",
        type=Path,
        default=Path(__file__).resolve().parents[1] / "openapi" / "openapi.yaml",
        help="OpenAPI spec path",
    )
    ap.add_argument("--out", type=Path, required=True, help="NDJSON output path (append unless --truncate)")
    ap.add_argument("--truncate", action="store_true", help="Overwrite output file")
    ap.add_argument("--base-url", default=os.environ.get("XURRENT_API_BASE", "https://api.xurrent.com").rstrip("/"))
    ap.add_argument(
        "--delay",
        type=float,
        default=float(os.environ.get("XURRENT_SWEEP_DELAY", "0.12")),
        help="Sleep between HTTP calls (seconds)",
    )
    ap.add_argument("--max", type=int, default=0, help="Stop after N GETs (0 = all; ignored with --gap-gets-only)")
    ap.add_argument("--dry-run", action="store_true", help="Print planned URLs only")
    ap.add_argument(
        "--no-seed",
        action="store_true",
        help="Do not POST bootstrap rows when collections are empty (old sweep behavior)",
    )
    ap.add_argument(
        "--gap-gets-only",
        action="store_true",
        help="Only GET paths in GAP_GET_ONLY_PATHS (implies --no-seed); see gap report under docs/internal/generated/",
    )
    ap.add_argument(
        "-v",
        "--verbose",
        action="store_true",
        help="Log each gap-get path result (status / skip reason) to stderr",
    )
    args = ap.parse_args()
    if args.gap_gets_only:
        args.no_seed = True

    token = os.environ.get("XURRENT_TOKEN", "").strip()
    account = os.environ.get("XURRENT_ACCOUNT", "").strip()
    if not args.dry_run and (not token or not account):
        print("Set XURRENT_TOKEN and XURRENT_ACCOUNT", file=sys.stderr)
        return 1

    headers: dict[str, str] = {
        "Accept": "application/json",
        "User-Agent": "go-xurrent-structure-sweep/1.0",
        "Authorization": f"Bearer {token}",
        "X-4me-Account": account,
    }
    unauth_headers = {
        "Accept": "application/json",
        "User-Agent": "go-xurrent-structure-sweep/1.0",
    }

    if args.gap_gets_only:
        ops = [("GET", p) for p in GAP_GET_ONLY_PATHS]
        print(
            f"gap-gets-only: {len(ops)} GET paths; delay={args.delay}s; "
            "429 handling via http_get_json; consider XURRENT_SWEEP_DELAY>=0.35 on strict tenants.",
            file=sys.stderr,
        )
    else:
        ops = load_get_operations(args.openapi)
        if args.max:
            ops = ops[: args.max]

    args.out.parent.mkdir(parents=True, exist_ok=True)
    mode = "w" if args.truncate else "a"
    stats = {"ok": 0, "skip": 0, "fail": 0}

    if not args.dry_run and token and not args.no_seed:
        run_bootstrap_seeds(args.base_url, headers, args.delay)

    with args.out.open(mode, encoding="utf-8") as sink:

        def gap_vlog(msg: str) -> None:
            if args.gap_gets_only and args.verbose:
                print(msg, file=sys.stderr)

        def emit_response(method: str, url_path: str, status: int, data: Any | None) -> None:
            if status not in (200, 201) or data is None:
                return
            norm = normalize_openapi_path(url_path)
            rec = {
                "ts": time.strftime("%Y-%m-%dT%H:%M:%SZ", time.gmtime()),
                "kind": "response",
                "method": method,
                "path": norm,
                "status": status,
                "content_type": "application/json",
                "json_skeleton": json_skeleton(data),
            }
            sink.write(json.dumps(rec, separators=(",", ":")) + "\n")
            stats["ok"] += 1

        for method, path in ops:
            if args.dry_run:
                print(method, path)
                continue

            # Rate limit is callable without tenant headers
            if path == "/v1/rate_limit":
                time.sleep(args.delay)
                url = f"{args.base_url}{path}"
                st, _, data = http_get_json(url, unauth_headers)
                if st == 200 and data is not None:
                    emit_response("GET", path, st, data)
                else:
                    stats["fail"] += 1
                continue

            if not token:
                stats["skip"] += 1
                continue

            # GET /v1/import/{token} — token string, not numeric id discovery on /v1/cis.
            if path == "/v1/import/{token}":
                tok = os.environ.get("XURRENT_GAPTEST_IMPORT_TOKEN", "").strip()
                if not tok:
                    tok = discover_import_job_token(args.base_url, headers, args.delay)
                if not tok:
                    stats["skip"] += 1
                    gap_vlog(f"gap-get: {path} skip (no import job token; set XURRENT_GAPTEST_IMPORT_TOKEN or ensure GET /v1/import returns jobs)")
                    continue
                resolved = path.replace("{token}", tok)
                time.sleep(args.delay)
                url = f"{args.base_url}{resolved}"
                st, _, data = http_get_json(url, headers)
                if st == 200 and data is not None:
                    emit_response("GET", resolved, st, data)
                    gap_vlog(f"gap-get: {path} ok HTTP {st}")
                else:
                    stats["fail"] += 1
                    gap_vlog(f"gap-get: {path} fail HTTP {st} url={url}")
                continue

            prefix = collection_prefix(path)
            resolved = path
            if prefix is not None:
                gap_raw = gap_id_from_env(prefix)
                if gap_raw is not None:
                    rid: Any = gap_raw
                    try:
                        rid = int(gap_raw, 10)
                    except ValueError:
                        pass
                else:
                    rid = discover_resource_id(prefix, args.base_url, headers, args.delay)
                if rid is None:
                    stats["skip"] += 1
                    gap_vlog(f"gap-get: {path} skip (no id for prefix {prefix}; set XURRENT_GAPTEST_* or seed data)")
                    continue
                resolved = re.sub(r"\{[^}]+\}", str(rid), path, count=1)

            time.sleep(args.delay)
            url = f"{args.base_url}{resolved}"
            st, _, data = http_get_json(url, headers)
            if st == 200 and data is not None:
                emit_response("GET", resolved, st, data)
                gap_vlog(f"gap-get: {path} ok HTTP {st}")
            else:
                stats["fail"] += 1
                gap_vlog(f"gap-get: {path} fail HTTP {st} url={url}")

    if args.dry_run:
        print(f"Would sweep {len(ops)} GET operations", file=sys.stderr)
        return 0

    print(
        f"Wrote skeleton responses to {args.out} (ok={stats['ok']} skip={stats['skip']} fail={stats['fail']})",
        file=sys.stderr,
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
