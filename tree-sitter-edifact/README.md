# tree-sitter-edifact

A [tree-sitter](https://tree-sitter.github.io/tree-sitter/) grammar for
UN/EDIFACT, used for syntax highlighting independent of the LSP server's own
parser (`internal/edifact` in the repo root).

**Scope:** default delimiters only (component `:`, element `+`, release
`?`, terminator `'`). A `UNA` service string advice segment is recognized
and highlighted as its own node, but a *custom* delimiter set it might
define is not honored — see the comment at the top of `grammar.js` for why.

## Build & test

Requires Node.js/npm (for `tree-sitter-cli`, a devDependency) and a C
compiler (to compile the generated `src/parser.c`).

```sh
npm install
npx tree-sitter generate   # regenerate src/ after editing grammar.js
npx tree-sitter test       # run test/corpus/*.txt
npx tree-sitter build -o edifact.so   # compile the loadable parser
```

## Using it in Neovim

The in-repo dev harness (`editors/nvim/init.lua`) picks this up
automatically when `EDIFACT_TS_PARSER` (and optionally
`EDIFACT_TS_HIGHLIGHTS`) are set — `scripts/e2e.sh` does this for you. To
try it manually:

```sh
npm install
npx tree-sitter build -o edifact.so
cd ..
EDIFACT_LS_BIN="$(pwd)/dist/edifact-ls" \
EDIFACT_TS_PARSER="$(pwd)/tree-sitter-edifact/edifact.so" \
EDIFACT_TS_HIGHLIGHTS="$(pwd)/tree-sitter-edifact/queries/highlights.scm" \
nvim -u editors/nvim/init.lua testdata/minimal.edi
```
