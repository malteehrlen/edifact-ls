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

-- Check: vim.lsp.buf.format() on a single-line "wire" fixture produces the
-- expected multiline layout.
local function check_formatting()
  local fixture = vim.fn.fnamemodify("testdata/unformatted.edi", ":p")
  vim.cmd.edit(fixture)

  local attached = vim.wait(3000, function()
    local clients = vim.lsp.get_clients({ name = "edifact_ls", bufnr = 0 })
    return #clients > 0 and clients[1].initialized == true
  end, 50)
  if not attached then
    fail("edifact_ls LSP client did not attach to " .. fixture .. " within timeout")
    return
  end

  local before = table.concat(vim.api.nvim_buf_get_lines(0, 0, -1, false), "\n")
  vim.lsp.buf.format({ bufnr = 0 })

  local changed = vim.wait(3000, function()
    local current = table.concat(vim.api.nvim_buf_get_lines(0, 0, -1, false), "\n")
    return current ~= before
  end, 50)
  if not changed then
    fail("formatting " .. fixture .. " did not change the buffer within timeout")
    return
  end

  local lines = vim.api.nvim_buf_get_lines(0, 0, -1, false)
  local want = { "UNH+1'", "BGM+220'" }
  if #lines ~= #want then
    fail("formatted " .. fixture .. " has " .. #lines .. " lines, want " .. #want .. ": " .. vim.inspect(lines))
    return
  end
  for i, w in ipairs(want) do
    if lines[i] ~= w then
      fail("formatted " .. fixture .. " line " .. i .. " = " .. vim.inspect(lines[i]) .. ", want " .. vim.inspect(w))
      return
    end
  end

  pass("formatting " .. fixture .. " produced the expected multiline layout")
end

-- Check: the :EdifactMinify user command collapses a multiline fixture to
-- the expected single-line wire format via edifact-ls.minify.
local function check_minify()
  local fixture = vim.fn.fnamemodify("testdata/minimal.edi", ":p")
  vim.cmd.edit(fixture)

  local attached = vim.wait(3000, function()
    local clients = vim.lsp.get_clients({ name = "edifact_ls", bufnr = 0 })
    return #clients > 0 and clients[1].initialized == true
  end, 50)
  if not attached then
    fail("edifact_ls LSP client did not attach to " .. fixture .. " within timeout")
    return
  end

  local before = table.concat(vim.api.nvim_buf_get_lines(0, 0, -1, false), "\n")
  vim.cmd("EdifactMinify")

  local changed = vim.wait(3000, function()
    local current = table.concat(vim.api.nvim_buf_get_lines(0, 0, -1, false), "\n")
    return current ~= before
  end, 50)
  if not changed then
    fail("EdifactMinify did not change the buffer for " .. fixture .. " within timeout")
    return
  end

  local lines = vim.api.nvim_buf_get_lines(0, 0, -1, false)
  local want = "UNB+UNOA:1+SENDER:ZZ+RECEIVER:ZZ+201001:1200+1'" ..
    "UNH+1+ORDERS:D:96A:UN'BGM+220+ORDER123+9'DTM+137:20100101:102'UNT+4+1'UNZ+1+1'"
  if #lines ~= 1 or lines[1] ~= want then
    fail("EdifactMinify on " .. fixture .. " produced " .. vim.inspect(lines) .. ", want a single line " .. vim.inspect(want))
    return
  end

  pass("EdifactMinify on " .. fixture .. " produced the expected single-line wire format")
end

check_lsp_attaches()
check_diagnostic("testdata/syntax-error.edi", "invalid segment tag")
check_diagnostic("testdata/envelope-error.edi", "missing UNZ")
check_formatting()
check_minify()

if failed then
  vim.cmd("cquit 1")
else
  -- Force-quit: checks like check_formatting intentionally leave buffers
  -- modified (we never save), so a plain :qa would refuse with E37.
  vim.cmd("qa!")
end
