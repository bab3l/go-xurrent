# Parallel workstreams (multi-agent / multi-owner)

Use this to assign **non-overlapping** ownership of `openapi/openapi.yaml` by **tag** (each tag maps to one root `api_*.go` file). Merge **one stream at a time** into **`main`**, then run `utils/generate_api.ps1` (or `utils/generate_client.ps1` after `fix_spec.py` if you edit raw exports).

**Roadmap context:** **[ROADMAP.md](ROADMAP.md)** — Phase **1c** (documentation parity) is the active parallel track; see **gap analysis** there for what remains vs [developer.xurrent.com/v1](https://developer.xurrent.com/v1/).

**API limits:** Prefer doc-first validation; batch live checks; use env-gated tests (`XURRENT_*`) and spacing between runs. **Do not** run unbounded parallel smoke tests against production.

---

## Streams A–G (tag ownership, stable)

| Stream | Tags / API files | Primary work |
|--------|------------------|--------------|
| **A — Core** | General, Account, Search (`api_general.go`, `api_account.go`, `api_search.go`) | Enums, `/v1/me`, search; **gap:** `/account/*` subpaths, `rate_limit`, global **`security`** + **`X-Xurrent-Account`** alignment. |
| **B — Requests & workflow** | Requests, Notes, Tasks, Problems, Events | Filters, nested routes; **gaps:** `GET /problems/{id}`, `POST /workflows/{id}/tasks`, task filters `finished` / `approval_by_me` vs spec. |
| **C — People & org** | People, Organizations, Teams, Sites | Nested `/people/{id}/…`; **gaps:** org lifecycle if documented; teams/sites backlog validation. |
| **D — CMDB** | Configuration Items, Configuration Item Relations, Products, Product Categories | **Gaps:** product **POST/PATCH** + filters (`disabled`, `enabled`, `supported_by_my_teams`); CI doc cross-check. |
| **E — Service catalog** | Services, Service Offerings, SLAs, Request Templates, Calendars, Workflows | Read/write parity; workflow **POST/PATCH** + lifecycle (in backlog); service instances / categories **not in spec** (future vertical). |
| **F — Platform** | Automation Rules, Custom Collections, UI Extensions, Import, Export | **Gaps:** PATCH/write paths for automation rules, custom collections, UI extensions per docs. |
| **G — Attachments & reservations** | Attachments, Reservations | Multipart; reservation paths vs docs. |

---

## Phase 1c — Parallel implementation batches (gap closure)

Work in **separate branches/PRs** by batch so merges stay reviewable. Each batch: `openapi/backlog/NN_*.yaml` → merge → `add_operation_ids` → validate → generate → test.

| Batch | Owner focus | Doc / spec gap theme | Suggested stream mapping |
|-------|-------------|----------------------|---------------------------|
| **1c-1** | **P1 correctness** | `GET /problems/{id}`; `POST /workflows/{id}/tasks`; task filters (`/tasks/finished`, doc `approval_by_me`) | **B**, **E** |
| **1c-2** | **P2 CMDB & catalog** | Product **POST/PATCH** + list filters; product category writes if required | **D** |
| **1c-3** | **P3 account & general** | `/account/usage_statements`, `/account/billable_users`; `/rate_limit` if applicable | **A** |
| **1c-4** | **P4 lifecycle** | Automation rules, custom collections, UI extensions — PATCH/archive patterns per docs | **F** |
| **1c-5** | **P5 new families** | Pick **one** vertical (e.g. knowledge articles **or** webhooks) — full subtree only when needed | **New tag / new `api_*.go`** |

Batches **1c-1** and **1c-2** can run in **parallel** (disjoint path prefixes: `problems`/`workflows`/`tasks` vs `products`). Batches touching **`ci.yml`** or **generator** tooling stay **serial** with repo maintainers.

---

## Foundation (shared)

- CI (`.github/workflows/ci.yml`), pinned generator image in `utils/generate_client.ps1`, `isGoSubmodule=false` for the root Go module.  
- Dependabot: `.github/dependabot.yml` (Go modules + GitHub Actions).
