-- Headless e2e checks for edifact-ls, run via scripts/e2e.sh under
-- `nvim --headless -u editors/nvim/init.lua`. Add new checks as new features
-- land (highlighting, diagnostics, formatting); each should call fail() on
-- the first problem it finds so the harness exits non-zero with a clear
-- reason, and pass() should only be called once, at the very end.

local failed = false

local function fail(msg)
  io.stderr:write("FAIL: " .. msg .. "\n")
  failed = true
end

local function pass(msg)
  print("PASS: " .. msg)
end

-- Check: opening a valid EDIFACT fixture attaches the edifact_ls LSP client.
local function check_lsp_attaches()
  local fixture = vim.fn.fnamemodify("testdata/minimal.edi", ":p")
  vim.cmd.edit(fixture)

  local attached = vim.wait(3000, function()
    local clients = vim.lsp.get_clients({ name = "edifact_ls" })
    return #clients > 0 and clients[1].initialized == true
  end, 50)

  if not attached then
    fail("edifact_ls LSP client did not attach to " .. fixture .. " within timeout")
    return
  end

  pass("edifact_ls attached to " .. fixture)
end

check_lsp_attaches()

if failed then
  vim.cmd("cquit 1")
else
  vim.cmd("qa")
end
