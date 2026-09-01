---
# edifact-ls-u9x8
title: Release automation & CI
status: completed
type: epic
priority: normal
created_at: 2026-09-01T15:38:19Z
updated_at: 2026-09-01T16:15:04Z
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
- [x] A LICENSE file exists
- [x] `lspserver.Version` is set at build time (ldflags), not hardcoded
- [x] Pushing a `vX.Y.Z` tag automatically cross-compiles Linux amd64+arm64
      binaries and publishes them as GitHub Release assets
- [x] `make test`/`make test-e2e` run automatically on every push/PR via
      GitHub Actions
- [x] A real tagged release has been cut and its binary verified to run

## Summary of Changes
edifact-ls now has real, versioned, downloadable releases: an MIT LICENSE,
build-time version injection, a goreleaser-driven release workflow (Linux
amd64+arm64), and a CI workflow testing every push/PR. v0.1.1 is live on
GitHub with verified-working binaries.

## Retro
- Straightforward once each piece was scoped concretely up front (Linux
  only, no signing, custom Mason registry not upstream) -- no scope
  creep or rework needed mid-epic.
- Two real bugs surfaced by actually exercising the pipeline rather than
  just writing config and trusting it: the binary had no --version/--help
  handling at all (any arg fell through to RunStdio() and hung -- exactly
  what you'd hit trying to smoke-test a downloaded release binary), and the
  initially-pinned Action versions were already on a deprecated Node.js 20
  runtime. Both were things `actionlint` and local `goreleaser check`
  couldn't have caught -- only running the real thing did.
- goreleaser's default dist/ output directory would have silently mixed
  with make build's dist/edifact-ls; separated it (dist/goreleaser/) before
  it caused confusion, not after.
- SSH-vs-HTTPS remote auth from inside the sandbox was the one real
  friction point -- couldn't push directly (no SSH key/known_hosts in the
  sandbox), so the user pushed both `main` and the `v0.1.1` tag themselves.
  Worth remembering for any future work needing a real push from this repo.
- Mason packaging (edifact-ls-vxf5) is next and now unblocked -- real
  release assets exist for it to reference.
