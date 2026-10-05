# Documentation

## Using this SDK

- **[../README.md](../README.md)** — install, configuration, authentication.
- **Official API reference:** [developer.xurrent.com/v1](https://developer.xurrent.com/v1/) (HTTP semantics, filters, examples).
- **Generated Markdown** (from OpenAPI Generator, same run as the Go client): `*API.md` and model `Get*` / `Post*` / `Patch*` pages in this folder — convenience only; the product docs site is authoritative.

## In this repository

| Document | Purpose |
|----------|---------|
| [ROADMAP.md](ROADMAP.md) | Project direction, spec vs docs parity, live-shape pipeline summary. |
| [PARALLEL_WORKSTREAMS.md](PARALLEL_WORKSTREAMS.md) | Parallel workstreams for documentation parity. |
| [ROUTE_VALIDATION_CHECKLIST.md](ROUTE_VALIDATION_CHECKLIST.md) | Operation inventory table (regenerate: `python utils/export_route_checklist.py > docs/ROUTE_VALIDATION_CHECKLIST.md`). |
| [internal/README.md](internal/README.md) | Maintainer-only: live capture, gap reports, generated tables under `internal/generated/`. |
