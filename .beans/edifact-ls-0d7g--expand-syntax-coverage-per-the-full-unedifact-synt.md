---
# edifact-ls-0d7g
title: Expand syntax coverage per the full UN/EDIFACT syntax rules
status: todo
type: epic
priority: normal
created_at: 2026-09-01T16:13:25Z
updated_at: 2026-09-01T16:13:25Z
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
- [ ] Functional groups (UNG/UNE) are recognized by the parser/grammar and
      cross-checked by envelope validation (including fixing UNZ's count
      semantics to depend on whether grouping is used)
- [ ] Diagnostics support info and warning severities, not just error, with
      at least one real check at each new severity level
- [ ] The segment tag's optional control-number components (explicit
      repeat representation, e.g. `GDS:1'`) are parsed correctly instead
      of being misinterpreted as segment data
- [ ] Tree-sitter highlighting distinguishes the newly-recognized service
      segments (UNG, UNE, UNS) sensibly
