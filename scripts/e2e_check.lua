-- Headless e2e checks for edifact-ls, run via scripts/e2e.sh under
-- `nvim --headless -u editors/nvim/init.lua`. Add new checks as new features
-- land (highlighting, formatting); each should call fail() on the first
-- problem it finds so the harness exits non-zero with a clear reason, and
-- pass() may be called once per check on success.

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

-- Check: opening a fixture with a known problem surfaces a diagnostic whose
-- message contains the given substring.
local function check_diagnostic(fixture, expect_substring)
  local path = vim.fn.fnamemodify(fixture, ":p")
  vim.cmd.edit(path)

  local diags
  local ok = vim.wait(3000, function()
    diags = vim.diagnostic.get(0)
    return #diags > 0
  end, 50)

  if not ok then
    fail("no diagnostics appeared for " .. path .. " within timeout")
    return
  end

  for _, d in ipairs(diags) do
    if d.message:find(expect_substring, 1, true) then
      pass("diagnostics for " .. path .. " include a message containing " .. vim.inspect(expect_substring))
      return
    end
  end

  local messages = vim.tbl_map(function(d) return d.message end, diags)
  fail("diagnostics for " .. path .. " did not include a message containing " ..
    vim.inspect(expect_substring) .. "; got: " .. vim.inspect(messages))
end

check_lsp_attaches()
check_diagnostic("testdata/syntax-error.edi", "invalid segment tag")
check_diagnostic("testdata/envelope-error.edi", "missing UNZ")

if failed then
  vim.cmd("cquit 1")
else
  vim.cmd("qa")
end
