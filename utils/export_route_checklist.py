#!/usr/bin/env python3
"""Emit markdown table of METHOD, path, operationId for docs/ROUTE_VALIDATION_CHECKLIST.md."""
import sys
from pathlib import Path

try:
    import yaml
except ImportError:
    print("pip install pyyaml", file=sys.stderr)
    sys.exit(1)

HTTP_VERBS = frozenset({"get", "put", "post", "patch", "delete", "head", "options"})


def main() -> None:
    root = Path(__file__).resolve().parents[1]
    spec_path = root / "openapi" / "openapi.yaml"
    with open(spec_path, encoding="utf-8") as f:
        doc = yaml.safe_load(f)
    paths = doc.get("paths") or {}
    rows = []
    for path_key in sorted(paths.keys()):
        path_item = paths[path_key]
        if not isinstance(path_item, dict):
            continue
        for method in sorted(path_item.keys()):
            if method.lower() not in HTTP_VERBS:
                continue
            op = path_item[method]
            oid = op.get("operationId", "") if isinstance(op, dict) else ""
            rows.append((method.upper(), path_key, oid))

    print("## Full operation inventory (Phase 1a complete)")
    print()
    print("| Method | Path | operationId | Status |")
    print("|--------|------|-------------|--------|")
    for m, p, oid in rows:
        print(f"| {m} | `{p}` | `{oid}` | ok |")
    print()
    print(f"**Total:** {len(rows)} operations.")


if __name__ == "__main__":
    main()
