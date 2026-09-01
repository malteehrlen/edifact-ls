-- Example standalone LSP config for edifact-ls, meant to be copied into (or
-- required from) YOUR OWN Neovim config -- unlike editors/nvim/init.lua
-- (this repo's dev/test harness, which points at an unstinstalled local
-- build via $EDIFACT_LS_BIN), this expects `edifact-ls` to already be
-- installed and discoverable on $PATH.
--
-- Install it first (from a checkout of the edifact-ls repo):
--
--   go install ./cmd/edifact-ls        (or: make install)
--
-- This uses your Go toolchain's normal $GOBIN (usually already on $PATH).
-- If you're building from an activated Hermit shell in this repo
-- specifically, note that Hermit points GOBIN at its own .hermit/go/bin/,
-- which is only on $PATH while that shell is active -- not useful for a
-- persistent global config. Either install with your own (non-Hermit) Go
-- toolchain, or point GOBIN somewhere already on your $PATH, e.g.:
--
--   GOBIN="$HOME/.local/bin" go install ./cmd/edifact-ls

vim.filetype.add({
  extension = {
    edi = "edifact",
    edifact = "edifact",
  },
})

local bin = vim.fn.exepath("edifact-ls")
if bin == "" then
  vim.notify(
    "edifact-ls not found on $PATH -- install it with 'go install ./cmd/edifact-ls' " ..
    "from a checkout of the edifact-ls repo (see editors/nvim/standalone_lsp.lua)",
    vim.log.levels.WARN
  )
  return
end

vim.lsp.config("edifact_ls", {
  cmd = { bin },
  filetypes = { "edifact" },
  root_dir = function(bufnr, on_dir)
    local fname = vim.api.nvim_buf_get_name(bufnr)
    on_dir(vim.fs.root(fname, ".git") or vim.fn.fnamemodify(fname, ":p:h"))
  end,
})

vim.lsp.enable("edifact_ls")

vim.api.nvim_create_user_command("EdifactMinify", function()
  local client = vim.lsp.get_clients({ name = "edifact_ls", bufnr = 0 })[1]
  if not client then
    vim.notify("edifact_ls is not attached to this buffer", vim.log.levels.WARN)
    return
  end
  client:exec_cmd({
    command = "edifact-ls.minify",
    arguments = { vim.uri_from_bufnr(0) },
  }, { bufnr = 0 })
end, { desc = "Collapse the current EDIFACT buffer to single-line wire format" })
