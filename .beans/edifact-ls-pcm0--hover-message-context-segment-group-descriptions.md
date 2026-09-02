---
# edifact-ls-pcm0
title: 'Hover: message-context (segment-group) descriptions'
status: completed
type: epic
priority: deferred
created_at: 2026-09-01T17:53:55Z
updated_at: 2026-09-02T12:33:20Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-tnp9
    - edifact-ls-3uzr
---

# Goal

Hover support, tier 3: hovering a segment *occurrence* shows its
message-specific context -- e.g. hovering an `NAD` that falls inside
IFTMCS's segment group 12 shows "Segment group 12: Party identification
and references", not just NAD's generic tier-1 description. This
requires knowing, for a given parsed message, which segment group a
specific segment occurrence belongs to.

# Strategy / status

Deferred. Blocked by both edifact-ls-<epic1> (segment-tag hover -- reuses
its `textDocument/hover` wiring) and edifact-ls-3uzr (structural
spec validation), since group membership for a given occurrence is
exactly what that epic's schema/validator already has to compute in
order to validate structure. This epic is really "surface something the
structural validator already knows" rather than new sourcing work.

No stories broken out yet -- design depends on the concrete shape
edifact-ls-3uzr's schema/validator end up taking.

# Scope decision (resolved)

Number only, not purpose text. UN/EDIFACT branching diagrams number
segment groups sequentially ("Segment group N") right in the table
itself -- already visible as comments in this project's generated
schema data (e.g. `iftmcs_d21a.go`'s `// Segment group 12`) -- but don't
reliably carry a separate descriptive *name* for what a group is for
("Party identification and references") anywhere in the same
machine-parseable table; that kind of description, where it exists at
all, lives in prose elsewhere in a spec page, not sourced by this
project's extraction pipeline. Sourcing purpose text for groups across
198 registered schemas would be a materially bigger, separate
undertaking (per-schema prose extraction, not mechanical table
parsing) -- deferred indefinitely, same call edifact-ls-9ger made for
code-list values vs. element presence.

A better-than-expected finding while implementing: the group *number*
needs no new sourcing at all. It's fully derivable from a Schema's
existing tree shape alone (SchemaNode already has everything needed;
no new field, no re-fetching any source page) -- see
`numberGroups`/`GroupPathAt` in schema.go. That means this feature
works identically for all 198 currently-registered schemas, not just
IFTMCS -- the PoC requirement below is about depth of *verification*,
not the feature's actual reach.

# Acceptance Criteria

[x] Design for exposing "which schema group does this segment occurrence
belong to" from the structural validator to a hover handler
[x] Hovering a segment inside a known group shows the group's
number (not purpose -- see scope decision) in addition to (or instead
of) the tier-1 description
[x] At least IFTMCS validated end-to-end as a proof of concept

## Summary of Changes

internal/edifact/schema.go: added `GroupPathAt(schema, segments,
targetIndex) []int`, returning the sequence of segment-group numbers
(outermost first) a given segment occurrence falls within -- the exact
numbering UN/EDIFACT branching diagrams themselves use, computed once
per call via a new `numberGroups` structural preorder pass (independent
of any message instance or repeat count). Rather than a second,
separately-maintained matching implementation that could drift from
`ValidateSchema`'s real one, `matchSequence`/`matchOnce` were refactored
to share one implementation via a new `matchContext` (carrying `errs`
plus an optional `visit` callback) and a precomputed `numbering` tree
threaded alongside the schema tree -- `ValidateSchema` behaves exactly
as before (all pre-existing tests pass unchanged), while `GroupPathAt`
rides the same matching decisions (including the same
overflow/ambiguity handling already hardened by edifact-ls-13gu) via
`visit`.

internal/edifact/schema_registry.go: exported `MessageIDOf` (wrapping
the existing unexported `messageIDOf`) and added `LookupSchema(id)
(Schema, bool)`, so a caller outside the package (the hover handler)
can identify a message and fetch its registered schema without new
plumbing.

internal/lspserver/hover.go: `textDocumentHover` now combines tier-1
tag description and tier-2 group context, returning nil only when
neither is available. New `groupPathAt` walks back from the hovered
segment to its enclosing UNH, finds the paired UNT, looks up the
message's registered schema, and calls `GroupPathAt` against the body
span; `formatGroupPath` renders it as "segment group N[, > segment
group M, ...]".

Testing: 5 new `GroupPathAt` unit tests in schema_test.go (including
`TestGroupPathAtNumberingStableAcrossRepeats`, a regression test for a
real bug this caught -- see Retro), a new generic test
(`TestAllRegisteredSchemasGroupPathAtStaysInRange`) smoke-testing every
registered schema's own minimal message, 4 new hover_test.go tests
using a real IFTMCS fixture (group context alone, tier-1 alone,
combined, and no-group-context cases), and a new e2e check
(`testdata/iftmcs-group-context.edi`) confirming "Part of segment group
2" appears in a live nvim hover. Full suite (`make test`) and e2e
harness (`make test-e2e`, 38 checks) pass.

## Retro

- Writing the dedicated nested/repeating-group unit tests *before*
  wiring anything into hover paid off immediately, the same way it did
  for the original schema engine (edifact-ls-bygc): a real bug surfaced
  before any real message data could hide it. `numberGroups`'s first
  version read `*counter` directly inside a struct literal alongside a
  recursive call in the same literal's `Children` field --
  `groupNumbering{Number: *counter, Children: numberGroups(...)}` --
  and Go does not guarantee a struct literal's field expressions
  evaluate left-to-right the way function-call arguments do. The
  recursive call (which itself advances the shared counter for any
  nested groups) could run before `*counter` was read for `Number`,
  silently assigning a nested group's *already-advanced* count to its
  parent instead of the parent's own number. `TestGroupPathAtNestedGroups`
  and `TestGroupPathAtNumberingStableAcrossRepeats` caught this
  immediately via plain assertion failures, not a crash -- worth
  remembering generally: this exact composite-literal-with-a-side-
  -effecting-call pattern is a real, non-obvious Go footgun, not
  hypothetical.
- Sharing one matching implementation (via `matchContext` + `visit`)
  instead of writing a second, hover-specific tree walker was the
  right call specifically because `matchSequence` had *already* had
  two real ambiguity bugs found and fixed this session (edifact-ls-bygc,
  edifact-ls-13gu) -- a hand-duplicated parallel walker would have had
  to independently re-discover and correctly replicate both of those
  fixes (the `insideGroup && i==0` exemption and
  `laterSiblingSharesLeadingTag`) to stay correct on real messages that
  exercise them, with no shared test coverage to catch drift.
- The scope turned out better than planned, not worse: group numbering
  needed zero new data sourcing and covers all 198 registered schemas
  for free, not just IFTMCS. Worth remembering when scoping future
  "surface something the validator already knows" epics -- the
  original scope note in this epic ("no sourcing work") turned out to
  be even more true than it looked at write time, once the concrete
  design was worked out.
- No adjustments needed before starting edifact-ls-6xaz (tier 2, the
  other hover epic queued alongside this one) -- it's a materially
  different kind of work (real UNCL code-list sourcing), not blocked
  or informed by anything built here beyond reusing the same
  `textDocument/hover` handler shape.
