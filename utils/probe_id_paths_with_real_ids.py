#!/usr/bin/env python3
"""
Validate id-based GET paths by first discovering real IDs from list endpoints.

Usage:
  python utils/probe_id_paths_with_real_ids.py

Outputs:
  - Console summary
  - docs/id_probe_report.md
"""
from __future__ import annotations

import json
import os
import re
import time
import urllib.error
import urllib.parse
import urllib.request
from collections import Counter
from datetime import datetime, timezone
from pathlib import Path

BASE_URL = "https://api.xurrent.com"
REQUEST_DELAY_SEC = 0.45

TARGETS = [
    ("/v1/account/usage_statements/{id}", "id", "/v1/account/usage_statements", ("id",)),
    ("/v1/automation_rules/{id}", "id", "/v1/automation_rules", ("id",)),
    ("/v1/cis/{id}", "id", "/v1/cis", ("id",)),
    ("/v1/cis/{id}/ci_relations", "id", "/v1/cis", ("id",)),
    ("/v1/cis/{id}/users", "id", "/v1/cis", ("id",)),
    ("/v1/custom_collection_elements/{id}", "id", "/v1/custom_collection_elements", ("id",)),
    ("/v1/custom_collections/{id}", "id", "/v1/custom_collections", ("id",)),
    ("/v1/import/{token}", "token", "/v1/import", ("token", "id")),
    ("/v1/organizations/{id}", "id", "/v1/organizations", ("id",)),
    ("/v1/people/{id}", "id", "/v1/people", ("id",)),
    ("/v1/people/{id}/cis", "id", "/v1/people", ("id",)),
    ("/v1/people/{id}/contacts", "id", "/v1/people", ("id",)),
    ("/v1/people/{id}/permissions", "id", "/v1/people", ("id",)),
    ("/v1/people/{id}/service_coverages", "id", "/v1/people", ("id",)),
    ("/v1/people/{id}/skill_pools", "id", "/v1/people", ("id",)),
    ("/v1/people/{id}/teams", "id", "/v1/people", ("id",)),
    ("/v1/problems/{id}", "id", "/v1/problems", ("id",)),
    ("/v1/product_categories/{id}", "id", "/v1/product_categories", ("id",)),
    ("/v1/products/{id}", "id", "/v1/products", ("id",)),
    ("/v1/request_templates/{id}", "id", "/v1/request_templates", ("id",)),
    ("/v1/requests/{id}", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/audit", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/cis", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/cis/active", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/cis/inactive", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/grouped_requests", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/notes", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/notes/internal", "id", "/v1/requests", ("id",)),
    ("/v1/requests/{id}/notes/public", "id", "/v1/requests", ("id",)),
    ("/v1/invoices/{id}/audit", "id", "/v1/invoices", ("id",)),
    ("/v1/invoices/{id}/cis", "id", "/v1/invoices", ("id",)),
    ("/v1/contracts/{id}/audit", "id", "/v1/contracts", ("id",)),
    ("/v1/contracts/{id}/cis", "id", "/v1/contracts", ("id",)),
    ("/v1/projects/{id}/audit", "id", "/v1/projects", ("id",)),
    ("/v1/flsas/{id}/audit", "id", "/v1/flsas", ("id",)),
    ("/v1/products/{id}/cis", "id", "/v1/products", ("id",)),
    ("/v1/surveys/{id}/survey_questions", "id", "/v1/surveys", ("id",)),
    ("/v1/reservations/{id}", "id", "/v1/reservations", ("id",)),
    ("/v1/service_offerings/{id}", "id", "/v1/service_offerings", ("id",)),
    ("/v1/service_offerings/{id}/audit", "id", "/v1/service_offerings", ("id",)),
    ("/v1/services/{id}", "id", "/v1/services", ("id",)),
    ("/v1/sites/{id}", "id", "/v1/sites", ("id",)),
    ("/v1/slas/{id}", "id", "/v1/slas", ("id",)),
    ("/v1/tasks/{id}", "id", "/v1/tasks", ("id",)),
    ("/v1/teams/{id}", "id", "/v1/teams", ("id",)),
    ("/v1/teams/{id}/members", "id", "/v1/teams", ("id",)),
    ("/v1/ui_extensions/{id}", "id", "/v1/ui_extensions", ("id",)),
    ("/v1/workflows/{id}", "id", "/v1/workflows", ("id",)),
]

SOURCE_ALIASES: dict[str, tuple[str, ...]] = {
    "/v1/requests": (
        "/v1/requests/waiting_for_me",
        "/v1/requests/problem_management_review",
        "/v1/requests/sla_accountability",
        "/v1/requests",
    ),
    "/v1/tasks": (
        "/v1/tasks/finished",
        "/v1/tasks/approval_by_me",
        "/v1/tasks",
    ),
    "/v1/products": (
        "/v1/products/enabled",
        "/v1/products/disabled",
        "/v1/products/supported_by_my_teams",
        "/v1/products",
    ),
    "/v1/problems": (
        "/v1/problems/active",
        "/v1/problems/known_errors",
        "/v1/problems/progress_halted",
        "/v1/problems/solved",
        "/v1/problems",
    ),
}


def load_dotenv() -> None:
    env_path = Path(__file__).resolve().parents[1] / ".env"
    if not env_path.is_file():
        return
    for line in env_path.read_text(encoding="utf-8").splitlines():
        line = line.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        k, v = line.split("=", 1)
        k, v = k.strip(), v.strip().strip('"').strip("'")
        if k:
            # Prefer the active .env file for deterministic runs across persistent shell sessions.
            os.environ[k] = v


def request(path: str, token: str, account: str, query: dict[str, str] | None = None) -> tuple[int | str, bytes]:
    if query:
        qs = urllib.parse.urlencode(query)
        url = f"{BASE_URL}{path}?{qs}"
    else:
        url = f"{BASE_URL}{path}"

    max_rounds = 12
    for attempt in range(max_rounds):
        req = urllib.request.Request(
            url,
            headers={
                "Authorization": f"Bearer {token}",
                "X-4me-Account": account,
                "Accept": "application/json",
            },
            method="GET",
        )
        try:
            with urllib.request.urlopen(req, timeout=90) as resp:
                return resp.status, resp.read()
        except urllib.error.HTTPError as e:
            if e.code == 429 and attempt + 1 < max_rounds:
                retry_after = e.headers.get("Retry-After") if e.headers else None
                sleep_s = min(120.0, 3.0 * (1.5 ** min(attempt, 8)))
                if retry_after:
                    try:
                        sleep_s = min(120.0, max(sleep_s, float(int(retry_after))))
                    except ValueError:
                        pass
                time.sleep(sleep_s)
                continue
            return e.code, e.read() if hasattr(e, "read") else b""
        except Exception as e:
            return f"err:{type(e).__name__}", b""
    return 429, b""


def first_id_from_payload(body: bytes, candidate_keys: tuple[str, ...]) -> str | None:
    try:
        data = json.loads(body.decode("utf-8"))
    except Exception:
        return None

    first = None
    if isinstance(data, list):
        first = data[0] if data else None
    elif isinstance(data, dict):
        if isinstance(data.get("data"), list) and data["data"]:
            first = data["data"][0]
        else:
            first = data

    if isinstance(first, str) and "token" in candidate_keys:
        return first
    if not isinstance(first, dict):
        return None
    for k in candidate_keys:
        v = first.get(k)
        if v is not None and str(v).strip() != "":
            return str(v)
    return None


def discover_identifier(source: str, token: str, account: str, keys: tuple[str, ...]) -> tuple[str | None, str]:
    # Try progressively broader calls to maximize chances of finding at least one entity in tenant data.
    attempts: list[tuple[dict[str, str], str]] = [
        ({"per_page": "1", "fields": ",".join(keys)}, "fields"),
        ({"per_page": "1"}, "plain"),
        ({"per_page": "1", "state": "all"}, "state_all"),
        ({"per_page": "1", "fields": ",".join(keys), "state": "all"}, "fields_state_all"),
    ]

    source_variants = SOURCE_ALIASES.get(source, (source,))
    saw_200 = False
    last_status: int | str | None = None
    for source_variant in source_variants:
        for query, label in attempts:
            st, body = request(source_variant, token, account, query)
            last_status = st
            if st != 200:
                time.sleep(REQUEST_DELAY_SEC)
                continue
            saw_200 = True
            found = first_id_from_payload(body, keys)
            time.sleep(REQUEST_DELAY_SEC)
            if found is not None:
                return found, f"ok:{source_variant}:{label}"
    if not saw_200:
        if last_status is None:
            return None, "source_unknown"
        return None, f"source_status={last_status}"
    return None, "source_no_id"


def main() -> int:
    load_dotenv()
    token = os.environ.get("XURRENT_TOKEN", "").strip()
    account = os.environ.get("XURRENT_ACCOUNT", "").strip()
    if not token or not account:
        print("Set XURRENT_TOKEN and XURRENT_ACCOUNT in .env or environment.")
        return 1

    checked_at = datetime.now(timezone.utc).replace(microsecond=0).isoformat().replace("+00:00", "Z")
    id_cache: dict[tuple[str, tuple[str, ...]], tuple[str | None, str]] = {}
    rows: list[dict[str, str | int]] = []

    for target, placeholder, source, keys in TARGETS:
        cache_key = (source, keys)
        if cache_key not in id_cache:
            id_cache[cache_key] = discover_identifier(source, token, account, keys)

        value, source_note = id_cache[cache_key]
        if not value:
            rows.append(
                {
                    "target": target,
                    "source": source,
                    "placeholder": placeholder,
                    "idValue": "",
                    "status": "not_tested",
                    "note": source_note,
                }
            )
            continue

        path = target.replace("{" + placeholder + "}", value)
        st, _ = request(path, token, account, None)
        rows.append(
            {
                "target": target,
                "source": source,
                "placeholder": placeholder,
                "idValue": value,
                "status": st,
                "note": "",
            }
        )
        time.sleep(REQUEST_DELAY_SEC)

    counts = Counter(str(r["status"]) for r in rows)

    report_path = Path(__file__).resolve().parents[1] / "docs" / "id_probe_report.md"
    lines = [
        "# ID Endpoint Validation Report",
        "",
        f"- Checked at: `{checked_at}`",
        f"- Base URL: `{BASE_URL}`",
        f"- Paths attempted: `{len(rows)}`",
        "",
        "## Summary by status",
        "",
    ]
    for k, v in sorted(counts.items(), key=lambda kv: (len(kv[0]), kv[0])):
        lines.append(f"- `{k}`: {v}")

    lines += [
        "",
        "## Detailed results",
        "",
        "| target | source_for_id | id_used | status | note |",
        "|---|---|---:|---:|---|",
    ]
    for r in rows:
        lines.append(
            f"| `{r['target']}` | `{r['source']}` | `{r['idValue']}` | `{r['status']}` | {r['note']} |"
        )

    report_path.write_text("\n".join(lines) + "\n", encoding="utf-8")

    print("=== id probe summary ===")
    for k, v in sorted(counts.items(), key=lambda kv: (len(kv[0]), kv[0])):
        print(f"  {k}: {v}")
    print(f"Wrote {report_path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
