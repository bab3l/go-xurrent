# Maintainer docs

These pages support **working on the OpenAPI spec**, the **generated Go client**, and **live-response capture** (NDJSON → merge → spec). They are not required to *use* the client as a library.

If you are new here, start with the repo **[ROADMAP.md](../ROADMAP.md)** for priorities and **[utils/README.md](../../utils/README.md)** at the repo root for scripts.

## Generated reports (`generated/`)

Files under **`generated/`** are **outputs of tools**. Prefer regenerating them over hand-editing:

| File | Regenerate |
|------|------------|
| `generated/NDJSON_OPENAPI_GAP.md` | `python utils/report_ndjson_openapi_gaps.py --write` |
| `generated/ENDPOINT_TEST_AUDIT.md` | `python utils/audit_endpoint_test_coverage.py --write` |

Paths are defined in **`utils/doc_paths.py`** so scripts stay consistent.

## Hand-maintained Go (not generated)

- **`client.go`**, **`api_*.go`**, **`model_*.go`** are produced by OpenAPI Generator — do not edit them for permanent features.
- Extensions live in paths listed in **`.openapi-generator-ignore`** (e.g. **`collection_filter.go`**, **`collection_get.go`**, matching `*_test.go`, **`README.md`**, **`go.mod`**). Regenerate with **`utils/generate_api.ps1`** (or **`utils/generate_client.ps1`**); those files are preserved.

## Guides (hand-written)

| Doc | Purpose |
|-----|---------|
| [LIVE_CAPTURE_GAP_CLOSURE_PLAN.md](LIVE_CAPTURE_GAP_CLOSURE_PLAN.md) | Why NDJSON gaps exist and how we close them (permissions, fixtures, sweep). |
| [GAP_GET_CAPTURE_IDS.md](GAP_GET_CAPTURE_IDS.md) | Optional `XURRENT_GAPTEST_*` ids for GET paths with `{id}`. |
| [GET_SEARCH_AND_NDJSON_GAPS.md](GET_SEARCH_AND_NDJSON_GAPS.md) | GetSearch quirks, 400s, and NDJSON gap categories. |
