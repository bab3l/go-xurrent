#!/usr/bin/env python3
"""
Probe GET endpoints on https://api.xurrent.com using .env (XURRENT_TOKEN, XURRENT_ACCOUNT).
Throttled to reduce rate-limit pressure. Path params are replaced with 1.

Usage (repo root; requires .env with XURRENT_TOKEN, XURRENT_ACCOUNT):
  python utils/probe_rest_read.py              # all GET paths from openapi.yaml
  python utils/probe_rest_read.py --extra      # also probe guessed /v1/* not in spec
  python utils/probe_rest_read.py --extra-only # only guessed paths (faster)

Does not print tokens. Default delay 0.5s between requests.
"""
from __future__ import annotations

import argparse
import os
import re
import sys
import time
import urllib.error
import urllib.request
from pathlib import Path

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    sys.exit(1)

DELAY_SEC = 0.5

# Guessed REST collection paths for UI permission names not covered by our OpenAPI (trial probes).
EXTRA_CANDIDATES = [
    "/v1/affected_slas",
    "/v1/agile_boards",
    "/v1/app_instances",
    "/v1/app_offerings",
    "/v1/audit_entries",
    "/v1/broadcasts",
    "/v1/contracts",
    "/v1/first_line_support_agreements",
    "/v1/invoices",
    "/v1/knowledge_articles",
    "/v1/knowledge_article_templates",
    "/v1/out_of_office_periods",
    "/v1/pdf_designs",
    "/v1/product_backlogs",
    "/v1/projects",
    "/v1/project_categories",
    "/v1/project_risk_levels",
    "/v1/project_tasks",
    "/v1/project_task_templates",
    "/v1/project_templates",
    "/v1/releases",
    "/v1/reservation_offerings",
    "/v1/rfc_types",
    "/v1/risks",
    "/v1/risk_severities",
    "/v1/scrum_workspaces",
    "/v1/service_categories",
    "/v1/service_instances",
    "/v1/shop_articles",
    "/v1/shop_article_categories",
    "/v1/shop_order_lines",
    "/v1/short_urls",
    "/v1/skill_pools",
    "/v1/sla_coverage_groups",
    "/v1/sla_notification_schemes",
    "/v1/sprints",
    "/v1/surveys",
    "/v1/survey_answers",
    "/v1/survey_questions",
    "/v1/survey_responses",
    "/v1/sync_sets",
    "/v1/system_logs",
    "/v1/task_templates",
    "/v1/time_allocations",
    "/v1/time_entries",
    "/v1/timesheets",
    "/v1/timesheet_settings",
    "/v1/translations",
    "/v1/waiting_for_customer_follow_ups",
    "/v1/webhooks",
    "/v1/webhook_policies",
    "/v1/workflow_templates",
    "/v1/workflow_types",
    "/v1/closure_codes",
    "/v1/holidays",
    "/v1/effort_classes",
    "/v1/email_templates",
    "/v1/note_reactions",
    "/v1/permission_events",
    "/v1/colour_palettes",
    "/v1/custom_filters",
    "/v1/custom_views",
    "/v1/ci_staged_changes",
    "/v1/account_designs",
    "/v1/self_service_designs",
    "/v1/standard_service_requests",
]


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
        if k and k not in os.environ:
            os.environ[k] = v


def fill_path(path: str) -> str:
    return re.sub(r"\{[^}]+\}", "1", path)


def probe_url(url: str, token: str, account: str) -> int | str:
    req = urllib.request.Request(
        url,
        headers={
            "Authorization": f"Bearer {token}",
            "X-Xurrent-Account": account,
            "Accept": "application/json",
        },
        method="GET",
    )
    try:
        with urllib.request.urlopen(req, timeout=45) as resp:
            return resp.status
    except urllib.error.HTTPError as e:
        return e.code
    except Exception as e:
        return f"err:{type(e).__name__}"


def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument("--extra", action="store_true", help="Probe EXTRA_CANDIDATES not in spec")
    ap.add_argument(
        "--extra-only",
        action="store_true",
        help="Only probe EXTRA_CANDIDATES (skip openapi GET paths)",
    )
    args = ap.parse_args()

    load_dotenv()
    token = os.environ.get("XURRENT_TOKEN", "").strip()
    account = os.environ.get("XURRENT_ACCOUNT", "").strip()
    if not token or not account:
        print("Set XURRENT_TOKEN and XURRENT_ACCOUNT in .env or environment.", file=sys.stderr)
        sys.exit(1)

    root = Path(__file__).resolve().parents[1]
    doc = yaml.safe_load((root / "openapi" / "openapi.yaml").read_text(encoding="utf-8"))
    paths = doc.get("paths", {})

    rows: list[tuple[str, str, int | str]] = []

    spec_path_keys = {pk for pk, pv in paths.items() if "get" in (pv or {})}

    if not args.extra_only:
        for path_key in sorted(spec_path_keys):
            filled = fill_path(path_key)
            no_qs = {
                "/v1/rate_limit",
                "/v1/me",
                "/v1/enums",
                "/v1/account",
            }
            if filled in no_qs:
                url = f"https://api.xurrent.com{filled}"
            elif re.match(r"^/v1/[^/]+/\{[^}]+\}$", path_key):
                url = f"https://api.xurrent.com{filled}"
            else:
                url = f"https://api.xurrent.com{filled}?per_page=1&fields=id"
            status = probe_url(url, token, account)
            rows.append(("spec", path_key, status))
            time.sleep(DELAY_SEC)

    covered = spec_path_keys if args.extra_only else {r[1] for r in rows}

    if args.extra or args.extra_only:
        seen: set[str] = set()
        for p in EXTRA_CANDIDATES:
            if p in seen:
                continue
            seen.add(p)
            if any(sp.startswith(p.rstrip("/") + "/") or sp == p for sp in covered):
                continue
            url = f"https://api.xurrent.com{p}?per_page=1&fields=id"
            status = probe_url(url, token, account)
            rows.append(("extra", p, status))
            time.sleep(DELAY_SEC)

    # summary counts
    by = {}
    for src, path, st in rows:
        key = str(st)
        by[key] = by.get(key, 0) + 1

    print("=== probe_rest_read summary (counts by HTTP status) ===")
    for k in sorted(by.keys(), key=lambda x: (len(str(x)), str(x))):
        print(f"  {k}: {by[k]}")
    print()
    print("=== detail (source\\tpath\\tstatus) ===")
    for src, path, st in rows:
        print(f"{src}\t{path}\t{st}")


if __name__ == "__main__":
    main()
