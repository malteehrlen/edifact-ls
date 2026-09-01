-- Example standalone tree-sitter syntax highlighting config for edifact-ls,
-- meant to be copied into (or required from) YOUR OWN Neovim config.
-- Independent of any `nvim-treesitter` plugin -- uses Neovim's built-in
-- `vim.treesitter.language.add`/`query.set`, same as this repo's dev
-- harness (editors/nvim/init.lua), so it works regardless of your plugin
-- manager.
--
-- Build the parser first (from tree-sitter-edifact/ in a checkout of the
-- edifact-ls repo):
--
--   npm install
--   npx tree-sitter build -o edifact.so
--
-- Then copy edifact.so and queries/highlights.scm somewhere stable outside
-- your repo checkout -- adjust EDIFACT_TS_DIR below to wherever you put
-- them, e.g.:
--
--   mkdir -p ~/.local/share/edifact-ls
--   cp edifact.so queries/highlights.scm ~/.local/share/edifact-ls/

local EDIFACT_TS_DIR = vim.fn.expand("~/.local/share/edifact-ls")
local parser_path = EDIFACT_TS_DIR .. "/edifact.so"
local highlights_path = EDIFACT_TS_DIR .. "/highlights.scm"

if vim.fn.filereadable(parser_path) == 0 then
  vim.notify(
    "edifact tree-sitter parser not found at " .. parser_path ..
    " -- see editors/nvim/standalone_treesitter.lua for how to build and install it",
    vim.log.levels.WARN
  )
  return
end

vim.treesitter.language.add("edifact", { path = parser_path })

if vim.fn.filereadable(highlights_path) == 1 then
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
