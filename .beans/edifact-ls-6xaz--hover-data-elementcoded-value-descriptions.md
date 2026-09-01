---
# edifact-ls-6xaz
title: 'Hover: data-element/coded-value descriptions'
status: todo
type: epic
priority: deferred
created_at: 2026-09-01T17:53:47Z
updated_at: 2026-09-01T17:53:47Z
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

# Acceptance Criteria

[ ] Scope decided: which segments/elements get coded-value hover first
(likely driven by whatever's already covered by tier 1 and/or IFTMCS)
[ ] Element/composite structure sourced for at least those segments
[ ] Code list (UNCL) data sourced for at least the coded elements among them
[ ] Hovering a coded value shows its meaning; hovering a non-coded value
falls back to the element's name/description (or no hover, if not
worth the complexity -- decide when picked up)
