---
# edifact-ls-ndhz
title: CLI check subcommand
status: completed
type: feature
priority: low
created_at: 2026-09-01T17:03:46Z
updated_at: 2026-09-01T18:59:17Z
parent: edifact-ls-3uzr
blocked_by:
    - edifact-ls-ogqj
---

# Description

Add a `edifact-ls check <file>` subcommand (alongside the existing
--version/--help flag handling in cmd/edifact-ls/main.go) that parses
the file, runs the same diagnostics/schema-validation path used by the
LSP, prints violations, and exits non-zero if any error-severity
diagnostic was produced -- for use in CI/scripts independent of the
editor. Spec compliance is naturally something people want to check in
a pipeline, not just see as squiggles while editing.

# Acceptance Criteria

[x] `edifact-ls check <file>` parses and validates a file without
starting the LSP server
[x] Exit code 0 when no error-severity diagnostics, 1 when at least one
[x] Violations printed in a human-readable form (position + message)
[x] Basic test coverage for both exit codes

## Summary of Changes

internal/edifact/edifact.go: added Validate(src) (*Interchange,
ErrorList), assembling the Parse -> ValidateEnvelopes -> Lint ->
ValidateMessageSchemas pipeline in one place. internal/lspserver/
diagnostics.go's diagnosticsForText now calls it instead of repeating
the four calls itself -- so the LSP and the new CLI command share one
canonical validation sequence and can't drift apart when a future
check is added to only one of them.

cmd/edifact-ls/main.go: new `check` subcommand alongside the existing
--version/--help handling. runCheck(w, path) reads the file, runs
edifact.Validate, prints each diagnostic to w via its existing
Error() formatting ("line:col: severity: message", one per line, no
new formatting logic needed), and returns 1 if any error-severity
diagnostic was found (or the file couldn't be read), 0 otherwise.
Factored as a plain function taking an io.Writer specifically so it's
testable without touching os.Exit or process I/O.

cmd/edifact-ls/main_test.go: 3 tests (clean file -> 0, a real IFTMCS
violation fixture -> 1, missing file -> 1). Also manually smoke-tested
the built binary directly against clean/violation/syntax-error/missing-
file/no-arg cases -- all behaved as designed, including a clean
`usage: edifact-ls check <file>` + exit 2 for the no-argument case.

Full suite (`make test`, now including cmd/edifact-ls) and e2e harness
(`make test-e2e`) still green.
