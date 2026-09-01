---
# edifact-ls-sbsx
title: LICENSE + build-time version injection
status: completed
type: feature
priority: normal
created_at: 2026-09-01T15:38:35Z
updated_at: 2026-09-01T15:40:16Z
parent: edifact-ls-u9x8
---

# Description
Two small prerequisites for a real release: a LICENSE file (needed before
distributing binaries or listing on any registry), and making
`lspserver.Version` reflect the actual release tag instead of the
hardcoded "0.0.1".

Default to MIT unless told otherwise -- it's the de facto standard for
small Go CLI/LSP tools and everything else in this ecosystem (nvim, most
Mason packages) assumes a permissive license. Easy to swap later; the
license text doesn't couple to anything else in this epic.

# Acceptance Criteria
- [x] `LICENSE` file added at repo root (MIT, unless directed otherwise)
- [x] `go build -ldflags "-X .../internal/lspserver.Version=$TAG"` (or
      equivalent) sets the version reported in `initialize`'s `serverInfo`;
      `make build` still produces a sane default (e.g. "dev") when built
      without that flag
- [x] Verified: a binary built with an explicit version flag reports it via
      the LSP `initialize` response

## Summary of Changes
Added `LICENSE` (MIT, defaulted per the story's own rationale -- not
explicitly directed otherwise). `internal/lspserver.Version` now defaults to
"dev" and is set via `-ldflags "-X .../internal/lspserver.Version=$(VERSION)"`,
wired into `Makefile`'s new `VERSION` variable (`make build VERSION=v1.2.3`)
and used by both `build` and `install` targets.

Verified end-to-end with a real LSP handshake (a small Python script
speaking the Content-Length-framed JSON-RPC protocol directly to a binary
built with `VERSION=v1.2.3-test`): `initialize`'s response reported
`serverInfo.version == "v1.2.3-test"`. Full test suite and e2e harness both
still pass with the default `VERSION=dev` build.
