---
# edifact-ls-lg15
title: 'Parser: token stream -> AST with structured errors'
status: todo
type: feature
priority: high
created_at: 2026-09-01T13:02:42Z
updated_at: 2026-09-01T13:02:42Z
parent: edifact-ls-btas
blocked_by:
    - edifact-ls-tw77
---

# Description
Build an AST (interchange -> messages -> segments -> elements/components)
from the token stream, with structured, positioned syntax errors for
malformed input and recovery so parsing can continue past an error.

# Acceptance Criteria
- [ ] AST types for interchange/message/segment/element/component, each
      carrying source position
- [ ] Structured error type (message + position + severity) collected during
      parse rather than causing panics
- [ ] Recovers past a malformed segment (e.g. resyncs at next segment
      terminator) so later valid content is still parsed
- [ ] Unit tests: valid realistic interchange round-trips into an AST; a
      handful of malformed inputs produce the expected errors at expected
      positions without panicking
