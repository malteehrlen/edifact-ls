---
# edifact-ls-tw77
title: Lexer/tokenizer for segments & delimiters
status: todo
type: feature
priority: high
created_at: 2026-09-01T13:02:42Z
updated_at: 2026-09-01T13:02:42Z
parent: edifact-ls-btas
---

# Description
Tokenize raw EDIFACT bytes/text into segments, elements, components, and
delimiters, honoring an optional `UNA` service string advice segment that
redefines the default delimiter set.

# Acceptance Criteria
- [ ] Parses `UNA` when present and applies its custom delimiters for the
      rest of the interchange; falls back to documented defaults when absent
- [ ] Splits input into segments (terminated by segment terminator), each
      segment into elements (element separator), composite elements into
      components (component separator), honoring the release/escape character
- [ ] Emits tokens with byte/line/column position information
- [ ] Unit tests: default delimiters, custom `UNA` delimiters, escaped
      delimiter characters inside data, empty elements/components
