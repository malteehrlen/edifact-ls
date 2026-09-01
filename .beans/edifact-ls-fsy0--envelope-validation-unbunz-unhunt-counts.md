---
# edifact-ls-fsy0
title: Envelope validation (UNB/UNZ, UNH/UNT counts)
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:02:42Z
updated_at: 2026-09-01T13:21:11Z
parent: edifact-ls-btas
blocked_by:
    - edifact-ls-lg15
---

# Description
Semantic-level validation on top of the AST: interchange and message
envelopes must be well-formed and internally consistent (e.g. `UNZ` segment
count matches number of messages, `UNT` segment count matches number of
segments in that message, matching control reference numbers).

# Acceptance Criteria
- [x] Validates `UNB`/`UNZ` pairing and that `UNZ`'s count/control reference
      match the interchange
- [x] Validates `UNH`/`UNT` pairing per message similarly
- [x] Produces the same structured, positioned error type as the parser
- [x] Unit tests: correct interchange passes; mismatched counts, missing
      `UNZ`/`UNT`, and mismatched control references each produce the
      expected error

## Summary of Changes
`internal/edifact/envelope.go`: `ValidateEnvelopes(*Interchange) ErrorList`
scans the flat segment list tracking UNB/UNZ and each UNH..UNT span, then
checks: UNZ's control count equals the number of messages and its control
reference matches UNB's; each UNT's segment count equals the number of
segments from its UNH to itself inclusive, and its message reference matches
UNH's. Also flags duplicate UNB/UNZ, a UNT with no preceding UNH, and a UNH
never closed by a UNT. Covered by `envelope_test.go` (valid interchange,
each mismatch/missing case, and multiple messages in one interchange).
