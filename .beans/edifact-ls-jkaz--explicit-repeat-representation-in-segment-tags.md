---
# edifact-ls-jkaz
title: Explicit repeat representation in segment tags
status: todo
type: feature
priority: low
created_at: 2026-09-01T16:14:05Z
updated_at: 2026-09-01T16:14:05Z
parent: edifact-ls-0d7g
---

# Description
Per section 9.5.1 of https://unece.org/DAM/trade/untdid/texts/d423.htm: "the standard specification for all
segments is that the segment tag comprises 10 component data elements. The
first is mandatory and contains the unique code to identify the segment.
The remainder are conditional and are available to carry control numbers
for use when required with repeating segments" -- e.g. `GDS:1+...'`,
`GDS:2+...'` for the 1st/2nd occurrence of a repeating GDS segment
("explicit representation"), as opposed to the far more common
"implicit representation" (`GDS+...'`, `GDS+...'`, no control numbers)
that current UNSMs and this project's parser both already assume.

Right now the lexer/parser have no concept of this: `GDS:1+...'` would be
misparsed, with "1" silently swallowed as if it were segment data rather
than part of the tag. Low priority -- the guide itself notes modern UNSMs
use implicit representation -- but it's part of the formal syntax and
worth being correct about rather than silently mangling real (if rare)
input.

# Acceptance Criteria
- [ ] Lexer/parser recognize up to 9 colon-separated numeric control-number
      components immediately following the base tag, before the first
      element separator, as part of the segment tag rather than data
- [ ] `Segment.Tag` stays just the base code (e.g. "GDS"); the control
      numbers are exposed separately (e.g. `Segment.TagControlNumbers
      []string`) rather than silently dropped or merged into Elements
- [ ] tree-sitter grammar updated to parse and highlight this form without
      producing ERROR nodes
- [ ] Unit tests: both implicit (`GDS+...'`) and explicit
      (`GDS:1+...'`, `GDS:2+...'`) forms parse correctly and don't
      regress existing behavior for plain 3-letter tags
