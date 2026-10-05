# Endpoint test coverage audit

Generated from `openapi/openapi.yaml`. **Heuristic:** *tested* if `operationId` or its PascalCase Go-style name (e.g. `get_products` → `GetProducts`) appears under `test/` (including skipped generated stubs).

- **Total operations:** 187
- **With operationId:** 187
- **Matched in test files:** 187
- **Not matched:** 0

## OpenAPI, client, and live-response shapes

- **Operations in the published spec** (`paths` × HTTP methods): **187**. The Go client in this repo is generated from the same document, so **implemented API operations match that count** (one method per `operationId`).
- **Live-captured anonymized JSON** (`x-structure-from-redacted-live-run: true` under `application/json` on a **200**/**201** response, from NDJSON merge): **96** / **187** operations have at least one such block; **119** total marker occurrences in the file (some operations have more than one marked block).
- **Operations without any such marker yet:** **91** (see *Gaps to full live-shape coverage*). For **mergeable** `200`/`201` JSON operations vs NDJSON keys, see `docs/internal/generated/NDJSON_OPENAPI_GAP.md`.

### Current status (short)

- Spec + generated client surface: **187** operations.
- Test heuristic (string match in `test/**/*.go`): **187** referenced, **0** missing references.
- Marker occurrences in YAML (includes multi-block operations): **119**.

### Gaps to full live-shape coverage

“Full” here means every operation’s successful **`application/json`** **`200` or `201`** response has been logged, merged, and marked with `x-structure-from-redacted-live-run`. Typical reasons the remaining **91** operations are not marked:

- Operation never called during integration + sweep (or only with **4xx** / **401** in your tenant).
- GET sweep **skip** (no path match / not in sweep list) or **fail** (network, rate limit, permissions).
- Response is **not JSON** or success code is not 200/201 (e.g. **204**, **302**, binary, HTML error pages).
- **Writes** not executed: opt-in `XURRENT_ALLOW_MUTATIONS=1` (and related flags) adds POST/PATCH bodies to NDJSON.
- Merge only applies when OpenAPI already has `content` → `application/json` with `schema:` or `application/json: {}` under that status.

Doc parity, **`components/schemas`**, and strict typing are separate from this counter; see `docs/ROADMAP.md`.

### Suggested next steps

1. Re-run integration tests with `XURRENT_STRUCTURE_LOG_FILE` and env flags for mutations/coverage as needed.
2. Run `python utils/collect_live_response_shapes.py` (see script help); append to NDJSON, then `python utils/merge_skeleton_into_openapi.py …`.
3. Fix tenant/token gaps that cause sweep **fail** (e.g. service create 401); set optional test IDs from integration test headers.
4. After merge: `python -c "import yaml; yaml.safe_load(open('openapi/openapi.yaml', encoding='utf-8'))"` to confirm YAML is valid.
5. List exact operations still missing from the capture: `python utils/report_ndjson_openapi_gaps.py --ndjson test/integration/_live_shapes/combined_for_gap_report.jsonl --write --only-mergeable` (build `combined_for_gap_report.jsonl` by concatenating NDJSON logs such as `full_run.jsonl` + `entity_get_by_id.jsonl`).
6. Service spine check (permissions): `XURRENT_PHASE0_POST_SERVICES=1 go test -tags=integration ./test/integration/ -run TestPhase0_PostServicesSpine -v` (see `docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md`).
7. Full live capture: `python utils/live_capture_pipeline.py --fresh -v` (loads `.env`; merges into OpenAPI; refreshes `docs/internal/generated/NDJSON_OPENAPI_GAP.md`).

Latest local capture artifacts (paths under `test/integration/_live_shapes/`, gitignored):

- `test/integration/_live_shapes/integration_test_run.log` — integration tests with structure logging.
- `test/integration/_live_shapes/collect_sweep.log` — `utils/collect_live_response_shapes.py` GET sweep.
- `test/integration/_live_shapes/full_run.jsonl` — combined NDJSON input to `utils/merge_skeleton_into_openapi.py`.
- `test/integration/_live_shapes/entity_get_by_id.jsonl` — optional focused run for `GetTeamsId` / `GetSitesId` / `GetSlasId` / `GetTasksId` list→id capture.

## Untested operationIds

| Method | Path | operationId |
|--------|------|-------------|

## Integration coverage

- `test/integration/production_smoke_test.go` — subset of authenticated GETs + rate limit; `GET /v1/problems/{id}` when a problem id is available (env, `parent_fixture.env`, or problem list discovery).
- `test/integration/mutations_test.go` — opt-in `XURRENT_ALLOW_MUTATIONS=1` for POST/PATCH (person + product).
- `test/integration/high_value_mutations_test.go` — opt-in `XURRENT_ALLOW_MUTATIONS=1` for teams, sites, services, SLAs (POST/PATCH + disable cleanup); chained calendar → service → service offering → request template (`TestMutation_Chained_CalendarServiceOfferingRequestTemplate`); optional workflow create/patch/trash when `XURRENT_TEST_WORKFLOW_TEMPLATE_ID` is set (`TestMutation_Chained_WorkflowCreateTrash`).
- `test/integration/lifecycle_mutations_test.go` — opt-in `XURRENT_ALLOW_MUTATIONS=1` and `XURRENT_ALLOW_LIFECYCLE_MUTATIONS=1` for archive/trash/restore on workflows (incl. `PostWorkflowsIdTasks`), problems, requests, people; POST→GET fixtures with cleanup.
- `test/integration/coverage_get_test.go` — opt-in `XURRENT_ENDPOINT_COVERAGE=1` for broad GET smoke (`TestCoverage_HighValueGETs`: workflows with list/env/fixture id → `GetWorkflowsId`, calendars, teams, sites, services, offerings, request templates, SLA lists; `TestCoverage_EntityGETByID`: `GetTeamsId`, `GetSitesId`, `GetSlasId`, `GetTasksId` via list discovery; filters for problems/requests/people; usage statements and product CIs where permitted).
- `utils/collect_live_response_shapes.py` — optional authenticated GET sweep → NDJSON; merge with `utils/merge_skeleton_into_openapi.py` (fills `application/json: {}` and existing `schema:` blocks with anonymized shapes). **Merge only within the same OpenAPI operation** (verb block); otherwise a `POST` without `200` could incorrectly match a sibling `GET` `200` and corrupt YAML.

Regenerate this file:

```bash
python utils/audit_endpoint_test_coverage.py --write
```

