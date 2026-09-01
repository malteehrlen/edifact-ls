---
# edifact-ls-upnf
title: Go LSP server skeleton (stdio, initialize/shutdown)
status: completed
type: feature
priority: high
created_at: 2026-09-01T13:02:21Z
updated_at: 2026-09-01T13:09:30Z
parent: edifact-ls-untk
---

# Description
Set up the Go module and a minimal LSP server binary speaking JSON-RPC over
stdio. No language features yet — just a correct handshake.

# Acceptance Criteria
- [x] `go.mod` + `cmd/edifact-ls` (or similar) builds a binary via `go build`
- [x] Server handles `initialize` (responds with capabilities, even if empty),
      `initialized`, `shutdown`, `exit` per LSP spec
- [x] Unhandled requests don't crash the process (return method-not-found)
- [x] Unit test drives the server over an in-memory pipe through the handshake
      and asserts correct responses
- [x] `go vet`/`go build` clean; basic `Makefile` or `go run` target to build

## Summary of Changes
Implemented `internal/lspserver` wrapping `github.com/tliron/glsp`: handles
`initialize`/`initialized`/`shutdown`/`exit`, tracks whether shutdown preceded
exit to compute the process exit code per spec. `cmd/edifact-ls` is the stdio
entrypoint. Added a `Makefile` with `build`/`test`/`test-e2e` targets.
Handshake is covered by an in-memory-pipe unit test in lspserver_test.go
(uses two os.Pipe pairs standing in for stdin/stdout, a real jsonrpc2 client
on the other end) that also asserts unhandled methods return method-not-found
instead of crashing.
