---
# edifact-ls-9117
title: Schema registry documentation + additional D.96A releases
status: in-progress
type: epic
priority: normal
created_at: 2026-09-02T08:43:30Z
updated_at: 2026-09-02T08:43:30Z
parent: edifact-ls-gdt6
---

# Goal

Two related asks from the user, both about the structural schema
registry (edifact-ls-3uzr onward):

1. Document exactly which message specifications (type/version/release/
   agency) are actually supported -- currently discoverable only by
   reading source or triggering the "no schema registered" info
   diagnostic by trial and error.
2. Register multiple releases of the same "popular" message type, not
   just whichever single release happened to get sourced first.

# Design: the registry is already the right mechanism for #2

`schemaRegistry` is keyed on the *full* MessageID tuple (type, version,
release, agency), not just type -- adding a second release of ORDERS is
already just "another file, another RegisterSchema call, different
Release string," the same plug-and-play mechanism verified back in
edifact-ls-ogqj. No registry/validator/diagnostics changes needed for
breadth; #2 is purely sourcing more pages through the existing pipeline.

# Design: generate documentation from the registry, don't hand-write it

Hand-maintained docs drift from the actual registry contents (we just
hit a case -- CUSCAR -- where even *we* didn't know a message's real
release until opening the file). The registry has all the data needed;
it just isn't queryable or exposed yet:

1. `RegisterSchema` gains a third parameter, `source string` (the
   canonical spec URL) -- carried as data instead of living only in a
   Go doc comment. Touches all 12 existing call sites (mechanical,
   one line each) plus the 4 new D.96A ones land using the new
   signature from day one.
2. `ListRegisteredSchemas() []SchemaInfo` -- the single source of truth
   for "what's supported," sorted for determinism.
3. A `edifact-ls schemas` CLI subcommand using it directly.
4. A generated `docs/SUPPORTED_MESSAGES.md`, produced by a small
   `tools/gendocs` program (`make docs`) from the same
   `ListRegisteredSchemas()` data, with a test
   (`TestSupportedMessagesDocIsUpToDate`) that fails if the checked-in
   file doesn't match what regenerating it would produce -- the same
   idea as a `gofmt -l` check, so docs and code structurally cannot
   disagree.
5. One line in the README's Features section pointing at the generated
   doc.

# Scope for #2 (this round)

D.96A for ORDERS, ORDRSP, INVOIC, DESADV -- picked by the user as a
widely-used real-world release (same era as their own CUSCAR file's
D:96B) for the four most commonly exchanged commercial documents. Not
exhaustive; more releases/types can follow the same pattern later.

# Acceptance Criteria

[ ] RegisterSchema carries a source URL; all existing + new call sites
updated
[ ] ListRegisteredSchemas() and a `edifact-ls schemas` CLI subcommand
exist and reflect the real registry
[ ] docs/SUPPORTED_MESSAGES.md is generated (via `make docs`) from the
registry, with a test enforcing it stays in sync
[ ] README links to the generated doc
[ ] ORDERS, ORDRSP, INVOIC, DESADV each also registered for D.96A
(real sourced data, tested, e2e-verified), alongside their existing
D.20A registrations
