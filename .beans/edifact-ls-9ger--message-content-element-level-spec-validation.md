---
# edifact-ls-9ger
title: Message content (element-level) spec validation
status: completed
type: epic
priority: normal
created_at: 2026-09-01T17:03:46Z
updated_at: 2026-09-01T20:01:28Z
parent: edifact-ls-gdt6
---

# Goal

Beyond structural validation (segment/group presence, order, cardinality
-- edifact-ls-3uzr) and hover (edifact-ls-tnp9), this epic validates
*content* within a segment: which of its data elements/components are
mandatory, and whether they're actually present in a given occurrence.

# Scope decision (resolved)

Element/component **presence** only, not code-list/value validation
(e.g. checking that a coded value like a transport-service-requirement
code is actually on its code list). That's a materially bigger,
separate undertaking (needs UNCL code list data on top of element
structure) and isn't attempted here -- if wanted later, it deserves its
own epic informed by whatever this one's implementation looks like.

# Design decision (resolved) -- the big simplification

Originally this epic assumed content rules would need to "attach to or
share data with" the structural per-message-type schema from
edifact-ls-3uzr (hence the original blocked-by). That assumption was
wrong, confirmed against the real source: fetching BGM's actual UNSD
segment definition
(https://service.unece.org/trade/untdid/d21a/trsd/trsdbgm.htm, via the
same Cloudflare/Wayback pattern as iftmcs_d21a.go) shows one single
element/component structure for BGM, followed by "Segment BGM is used
in the following Messages:" listing ~150 message types that all share
it. A segment's element/component structure is intrinsic to its tag,
not per-message-type.

That means this validation is a fully independent, generic pass --
architecturally parallel to Lint, not layered onto SchemaNode/Schema
from edifact-ls-3uzr. No blocked-by; it never actually depended on that
epic.

# Proof of concept (resolved)

Three segments, chosen because real fetched data showed a useful mix
and because descriptions for all three already exist (segments.go, from
the hover epic):
- **DTM**: has genuine mandatory structure -- its one element (C507) is
  itself mandatory, and within it, component 2005 (date/time/period
  function code qualifier) is mandatory too. The real "missing
  mandatory" case.
- **BGM** and **CTA**: every element/component in both is conditional
  (confirmed from their real UNSD pages) -- proves the validator
  doesn't false-positive on segments with no mandatory content at all,
  the same kind of finding edifact-ls-7uhx hit with IFTMCS's groups.

# Acceptance Criteria

[x] Scope decided: mandatory/conditional element presence only, vs. also
code-list/value validation
[x] Design for how per-segment content rules attach to (or are shared
with) the structural schema from the other message-spec epic
[x] At least one segment/message-type pair validated end-to-end as a
proof of concept

## Summary of Changes

All three stories completed: a generic element/component presence
engine (SegmentElementSchema/ElementSchema/ComponentSchema +
ValidateSegmentElements), a tag-keyed registry + ValidateSegmentContent
wired into edifact.Validate (reaching both LSP diagnostics and the CLI
`check` command automatically), and real UNSD data for BGM, DTM, and
CTA. Verified end-to-end: a DTM missing its mandatory function-code-
qualifier component shows the diagnostic in a live nvim + LSP session
and via the CLI.

## Retro

- The epic's own pre-work assumption -- that content rules would need
  to attach to or share data with the structural per-message-type
  schema -- turned out to be wrong, and checking it against the real
  source before writing any code avoided building the wrong
  architecture. BGM's real UNSD page settled it directly: one element
  structure, then a list of ~150 message types that all share it. That
  meant this whole epic could be a small, independent, message-type-
  agnostic pass (parallel to Lint), not an extension of
  Schema/SchemaNode -- and the implementation ended up genuinely
  simple as a result: three short stories, no surprises, no rework.
- Deliberately fetched real data for three segments instead of one,
  specifically to include a segment (BGM) known to be entirely
  conditional alongside one with genuine mandatory structure (DTM).
  That paid for itself as a real test of "does the validator
  false-positive on an all-conditional segment," the same shape of
  bug edifact-ls-7uhx hit with IFTMCS's groups -- catching that class
  of mistake by construction rather than by luck.
- Continues the pattern from every prior message-spec epic this
  session: source real UN/EDIFACT data (via Wayback when Cloudflare
  blocks the direct fetch) rather than reasoning from memory, and
  verify a specific factual claim (here: "is segment structure really
  message-independent?") before designing around it. That habit has
  caught a real mistake in every epic it's been applied to so far.
- No adjustments needed before starting whatever's next.
