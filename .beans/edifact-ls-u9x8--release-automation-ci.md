---
# edifact-ls-u9x8
title: Release automation & CI
status: in-progress
type: epic
priority: normal
created_at: 2026-09-01T15:38:19Z
updated_at: 2026-09-01T15:39:22Z
parent: edifact-ls-gdt6
---

# Goal
Give edifact-ls real, versioned, downloadable releases and continuous
testing -- the prerequisite for any distribution channel (Mason and
beyond), and good hygiene on its own regardless of Mason.

Scope decision: Linux only for now (amd64 + arm64). No code signing.
Cross-platform (macOS/Windows) support is explicitly out of scope until
asked for.

# Acceptance Criteria
- [ ] A LICENSE file exists
- [ ] `lspserver.Version` is set at build time (ldflags), not hardcoded
- [ ] Pushing a `vX.Y.Z` tag automatically cross-compiles Linux amd64+arm64
      binaries and publishes them as GitHub Release assets
- [ ] `make test`/`make test-e2e` run automatically on every push/PR via
      GitHub Actions
- [ ] A real tagged release has been cut and its binary verified to run
