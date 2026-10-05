# go-xurrent roadmap

This project has **no downstream SDK users yet**. The OpenAPI spec and generated Go API may **change freely** until you publish a **v1** (or first stable) release and document compatibility expectations.

**Official API reference:** [developer.xurrent.com/v1](https://developer.xurrent.com/v1/)

**Documentation index:** **`docs/README.md`**. **Maintainers:** capture tooling, gap reports, and related notes under **`docs/internal/`** (see **`docs/internal/README.md`**). Regenerated tables use **`docs/internal/generated/`**; paths are centralized in **`utils/doc_paths.py`**.

---

## Snapshot: spec vs documentation (2026-04-11)

| Metric | Current (`openapi/openapi.yaml`) | Notes |
|--------|----------------------------------|--------|
| **Path templates** | **141** distinct `paths` keys | Run `python utils/extract_openapi_paths.py` for METHOD + path list |
| **Operations** | **179** (HTTP methods on those paths) | Each has `operationId`; regenerate checklist when the spec changes |
| **OpenAPI `tags`** | (varies with spec) | Maps to generated `api_*.go` files |
| **Doc nav coverage** | Partial | The public docs list **many** resource families that have **no** `/v1/...` entry in this spec yet (see *Missing resource families* below) |
| **Live response shapes in spec** | **118 / 179** operations | `x-structure-from-redacted-live-run` from NDJSON merge (see below); **61** operations not yet enriched from a captured 200/201 JSON body |

Regenerate the route checklist table after large spec edits:

`python utils/export_route_checklist.py`

**Endpoint test + shape audit:** `docs/internal/generated/ENDPOINT_TEST_AUDIT.md` (regenerate: `python utils/audit_endpoint_test_coverage.py --write`).

### Status: identification, implementation, testing, and OpenAPI “examples” (2026-04-11)

| Area | What it means here | Current state |
|------|--------------------|---------------|
| **Identification** | Finding **which** routes lack live shapes, tests, or doc parity | **Gaps:** `python utils/report_ndjson_openapi_gaps.py --write` compares NDJSON **(method,path)** keys to **OpenAPI** (179 ops). **Test coverage heuristic:** `python utils/audit_endpoint_test_coverage.py` — string match of `operationId` / Go names under `test/` (**179/179** referenced; stubs count). **Doc vs spec:** qualitative gap lists in this file (*Gap analysis*, *Missing resource families*). |
| **Implementation** | Spec + generated client + merge tooling | **179** ops in `openapi/openapi.yaml`; **179** generated methods. **Live-shape merge:** `utils/merge_skeleton_into_openapi.py` writes **redacted JSON skeletons** into `application/json` **response** `schema:` (and `application/json: {}` placeholders) and sets **`x-structure-from-redacted-live-run: true`**. It does **not** add OpenAPI **`example:`** fields — those remain hand-maintained in YAML where present. |
| **Testing** | Exercising the API so NDJSON captures 200/201 JSON | **Default:** offline unit tests; **integration** is env-gated. **Structure logging:** set `XURRENT_STRUCTURE_LOG_FILE` on integration runs. **Mutations:** `XURRENT_ALLOW_MUTATIONS=1` (+ lifecycle / coverage flags as documented). **Extra write/read shapes:** `XURRENT_WRITE_SHAPE_COVERAGE=1` runs `test/integration/write_shape_coverage_test.go` (search, PATCH maps, PUT request, notes, contacts, calendars; best-effort). **`live_capture_pipeline.py`** sets mutations, lifecycle, endpoint coverage, and **write-shape coverage** when running integration. |
| **Examples in `openapi.yaml`** | Two different concepts | **(1) Static** `example` / rich `schema` in YAML — **manual** (docs parity, backlog). **(2) Live-derived** response **schemas** — from **NDJSON merge** above (**118** ops as of last merge on disk); re-run capture + merge to raise toward **179** where the tenant returns JSON 200/201. |

### Live response shape pipeline (current status)

- **Capture:** Integration tests (`XURRENT_STRUCTURE_LOG_FILE`) plus optional `utils/collect_live_response_shapes.py` GET sweep → append NDJSON (e.g. `test/integration/_live_shapes/full_run.jsonl`; directory is gitignored).
- **Merge:** `utils/merge_skeleton_into_openapi.py` — fills `application/json` response `schema:` (and `application/json: {}` placeholders) from redacted skeletons; adds `x-structure-from-redacted-live-run: true`. **Patched 2026-04-10:** merge is scoped to a single OpenAPI **verb block** so a method without its own `200` cannot match a sibling `GET` `200` (avoids invalid YAML).
- **2026-04-11 capture:** GET sweep with local `XURRENT_GAPTEST_*` ids and **`--delay 2.0`** (after hourly burst recovery if needed) → **105** NDJSON rows merged; **118** operations with `x-structure-from-redacted-live-run` (see `docs/internal/generated/ENDPOINT_TEST_AUDIT.md`).
- **Validate:** `python -c "import yaml; yaml.safe_load(open('openapi/openapi.yaml'))"` after merge; copy stays in sync via the merge script (`api/openapi.yaml`).

### Next steps (live shapes & validation)

1. **Re-run capture** when the spec or tenant data changes; keep appending or truncating NDJSON as you prefer, then merge again.
2. **Reduce sweep `fail` / `skip`:** fix auth/permissions for the test account (e.g. `PostServices` 401), supply optional env IDs (`XURRENT_TEST_WORKFLOW_TEMPLATE_ID`, request template / service instance IDs where tests expect them), and extend the sweep list if new paths are added.
3. **Writes:** opt-in mutation, lifecycle, and **`write_shape_coverage`** tests add **POST/PATCH/PUT** responses to the log; many operations will stay uncovered until exercised or until you add targeted calls.
4. **Gap report:** `python utils/report_ndjson_openapi_gaps.py --write` compares OpenAPI `(method, path)` to NDJSON rows (200/201 + `json_skeleton`); use `--only-mergeable` to restrict to operations the merge script can patch. Regenerate after each **fresh** NDJSON capture so counts match disk.
5. **Live capture pipeline:** `python utils/live_capture_pipeline.py --fresh -v` runs integration (structure log + mutations + lifecycle + coverage + **write-shape coverage**) → GET sweep → merge → YAML check → gap report. Maintainer context: `docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md`.

### Gaps to “full” coverage (what that means)

| Layer | Full coverage means | Where we are | Main gaps |
|-------|---------------------|--------------|-----------|
| **Spec operations** | Every doc-relevant route exists in OpenAPI | **179** ops today | Doc parity (Phase 1c): missing families, weak schemas — see *Gap analysis* below |
| **Generated Go client** | One method per `operationId` | **179** methods | Same as spec; wrong/incomplete **types** where spec uses loose `object` |
| **Test file references** | Heuristic: name appears under `test/` | **179 / 179** | Stubs ≠ executed integration tests; CI is mostly offline |
| **Live JSON shape in spec** | 200/201 `application/json` enriched from real responses | **91 / 179** | Not called; non-JSON; only error status; sweep fail; tenant empty; **HEAD/OPTIONS**; duplicate path keys rare edge cases |

Treat **91/179** as **enrichment of response schemas** (from live JSON skeletons), not as “half the product tested.”

---

## Gap analysis (depth)

### 1. Missing or incomplete **core** resources (documented, high usage)

These appear in the official docs with full CRUD-style sections; our spec is **read-only**, **missing `GET /{id}`**, **missing write verbs**, or **missing nested routes**.

| Area | Doc highlights | Spec gap |
|------|----------------|----------|
| **Problems** | `GET /problems/:id`, lifecycle | **`GET /v1/problems/{id}`** absent (only `PATCH` on that path). Backlog added filters + archive/trash/restore; **single-resource GET** still missing. |
| **Tasks** | `POST /workflows/:workflow_id/tasks`, filters `finished`, `approval_by_me` | **No `POST /v1/workflows/{id}/tasks`**. Missing **`GET /v1/tasks/finished`**. Doc filter **`/tasks/approval_by_me`** vs spec **`assigned_by_me`** — verify naming against current API. |
| **Products** | `POST`, `PATCH`, filters `disabled` / `enabled` / `supported_by_my_teams` | Only **`GET /products`**, **`GET /products/{id}`** — **no POST/PATCH**, **no collection filters**. |
| **Product categories** | Typically list + get + writes | Spec is **GET-only**; missing writes if required by your integration. |
| **Configuration items (`/cis`)** | Large surface: writes, relations, users | Spec is relatively strong; cross-check **archive/trash/restore**, **audit** subpaths, and **predefined filters** against docs. |
| **Organizations** | May include archive/trash patterns | Spec has GET/POST/PATCH; confirm **lifecycle** endpoints from docs if needed. |
| **Teams / sites** | List + detail + PATCH | Backlog added **`GET/POST /teams`**, **`PATCH /teams/{id}`**, **`GET/POST /sites`**, **`PATCH /sites/{id}`**; validate **nested** routes (members, etc.) vs docs. |
| **Account** | Sub-resources **usage_statements**, **billable_users** | Spec has **`GET /account`** only — **no** `/account/usage_statements` or `/account/billable_users` paths. |
| **Automation rules** | Often `PATCH` / lifecycle | Spec is **GET-only** for collection + item. |
| **Custom collections / elements** | Writes + audit in docs | Spec is **GET-only**; docs reference **audit** and **collection_elements** subtrees. |
| **UI extensions** | Writes | Spec is **GET-only**. |
| **Notes** (standalone) | [Notes API](https://developer.xurrent.com/v1/notes/) | Request/task notes exist; **global Notes** resource may differ — verify. |

### 2. **Missing resource families** (listed in doc sidebar, **no** first-class path in spec)

The following are **not** represented as top-level `/v1/...` APIs in `openapi.yaml` today. Adding them is **large, parallelizable** work (new tags, new `api_*.go`, or grouped by domain):

- **Project / agile / collaboration:** Projects, Project templates, Project tasks, Sprints, Scrum workspaces, Product backlogs, Agile boards, Releases  
- **Knowledge:** Knowledge articles, Knowledge article templates  
- **Service model extensions:** Service instances, Service categories, Affected SLAs, Reservation offerings  
- **Financial / commercial:** Invoices, Contracts, Shop articles / shop order lines, Short URLs  
- **People & time:** Skill pools (API resource), Timesheets, Time entries, Time allocations, Out-of-office periods, Effort classes  
- **Governance:** Webhooks, Webhook policies, Broadcasts, Surveys, SLA coverage groups, SLA notification schemes, RFC types, Risks, Closure codes  
- **Account / platform:** PDF designs, App offerings / app instances (beyond links in docs), CTI, Mail (as API)  

*Strategy:* Treat these as **Phase 2+ product verticals** unless a consumer explicitly needs them; each vertical can be a **separate backlog file** under `openapi/backlog/` and a **separate workstream**.

### 3. **Schema & contract quality** (not route count)

| Topic | Issue |
|-------|--------|
| **Request/response bodies** | Many operations use anonymous `type: object` or empty `application/json: {}` — weak typing vs rich docs **Fields** sections. |
| **`components/schemas`** | Only a small reuse (e.g. `Organization`); most entities are not shared components — harder client ergonomics and doc sync. |
| **Account header** | Docs use **`X-Xurrent-Account`**; much of the spec still shows **`X-4me-Account`** (legacy). Align names + examples for generated parameter docs. |
| **Global `security`** | Bearer auth is not applied globally; many operations duplicate header parameters. |
| **Pagination / filtering** | Documented query params (`per_page`, `page`, filters) are **inconsistently** modeled per operation. |

### 4. **Generated library**

The Go client is **generated** from the spec: any **missing path** or **wrong method** in OpenAPI is a **missing or wrong method** on `APIClient`. **Hand-written** helpers belong in **`.openapi-generator-ignore`** paths (e.g. **`collection_filter.go`**, **`collection_get.go`**), not in **`client.go`** / **`api_*.go`**.

---

## Completed (foundation)

- **Spec hygiene**: Path normalization (`utils/fix_spec.py`), optional localization (`utils/localize_spec.py`), duplicate path parameters fixed.
- **Backlog merge**: `utils/merge_backlog_paths.py` + `openapi/backlog/*.yaml` for incremental doc parity.
- **Generation**: Docker **OpenAPI Generator** pinned in `utils/generate_client.ps1` (`v7.8.0`), `packageName=xurrent`.
- **Quality gate**: GitHub Actions — `go vet`, `go test`, OpenAPI `validate` (`.github/workflows/ci.yml`).
- **Tests**: Generated test stubs; `testify`; integration skipped by default.
- **`operationId`**: `utils/add_operation_ids.py` — unique snake_case IDs.

---

## Phase 1a — Parallel route validation — **baseline complete**

*(Historical)* Batches V1–V7 doc-cross-checked; reservations and core routes validated.

---

## Phase 1b — `operationId` — **complete**

---

## Phase 1c — Documentation parity (parallel implementation) — **in progress**

**Goal:** Close high-value gaps between [developer.xurrent.com](https://developer.xurrent.com/v1/) and `openapi/openapi.yaml`, using **`openapi/backlog/`** fragments + merge script, then regenerate.

**Shipped in spec (backlog `07`–`09`):** **P1** — `GET /v1/problems/{id}`, `POST /v1/workflows/{id}/tasks`, `GET /v1/tasks/finished`, `GET /v1/tasks/approval_by_me`. **P2** — `POST/PATCH /v1/products`, `GET /v1/products/disabled|enabled|supported_by_my_teams`. **P3** — `GET /v1/account/usage_statements`, `GET /v1/account/usage_statements/{id}`, `GET /v1/account/billable_users`, `GET /v1/rate_limit`. **Production smoke:** `go test -tags=integration ./test/integration/ -v` (see file header for env vars; default CI runs only offline tests).

**Collection conventions (reduce payload & calls):** `python utils/inject_api_conventions.py` adds optional **`per_page`**, **`search_after`**, **`search_before`**, **`fields`**, **`sort`**, **`state`** on collection `GET`s; **`X-Xurrent-Language`** on all `GET`s; **`If-None-Match`** on single-resource `GET /v1/{tag}/{id}`; response **`Link`** + **`X-Pagination-Throttled`** where applicable. Hand helpers in **`collection_filter.go`** build ad-hoc [filtering](https://developer.xurrent.com/v1/general/filtering/) query strings. Wired into **`utils/generate_api.ps1`** after `fix_spec.py`.

**Priority tiers (suggested):**

| Tier | Focus | Examples |
|------|--------|----------|
| **P1 — Correctness** | Single-resource reads + doc-named filters where behavior is unclear | `GET /v1/problems/{id}`; tasks `finished` vs `approval_by_me`; `POST /v1/workflows/{id}/tasks` |
| **P2 — CMDB & catalog writes** | Products, product categories, CIs as per docs | `POST/PATCH /products`, product filters; CI lifecycle if missing |
| **P3 — Platform reads** | Account sub-resources, rate limit helper | `GET /account/usage_statements`, `GET /account/billable_users`; `GET /rate_limit` if exposed |
| **P4 — Lifecycle & writes** | Archive/trash/restore + PATCH where docs say so | Organizations, automation rules, custom collections, UI extensions (batch by tag) |
| **P5 — New families** | Optional verticals | Knowledge articles, projects, webhooks — **one vertical per PR/stream** |

**Process (each batch):**

1. Add or extend `openapi/backlog/NN_name.yaml` (paths only).  
2. `python utils/merge_backlog_paths.py` → `python utils/inject_api_conventions.py` → `python utils/add_operation_ids.py`  
3. `openapi-generator-cli validate` → `utils/generate_client.ps1`  
4. Restore `go.mod` / `go.sum` if generator clobbers them; `go test ./...`  
5. Refresh `docs/ROUTE_VALIDATION_CHECKLIST.md` table (`export_route_checklist.py`) when the inventory should match CI.

---

## Phase 2 — Developer experience

- **README** aligned with go-netbox style: install, auth (**`X-Xurrent-Account`**), example snippet.
- **Versioning policy**: SDK tags `v0.x` track spec iterations.
- **Linting** (optional): `golangci-lint` in CI.

### GraphQL and other query interfaces (optional, out of band)

The **canonical** Xurrent HTTP API is **REST** with **OpenAPI**; **go-xurrent** stays centered on that contract (generated client + spec parity). There is **no** plan to replace REST with GraphQL inside this repository unless the **vendor** ships a supported GraphQL endpoint.

**If** you need GraphQL (e.g. aggregated reads, fewer round-trips for a custom UI or BFF):

| Approach | Role | Tradeoffs |
|----------|------|-----------|
| **REST + this SDK** | Default for integrations and MCP tools that map 1:1 to documented routes | Full parity work lives in **OpenAPI**; predictable for agents. |
| **Custom GraphQL gateway (BFF)** | Your service fronts **developer.xurrent.com** REST, exposes a schema you own | You maintain resolvers, auth token forwarding, **rate limits**, and **schema drift** when REST changes; does **not** remove the need for a correct REST/OpenAPI client for mutations and parity testing. |
| **MCP “query / discover” layer** | Described in **`xurrent-mcp/docs/PROJECT_PLAN_AND_ROADMAP.md`** — agent-optimized tools can wrap REST **`fields=`** + filters + pagination; GraphQL could **back** a single `query()`-style tool **only if** you already run a BFF | Same drift and security review as any gateway; prefer REST-first tools until a BFF exists. |

**Recommendation:** keep **go-xurrent** and spec work **REST-first**. Treat GraphQL as a **separate product** (optional BFF) with its own repo, tests, and versioning — not a second generated output from this OpenAPI spec unless you explicitly add codegen for it.

---

## Phase 3 — Validation without burning API quota

- **CI:** Offline tests only.
- **Manual / scheduled:** Env-gated integration tests; throttle; read-only first.
- **Live shapes:** Use NDJSON capture + merge (see *Live response shape pipeline*) to tighten response schemas without committing secrets; aim to raise **91 → 179** only where JSON 200/201 exists and is reachable from your tenant. **Targeted run:** `python utils/live_capture_pipeline.py --integration-profile capture-writes --go-test-run TestWriteShape_ …` (see `docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md`).

---

## Phase 4 — Release

- First **tagged release** when ready; **CHANGELOG** with generator version and spec milestones.

---

## Recommended order (updated)

1. **Phase 1c P1–P2** — fixes and CMDB/catalog writes that integrations hit most often.  
2. **Components/schemas** for shared entities (Request, Person, CI, Service) — reduces drift.  
3. **Header + security cleanup** — `X-Xurrent-Account`, global `security`, fewer duplicated parameters.  
4. **Live response shapes** — re-run NDJSON capture + merge when you extend tests or tenant access; see `docs/internal/generated/ENDPOINT_TEST_AUDIT.md` for counts and next steps.  
5. **Phase 1c P5** — new API families only when product requires them.  
6. **README / semver** when publishing.  
7. **GraphQL / BFF** — only if a product requirement appears; see *GraphQL and other query interfaces* under Phase 2 (do not block SDK work on it).

---

## Related

- **`xurrent-mcp/docs/PROJECT_PLAN_AND_ROADMAP.md`** — MCP agent patterns (layered tools vs code mode); **GraphQL** discussed only as an optional BFF layer — see *GraphQL and other query interfaces* above.  
- `docs/README.md` — what lives in `docs/` (roadmap, checklists, internal).  
- `SUMMARY.md` — generation pipeline.  
- `docs/PARALLEL_WORKSTREAMS.md` — **Phase 1c** parallel streams (updated).  
- `docs/ROUTE_VALIDATION_CHECKLIST.md` — operation inventory (**regenerate** after spec changes).  
- `docs/internal/generated/ENDPOINT_TEST_AUDIT.md` — operations vs tests vs live-shape merge (**regenerate:** `python utils/audit_endpoint_test_coverage.py --write`).  
- `docs/internal/generated/NDJSON_OPENAPI_GAP.md` — operations missing from NDJSON capture vs spec (**regenerate:** `python utils/report_ndjson_openapi_gaps.py --write`).  
- `docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md` — how we close capture gaps (fixtures, service spine, permissions).  
- `docs/internal/GET_SEARCH_AND_NDJSON_GAPS.md` — **GetSearch** quirks, 400 troubleshooting, NDJSON gap categories.  
- `utils/live_capture_pipeline.py` — integration + sweep + merge + gap report; `utils/live_capture_pipeline.ps1` on Windows.  
- `utils/README.md` — what each utility does and when to run it.
