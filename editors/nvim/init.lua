-- Minimal, in-repo Neovim config for developing/testing edifact-ls against a
-- local build. Not meant for end users — see README.md for that. Launch it
-- with:
--
--   EDIFACT_LS_BIN=$(pwd)/dist/edifact-ls nvim -u editors/nvim/init.lua some.edi
--
-- It registers the server using Neovim's built-in vim.lsp.config/vim.lsp.enable
-- (stable since 0.11), which is what nvim-lspconfig itself now builds on. This
-- keeps the dev/test harness dependency-free; a real nvim-lspconfig-based
-- setup for end users is tracked separately (editor packaging & distribution).

vim.filetype.add({
  extension = {
    edi = "edifact",
    edifact = "edifact",
  },
})

local bin = os.getenv("EDIFACT_LS_BIN")
if not bin or bin == "" then
  error("EDIFACT_LS_BIN must point at a built edifact-ls binary")
end

vim.lsp.config("edifact_ls", {
  cmd = { bin },
  filetypes = { "edifact" },
  root_dir = function(_, on_dir)
    on_dir(vim.fn.getcwd())
  end,
})

vim.lsp.enable("edifact_ls")
