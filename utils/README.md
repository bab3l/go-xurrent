# Utilities

Scripts used to **validate the spec**, **generate the Go client**, **capture live JSON shapes** into OpenAPI (safely), and **audit** coverage. Run from the **repository root** unless noted.

**CI quality gates** (see `.github/workflows/ci.yml`): `go vet`, `go test`, `golangci-lint` (see `.golangci.yml` — skips generated `api_*` / `model_*` / stub tests), `govulncheck`, OpenAPI `validate`. **Dependabot** (`.github/dependabot.yml`): grouped updates for Go modules and Actions.

## Spec and client generation

| Script | Role |
|--------|------|
| `generate_client.ps1` | Generate the Go client from `openapi/openapi.yaml` (OpenAPI Generator, pinned version). |
| `generate_api.ps1` | API generation workflow used by this repo. |
| `fix_spec.py` | Spec hygiene (paths, known corrections). |
| `localize_spec.py` | Optional localization pass. |
| `inject_api_conventions.py` | Adds collection/query conventions per project rules. |
| `add_operation_ids.py` | Fills missing `operationId` values. |
| `merge_backlog_paths.py` | Merges `openapi/backlog/*.yaml` fragments into the main spec. |

## Live response capture (NDJSON → OpenAPI)

**Flow:** integration tests and/or GET sweep append **NDJSON** lines → `merge_skeleton_into_openapi.py` writes **redacted** JSON skeletons into response schemas and sets `x-structure-from-redacted-live-run`.

| Script | Role |
|--------|------|
| `live_capture_pipeline.py` | **Main entry:** optional integration tests, GET sweep, merge into `openapi/openapi.yaml`, YAML check, refresh gap report. **PowerShell:** `live_capture_pipeline.ps1`. |
| `run_write_capture_cycle.py` | Sets conservative rate-limit defaults, then forwards to `live_capture_pipeline.py` (same flags). |
| `collect_live_response_shapes.py` | Authenticated GET sweep → NDJSON (append). |
| `merge_skeleton_into_openapi.py` | Merges NDJSON skeletons into the spec (and `api/openapi.yaml` as implemented). |
| `report_ndjson_openapi_gaps.py` | Lists **OpenAPI operations** without a matching NDJSON capture row; writes `docs/internal/generated/NDJSON_OPENAPI_GAP.md` (default `--write`). |

**Environment:** `XURRENT_TOKEN`, `XURRENT_ACCOUNT`, optional `.env`; sweep spacing uses `XURRENT_SWEEP_DELAY` and `XURRENT_HTTP_MIN_INTERVAL`. Optional cooldown before a long run: `XURRENT_LIVE_CAPTURE_PRE_SLEEP` (legacy alias: `XURRENT_PHASE1_PRE_SLEEP`).

## Audits and inventories

| Script | Role |
|--------|------|
| `audit_endpoint_test_coverage.py` | Heuristic: `operationId` references in `test/**/*.go`; live-shape counts; writes `docs/internal/generated/ENDPOINT_TEST_AUDIT.md` (default `--write`). |
| `export_route_checklist.py` | Emits `docs/ROUTE_VALIDATION_CHECKLIST.md` operation table. |
| `extract_openapi_paths.py` | Lists METHOD + path from the spec. |

## OpenAPI hygiene and publishing

| Script | Role |
|--------|------|
| `sanitize_openapi_redaction.py` | Replaces tenant-specific hostnames and account examples in **OpenAPI YAML**; run before wide publication or when merging live examples. Keeps `https://api.xurrent.com` as the standard API host. |

## Probes and one-offs

| Script | Role |
|--------|------|
| `probe_rest_read.py`, `probe_nested_doc_candidates.py`, `probe_id_paths_with_real_ids.py` | exploratory HTTP probes (use with care; rate limits). |

## Configuration

**Single source of truth** for generated doc paths: **`doc_paths.py`** (`docs/internal/generated/…`).
