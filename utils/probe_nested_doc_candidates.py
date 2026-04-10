#!/usr/bin/env python3
"""
Inventory nested routes in openapi/openapi.yaml and probe doc-mentioned paths
that may not be in the spec (audit subroutes, sprint backlog_items, etc.).

Uses .env: XURRENT_TOKEN, XURRENT_ACCOUNT. Throttled GETs.

Usage (repo root):
  python utils/probe_nested_doc_candidates.py --inventory
  python utils/probe_nested_doc_candidates.py --probe
  python utils/probe_nested_doc_candidates.py              # both
"""
from __future__ import annotations

import argparse
import json
import os
import re
import sys
import time
import urllib.error
import urllib.parse
import urllib.request
from pathlib import Path

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    raise SystemExit(1)

BASE_URL = "https://api.xurrent.com"
REQUEST_DELAY_SEC = 0.45

# Paths from public docs / gap notes; not necessarily in OpenAPI. (template, list_path, id_keys)
# Still-not-in-spec nested paths (promoted routes removed as they land in openapi/openapi.yaml).
DOC_PROBE_CANDIDATES: list[tuple[str, str, tuple[str, ...]]] = [
    ("/v1/releases/{id}/audit", "/v1/releases", ("id",)),
    ("/v1/sprints/{id}/audit", "/v1/sprints", ("id",)),
    ("/v1/sprints/{id}/backlog_items", "/v1/sprints", ("id",)),
    ("/v1/surveys/{id}/questions", "/v1/surveys", ("id",)),
    ("/v1/shop_articles/{id}/cis", "/v1/shop_articles", ("id",)),
    ("/v1/knowledge_articles/{id}/cis", "/v1/knowledge_articles", ("id",)),
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
        if k:
            os.environ[k] = v


def request(path: str, token: str, account: str, query: dict[str, str] | None = None) -> tuple[int | str, bytes]:
    if query:
        url = f"{BASE_URL}{path}?{urllib.parse.urlencode(query)}"
    else:
        url = f"{BASE_URL}{path}"
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
        return e.code, e.read() if hasattr(e, "read") else b""
    except Exception as e:
        return f"err:{type(e).__name__}", b""


def openapi_paths(repo_root: Path) -> list[str]:
    p = repo_root / "openapi" / "openapi.yaml"
    data = yaml.safe_load(p.read_text(encoding="utf-8"))
    return sorted(data.get("paths", {}).keys())


def nested_paths_in_spec(paths: list[str]) -> list[str]:
    """Paths with at least one /{param}/ segment after /v1/<collection> (nested sub-resource)."""
    nested: list[str] = []
    for path in paths:
        if not path.startswith("/v1/"):
            continue
        rest = path[len("/v1/") :]
        if rest.count("/") < 2:
            continue
        if "{" not in rest:
            continue
        nested.append(path)
    return sorted(nested)


def first_id_from_payload(body: bytes, candidate_keys: tuple[str, ...]) -> str | None:
    try:
        data = json.loads(body.decode("utf-8"))
    except Exception:
        return None
    first = None
    if isinstance(data, list):
        first = data[0] if data else None
    elif isinstance(data, dict) and isinstance(data.get("data"), list) and data["data"]:
        first = data["data"][0]
    if not isinstance(first, dict):
        return None
    for k in candidate_keys:
        v = first.get(k)
        if v is not None and str(v).strip() != "":
            return str(v)
    return None


def discover_id(list_path: str, token: str, account: str, keys: tuple[str, ...]) -> tuple[str | None, str]:
    for query in (
        {"per_page": "1", "fields": ",".join(keys)},
        {"per_page": "1"},
    ):
        st, body = request(list_path, token, account, query)
        time.sleep(REQUEST_DELAY_SEC)
        if st != 200:
            continue
        found = first_id_from_payload(body, keys)
        if found:
            return found, "ok"
    return None, "no_id_or_non_200"


def run_inventory(repo_root: Path) -> None:
    paths = openapi_paths(repo_root)
    nested = nested_paths_in_spec(paths)
    print(f"=== OpenAPI nested paths ({len(nested)} under /v1/.../{{param}}/...) ===")
    for p in nested:
        print(p)
    per_audit = [p for p in nested if "/audit" in p]
    print()
    print(f"--- Subset: .../audit ({len(per_audit)}) ---")
    for p in per_audit:
        print(p)


def run_probe(repo_root: Path) -> None:
    load_dotenv()
    token = os.environ.get("XURRENT_TOKEN", "").strip()
    account = os.environ.get("XURRENT_ACCOUNT", "").strip()
    if not token or not account:
        print("Set XURRENT_TOKEN and XURRENT_ACCOUNT in .env", file=sys.stderr)
        raise SystemExit(1)

    spec_paths = set(openapi_paths(repo_root))
    print("=== Doc-candidate nested routes (live GET) ===")
    print("(template | in_openapi? | list_ok | status | note)")
    for template, list_path, keys in DOC_PROBE_CANDIDATES:
        in_spec = "yes" if template in spec_paths else "no"
        vid, note = discover_id(list_path, token, account, keys)
        if not vid:
            print(f"{template} | {in_spec} | list_fail | - | {note} ({list_path})")
            continue
        resolved = re.sub(r"\{id\}", vid, template)
        st, _ = request(resolved, token, account, None)
        time.sleep(REQUEST_DELAY_SEC)
        print(f"{template} | {in_spec} | ok | {st} | id={vid}")


def main() -> int:
    repo_root = Path(__file__).resolve().parents[1]
    ap = argparse.ArgumentParser(description="Nested route inventory and doc-candidate probe")
    ap.add_argument("--inventory", action="store_true", help="List nested paths from openapi.yaml")
    ap.add_argument("--probe", action="store_true", help="Probe doc candidates against API")
    args = ap.parse_args()
    if not args.inventory and not args.probe:
        args.inventory = args.probe = True
    if args.inventory:
        run_inventory(repo_root)
    if args.probe:
        if args.inventory:
            print()
        run_probe(repo_root)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
