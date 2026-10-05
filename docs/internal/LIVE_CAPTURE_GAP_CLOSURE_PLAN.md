# Closing NDJSON and live-shape gaps

This note sits next to the **machine-generated** gap table (`docs/internal/generated/NDJSON_OPENAPI_GAP.md`). That file tells you **what** is missing; this one explains **why** and **how** we usually fix it—without drowning in detail.

## Why gaps show up

1. **Writes** — The GET sweep only performs **GET**. Many operations are POST/PATCH/PUT/DELETE. They appear in NDJSON only when integration tests run with `XURRENT_STRUCTURE_LOG_FILE` and return **200/201** JSON.

2. **Empty lists** — The sweep resolves `/v1/foo/{id}` by listing `GET /v1/foo?per_page=1`. If nothing exists yet, paths are **skipped**. Seeding minimal rows (tests or sweep bootstrap) helps.

3. **Permissions** — If **`PostServices`** returns **401/403**, fix **roles** on the account for the API identity (e.g. configuration rights). Ordering alone cannot fix auth failures.

So: **ordering** reduces *skip*; **roles** reduce *fail*; **integration logging** reduces *write* gaps.

## Fixture order (service-centric)

A practical dependency chain (see `TestMutation_Chained_CalendarServiceOfferingRequestTemplate` in `test/integration/high_value_mutations_test.go`):

```text
Organization (existing or created)
  → Service
  → Service offering
  → Request template
  → Request / workflow / task …
```

Starting from a **service** is a good spine once you have a provider organization. A calendar is not required for `PostServices`—tests may create one for their own coverage.

**Typical order for capture:** run mutation + lifecycle + coverage tests with structure logging, **then** run the GET sweep with the same token so list endpoints return ids.

## Service spine check (optional but high leverage)

| Step | Action | Success |
|------|--------|---------|
| A | `XURRENT_PHASE0_POST_SERVICES=1` and valid `XURRENT_TOKEN` / `XURRENT_ACCOUNT`, then `go test -tags=integration ./test/integration/ -run TestPhase0_PostServicesSpine -v -count=1` | PASS with a created service id in logs |
| B | On 401/403, grant roles (e.g. `configuration_manager`) via UI or [People permissions](https://developer.xurrent.com/v1/people/permissions/) | Re-run A → PASS |

See `.env.example` for `XURRENT_PHASE0_POST_SERVICES`. Without A–B, integration can still log **writes**, but the **GET sweep** may keep skipping service-scoped paths that need a real service id.

## Full live capture pipeline (automated)

From the repo root (with `.env` containing token + account; service spine check should pass if you rely on `PostServices` during tests):

```bash
python utils/live_capture_pipeline.py --fresh -v
```

Windows: `.\utils\live_capture_pipeline.ps1 --fresh -v`

The pipeline:

1. **`--fresh`** — removes the NDJSON file so the run is not mixed with older captures.
2. Sets `XURRENT_STRUCTURE_LOG_FILE`, `XURRENT_ALLOW_MUTATIONS=1`, `XURRENT_WRITE_SHAPE_COVERAGE=1`, and (unless **`--integration-profile capture-writes`**) `XURRENT_ALLOW_LIFECYCLE_MUTATIONS=1` and `XURRENT_ENDPOINT_COVERAGE=1`. Runs **`go test -tags=integration ./test/integration/ -count=1`**, optionally **`-run REGEX`** from **`--go-test-run`**.
3. Optional **`--parent-fixture`** / **`--load-parent-fixture`** — writes or loads `test/integration/_live_shapes/parent_fixture.env`.
4. Appends **`collect_live_response_shapes.py`** GET sweep to the same NDJSON.
5. Runs **`merge_skeleton_into_openapi.py`** → `openapi/openapi.yaml`, then validates YAML.
6. Regenerates **`docs/internal/generated/NDJSON_OPENAPI_GAP.md`**.

**Faster, targeted runs:**

```bash
python utils/live_capture_pipeline.py --integration-profile capture-writes \
  --go-test-run "TestFixture_ParentIDsForCapture|TestWriteShape_" --parent-fixture \
  --skip-sweep -v
```

Use `--skip-integration`, `--skip-sweep`, `--skip-merge`, `--skip-report` for partial steps; `--ndjson` / `--env-file` to override paths.

## GET sweep improvements

`collect_live_response_shapes.py` uses list fallbacks for some collections, optional **bootstrap POSTs** when lists are empty (`--no-seed` disables), and optional env vars for workflows/products. Remaining skips may need better **multi-segment** `{id}` discovery or tenant-specific data (import/export, attachments, audit).

## Write-heavy operations

For each missing **POST/PATCH/PUT/DELETE** in the gap report: extend integration tests, add a small fixture test, or use `test/integration/write_shape_coverage_test.go` with `XURRENT_WRITE_SHAPE_COVERAGE=1`. Prioritize high-value resources first; some endpoints stay hard to capture (long-running jobs, special permissions).

## When to accept “still missing”

Some routes may never show up in a quick capture: import/export jobs, certain audit views, or search without a safe query. That is fine; the gap report is a **tooling aid**, not a scorecard.

## Quick reference (manual)

```bash
set XURRENT_STRUCTURE_LOG_FILE=test/integration/_live_shapes/full_run.jsonl
set XURRENT_ALLOW_MUTATIONS=1
set XURRENT_ALLOW_LIFECYCLE_MUTATIONS=1
set XURRENT_ENDPOINT_COVERAGE=1
set XURRENT_WRITE_SHAPE_COVERAGE=1
go test -tags=integration ./test/integration/ -v -count=1

python utils/collect_live_response_shapes.py --out test/integration/_live_shapes/full_run.jsonl
python utils/merge_skeleton_into_openapi.py test/integration/_live_shapes/full_run.jsonl openapi/openapi.yaml
python -c "import yaml; yaml.safe_load(open('openapi/openapi.yaml', encoding='utf-8'))"
python utils/report_ndjson_openapi_gaps.py --write
```

## Related

- `docs/internal/generated/NDJSON_OPENAPI_GAP.md` — per-operation gap table.
- `docs/internal/generated/ENDPOINT_TEST_AUDIT.md` — coverage and live-shape counts.
- `test/integration/high_value_mutations_test.go` — chained calendar → service → offering → template.
- `test/integration/write_shape_coverage_test.go` — optional write/read shapes.
- `utils/collect_live_response_shapes.py` — GET sweep implementation.
