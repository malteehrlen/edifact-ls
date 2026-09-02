---
# edifact-ls-6xaz
title: 'Hover: data-element/coded-value descriptions'
status: completed
type: epic
priority: deferred
created_at: 2026-09-01T17:53:47Z
updated_at: 2026-09-02T12:40:13Z
parent: edifact-ls-gdt6
blocked_by:
    - edifact-ls-tnp9
---

# Goal

Hover support, tier 2: hovering a specific element/component *value*
(e.g. `220` in `BGM+220`) shows its coded meaning (e.g. "Document/message
name, coded: 220 = Order"). This goes deeper than segment-tag hover
(edifact-ls-<epic1>) -- it needs each segment's composite/element
structure (which element is which, from the UN Data Element Directory)
plus code-list lookups (UNCL) for coded values.

# Strategy / status

Deferred. Blocked by edifact-ls-<epic1> (segment-tag hover) since it
reuses that story's `textDocument/hover` wiring -- this epic only adds a
deeper lookup, not a new LSP capability. The UN directory data this
needs (element structure + code lists) is the same flavor of sourcing
work as the content-validation epic (edifact-ls-9ger); when this is
picked up, check whether that work can be shared rather than sourced
twice.

No stories broken out yet -- scope (which segments/elements to cover
first, how much of UNCL to embed) is intentionally left open until
tier 1 ships and tier 2 is actually picked up.

# Scope decision (resolved)

Started from the element/composite structure edifact-ls-9ger already
sourced (BGM, DTM, CTA), rather than sourcing new segments from
scratch, per this epic's own note about sharing rather than
re-sourcing. Of those, picked exactly two real, moderately-sized code
lists to actually wire up as a proof of concept:

- BGM's "Message function code" component -> UN/EDIFACT code list
  1225 (69 real entries, e.g. "9" = Original), source:
  https://service.unece.org/trade/untdid/d21a/tred/tred1225.htm
- CTA's "Contact function code" component -> code list 3139 (103
  entries), source:
  https://service.unece.org/trade/untdid/d21a/tred/tred3139.htm

Both 403 directly via Cloudflare; sourced via the Wayback Machine, same
caveat as every other data source in this project.

Deliberately NOT attempted: BGM's "Document name code" (list 1001) --
fetched and inspected, found to have ~795 entries, an order of
magnitude bigger than either list actually wired up. That's a
materially bigger undertaking on its own and doesn't change this
epic's core design at all (the CodeList/CodedValue plumbing is
identical regardless of a list's size) -- deferred as a natural,
not-yet-requested follow-up, the same way edifact-ls-fh22's CONTRL gap
and D.21A/D.99B/D.01B releases beyond what's registered were deferred
in earlier epics. DTM's date/time function code qualifier (list 2005)
was also left unattempted for the same reason (large list, no design
value added by doing it now).

Fallback behavior (resolved): hovering a non-coded component, or a
coded component whose actual value isn't a recognized code, falls back
to the segment's ordinary tier-1 description rather than showing
nothing -- and, importantly, rather than asserting a fabricated
meaning for a value that was never actually looked up successfully.

# Acceptance Criteria

[x] Scope decided: which segments/elements get coded-value hover first
(see Scope decision above)
[x] Element/composite structure sourced for at least those segments
(reused from edifact-ls-9ger)
[x] Code list (UNCL) data sourced for at least the coded elements among
them (1225, 3139)
[x] Hovering a coded value shows its meaning; hovering a non-coded
value (or a coded one with an unrecognized value) falls back to the
element's tier-1 tag description

## Summary of Changes

internal/edifact/codelist.go (new): `CodedValue{Name, Description}` and
a registry (`RegisterCodeList`/`LookupCode`), keyed by UN Trade Data
Element Directory data-element number then by code -- deliberately
independent of `SegmentElementSchema`/`ComponentSchema`, the same way
`schema_registry.go` is independent of `segment_content.go` (different
concerns, different registries, matching this project's established
pattern of small, focused registries rather than one large coupled
one).

internal/edifact/segment_content.go: added `ComponentSchema.CodeList`
(a data-element number, or empty if not coded / not sourced yet).
Added `SegmentElementSchemaFor(tag)`, an exported lookup hover needs
from outside the package.

internal/edifact/segment_elements_data.go: marked BGM's "Message
function code" and CTA's "Contact function code" components with their
real `CodeList` ids.

internal/edifact/codelist_1225.go, codelist_3139.go (new, generated):
the two real code lists, extracted mechanically (not hand-transcribed)
from their Wayback-archived source pages by a small script -- caught
and fixed a real extraction bug in review before committing: an early
version didn't stop before the source page's own "Data Element Cross
Reference" footer section and picked up two footer tokens as a bogus
extra "code" entry.

internal/lspserver/hover.go: added `codedValueHoverAt`, checked first
(most specific) in `textDocumentHover`'s now-three-tier priority
(coded value > tag description > group context, falling through to
whichever combination is actually available). Locates the exact
component under the cursor via its real `Pos`/`Raw` span, looks up its
`ComponentSchema.CodeList` if the segment's structure is known, and
only returns coded-value content when the actual value present is a
real, recognized entry in that list -- anything short of that (no
structure, no code list, or an unrecognized value) falls through
rather than asserting a wrong or fabricated meaning.

Testing: 7 new unit tests in codelist_test.go (including exact
extracted-entry counts, 69 and 103, as a lightweight guard against
silent extraction corruption without re-fetching the source), 4 new
hover_test.go tests (a real coded lookup for each of BGM/CTA, plus two
fallback-correctness tests for an unrecognized code and a non-coded
component), and a new e2e check confirming "9" in a real BGM segment
resolves to "Original" in a live nvim hover. Full suite (`make test`)
and e2e harness (`make test-e2e`, 39 checks) pass.

## Retro

- Fetching the actual source pages before committing to a scope
  decision (rather than assuming code-list sizes) was the right
  sequencing: 1001 turning out to be ~800 entries vs. 1225's 69 and
  3139's 103 is exactly the kind of thing that's cheap to check first
  and expensive to discover mid-implementation. Matches this session's
  established discipline (e.g. edifact-ls-13gu's CDX survey before
  committing to "fetch everything") of checking real scale before
  scoping, not after.
- The extraction bug (footer tokens leaking in as a bogus code) was
  caught by eyeballing the generated output's tail before writing it
  into the committed data file, not by any test -- worth remembering
  this project's own established discipline (extract_msgtype.py's
  balance verification, this session's earlier wrapped-segment-name
  fix from edifact-ls-13gu) applies just as much to code-list
  extraction as message-schema extraction: mechanical extraction still
  needs its output inspected, not just trusted because it's mechanical.
- Keeping the fallback strict (only show coded-value content for an
  actually-recognized code, never for "this component is coded but I
  don't know this specific value") was a deliberate, and in hindsight
  necessary, design choice -- with only 2 of BGM's 4 coded components
  wired up (and none of DTM's), an unrecognized-value case is the
  common case for any segment/component this epic didn't specifically
  source, not an edge case. Explicit fallback tests for both "non-coded
  component" and "coded but unrecognized value" catch a regression
  here that a single happy-path test wouldn't.
- No adjustments needed before further work. The next natural,
  not-yet-requested step in this area would be sourcing 1001 (BGM's
  document name code) or 2005 (DTM's qualifier) -- both real, both
  larger, neither blocked by anything built here.
