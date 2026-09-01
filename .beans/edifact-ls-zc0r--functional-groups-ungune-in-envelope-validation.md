---
# edifact-ls-zc0r
title: Functional groups (UNG/UNE) in envelope validation
status: completed
type: feature
priority: normal
created_at: 2026-09-01T16:13:41Z
updated_at: 2026-09-01T16:18:12Z
parent: edifact-ls-0d7g
---

# Description
Add functional group support to `internal/edifact`, per section 8.3.5-8.3.7
of https://unece.org/DAM/trade/untdid/texts/d423.htm:

- Messages can optionally be wrapped in `UNG...UNE` groups between `UNB`
  and `UNZ`. It is invalid to mix grouped and ungrouped messages in the
  same interchange (explicitly stated in the guide).
- `UNG` and `UNE` pair up the same way `UNH`/`UNT` already do: `UNE`'s
  first element is a count of messages in that group, its second element
  must match `UNG`'s reference (its own analog of message reference).
- **Real bug to fix**: `UNZ`'s control count means "number of messages"
  *only when functional grouping isn't used* -- when it is, it means
  "number of functional groups" instead. `ValidateEnvelopes` currently
  always assumes messages.
- Also recognize `UNS` (section control segment, section 8.3.11): its
  single data element must be exactly `D` or `S`.

# Acceptance Criteria
- [x] Parser/grammar recognize `UNG`, `UNE`, `UNS` as ordinary segments
      (no special lexer changes needed, just envelope-level semantics)
- [x] `ValidateEnvelopes` groups messages under their enclosing `UNG`/`UNE`
      when present, validates each group's count/reference like `UNH`/`UNT`
      messages already are, and errors if grouped and ungrouped messages
      are mixed in one interchange
- [x] `UNZ`'s control count is checked against the number of functional
      groups when any `UNG` is present, or the number of messages
      otherwise
- [x] `UNS`'s single component must be `D` or `S`, else an error
- [x] Unit tests: a grouped interchange (multiple `UNG`/`UNE` pairs) passes
      cleanly; each new error case (mismatched group count/reference,
      mixed grouped+ungrouped, bad `UNS` value) is covered

## Summary of Changes
`internal/edifact/envelope.go`: `ValidateEnvelopes` now tracks `UNG..UNE`
functional groups alongside `UNH..UNT` messages, attributing each message
to its enclosing group (if any). New checks: UNE's message count and group
reference against UNG's, a duplicate/stray/unclosed UNG-UNE error family
matching the existing UNH/UNT one, an error when grouped and ungrouped
messages are mixed in the same interchange, and UNS's value must be
exactly "D" or "S". Fixed the real bug the story called out: UNZ's control
count now means "number of functional groups" when any UNG is present,
"number of messages" otherwise -- it previously always assumed messages.

12 new tests (valid grouped interchange, the UNZ group-vs-message-count
distinction, each new error case) all pass; full suite + e2e harness still
green.
