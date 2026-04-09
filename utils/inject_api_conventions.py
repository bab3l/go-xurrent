#!/usr/bin/env python3
"""
Inject Xurrent API conventions from public docs into openapi/openapi.yaml:

  Pagination: https://developer.xurrent.com/v1/general/pagination/
  Filtering / predefined state: https://developer.xurrent.com/v1/general/filtering/
  Field selection: https://developer.xurrent.com/v1/general/field_selection/
  Ordering: https://developer.xurrent.com/v1/general/ordering/

- Collection GETs (not single-resource): per_page, search_after, search_before, fields, sort, state
- Single-resource GETs: If-None-Match (conditional), X-Xurrent-Language
- All GETs except denylist: X-Xurrent-Language (optional)
- Collection GET responses (200): Link, X-Pagination-Throttled, merge pagination headers if missing

Run from repo root: python utils/inject_api_conventions.py
"""
from __future__ import annotations

import re
import sys
from copy import deepcopy
from pathlib import Path

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    sys.exit(1)

# No pagination / list params (global or special endpoints)
DENY_PAGINATION = frozenset(
    {
        "/v1/me",
        "/v1/rate_limit",
        "/v1/enums",
        "/v1/account",
        "/v1/attachments/upload",
        "/v1/attachments/storage",
    }
)

# Single top-level resource: GET /v1/{tag}/{id} only — no trailing subpath
_SINGLE_RESOURCE = re.compile(r"^/v1/[^/]+/\{[^}]+\}$")


def _param_names(params: list) -> set[str]:
    out = set()
    for p in params or []:
        if isinstance(p, dict) and "name" in p:
            out.add(p["name"].lower())
        elif isinstance(p, dict) and "$ref" in p:
            out.add("$ref")
    return out


def _is_collection_get(path: str) -> bool:
    if path in DENY_PAGINATION:
        return False
    if _SINGLE_RESOURCE.match(path):
        return False
    return True


def _collection_query_params() -> list[dict]:
    return [
        {
            "name": "per_page",
            "in": "query",
            "description": "Page size (max 100). See Pagination docs.",
            "schema": {"type": "integer", "minimum": 1, "maximum": 100},
        },
        {
            "name": "search_after",
            "in": "query",
            "description": "Cursor for next page (from Link rel=next).",
            "schema": {"type": "string"},
        },
        {
            "name": "search_before",
            "in": "query",
            "description": "Cursor for previous page (from Link rel=prev).",
            "schema": {"type": "string"},
        },
        {
            "name": "fields",
            "in": "query",
            "description": "Comma-separated fields for collections (field selection).",
            "schema": {"type": "string"},
        },
        {
            "name": "sort",
            "in": "query",
            "description": "Sort fields, comma-separated; prefix with - for descending. See Ordering docs.",
            "schema": {"type": "string"},
        },
        {
            "name": "state",
            "in": "query",
            "description": "Predefined filter name (alternative to path /state). See Filtering docs.",
            "schema": {"type": "string"},
        },
    ]


def _language_header() -> dict:
    return {
        "name": "X-Xurrent-Language",
        "in": "header",
        "description": "Override response language for enums/errors (e.g. nl, fr). See API introduction.",
        "schema": {"type": "string"},
    }


def _if_none_match() -> dict:
    return {
        "name": "If-None-Match",
        "in": "header",
        "description": "Conditional GET; returns 304 Not Modified when unchanged.",
        "schema": {"type": "string"},
    }


def _collection_response_headers() -> dict[str, dict]:
    return {
        "Link": {
            "description": "RFC5988 links for pagination (rel first, prev, next).",
            "schema": {"type": "string"},
        },
        "X-Pagination-Throttled": {
            "description": "true when collection capped (e.g. 100k audit cap).",
            "schema": {"type": "string", "enum": ["true", "false"]},
        },
    }


def inject(doc: dict) -> tuple[int, int]:
    paths = doc.setdefault("paths", {})
    qp_added = 0
    hdr_added = 0
    resp_hdr_added = 0

    for path_key, path_item in paths.items():
        if not isinstance(path_item, dict):
            continue
        op = path_item.get("get")
        if not isinstance(op, dict):
            continue

        params = op.setdefault("parameters", [])
        if not isinstance(params, list):
            params = []
            op["parameters"] = params

        names = _param_names(params)

        # Optional language on all GETs
        if "x-xurrent-language" not in names:
            params.append(_language_header())
            hdr_added += 1

        if _is_collection_get(path_key):
            for p in _collection_query_params():
                if p["name"].lower() not in names:
                    params.append(deepcopy(p))
                    qp_added += 1
                    names.add(p["name"].lower())
            # Response headers for collections
            responses = op.setdefault("responses", {})
            r200 = responses.get("200")
            if isinstance(r200, dict):
                hdrs = r200.setdefault("headers", {})
                if isinstance(hdrs, dict):
                    for hk, hv in _collection_response_headers().items():
                        if hk not in hdrs:
                            hdrs[hk] = deepcopy(hv)
                            resp_hdr_added += 1
        else:
            # Single-resource style path: conditional GET
            if "if-none-match" not in names:
                params.append(_if_none_match())
                hdr_added += 1

    return qp_added, hdr_added + resp_hdr_added


def main() -> None:
    root = Path(__file__).resolve().parents[1]
    spec_path = root / "openapi" / "openapi.yaml"
    doc = yaml.safe_load(spec_path.read_text(encoding="utf-8"))
    info = doc.setdefault("info", {})
    info["description"] = (
        "Xurrent REST API. Global conventions (injected on collection GETs): "
        "pagination (`per_page`, `search_after`, `search_before`), "
        "field selection (`fields`), ordering (`sort`), predefined state (`state`), "
        "and `X-Xurrent-Language`. "
        "Resource-specific filters use additional query parameters per "
        "https://developer.xurrent.com/v1/general/filtering/ — use `collection_query.go` "
        "helpers or append encoded query pairs when filters are not modeled as named parameters."
    )
    qp, rest = inject(doc)
    spec_path.write_text(
        yaml.dump(doc, default_flow_style=False, allow_unicode=True, sort_keys=False, width=120),
        encoding="utf-8",
    )
    print(f"Wrote {spec_path}: added {qp} query params, {rest} header/response header blocks (approx)")


if __name__ == "__main__":
    main()
