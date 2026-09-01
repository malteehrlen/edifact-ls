---
# edifact-ls-ogqj
title: Schema registry keyed by UNH's message identifier
status: completed
type: feature
priority: normal
created_at: 2026-09-01T17:03:46Z
updated_at: 2026-09-01T18:46:50Z
parent: edifact-ls-3uzr
blocked_by:
    - edifact-ls-bygc
---

# Description

Extract the message identifier tuple (type, version, release, agency)
from UNH's S009 composite (element 1: data elements 0065/0052/0054/0051)
and use it to look up a registered schema. Wire the lookup into the
existing diagnostics pipeline so validation runs automatically, same as
every other check -- no new nvim command. When the type is recognized
but no schema matches the exact version/release, emit an info diagnostic
naming what schema(s) are available instead of silently skipping or
guessing at compatibility.

# Acceptance Criteria

[x] Accessor added (or reused, alongside the existing Component0 helper)
to read the four S009 components from UNH
[x] In-memory schema registry mapping (type, version, release, agency)
-> schema, with a registration API a later story uses to add IFTMCS
[x] Diagnostics run schema validation automatically when UNH's tuple
matches a registered schema, surfacing violations with the same
severity levels used elsewhere (error/warning/info)
[x] Known type + unmatched version/release produces an info diagnostic
listing the registered alternative(s); unknown type produces no
diagnostic at all
[~] e2e check -- not added here after all: with zero message types
actually registered in production code yet, there's nothing real for
an e2e check to exercise. Moved to edifact-ls-7uhx, which registers
IFTMCS and can exercise it against a genuine violation.
[x] Plug-and-play guardrail: registering a second, entirely different
message type (e.g. a small hand-built ORDERS-shaped schema in a
test) requires touching only a new schema-data file + one
registration call -- no changes to the registry, validator, or
diagnostics wiring

## Summary of Changes

internal/edifact/ast.go: added Segment.ComponentN(elementIndex,
componentIndex, d) to reach any component of an element, not just the
first; Component0 now just calls ComponentN(i, 0, d).

internal/edifact/schema_registry.go (new): MessageID{Type, Version,
Release, Agency}; messageIDOf(unh, d) reads it from UNH's S009 (element
1, components 0-3). schemaRegistry is a package-level map[MessageID]Schema
with RegisterSchema(id, schema) to populate it -- the intended pattern
is a message-type-specific file calling RegisterSchema from its own
init(), so adding a type never touches this file.
ValidateMessageSchemas(ic) scans ic.Segments for UNH..UNT spans (a
small dedicated scan, not reusing envelope.go's messageSpan -- that
type also tracks UNB/UNG/UNE concerns this doesn't need), looks up each
message's MessageID, and either runs ValidateSchema against a match,
emits an info diagnostic naming registered alternatives for a
known-type/wrong-version mismatch (registeredVersionsOf, sorted for
determinism), or does nothing for a wholly unknown type. A message
missing its UNT is skipped -- ValidateEnvelopes already reports that,
and the body's extent would be ambiguous anyway.

internal/lspserver/diagnostics.go: diagnosticsForText now also calls
edifact.ValidateMessageSchemas(ic), appended after Lint. No new nvim
command -- ties into the existing didOpen/didChange diagnostics
pipeline like everything else.

10 new unit tests (schema_registry_test.go), including one that
registers two unrelated message types back to back and validates
messages against each without touching any shared code path --
verifying the plug-and-play guardrail directly. Full suite (`make
test`) and e2e harness (`make test-e2e`) still green; the existing
fixtures (all ORDERS, unregistered) produce no spurious diagnostics
from the new pipeline stage, confirming the "unknown type -> silent"
path in the real running server, even without a dedicated e2e
assertion for it.
