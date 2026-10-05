"""Single source of truth for paths to machine-generated Markdown under docs/internal/generated/."""

from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parents[1]
GENERATED_DOCS_DIR = REPO_ROOT / "docs" / "internal" / "generated"
NDJSON_GAP_REPORT_MD = GENERATED_DOCS_DIR / "NDJSON_OPENAPI_GAP.md"
ENDPOINT_TEST_AUDIT_MD = GENERATED_DOCS_DIR / "ENDPOINT_TEST_AUDIT.md"
