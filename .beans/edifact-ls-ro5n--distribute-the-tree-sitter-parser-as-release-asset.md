---
# edifact-ls-ro5n
title: Distribute the tree-sitter parser as release assets
status: in-progress
type: feature
priority: normal
created_at: 2026-09-01T16:01:42Z
updated_at: 2026-09-01T16:03:07Z
parent: edifact-ls-u9x8
blocked_by:
    - edifact-ls-f31a
---

# Description
edifact.so is a compiled C shared library, not a portable Go binary, so it
can't be cross-compiled from one machine the way the LSP binary is --
needs building natively per architecture. Add a matrix job to
release.yml (linux/amd64 on ubuntu-latest, linux/arm64 on the hosted
ubuntu-24.04-arm runner -- no C cross-compiler or QEMU emulation needed)
that builds edifact.so + queries/highlights.scm and uploads them as
additional assets on the *same* GitHub Release the existing goreleaser job
creates for that tag -- the grammar is versioned in lockstep with the LSP
server, not on its own release cadence.

# Acceptance Criteria
- [x] New `treesitter` job in `.github/workflows/release.yml`, matrix over
      linux/amd64 + linux/arm64, `needs: goreleaser` (so the release already
      exists), builds `edifact.so` and packages it with
      `queries/highlights.scm` into
      `tree-sitter-edifact_<version>_linux_<arch>.tar.gz`
- [x] Uploads each archive plus a per-archive `.sha256` sidecar file (not
      one shared checksums file -- the two matrix jobs run in parallel and
      would race on a single shared file) to the same release via
      `gh release upload`
- [x] README: downloading the tree-sitter parser from a release becomes the
      documented default in the syntax-highlighting install step, with
      building it locally (`npm install` + `tree-sitter build`) kept as an
      alternative
- [ ] Verified end-to-end: a real tagged release has the tree-sitter
      archives attached alongside the LSP binaries, and a downloaded
      archive's `edifact.so` actually loads and highlights in nvim
