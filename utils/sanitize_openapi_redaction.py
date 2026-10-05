#!/usr/bin/env python3
"""
Replace tenant-specific hostnames and account slug examples in OpenAPI YAML with neutral placeholders.

Preserves https://api.xurrent.com (standard across customers).

Usage (repo root):
  python utils/sanitize_openapi_redaction.py openapi/openapi.yaml api/openapi.yaml

Requires: no extra deps (text replacement only). Re-run yaml validation after:
  python -c \"import yaml; yaml.safe_load(open('openapi/openapi.yaml', encoding='utf-8'))\"
"""

from __future__ import annotations

import argparse
import re
import sys
from pathlib import Path


def sanitize(text: str) -> str:
    # Order matters: replace longer / more specific hosts first.
    replacements: list[tuple[str, str]] = [
        # QA / regional API hosts -> standard public API hostname
        ("api.4me-qa.mos.ru", "api.xurrent.com"),
        # Tenant web / attachment hosts (signed URLs, notes links)
        ("moscow.4me-qa.mos.ru", "web.example.com"),
        ("sc-its.4me-qa.mos.ru", "web.example.com"),
        ("4me.4me-dev.mos.ru", "web.example.com"),
        ("api.4me-dev.mos.ru", "api.xurrent.com"),
        ("moscow.4me-dev.mos.ru", "web.example.com"),
        # Any remaining QA zone host under .mos.ru
        ("4me-qa.mos.ru", "example.com"),
        # Example account slugs in captured examples (not customer-specific standard values)
        ("id: moscow", "id: example_account"),
        ("id: it-aktiv", "id: example_account"),
    ]
    out = text
    for old, new in replacements:
        out = out.replace(old, new)

    # Remove presigned query parameters from example URLs (shared-secret signatures).
    out = re.sub(
        r"\?X-4me-Expiration=\d+&X-4me-Signature=[a-f0-9]+",
        "",
        out,
    )

    # Redact email addresses left in merged live-capture examples (regional domains).
    out = re.sub(
        r"\b[\w.+%-]+@(it|edu|transport)\.mos\.ru\b",
        "user@example.com",
        out,
        flags=re.IGNORECASE,
    )
    return out


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument(
        "paths",
        nargs="+",
        type=Path,
        help="YAML files to rewrite in place (e.g. openapi/openapi.yaml api/openapi.yaml)",
    )
    args = ap.parse_args()
    for p in args.paths:
        if not p.is_file():
            print(f"Missing: {p}", file=sys.stderr)
            return 1
        raw = p.read_text(encoding="utf-8")
        new = sanitize(raw)
        if new != raw:
            p.write_text(new, encoding="utf-8")
            print(f"Updated {p}", file=sys.stderr)
        else:
            print(f"No changes {p}", file=sys.stderr)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
