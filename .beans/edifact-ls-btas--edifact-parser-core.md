---
# edifact-ls-btas
title: EDIFACT parser core
status: completed
type: epic
priority: high
created_at: 2026-09-01T13:02:31Z
updated_at: 2026-09-01T13:21:32Z
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
- [x] Correctly parses well-formed EDIFACT interchanges per the public
      UN/EDIFACT syntax rules (ISO 9735 structure): `UNA` service string
      advice (and its absence, implying default delimiters), `UNB`/`UNZ`
      interchange envelope, `UNH`/`UNT` message envelope, segments, composite
      / simple data elements, component/element/segment separators, release
      (escape) character
- [x] AST nodes carry source position (line/column/byte offset) for every
      segment/element, needed later for diagnostics ranges and formatting
- [x] Malformed input produces structured, positioned syntax errors instead
      of panicking, and the parser recovers enough to keep parsing where
      reasonable
- [x] Unit test suite covering: default vs. custom delimiters via `UNA`,
      nested composite elements, escaped/release characters, missing/
      mismatched envelope segments, at least one realistic sample interchange

## Summary of Changes
`internal/edifact`: a lexer (`lexer.go`, byte-oriented, UNA-aware,
release-character-escaping), a syntactic parser (`parser.go`) producing a
flat `Interchange{UNA, Segments, Delimiters}` AST with full position
tracking and structured recoverable errors, and envelope validation
(`envelope.go`) checking UNB/UNZ and UNH/UNT pairing, counts, and control
references on top of that AST. 24 unit tests across three stories, all
passing; `go vet` clean.

## Retro
- Went well overall, but the parser had a real off-by-one bug (a phantom
  leading empty element from misreading the segment grammar as "elements
  between separators" instead of "elements introduced by a separator") that
  only became obvious once tests used realistic multi-element data
  (`BGM+220+ORDER123+9`) instead of synthetic single-character values. Two
  of my own early unit tests had matching wrong expectations and passed
  anyway; a third only checked token *count* instead of token *kind* and
  masked a mismatched test fixture. Net takeaway: for a hand-written
  tokenizer/parser, realistic sample data catches bugs synthetic minimal
  cases (and loosely-asserting tests) don't.
- Deliberate scope call: the AST has no `Message` (UNH/UNT grouping) type —
  that's semantic, not syntactic, so it lives in envelope validation instead
  of the parser. Keeps `Parse` usable even on input with broken/missing
  UNH/UNT.
- This epic had no dependency on the scaffolding/e2e-harness epic and indeed
  ran independently — no integration surprises since neither epic's code
  touches the other yet (the harness's e2e check only exercises the LSP
  handshake, not parsing). That integration point arrives with diagnostics.
- No adjustments needed before continuing. Parser core and harness are both
  done; ready for the diagnostics and/or formatting epics whenever you'd
  like to green-light one.
