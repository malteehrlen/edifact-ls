---
# edifact-ls-btas
title: EDIFACT parser core
status: todo
type: epic
priority: high
created_at: 2026-09-01T13:02:31Z
updated_at: 2026-09-01T13:02:31Z
parent: edifact-ls-gdt6
---

# Goal
A Go library that lexes and parses UN/EDIFACT interchanges into a structured
AST with positional information, plus structured syntax errors. This is the
shared foundation diagnostics, formatting, and (independently) the tree-sitter
grammar all build on.

Not blocked by the e2e harness epic — this is pure parsing logic and can
proceed in parallel, validated with unit tests.

# Acceptance Criteria
- [ ] Correctly parses well-formed EDIFACT interchanges per the public
      UN/EDIFACT syntax rules (ISO 9735 structure): `UNA` service string
      advice (and its absence, implying default delimiters), `UNB`/`UNZ`
      interchange envelope, `UNH`/`UNT` message envelope, segments, composite
      / simple data elements, component/element/segment separators, release
      (escape) character
- [ ] AST nodes carry source position (line/column/byte offset) for every
      segment/element, needed later for diagnostics ranges and formatting
- [ ] Malformed input produces structured, positioned syntax errors instead
      of panicking, and the parser recovers enough to keep parsing where
      reasonable
- [ ] Unit test suite covering: default vs. custom delimiters via `UNA`,
      nested composite elements, escaped/release characters, missing/
      mismatched envelope segments, at least one realistic sample interchange
