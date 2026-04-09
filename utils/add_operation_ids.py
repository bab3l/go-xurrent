#!/usr/bin/env python3
"""
Add unique operationId to every operation in openapi/openapi.yaml if missing.
Run from repo root: python utils/add_operation_ids.py
Requires: pip install pyyaml
"""
from __future__ import annotations

import re
import sys
from pathlib import Path

try:
    import yaml
except ImportError:
    print("PyYAML required: pip install pyyaml", file=sys.stderr)
    sys.exit(1)

HTTP_VERBS = frozenset({"get", "put", "post", "patch", "delete", "head", "options", "trace"})


def path_to_slug(path: str) -> str:
    p = path.strip()
    if p.startswith("/v1"):
        p = p[3:]
    p = p.strip("/")
    if not p:
        return "v1_root"
    parts: list[str] = []
    for seg in p.split("/"):
        seg = re.sub(r"^\{|\}$", "", seg)
        parts.append(seg)
    slug = "_".join(parts)
    slug = re.sub(r"[^a-zA-Z0-9_]+", "_", slug)
    slug = re.sub(r"_+", "_", slug).strip("_").lower()
    return slug or "path"


def make_operation_id(method: str, path: str, used: set[str]) -> str:
    base = f"{method.lower()}_{path_to_slug(path)}"
    if base[0].isdigit():
        base = "op_" + base
    oid = base
    n = 2
    while oid in used:
        oid = f"{base}_{n}"
        n += 1
    used.add(oid)
    return oid


def main() -> None:
    root = Path(__file__).resolve().parents[1]
    spec_path = root / "openapi" / "openapi.yaml"
    with open(spec_path, encoding="utf-8") as f:
        doc = yaml.safe_load(f)

    paths = doc.get("paths") or {}
    used: set[str] = set()
    # reserve existing ids
    for path_key, path_item in paths.items():
        if not isinstance(path_item, dict):
            continue
        for method, op in path_item.items():
            if method.lower() not in HTTP_VERBS:
                continue
            if not isinstance(op, dict):
                continue
            oid = op.get("operationId")
            if isinstance(oid, str) and oid.strip():
                used.add(oid.strip())

    added = 0
    for path_key in sorted(paths.keys()):
        path_item = paths[path_key]
        if not isinstance(path_item, dict):
            continue
        for method in sorted(path_item.keys(), key=lambda m: (m.lower() not in HTTP_VERBS, m)):
            if method.lower() not in HTTP_VERBS:
                continue
            op = path_item[method]
            if not isinstance(op, dict):
                continue
            if op.get("operationId"):
                continue
            op["operationId"] = make_operation_id(method, path_key, used)
            added += 1

    if added == 0:
        print(f"No changes ({spec_path} already has {len(used)} operationIds).")
        return

    with open(spec_path, "w", encoding="utf-8") as f:
        yaml.dump(
            doc,
            f,
            default_flow_style=False,
            allow_unicode=True,
            sort_keys=False,
            width=120,
        )

    print(f"Wrote {spec_path} ({added} operationIds added, total unique reserved: {len(used)})")


if __name__ == "__main__":
    main()
