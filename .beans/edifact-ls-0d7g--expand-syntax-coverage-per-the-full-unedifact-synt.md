---
# edifact-ls-0d7g
title: Expand syntax coverage per the full UN/EDIFACT syntax rules
status: completed
type: epic
priority: normal
created_at: 2026-09-01T16:13:25Z
updated_at: 2026-09-01T16:27:02Z
parent: edifact-ls-gdt6
---

# Goal
Our current parser/grammar cover a deliberately minimal subset of real
UN/EDIFACT syntax (UNB/UNZ + UNH/UNT only, no functional groups, no
severities beyond error). This epic closes the gap identified by reading
the UN/EDIFACT Syntax Implementation Guidelines (Chapter 2.3) in full:

https://unece.org/DAM/trade/untdid/texts/d423.htm

(Reference this URL in each story if something needs cross-checking later.
Note: Section 10 of that document, "EDIFACT Service & Control Messages" --
which would normally define the CONTRL message's standardized syntax-error
codes -- is a stub in this particular text ["being submitted initially as
a separate paper, which after approval, will be inserted"]. So this epic
does NOT source an error-code taxonomy from that URL; if we want to adopt
CONTRL's actual error codes later, that needs a different reference.)

# Explicit non-goal
Section 9 of the guidelines (segment construction/repetition/grouping
rules) is almost entirely about *per-message-type* structure -- which
segments a specific UNSM (e.g. ORDERS, INVOIC) allows, in what order, how
many times. Validating that requires a full UNSM message directory (per
message type, versioned) and is a fundamentally different, much larger
scope than generic interchange-syntax validation. Not attempted here.

# Acceptance Criteria
- [x] Functional groups (UNG/UNE) are recognized by the parser/grammar and
      cross-checked by envelope validation (including fixing UNZ's count
      semantics to depend on whether grouping is used)
- [x] Diagnostics support info and warning severities, not just error, with
      at least one real check at each new severity level
- [x] The segment tag's optional control-number components (explicit
      repeat representation, e.g. `GDS:1'`) are parsed correctly instead
      of being misinterpreted as segment data
- [x] Tree-sitter highlighting distinguishes the newly-recognized service
      segments (UNG, UNE, UNS) sensibly

## Summary of Changes
All four stories completed: functional groups (UNG/UNE) recognized and
cross-checked by envelope validation, including a real bug fix (UNZ's
count meant "messages" unconditionally; now correctly means "groups" when
grouping is used); UNS validated; info/warning diagnostic severities added
alongside error, each with a real check grounded in the spec; segment tags'
explicit-repeat-representation control numbers (e.g. `GDS:1'`) now parsed
and round-tripped correctly instead of being silently misinterpreted as
data; tree-sitter highlighting distinguishes service segments from user
data segments via a query-level `#match?` predicate.

## Retro
- Reading the full spec up front (previous conversation turn) paid off --
  every story had a concrete section/example to implement against and cite,
  rather than guessing at behavior.
- No real bugs found *during* this epic's implementation (unlike the
  parser-core and formatting epics) -- the design decisions from reading
  the spec carefully (e.g. realizing UNZ's count is conditional on
  grouping, or that tag control numbers need lexer-level rewind support)
  were validated by tests passing on the first or second try throughout.
  The one recurring mistake was the same as every previous epic: a
  hand-counted segment-count field in a test fixture being wrong (twice
  this time) -- not a design or implementation problem, just arithmetic.
- The `Lint` vs `ValidateEnvelopes` split (advisory/severity-graded checks
  vs structural correctness) feels like the right architectural line and
  should be where any future syntax-rule checks land, rather than growing
  either function past its current, coherent scope.
- Also confirmed empirically (not just assumed) that Neovim's highlighter
  resolves overlapping captures by declaration order in the query file --
  worth remembering if highlights.scm grows more overlapping patterns.
- No adjustments needed before continuing. This closes out everything
  currently in the backlog except the deliberately-deferred official
  mason-registry PR.
