#!/usr/bin/env python3
"""
Merge openapi/backlog/*.yaml fragments into openapi/openapi.yaml (paths only).
Each fragment must be a mapping with a top-level 'paths:' key.
Later files win on duplicate path+method (last merge wins).
Run from repo root: python utils/merge_backlog_paths.py
"""
from __future__ import annotations

import sys
from pathlib import Path

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    sys.exit(1)

HTTP_VERBS = frozenset({"get", "put", "post", "patch", "delete", "head", "options"})


def load(p: Path) -> dict:
    with open(p, encoding="utf-8") as f:
        return yaml.safe_load(f) or {}


def merge_paths(main_paths: dict, fragment_paths: dict, source: str) -> int:
    added = 0
    for path_key, path_item in fragment_paths.items():
        if not isinstance(path_item, dict):
            continue
        if path_key not in main_paths:
            main_paths[path_key] = {}
        for method, op in path_item.items():
            ml = method.lower()
            if ml not in HTTP_VERBS:
                continue
            if ml in main_paths[path_key]:
                print(f"skip duplicate {source}: {path_key} {ml}", file=sys.stderr)
                continue
            main_paths[path_key][method] = op
            added += 1
    return added


def main() -> None:
    root = Path(__file__).resolve().parents[1]
    spec_path = root / "openapi" / "openapi.yaml"
    backlog_dir = root / "openapi" / "backlog"
    if not backlog_dir.is_dir():
        print(f"No backlog dir: {backlog_dir}", file=sys.stderr)
        sys.exit(1)

    doc = load(spec_path)
    main_paths = doc.get("paths")
    if not isinstance(main_paths, dict):
        print("Invalid spec: missing paths", file=sys.stderr)
        sys.exit(1)

    total = 0
    for frag in sorted(backlog_dir.glob("*.yaml")):
        data = load(frag)
        frag_paths = data.get("paths")
        if not isinstance(frag_paths, dict):
            print(f"skip {frag.name}: no paths key", file=sys.stderr)
            continue
        n = merge_paths(main_paths, frag_paths, frag.name)
        print(f"{frag.name}: merged {n} operations")
        total += n

    doc["paths"] = main_paths

    with open(spec_path, "w", encoding="utf-8") as f:
        yaml.dump(
            doc,
            f,
            default_flow_style=False,
            allow_unicode=True,
            sort_keys=False,
            width=120,
        )

    print(f"Wrote {spec_path} ({total} new operations total)")


if __name__ == "__main__":
    main()
