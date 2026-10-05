# GET gaps — IDs to supply for readonly capture

Use this with a **readonly** production token and **`XURRENT_ACCOUNT`** for that account. Goal: one successful **200** JSON response per **`(method, path)`** so NDJSON + merge can fill live shapes.

**You do not need a separate ID per row** where paths share the same `{id}` placeholder — one numeric id per **entity type** covers every GET under that resource.

---

## A. No path `{id}` (3) — optional env / run as-is

These are **collection or predefined-filter** GETs; we mainly need **any 200** body in the log.

| operationId | Path |
|-------------|------|
| `get_problems` | `GET /v1/problems` |
| `get_requests_requests_of_my_organization` | `GET /v1/requests/requests_of_my_organization` |
| `get_tasks_assigned_by_me` | `GET /v1/tasks/assigned_by_me` |

**Your input:** not required unless calls fail (permissions). If they return **200** with `[]`, merge may still be weak — prefer tenants with data.

---

## B. Special — not a stable numeric “entity id”

| operationId | Path | What to supply |
|-------------|------|----------------|
| `get_import` | `GET /v1/import` | *(none)* — call as documented. |
| `get_import_token` | `GET /v1/import/{token}` | **Token string** from a real import job in that account (if you have one). |
| `get_attachments_upload` | `GET /v1/attachments/upload` | Often **query params** per docs — share an example URL or params if the API requires them. |

---

## C. One ID per entity type (grouped)

Provide **one numeric id** per line (or “none / skip” if the account has no such record).

| # | Your ID (numeric unless noted) | Covers these GET gaps (`operationId`) |
|---|------------------------------|----------------------------------------|
| 1 | **automation_rule_id** | `get_automation_rules_id` |
| 2 | **ci_id** | `get_cis_id`, `get_cis_id_ci_relations`, `get_cis_id_users` *(see note below)* |
| 3 | **contract_id** | `get_contracts_id_audit`, `get_contracts_id_cis` |
| 4 | **custom_collection_id** | `get_custom_collections_id` |
| 5 | **custom_collection_element_id** | `get_custom_collection_elements_id` |
| 6 | **flsa_id** | `get_flsas_id_audit` |
| 7 | **invoice_id** | `get_invoices_id_audit`, `get_invoices_id_cis` |
| 8 | **person_id** | `get_people_id_ci` *(person should have CI linkage if possible)* |
| 9 | **problem_id** | `get_problems_id` |
| 10 | **project_id** | `get_projects_id_audit` |
| 11 | **request_template_id** | `get_request_templates_id` |
| 12 | **request_id** | `get_requests_id`, `get_requests_id_audit`, `get_requests_id_cis`, `get_requests_id_cis_active`, `get_requests_id_cis_inactive`, `get_requests_id_grouped_requests`, `get_requests_id_notes`, `get_requests_id_notes_internal`, `get_requests_id_notes_public` |
| 13 | **reservation_id** | `get_reservations_id` |
| 14 | **service_offering_id** | `get_service_offerings_id`, `get_service_offerings_id_audit` |
| 15 | **service_id** | `get_services_id` |
| 16 | **sla_id** | `get_slas_id` |
| 17 | **survey_id** | `get_surveys_id_survey_questions` |
| 18 | **ui_extension_id** | `get_ui_extensions_id` |
| 19 | **workflow_id** | `get_workflows_id` |
| 20 | **service_instance_id** | `get_service_instances_id` |

**Minimum count:** up to **20 numeric ids** + optional **import token** + Section A/B notes.

**CI note:** With a **writable** token and `XURRENT_ALLOW_MUTATIONS=1`, `TestMutation_CIReservationAutomation` creates a CI and archives it after the run (no long-lived test CI). That run still **does not replace** pinning **`XURRENT_GAPTEST_CI_ID`** for a **readonly** sweep if you need stable `get_cis_*` rows without mutations, or if the mutation test skips (e.g. no service id — set `XURRENT_TEST_SERVICE_ID` or run `TestFixture_ParentIDsForCapture` with `PostServices` succeeding).

---

## D. Suggested env names (for scripts / sweep seeds)

Prefix is arbitrary; example:

