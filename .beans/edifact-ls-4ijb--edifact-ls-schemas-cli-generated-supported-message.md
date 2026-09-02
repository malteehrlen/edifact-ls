---
# edifact-ls-4ijb
title: edifact-ls schemas CLI + generated SUPPORTED_MESSAGES.md
status: completed
type: feature
priority: normal
created_at: 2026-09-02T08:44:04Z
updated_at: 2026-09-02T08:49:47Z
parent: edifact-ls-9117
blocked_by:
    - edifact-ls-0m41
---

# Description

Surface the registry's contents to both humans (a generated doc) and
scripts (a CLI subcommand), both built on
`edifact.ListRegisteredSchemas()` from the previous story so there's
exactly one source of truth.

- `edifact-ls schemas` CLI subcommand (alongside `check`/`--version`):
  prints every registered `Type Version:Release:Agency Source` line.
- `RenderSupportedMessagesDoc() string` in `internal/edifact`: renders
  the same data as a Markdown table.
- `tools/gendocs/main.go`: a tiny program printing
  `RenderSupportedMessagesDoc()`; a `make docs` target runs it into
  `docs/SUPPORTED_MESSAGES.md`.
- `TestSupportedMessagesDocIsUpToDate`: compares the checked-in
  `docs/SUPPORTED_MESSAGES.md` against `RenderSupportedMessagesDoc()`,
  failing with a clear "run `make docs`" message if they differ --
  makes it structurally impossible for the doc to silently drift from
  the registry, the same idea as a `gofmt -l` check.
- One line in README's Features section linking to the generated doc.

# Acceptance Criteria

[x] `edifact-ls schemas` prints every registered message identity +
source
[x] `docs/SUPPORTED_MESSAGES.md` exists, generated via `make docs`,
and matches `RenderSupportedMessagesDoc()` exactly
[x] `TestSupportedMessagesDocIsUpToDate` fails clearly if the doc and
registry disagree (verify this directly: add a schema, confirm the
test fails until `make docs` is rerun)
[x] README links to the generated doc
[x] Full suite (`make test`) green

## Summary of Changes

internal/edifact/supported_messages_doc.go: RenderSupportedMessagesDoc()
renders ListRegisteredSchemas() as a Markdown table with a short
explanatory preamble (what "registered" means, what happens for a
recognized-type-wrong-release message).

internal/edifact/supported_messages_doc_test.go:
TestSupportedMessagesDocIsUpToDate compares the checked-in
docs/SUPPORTED_MESSAGES.md against that function's live output.
Verified directly, not just trusted: appended a stray blank line to the
checked-in doc, confirmed the test fails with a clear "run `make docs`"
message, then regenerated and confirmed it passes again.

tools/gendocs/main.go: a minimal program printing
RenderSupportedMessagesDoc(); Makefile gained a `docs` target piping it
into docs/SUPPORTED_MESSAGES.md. Generated the doc for the first time --
all 12 currently-registered schemas present, correctly sorted.

cmd/edifact-ls/main.go: `schemas` subcommand (runSchemas) prints each
registered `TYPE VERSION:RELEASE:AGENCY SOURCE` line, built on the same
ListRegisteredSchemas() data as the generated doc. Manually smoke-tested
the real built binary's output, not just the unit test.

README.md: Features section's message-specification bullet no longer
hardcodes a type count/list (which had already gone stale once --
CUSCAR wasn't in it) -- now links to docs/SUPPORTED_MESSAGES.md and
mentions `edifact-ls schemas` instead. CLI section documents the new
subcommand. Development section documents `make docs` and notes the
staleness test.

Full suite (`make test`) and e2e harness (`make test-e2e`) green.
