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
--
-- Syntax highlighting (tree-sitter) is registered too, but only if
-- EDIFACT_TS_PARSER is set (path to the compiled tree-sitter-edifact/*.so --
-- see tree-sitter-edifact/README or scripts/e2e.sh for how to build it), so
-- plain LSP-only testing doesn't require the tree-sitter toolchain.

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

local ts_parser_path = os.getenv("EDIFACT_TS_PARSER")
if ts_parser_path and ts_parser_path ~= "" then
  vim.treesitter.language.add("edifact", { path = ts_parser_path })

  local highlights_path = os.getenv("EDIFACT_TS_HIGHLIGHTS")
  if highlights_path and highlights_path ~= "" then
    local f = assert(io.open(highlights_path, "r"))
    local query = f:read("*a")
    f:close()
    vim.treesitter.query.set("edifact", "highlights", query)
  end

  vim.api.nvim_create_autocmd("FileType", {
    pattern = "edifact",
    callback = function(args)
      vim.treesitter.start(args.buf, "edifact")
    end,
  })
end

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
