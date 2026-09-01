---
# edifact-ls-tw77
title: Lexer/tokenizer for segments & delimiters
status: completed
type: feature
priority: high
created_at: 2026-09-01T13:02:42Z
updated_at: 2026-09-01T13:19:44Z
parent: edifact-ls-btas
---

# Description
Tokenize raw EDIFACT bytes/text into segments, elements, components, and
delimiters, honoring an optional `UNA` service string advice segment that
redefines the default delimiter set.

# Acceptance Criteria
- [x] Parses `UNA` when present and applies its custom delimiters for the
      rest of the interchange; falls back to documented defaults when absent
- [x] Splits input into segments (terminated by segment terminator), each
      segment into elements (element separator), composite elements into
      components (component separator), honoring the release/escape character
      (this grouping lives in `parser.go`, built on the lexer's token stream)
- [x] Emits tokens with byte/line/column position information
- [x] Unit tests: default delimiters, custom `UNA` delimiters, escaped
      delimiter characters inside data, empty elements/components

## Summary of Changes
`internal/edifact/lexer.go`: byte-oriented lexer emitting Data/ComponentSep/
ElementSep/SegmentTerminator/EOF tokens with full position tracking, plus
`detectUNA` for the UNA service-string-advice delimiter override. Handles
release-character escaping (including a dangling-release-at-EOF lexical
error) and skips CRLF/LF between segments. Covered by
`internal/edifact/lexer_test.go`.
