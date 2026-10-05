# Changelog

All notable changes to this project will be documented here.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html) once a stable **v1** SDK release is published (see `docs/ROADMAP.md`).

## [Unreleased]

### Added

- CI: Go **1.26** (`check-latest: true`). `golangci-lint` (hand-maintained and integration-tagged code only; generated `api_*` / `model_*` skipped).
- CI: `govulncheck` for module vulnerability scanning.
- Dependabot: grouped updates for Go modules and GitHub Actions.
- `.openapi-generator-ignore`: preserve `README.md`, `go.mod`, `go.sum`, and hand-maintained `collection_*.go` (filter + HTTP helpers) across regenerations.

### Changed

- `docs/`: removed redundant OpenAPI Generator Markdown (per-endpoint and model pages); canonical reference is [developer.xurrent.com/v1](https://developer.xurrent.com/v1/). Added `docs/README.md` index.
- `README.md`: replaced generated endpoint/model link tables with short pointers to official docs and repo `docs/`.
