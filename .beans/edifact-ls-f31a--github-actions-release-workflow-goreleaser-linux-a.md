---
# edifact-ls-f31a
title: GitHub Actions release workflow (goreleaser, Linux amd64+arm64)
status: in-progress
type: feature
priority: normal
created_at: 2026-09-01T15:38:35Z
updated_at: 2026-09-01T15:45:56Z
parent: edifact-ls-u9x8
blocked_by:
    - edifact-ls-sbsx
---

# Description
Automate cutting a release: pushing a `vX.Y.Z` tag cross-compiles
linux/amd64 and linux/arm64 binaries and publishes them as GitHub Release
assets, using goreleaser (the standard tool for this in the Go ecosystem --
handles cross-compilation, archiving, checksums, and release creation from
one config file) driven by a GitHub Actions workflow.

# Acceptance Criteria
- [x] `.goreleaser.yaml` (goreleaser's own default naming; the story said
      `.yml`) builds `./cmd/edifact-ls` for linux/amd64 and linux/arm64,
      injecting the version (edifact-ls-sbsx) into each binary, producing
      archives + a checksums file
- [x] `.github/workflows/release.yml` runs goreleaser on `v*` tag pushes
      using the repo's default `GITHUB_TOKEN` (no extra secrets needed)
- [x] README documents how to cut a release (tag + push) and how to
      download/verify a release binary
- [ ] Verified end-to-end: a real tag pushed to the repo produces a GitHub
      Release with working Linux binaries attached
