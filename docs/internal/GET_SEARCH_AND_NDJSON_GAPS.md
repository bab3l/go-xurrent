# GET /v1/search — investigation and NDJSON gap categories

## GetSearch (`get_search`) — parameters and 400 responses

### Official documentation (authoritative)

Source: [Search | Xurrent API](https://developer.xurrent.com/v1/search/)

| Input | Purpose |
|--------|---------|
| **`q`** | Search string (required for a meaningful search). Examples: `conference`, `unix`, `windows`. |
| **`types`** | Optional. Comma-separated: `ci`, `knowledge-article`, `project-task`, `request`, `request-template`, `service`, `service-instance`, `task`. Shorthand e.g. `req`, `task`. |
| **`per_page`** | Page size; default 25; **max 100**. |
| **`on_behalf_of`** | Search as another person (id or email). |
| **Next page** | Uses header **`x-pagination-next-page`** with query param **`page=<token>`** — **not** the same as cursor pagination on most other collections. |

Official pagination example:

```http
GET /v1/search?per_page=2&q=windows
```

Then:

```http
GET /v1/search?page=<token>&per_page=2&q=windows
```

### Spec / client mismatch in this repo

- `openapi/openapi.yaml` lists **`search_after`** / **`search_before`** (generic collection conventions) on `/v1/search`. The **Search** doc describes **`page=`** continuation instead.
- The generated Go client has **`SearchAfter` / `SearchBefore`** on `GetSearch` but **no `Page`** field for the token described in the Search doc.
- For **first-page** calls, only **`q`** and **`per_page`** are needed; fixing `q` length is enough to avoid many 400s.

### Likely cause of previous `400 Bad Request` in tests

Integration used **`q=a`** (single character) and **`per_page=1`**. The public docs never use a one-letter query; some backends reject overly short or stopword-like queries. **Mitigation:** use documented-style queries (**`windows`**, **`test`**, **`unix`**) and **`per_page=25`** (or omit to use default).

If 400 persists after that:

1. Log the JSON **`message`** from the error body (the test now logs `GenericOpenAPIError` body when present).
2. Check tenant settings (search visibility, SLA coverage — search only returns records the API user can see per the Search doc introduction).

### Web / forums

No additional public detail was found beyond [developer.xurrent.com/v1/search/](https://developer.xurrent.com/v1/search/) and general [Pagination](https://developer.xurrent.com/v1/general/pagination/) (which describes **`search_after` / `search_before`** for **most** collections, **not** the Search-specific `page=` flow).

### Status

| Topic | Status |
|--------|--------|
| Core **`q`**, **`types`**, **`per_page`** | **Documented** — use as above. |
| Continuation via **`page=`** | **Documented** on Search page. Hand-maintained **`APIClient.GetSearchJSON`** (**`search_get.go`**) plus **`SearchNextPageToken`**, **`WithSearchPage`**, and **`APIClient.ForEachSearchPage`** (**`search_pagination.go`**) implement **`x-pagination-next-page`** → **`page=`**. The generated **`SearchAPIService.GetSearch`** still maps **SearchAfter** → **`search_after`** (OpenAPI parity) — use **`GetSearchJSON`** / **`ForEachSearchPage`** when you need **`page=`**. |
| Exact validation rules for **`q`** (min length, etc.) | **Not spelled out** in public docs — rely on **error `message`** from 400 responses. |
| If search still returns 400 after query fixes | Treat as **blocked for automation** until the API error body explains the rule — not “fully undocumented,” but **not machine-guessable**. |

---

## The ~89 NDJSON gaps — are most “chaining”?

**Short answer: a large share is chaining or write-heavy** (need a parent id, or a successful POST before a GET/PATCH/DELETE). The rest are special endpoints, audit, or tenant emptiness.

### Rough grouping (from `docs/internal/generated/NDJSON_OPENAPI_GAP.md`)

| Category | Approx. count | What it means | Practical mitigation |
|----------|---------------|---------------|------------------------|
| **Request subtree** | ~18 | `GET/POST/PATCH/...` under `/v1/requests/{id}/…` | Need a stable **open request id** (fixture or env). Sweep skips when no request in list. |
| **Problem subtree** | ~10 | Problems + lifecycle + link to request | Needs **POST problem** or env problem id; **PostServices 422** blocked chained tests that create a service first. |
| **Workflow subtree** | ~9 | `POST /workflows`, lifecycle, tasks | **`XURRENT_TEST_WORKFLOW_TEMPLATE_ID`** + token rights; or sweep seed workflow. |
| **CI / CMDB graph** | ~10 | `cis`, relations, users on CI | Needs **product**, **service**, **team** seeds; CMDB permissions. Mutation test **`TestMutation_CIReservationAutomation`** POSTs a CI and **archives** it in `defer` (after canceling the reservation); **readonly** gap sweeps still need **`XURRENT_GAPTEST_CI_ID`** or list→id discovery when the tenant has no suitable id (see `docs/internal/GAP_GET_CAPTURE_IDS.md`). |
| **Service / SLA / offering spine** | ~9 | `POST/PATCH/GET` services, slas, offerings | **Chaining** from org; blocked when **`PostServices` returns 422** until payload/account rules are fixed. |
| **Reservations** | ~3 | POST + GET/PATCH by id | Need reservation id from list or create. |
| **Audit under entity** | ~5 | `…/audit`, `…/cis` on contracts, invoices, FLSA, projects, service offerings | Need **existing entity id** with audit history (often non-empty tenant). |
| **Misc GET by id** | ~8 | automation_rules, custom_collections, ui_extensions, surveys questions, tasks `assigned_by_me`, etc. | **Discovery id** from list GET or env; some lists empty → skip. |
| **People lifecycle** | ~3 | archive / trash / restore | Needs **same-account** person (cross-account skipped in tests). |
| **Platform / batch** | ~5 | import, export, events, attachments | **Different semantics** (jobs, files); not fixable by simple GET sweep. |
| **Search** | 1 | `get_search` | Parameter / query shape (this doc). |

**Not primarily chaining:** import/export/events/attachments — **product workflows**, long-running or multipart.

### Conclusion

- **Yes — most gaps are either (1) chaining** (parent resource id, or POST before GET), **or (2) writes** not executed in integration, **or (3) empty tenant** for a collection.
- **Solutions** (in order): unblock **`PostServices`** (or provide env service id), set **env template/instance ids**, extend **integration tests** per subtree, improve **sweep seeds**, accept that **import/export/attachments** need dedicated harnesses.

See also: `docs/internal/LIVE_CAPTURE_GAP_CLOSURE_PLAN.md`, `docs/internal/generated/NDJSON_OPENAPI_GAP.md`.
