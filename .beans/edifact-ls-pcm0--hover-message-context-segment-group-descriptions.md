---
# edifact-ls-pcm0
title: 'Hover: message-context (segment-group) descriptions'
status: todo
type: epic
priority: deferred
created_at: 2026-09-01T17:53:55Z
updated_at: 2026-09-01T17:53:55Z
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

# Acceptance Criteria

[ ] Design for exposing "which schema group does this segment occurrence
belong to" from the structural validator to a hover handler
[ ] Hovering a segment inside a known group shows the group's
number/purpose in addition to (or instead of) the tier-1 description
[ ] At least IFTMCS validated end-to-end as a proof of concept
