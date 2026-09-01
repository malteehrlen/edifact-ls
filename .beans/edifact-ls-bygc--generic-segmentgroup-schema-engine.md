---
# edifact-ls-bygc
title: Generic segment/group schema engine
status: completed
type: feature
priority: normal
created_at: 2026-09-01T17:03:46Z
updated_at: 2026-09-01T18:36:45Z
parent: edifact-ls-3uzr
---

# Description

Build the schema type and validator described in the epic, independent
of any specific message type. A schema node is either a segment slot
(tag, mandatory bool, max repeat) or a group (mandatory bool, max
repeat, ordered children which are themselves nodes). The validator
walks the actual parsed segment stream for one UNH..UNT message against
the schema, matching in document order, and produces diagnostics for: a
mandatory segment/group missing before the sequence moves past its
position, a group's max repeat count exceeded, and (as a warning, not
error, matching the existing Lint/ValidateEnvelopes severity split) an
unrecognized segment tag appearing where the schema doesn't expect one.

Not wired to any real message type yet -- this story is the engine only,
exercised by hand-built test schemas.

# Acceptance Criteria

[x] Schema tree type defined (segment node: tag/mandatory/maxRepeat;
group node: mandatory/maxRepeat/children)
[x] Validator walks a message's segment stream against a schema and
reports missing-mandatory, repeat-exceeded, and unexpected-tag
violations with accurate positions
[x] Nested groups (a group containing child groups) validate correctly,
matching the kind of multi-level nesting real messages like IFTMCS use
[x] Unit tests using small hand-built schemas (not IFTMCS) covering:
clean pass, missing mandatory segment, missing mandatory group,
exceeded repeat, wrong order

## Summary of Changes

internal/edifact/schema.go: Schema/SchemaNode types (a node is either a
segment slot or a group of child nodes, each with Mandatory/MaxRepeat).
ValidateSchema walks a message's body segments against a Schema via
matchSequence/matchOnce, a recursive greedy matcher: for each node in
order, consume up to MaxRepeat matching occurrences (recursing into
matchSequence for a group's children), report "missing mandatory" if a
mandatory node matched zero times, report "exceeded max repeat" if more
matches follow after the cap, and warn on any segments left over once
the whole schema is exhausted.

Found and fixed a real correctness bug via the nested-group test: a
group's *first* child shares its leading tag with the group itself (by
construction -- see leadingTag), so naively overflow-checking that
child locally misfires whenever an outer, still-repeatable occurrence
of the group legitimately starts right after the inner child's own cap
is reached (a real shape IFTMCS uses, e.g. a repeating group whose
first child is itself a repeating sub-group). Fixed by exempting a
group's first child from local overflow detection (insideGroup
parameter in matchSequence) -- the ambiguity resolves one level up,
where the enclosing call correctly attributes the overflow to the
*group* repeating too many times instead of misattributing it to the
leaf. Covered by TestValidateSchemaNestedGroupExceedsOwnRepeat, which
asserts the diagnostic lands on the group, not the child.

8 new unit tests in schema_test.go, all passing; full suite (`make
test`) and e2e harness (`make test-e2e`) still green -- this story adds
no LSP-facing behavior, so e2e is an unaffected sanity check, not new
coverage.
