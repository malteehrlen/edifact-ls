# edifact-ls

A language server for [UN/EDIFACT](https://en.wikipedia.org/wiki/EDIFACT)
files, built for Neovim — real message-specification validation, not just
syntax checking.

Disclamer: This project is 100% generated code.

## Features

**Diagnostics**

- Syntax errors — malformed segments, tags, and delimiters.
- Envelope structure — UNB/UNZ, UNG/UNE, and UNH/UNT pairing, counts, and
  references.
- Linting — reserved `UN`-prefixed segment tags, redundant `UNA` service
  string advice (version-aware: distinguishes ISO 9735 versions 1–3 vs. 4
  conventions).
- Message-specification validation — segment/group presence, order, and
  repeat counts checked against the real UN/EDIFACT branching diagram for
  11 message types (IFTMCS, ORDERS, ORDRSP, INVOIC, DESADV, IFTMIN,
  IFTSTA, PRICAT, INVRPT, DELFOR, APERAK), not a hand-rolled
  approximation — each transcribed and verified against UNECE's own
  segment tables.
- Content validation — mandatory data element/component presence checked
  within individual segments (BGM, DTM, CTA), independent of message
  type.

**Editing**

- Hover — segment tag name and description, sourced from the UN/EDIFACT
  Segment Directory.
- Formatting — single-line "wire" format ⇄ human-readable multiline, via
  `textDocument/formatting` and the `:EdifactMinify` command.
- Syntax highlighting via a [tree-sitter](https://tree-sitter.github.io/tree-sitter/)
  grammar.

**Tooling**

- `edifact-ls check <file>` — a CLI subcommand running the same
  validation pipeline as the editor, for CI and scripts. Exits non-zero
  on any error-severity diagnostic.

A single static binary, no runtime dependencies, Linux amd64/arm64.

## Installation

**1. Install the binary onto `$PATH`.** Linux only for now (amd64/arm64) —
download the latest release from
[GitHub Releases](https://github.com/malteehrlen/edifact-ls/releases):

```sh
os=linux
arch=$(uname -m); case "$arch" in x86_64) arch=amd64 ;; aarch64) arch=arm64 ;; esac
version=$(curl -fsSL https://api.github.com/repos/malteehrlen/edifact-ls/releases/latest | grep -m1 '"tag_name"' | cut -d'"' -f4)
curl -fsSL -o edifact-ls.tar.gz \
  "https://github.com/malteehrlen/edifact-ls/releases/download/${version}/edifact-ls_${version#v}_${os}_${arch}.tar.gz"
tar -xzf edifact-ls.tar.gz
mv edifact-ls ~/.local/bin/   # or wherever's already on your $PATH
```

(Optionally verify against `checksums.txt`, also attached to the release.)

**Building from source** is also available, e.g. for a platform without a
release binary, or to build from a specific commit:

```sh
go install ./cmd/edifact-ls    # from a checkout of this repo (or: make install)
```

This uses your Go toolchain's `$GOBIN`, usually already on `$PATH`. From an
activated Hermit shell in this repo, note that Hermit points `$GOBIN` at its
own `.hermit/go/bin/`, which isn't useful for a persistent global config —
either use your own Go toolchain, or set `GOBIN` explicitly, e.g.
`GOBIN="$HOME/.local/bin" go install ./cmd/edifact-ls`.

**2. Copy `editors/nvim/standalone_lsp.lua` into your config** (e.g.
`require`d from your `init.lua`, or its contents pasted in directly). It
registers filetype detection, the LSP server (via built-in
`vim.lsp.config`/`vim.lsp.enable` — no `nvim-lspconfig` dependency required,
though it'll work fine alongside it), and the `:EdifactMinify` command.

**3. For syntax highlighting**, also install the tree-sitter parser, then
copy `editors/nvim/standalone_treesitter.lua` in the same way. Releases
include prebuilt `linux/amd64` and `linux/arm64` archives alongside the LSP
binary:

```sh
os=linux
arch=$(uname -m); case "$arch" in x86_64) arch=amd64 ;; aarch64) arch=arm64 ;; esac
version=$(curl -fsSL https://api.github.com/repos/malteehrlen/edifact-ls/releases/latest | grep -m1 '"tag_name"' | cut -d'"' -f4)
mkdir -p ~/.local/share/edifact-ls
curl -fsSL "https://github.com/malteehrlen/edifact-ls/releases/download/${version}/tree-sitter-edifact_${version#v}_${os}_${arch}.tar.gz" \
  | tar -xz -C ~/.local/share/edifact-ls
```

(Optionally verify against the matching `.tar.gz.sha256` file, also
attached to the release.)

**Building it locally** is also available (requires Node.js/npm and a C
compiler):

```sh
cd tree-sitter-edifact
npm install
npx tree-sitter build -o edifact.so
mkdir -p ~/.local/share/edifact-ls
cp edifact.so queries/highlights.scm ~/.local/share/edifact-ls/
```

(Adjust the target directory and `EDIFACT_TS_DIR` at the top of
`standalone_treesitter.lua` together if you'd rather put it somewhere else.)

Both files are self-contained and don't depend on the rest of this repo at
runtime.

### Commands

- `:EdifactMinify` — collapses the current buffer to single-line "wire"
  EDIFACT (segments joined by their terminator only, no newlines). The
  reverse direction is just `vim.lsp.buf.format()` (`textDocument/formatting`),
  since formatting already produces the human-readable multiline form.

### CLI

`edifact-ls check <file>` parses and validates a file — the same pipeline
the editor uses — without starting the language server. Prints each
diagnostic as `line:col: severity: message` and exits `1` if any is
error-severity, `0` otherwise. Useful in CI or a pre-commit hook.

## Development

Project status and planned work are tracked as "beans" (via the `beans`
CLI) in `.beans/` — run `beans list --json` or `beans roadmap` to see the
current backlog.

### Requirements

- [Hermit](https://cashapp.github.io/hermit/) manages this repo's toolchain
  (Go, Node.js/npm, `make`, `goreleaser`, `actionlint`, the `beans` CLI).
  Activate it once per shell:

  ```sh
  source bin/activate-hermit
  ```

  (or just prefix commands with `bin/go`, `bin/make`, etc.)

- A C compiler, to build the tree-sitter grammar (`tree-sitter-edifact/`) —
  only required for syntax highlighting; the LSP server, formatting, and
  diagnostics don't need it. Hermit doesn't manage a C toolchain, so this
  has to come from your system.

### Build & test

```sh
make build     # -> dist/edifact-ls
make test      # go vet + go test ./...
make test-e2e  # build + run the headless-nvim e2e harness (see below)
```

### Releasing

Pushing a `vX.Y.Z` tag triggers `.github/workflows/release.yml`, which:

1. Runs [goreleaser](https://goreleaser.com/) (config: `.goreleaser.yaml`)
   to cross-compile Linux amd64+arm64 LSP binaries, with the version baked
   in via `-ldflags`, and publish them as a GitHub Release with a
   checksums file.
2. Builds the tree-sitter parser natively on a linux/amd64 and a
   linux/arm64 runner (it's a compiled C shared library, so unlike the Go
   binary it can't be cross-compiled from one machine) and uploads those
   archives to the same release.

```sh
git tag vX.Y.Z
git push origin vX.Y.Z
```

To try the pipeline locally without publishing anything (useful when
editing `.goreleaser.yaml`):

```sh
goreleaser release --snapshot --clean --skip=publish
```

To verify a downloaded release binary:

```sh
tar -xzf edifact-ls_<version>_linux_<arch>.tar.gz
./edifact-ls --version
```

### Trying it in Neovim

`editors/nvim/` holds a minimal, in-repo Neovim config for developing
against a local build — separate from the installation method above (this
one points at an uninstalled local build via an env var instead of
`$PATH`). It registers edifact-ls via Neovim's built-in
`vim.lsp.config`/`vim.lsp.enable` (stable since 0.11), for
`*.edi`/`*.edifact` files.

```sh
make build
EDIFACT_LS_BIN="$(pwd)/dist/edifact-ls" nvim -u editors/nvim/init.lua testdata/minimal.edi
```

Syntax highlighting in this harness is the same
[tree-sitter](https://tree-sitter.github.io/tree-sitter/) grammar in
`tree-sitter-edifact/` described above — see `tree-sitter-edifact/README.md`
for how to build it and enable it here (via
`EDIFACT_TS_PARSER`/`EDIFACT_TS_HIGHLIGHTS`). `scripts/e2e.sh` builds and
wires it in automatically.

### e2e test harness

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
