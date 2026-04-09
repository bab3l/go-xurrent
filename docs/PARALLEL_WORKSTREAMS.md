# Parallel workstreams (multi-agent / multi-owner)

Use this to assign **non-overlapping** ownership of `openapi/openapi.yaml` by **tag** (each tag maps to one root `api_*.go` file). Merge **one stream at a time** into `master`, then run `utils/generate_api.ps1` (or `utils/generate_client.ps1` after `fix_spec.py` if you edit raw exports).

**Roadmap context:** **[ROADMAP.md](ROADMAP.md)** — Phase **1a/1b** baselines are complete; use **[ROUTE_VALIDATION_CHECKLIST.md](ROUTE_VALIDATION_CHECKLIST.md)** for the full operation inventory. Breaking changes to the spec and generated SDK are acceptable until a stable release is published.

| Stream | Tags / API files | Primary work |
|--------|------------------|--------------|
| **A — Core** | General, Account, Search (`api_general.go`, `api_account.go`, `api_search.go`) | Global enums, `/v1/me`, search params; candidate for **global `security`** later. |
| **B — Requests & workflow** | Requests, Notes, Tasks, Problems, Events | List filters, pagination headers, nested routes; highest call volume if live-tested—**throttle**. |
| **C — People & org** | People, Organizations, Teams, Sites | Nested `/people/{id}/…` schemas. |
| **D — CMDB** | Configuration Items, Configuration Item Relations, Products, Product Categories | Relation paths; coordinate with **shared `components/schemas`**. |
| **E — Service catalog** | Services, Service Offerings, SLAs, Request Templates, Calendars, Workflows | Read-heavy; safer for early integration checks. |
| **F — Platform** | Automation Rules, Custom Collections, UI Extensions, Import, Export | Import/export may be expensive—**few** live calls. |
| **G — Attachments & reservations** | Attachments, Reservations | Multipart; confirm **reservations** URL against official docs. |

**API limits:** Prefer doc-first validation; batch live checks; use env-gated tests (`XURRENT_*`) and spacing between runs. **Do not** run unbounded parallel smoke tests against production.

**Foundation (shared):** CI (`.github/workflows/ci.yml`), pinned generator image in `utils/generate_client.ps1`, `isGoSubmodule=false` for root module `github.com/bab3l/go-xurrent`.
