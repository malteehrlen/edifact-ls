---
# edifact-ls-lg15
title: 'Parser: token stream -> AST with structured errors'
status: completed
type: feature
priority: high
created_at: 2026-09-01T13:02:42Z
updated_at: 2026-09-01T13:20:02Z
parent: edifact-ls-btas
blocked_by:
    - edifact-ls-tw77
---

# Description
Build an AST (interchange -> messages -> segments -> elements/components)
from the token stream, with structured, positioned syntax errors for
malformed input and recovery so parsing can continue past an error.

# Acceptance Criteria
- [x] AST types for interchange/segment/element/component, each carrying
      source position (`ast.go`). Deliberately no `Message` type at this
      layer: message (UNH/UNT) grouping is envelope semantics, not syntax --
      it's built on top of the flat segment list in the envelope-validation
      story (edifact-ls-fsy0) instead, keeping this parser reusable even for
      input where UNH/UNT are absent or malformed.
- [x] Structured error type (message + position + severity) collected during
      parse rather than causing panics
- [x] Recovers past a malformed segment (e.g. resyncs at next segment
      terminator) so later valid content is still parsed
- [x] Unit tests: valid realistic interchange round-trips into an AST; a
      handful of malformed inputs produce the expected errors at expected
      positions without panicking

## Summary of Changes
`internal/edifact/parser.go`: builds the flat `Interchange{UNA, Segments,
Delimiters}` AST from the lexer's token stream. Invalid segment tags (not
exactly 3 uppercase letters) and a missing terminator at EOF are recorded as
structured, positioned errors; invalid tags resync at the next terminator so
later segments still parse. Found and fixed an off-by-one bug during testing
where the first separator after a tag produced a phantom leading empty
element (grammar is `tag (SEP element)*` -- the first separator introduces
the first element rather than closing an implicit one before it). Covered by
`parser_test.go`, including a no-panic sweep over adversarial inputs.
