# go-xurrent roadmap

This project has **no downstream SDK users yet**. The OpenAPI spec and generated Go API may **change freely** until you publish a **v1** (or first stable) release and document compatibility expectations.

**Official API reference:** [developer.xurrent.com/v1](https://developer.xurrent.com/v1/)

---

## Snapshot: spec vs documentation (2026-04)

| Metric | Current (`openapi/openapi.yaml`) | Notes |
|--------|----------------------------------|--------|
| **Path templates** | **123** distinct `paths` keys | Run `python utils/extract_openapi_paths.py` for METHOD + path list |
| **Operations** | **158** (HTTP methods on those paths) | Each has `operationId` via `utils/add_operation_ids.py` |
| **OpenAPI `tags`** | 28 tag groups | Maps 1:1 to generated `api_*.go` files |
| **Doc nav coverage** | Partial | The public docs list **many** resource families that have **no** `/v1/...` entry in this spec yet (see *Missing resource families* below) |

The embedded table in `docs/ROUTE_VALIDATION_CHECKLIST.md` is **stale** (it still says **112** operations). Regenerate the inventory with:

`python utils/export_route_checklist.py`

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

The Go client is **generated** from the spec: any **missing path** or **wrong method** in OpenAPI is a **missing or wrong method** on `APIClient`. Non-generated **hand-written** helpers are out of scope unless added deliberately.

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
2. `python utils/merge_backlog_paths.py` → `python utils/add_operation_ids.py`  
3. `openapi-generator-cli validate` → `utils/generate_client.ps1`  
4. Restore `go.mod` / `go.sum` if generator clobbers them; `go test ./...`  
5. Refresh `docs/ROUTE_VALIDATION_CHECKLIST.md` table (`export_route_checklist.py`) when the inventory should match CI.

---

## Phase 2 — Developer experience

- **README** aligned with go-netbox style: install, auth (**`X-Xurrent-Account`**), example snippet.
- **Versioning policy**: SDK tags `v0.x` track spec iterations.
- **Linting** (optional): `golangci-lint` in CI.

---

## Phase 3 — Validation without burning API quota

- **CI:** Offline tests only.
- **Manual / scheduled:** Env-gated integration tests; throttle; read-only first.

---

## Phase 4 — Release

- First **tagged release** when ready; **CHANGELOG** with generator version and spec milestones.

---

## Recommended order (updated)

1. **Phase 1c P1–P2** — fixes and CMDB/catalog writes that integrations hit most often.  
2. **Components/schemas** for shared entities (Request, Person, CI, Service) — reduces drift.  
3. **Header + security cleanup** — `X-Xurrent-Account`, global `security`, fewer duplicated parameters.  
4. **Phase 1c P5** — new API families only when product requires them.  
5. **README / semver** when publishing.

---

## Related

- `SUMMARY.md` — generation pipeline.  
- `docs/PARALLEL_WORKSTREAMS.md` — **Phase 1c** parallel streams (updated).  
- `docs/ROUTE_VALIDATION_CHECKLIST.md` — operation inventory (**regenerate** after spec changes).
