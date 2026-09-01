---
# edifact-ls-fsy0
title: Envelope validation (UNB/UNZ, UNH/UNT counts)
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:02:42Z
updated_at: 2026-09-01T13:02:42Z
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
- [ ] Validates `UNB`/`UNZ` pairing and that `UNZ`'s count/control reference
      match the interchange
- [ ] Validates `UNH`/`UNT` pairing per message similarly
- [ ] Produces the same structured, positioned error type as the parser
- [ ] Unit tests: correct interchange passes; mismatched counts, missing
      `UNZ`/`UNT`, and mismatched control references each produce the
      expected error
