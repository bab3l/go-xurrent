# go-xurrent roadmap

This project has **no downstream SDK users yet**. The OpenAPI spec and generated Go API may **change freely** until you publish a **v1** (or first stable) release and document compatibility expectations.

---

## Completed (foundation)

- **Spec hygiene**: Path normalization (`utils/fix_spec.py`), optional localization (`utils/localize_spec.py`), duplicate path parameters fixed (distinct names for nested `{id}` segments).
- **Generation**: Docker **OpenAPI Generator** pinned in `utils/generate_client.ps1` (`v7.8.0`), root module `github.com/bab3l/go-xurrent` with `isGoSubmodule=false`.
- **Quality gate**: GitHub Actions CI — `go vet`, `go test`, OpenAPI `validate` (`.github/workflows/ci.yml`).
- **Tests**: Generated test stubs compile; `testify` in `go.mod`; integration remains skipped by default.
- **Parallel ownership**: `docs/PARALLEL_WORKSTREAMS.md` (streams A–G by tag).
- **Route validation + inventory**: `docs/ROUTE_VALIDATION_CHECKLIST.md` (full **112**-operation table, **official API reference**: [developer.xurrent.com/v1](https://developer.xurrent.com/v1)).
- **Phase 1a (validation baseline):** Batches V1–V7 doc-cross-checked; reservations API completed in spec (`GET/POST /reservations`, filters `open`/`completed`, `GET/PATCH /reservations/{id}`).
- **Phase 1b (`operationId`):** `utils/add_operation_ids.py` assigns unique snake_case IDs; `utils/generate_api.ps1` runs it after `fix_spec.py`. Go methods renamed to match (e.g. `GetRequests`, `PostReservations`).

---

## Phase 1a — Parallel route validation — **baseline complete**

Further routes from [official docs](https://developer.xurrent.com/v1/) (e.g. request archive/trash/restore, extra filters, people lifecycle) remain **backlog** — see checklist intro.

**Tooling:** `python utils/extract_openapi_paths.py` — list paths; `python utils/export_route_checklist.py` — regenerate inventory markdown.

---

## Phase 1b — `operationId` — **complete**

- **Script:** `utils/add_operation_ids.py` (idempotent).
- **Convention:** `{method}_{path_slug}` (e.g. `get_requests_open`). Change an ID only when you accept Go client renames.

---

## Phase 2 — Developer experience

- **README** aligned with go-netbox style: install, auth, example snippet, link to roadmap and official docs.
- **Versioning policy**: e.g. SDK tags `v0.x` track spec iterations.
- **Linting** (optional): `golangci-lint` in CI.

---

## Phase 3 — Validation without burning API quota

- **CI:** Offline tests only.
- **Manual / scheduled:** Env-gated integration tests; throttle; read-only first.

---

## Phase 4 — Release

- First **tagged release** when ready; **CHANGELOG** with generator version and spec milestones.

---

## Recommended order (current)

1. **Backlog coverage:** Add high-value missing routes from official docs (see checklist / V2 notes).
2. **`components/schemas`** for shared entities; reduce anonymous `object` blobs.
3. **Global `security`** + header naming (`X-Xurrent-Account` vs examples).
4. **README / semver** when publishing a release.

---

## Related

- `SUMMARY.md` — generation pipeline.
- `docs/PARALLEL_WORKSTREAMS.md` — tag-based ownership (aligns with V1–V7 batches).
- `docs/ROUTE_VALIDATION_CHECKLIST.md` — validation status and findings.
