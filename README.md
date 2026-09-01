# edifact-ls

A language server for [UN/EDIFACT](https://en.wikipedia.org/wiki/EDIFACT)
files: formatting, syntax highlighting, and linting/validation, designed to
work with Neovim.

Project status and planned work are tracked as "beans" (via the `beans` CLI)
in `.beans/` — run `beans list --json` or `beans roadmap` to see the current
backlog.

## Requirements

- [Hermit](https://cashapp.github.io/hermit/) manages this repo's Go
  toolchain. Activate it once per shell:

  ```sh
  source bin/activate-hermit
  ```

  (or just prefix commands with `bin/go`, `bin/gofmt`, etc.)

- The `beans` CLI for backlog management (installed as a Hermit-managed Go
  binary) — see `AGENTS.md` for how this repo uses it (run `beans prime` for
  the full guide).

- Node.js/npm and a C compiler are needed to build the tree-sitter grammar
  (`tree-sitter-edifact/`) — only required for syntax highlighting; the LSP
  server, formatting, and diagnostics don't need them.

## Build & test

```sh
make build     # -> dist/edifact-ls
make test      # go vet + go test ./...
make test-e2e  # build + run the headless-nvim e2e harness (see below)
```

## Trying it in Neovim

`editors/nvim/` holds a minimal, in-repo Neovim config for developing against
a local build — not an end-user installation method (that's tracked
separately as the "editor packaging & distribution" epic). It registers
edifact-ls via Neovim's built-in `vim.lsp.config`/`vim.lsp.enable` (stable
since 0.11), for `*.edi`/`*.edifact` files.

```sh
make build
EDIFACT_LS_BIN="$(pwd)/dist/edifact-ls" nvim -u editors/nvim/init.lua testdata/minimal.edi
```

### Commands

- `:EdifactMinify` — collapses the current buffer to single-line "wire"
  EDIFACT (segments joined by their terminator only, no newlines). The
  reverse direction is just `vim.lsp.buf.format()` (`textDocument/formatting`),
  since formatting already produces the human-readable multiline form.

### Syntax highlighting

Highlighting is a separate [tree-sitter](https://tree-sitter.github.io/tree-sitter/)
grammar in `tree-sitter-edifact/`, independent of the LSP server's own
parser — see `tree-sitter-edifact/README.md` for how to build it and enable
it in the dev harness (via `EDIFACT_TS_PARSER`/`EDIFACT_TS_HIGHLIGHTS`).
`scripts/e2e.sh` builds and wires it in automatically.

## e2e test harness

`scripts/e2e.sh` builds nothing itself (run `make build` first, or use
`make test-e2e` which does both) but does:

1. Locate or download a pinned Neovim into `.tools/` (no manual setup
   needed — safe to run in CI).
2. Build the tree-sitter grammar (`tree-sitter-edifact/`), installing its
   npm dependency on first run.
3. Launch nvim `--headless` with `editors/nvim/init.lua`, pointed at
   `$EDIFACT_LS_BIN` and the built tree-sitter parser.
4. Run `scripts/e2e_check.lua`, which drives real editor behavior (opening
   files, waiting for the LSP client, etc.) and asserts on it.

Exit code is non-zero if any check fails, with a `FAIL: ...` reason on
stderr.

**Adding a new e2e check:** add a fixture under `testdata/` if needed, then
add a `check_*` function to `scripts/e2e_check.lua` following the existing
`check_lsp_attaches` example — call `fail("reason")` on the first problem
found, or `pass("...")` on success. The script calls `cquit 1` if anything
failed, `qa!` (exit 0) otherwise — the `!` is needed because checks may
leave buffers modified without saving (e.g. after formatting/minifying).