```text
XURRENT_GAPTEST_AUTOMATION_RULE_ID=
XURRENT_GAPTEST_CI_ID=
XURRENT_GAPTEST_CONTRACT_ID=
XURRENT_GAPTEST_CUSTOM_COLLECTION_ID=
XURRENT_GAPTEST_CUSTOM_COLLECTION_ELEMENT_ID=
XURRENT_GAPTEST_FLSA_ID=
XURRENT_GAPTEST_INVOICE_ID=
XURRENT_GAPTEST_PERSON_ID=
XURRENT_GAPTEST_PROBLEM_ID=
XURRENT_GAPTEST_PROJECT_ID=
XURRENT_GAPTEST_REQUEST_TEMPLATE_ID=
XURRENT_GAPTEST_REQUEST_ID=
XURRENT_GAPTEST_RESERVATION_ID=
XURRENT_GAPTEST_SERVICE_OFFERING_ID=
XURRENT_GAPTEST_SERVICE_INSTANCE_ID=
XURRENT_GAPTEST_SERVICE_ID=
XURRENT_GAPTEST_SLA_ID=
XURRENT_GAPTEST_SURVEY_ID=
XURRENT_GAPTEST_UI_EXTENSION_ID=
XURRENT_GAPTEST_WORKFLOW_ID=
XURRENT_GAPTEST_IMPORT_TOKEN=
```

---

## E. Flat list of GET endpoints (reference)

Total **39** GET rows in the current gap table; grouped above to reduce duplicate IDs.

<details>
<summary>All GET paths in gap list</summary>

- `GET /v1/attachments/upload`
- `GET /v1/automation_rules/{id}`
- `GET /v1/cis/{id}` — `GET /v1/cis/{id}/ci_relations` — `GET /v1/cis/{id}/users`
- `GET /v1/contracts/{id}/audit` — `GET /v1/contracts/{id}/cis`
- `GET /v1/custom_collection_elements/{id}` — `GET /v1/custom_collections/{id}`
- `GET /v1/flsas/{id}/audit`
- `GET /v1/import` — `GET /v1/import/{token}`
- `GET /v1/invoices/{id}/audit` — `GET /v1/invoices/{id}/cis`
- `GET /v1/people/{id}/ci`
- `GET /v1/problems` — `GET /v1/problems/{id}`
- `GET /v1/projects/{id}/audit`
- `GET /v1/request_templates/{id}`
- `GET /v1/requests/requests_of_my_organization`
- `GET /v1/requests/{id}` + audit, cis, cis/active, cis/inactive, grouped_requests, notes, notes/internal, notes/public
- `GET /v1/reservations/{id}`
- `GET /v1/service_offerings/{id}` — `GET /v1/service_offerings/{id}/audit`
- `GET /v1/service_instances/{id}`
- `GET /v1/services/{id}`
- `GET /v1/slas/{id}`
- `GET /v1/surveys/{id}/survey_questions`
- `GET /v1/tasks/assigned_by_me`
- `GET /v1/ui_extensions/{id}`
- `GET /v1/workflows/{id}`

</details>

---

## F. Rate limits and sweep timing

The API returns **429** when burst or hourly quotas are exhausted. If a sweep logs almost all **`fail`** rows, wait for **`Retry-After`** (or ~15 minutes after a heavy run), then re-run with a larger delay, e.g. **`--delay 2.0`** or **`XURRENT_SWEEP_DELAY=2`**. The collector also **retries 429** on each GET (sleeps per `Retry-After`, up to 8 attempts).

---

## G. Last full GET capture (2026-04-11, local)

Using **`XURRENT_GAPTEST_*`** in `.env` (see section D) and **`python utils/collect_live_response_shapes.py --truncate --delay 2.0 --out test/integration/_live_shapes/gap_run.jsonl`** (after token bucket recovery):

| Metric | Value |
|--------|------:|
| NDJSON rows (**200** + `json_skeleton`) | **105** |
| OpenAPI ops with merged live shape marker after merge | **118** (`x-structure-from-redacted-live-run`; see `docs/internal/generated/ENDPOINT_TEST_AUDIT.md`) |
| Sweep stats (stderr) | **ok=105**, **skip=2**, **fail=10** |

**Still missing from this NDJSON** (among others): writes (**POST/PATCH/…**), **`get_search`**, **`get_import`**, **`get_attachments_upload`**, several **GET `{id}`** rows that returned **4xx** in this tenant (e.g. **`get_automation_rules_id`**, **`get_cis_id`** and CI subpaths — list endpoints succeeded; pinned id may not resolve), **`get_people_id_ci`**, **`get_requests_requests_of_my_organization`**, **`get_tasks_assigned_by_me`**, **`get_reservations_id`**. Regenerate the exact list:  
`python utils/report_ndjson_openapi_gaps.py --ndjson test/integration/_live_shapes/gap_run.jsonl --write`

---

*Regenerate gap list after capture: `python utils/report_ndjson_openapi_gaps.py --write`*
