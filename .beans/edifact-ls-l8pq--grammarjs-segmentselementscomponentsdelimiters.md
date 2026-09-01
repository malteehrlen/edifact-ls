---
# edifact-ls-l8pq
title: 'grammar.js: segments/elements/components/delimiters'
status: completed
type: feature
priority: normal
created_at: 2026-09-01T13:03:03Z
updated_at: 2026-09-01T14:08:47Z
parent: edifact-ls-nr8g
---

# Description
Write the core tree-sitter grammar for EDIFACT syntax: interchange envelope,
segments, data elements, composite/component elements, and delimiters
(including handling a custom `UNA`-defined delimiter set, if feasible within
tree-sitter's parsing model — otherwise document the limitation).

# Acceptance Criteria
- [x] `grammar.js` builds cleanly with `tree-sitter generate`
- [x] Parses the same sample interchanges used by the parser core's unit
      tests with zero `ERROR`/`MISSING` nodes
- [x] Tree-sitter test corpus under `test/corpus/` (tree-sitter's standard
      location -- the bean said `corpus/`, adjusted to match the tool's own
      convention) covering default delimiters, composite elements, and at
      least one realistic full interchange
- [x] `tree-sitter test` passes locally

## Summary of Changes
New `tree-sitter-edifact/` package (own `package.json` with `tree-sitter-cli`
as a devDependency, `tree-sitter.json` grammar config). `grammar.js`
implements segments/elements/components/delimiters for the *default*
delimiter set only -- a UNA-redefined custom delimiter set isn't honored
(tree-sitter grammars are static/compiled-ahead-of-time; adapting to
content-defined delimiters would need an external C scanner, out of scope
here and documented as a limitation in the grammar file itself, per what
this story's description explicitly allowed).

6 corpus tests under `test/corpus/edifact.txt`: minimal segment, composite
element, escaped/release characters, empty components/elements, UNA advice,
and a realistic full interchange (same shape as `testdata/minimal.edi`).
Verified zero ERROR/MISSING nodes against all of `testdata/*.edi` except
`syntax-error.edi`, which correctly still errors (it's deliberately
malformed).

Environment note: needed a C compiler to run `tree-sitter test` (parser.c is
compiled to a shared library) -- installed `build-essential` via apt, since
it wasn't present.

Found via testing (fixture mistake, same pattern as prior epics): my first
"empty components and elements" corpus case used `A::B` as a bare top-level
"segment" with no valid 3-letter tag before it, so it correctly produced an
ERROR node -- not a grammar bug, a malformed test fixture. Fixed by giving
it a real tag (`ABC+A::B'`).
