#!/usr/bin/env python3
"""Emit METHOD path lines from openapi/openapi.yaml for route validation assignments."""
import sys
from pathlib import Path

try:
    import yaml  # type: ignore
except ImportError:
    print("PyYAML required: pip install pyyaml", file=sys.stderr)
    sys.exit(1)


def main() -> None:
    root = Path(__file__).resolve().parents[1]
    spec_path = root / "openapi" / "openapi.yaml"
    with open(spec_path, encoding="utf-8") as f:
        doc = yaml.safe_load(f)
    paths = doc.get("paths") or {}
    verbs = ("get", "put", "post", "patch", "delete", "head", "options")
    lines = []
    for path_key in sorted(paths.keys()):
        ops = paths[path_key]
        for verb in verbs:
            if verb in ops:
                lines.append(f"{verb.upper():7} {path_key}")
    for line in lines:
        print(line)


if __name__ == "__main__":
    main()
