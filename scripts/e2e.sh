#!/usr/bin/env bash
# Builds edifact-ls (unless EDIFACT_LS_BIN is already set) and runs the
# headless-nvim e2e checks in scripts/e2e_check.lua against it.
#
# Usage: scripts/e2e.sh   (or: make test-e2e)
#
# Downloads a pinned Neovim into .tools/ on first run if `nvim` isn't already
# on PATH and NVIM_BIN isn't set, so this works with no manual setup in CI.
set -euo pipefail
cd "$(dirname "$0")/.."

: "${EDIFACT_LS_BIN:=$(pwd)/dist/edifact-ls}"
export EDIFACT_LS_BIN
if [ ! -x "$EDIFACT_LS_BIN" ]; then
  echo "EDIFACT_LS_BIN not found/executable: $EDIFACT_LS_BIN (build it first, e.g. 'make build')" >&2
  exit 1
fi

NVIM_VERSION_TAG="v0.12.5"
NVIM_DIR=".tools/nvim-linux-x86_64"

if [ -n "${NVIM_BIN:-}" ]; then
  : # explicit override, use as-is
elif command -v nvim >/dev/null 2>&1; then
  NVIM_BIN="$(command -v nvim)"
else
  if [ ! -x "$NVIM_DIR/bin/nvim" ]; then
    echo "Neovim not found; downloading pinned $NVIM_VERSION_TAG into $NVIM_DIR..." >&2
    mkdir -p .tools
    curl -fsSL -o .tools/nvim-linux-x86_64.tar.gz \
      "https://github.com/neovim/neovim/releases/download/${NVIM_VERSION_TAG}/nvim-linux-x86_64.tar.gz"
    tar -xzf .tools/nvim-linux-x86_64.tar.gz -C .tools
    rm .tools/nvim-linux-x86_64.tar.gz
  fi
  NVIM_BIN="$NVIM_DIR/bin/nvim"
fi

echo "Using nvim: $NVIM_BIN ($("$NVIM_BIN" --version | head -1))" >&2
echo "Using edifact-ls: $EDIFACT_LS_BIN" >&2

# noswapfile: checks intentionally leave buffers modified without saving
# (e.g. after formatting), and a leftover swapfile from a prior run --
# crashed, killed, or interrupted -- would otherwise block a later run on
# the same fixture behind an interactive "ATTENTION" prompt headless nvim
# can't answer.
"$NVIM_BIN" --headless -u editors/nvim/init.lua --cmd "set noswapfile" -c "luafile scripts/e2e_check.lua"
