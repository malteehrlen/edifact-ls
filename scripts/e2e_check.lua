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
  vim.cmd.edit({ fixture, bang = true })

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
-- message contains the given substring. If expect_severity is given (one of
-- vim.diagnostic.severity.*), the matching diagnostic must also have that
-- exact severity.
local function check_diagnostic(fixture, expect_substring, expect_severity)
  local path = vim.fn.fnamemodify(fixture, ":p")
  vim.cmd.edit({ path, bang = true })

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
      if expect_severity and d.severity ~= expect_severity then
        fail("diagnostic for " .. path .. " matching " .. vim.inspect(expect_substring) ..
          " has severity " .. vim.diagnostic.severity[d.severity] ..
          ", want " .. vim.diagnostic.severity[expect_severity])
        return
      end
      pass("diagnostics for " .. path .. " include a message containing " .. vim.inspect(expect_substring) ..
        (expect_severity and (" with severity " .. vim.diagnostic.severity[expect_severity]) or ""))
      return
    end
  end

  local messages = vim.tbl_map(function(d) return d.message end, diags)
  fail("diagnostics for " .. path .. " did not include a message containing " ..
    vim.inspect(expect_substring) .. "; got: " .. vim.inspect(messages))
end

-- Check: opening a fixture that structurally validates cleanly produces no
-- error-severity diagnostic. Unlike check_diagnostic, absence can't be
-- detected by waiting for a diagnostic to appear -- instead waits for the
-- LSP client to attach (so we know the server has had a chance to publish
-- something), then asserts none of whatever it did publish is an error.
local function check_no_error_diagnostic(fixture)
  local path = vim.fn.fnamemodify(fixture, ":p")
  vim.cmd.edit({ path, bang = true })

  local attached = vim.wait(3000, function()
    local clients = vim.lsp.get_clients({ name = "edifact_ls", bufnr = 0 })
    return #clients > 0 and clients[1].initialized == true
  end, 50)
  if not attached then
    fail("edifact_ls LSP client did not attach to " .. path .. " within timeout")
    return
  end

  -- Give the server a moment to finish publishing after attach.
  vim.wait(500)

  local errors = vim.tbl_filter(function(d) return d.severity == vim.diagnostic.severity.ERROR end, vim.diagnostic.get(0))
  if #errors > 0 then
    local messages = vim.tbl_map(function(d) return d.message end, errors)
    fail("expected no error diagnostics for " .. path .. ", got: " .. vim.inspect(messages))
    return
  end

  pass("no error diagnostics for " .. path)
end

-- Check: vim.lsp.buf.format() on a single-line "wire" fixture produces the
-- expected multiline layout.
local function check_formatting()
  local fixture = vim.fn.fnamemodify("testdata/unformatted.edi", ":p")
  vim.cmd.edit({ fixture, bang = true })

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
  vim.cmd.edit({ fixture, bang = true })

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
    "UNH+1+TESTMSG:D:96A:UN'BGM+220+ORDER123+9'DTM+137:20100101:102'UNT+4+1'UNZ+1+1'"
  if #lines ~= 1 or lines[1] ~= want then
    fail("EdifactMinify on " .. fixture .. " produced " .. vim.inspect(lines) .. ", want a single line " .. vim.inspect(want))
    return
  end

  pass("EdifactMinify on " .. fixture .. " produced the expected single-line wire format")
end

-- Check: the tree-sitter parser is active for a valid fixture and produces
-- no ERROR/MISSING nodes. Only runs if EDIFACT_TS_PARSER is set (see
-- editors/nvim/init.lua); skipped otherwise rather than failing, so plain
-- LSP-only runs aren't forced to build the tree-sitter toolchain.
local function check_treesitter(fixture_path)
  if not os.getenv("EDIFACT_TS_PARSER") or os.getenv("EDIFACT_TS_PARSER") == "" then
    print("SKIP: tree-sitter check (EDIFACT_TS_PARSER not set)")
    return
  end

  local fixture = vim.fn.fnamemodify(fixture_path or "testdata/minimal.edi", ":p")
  vim.cmd.edit({ fixture, bang = true })

  local ok, parser_or_err = pcall(vim.treesitter.get_parser, 0, "edifact")
  if not ok then
    fail("no tree-sitter parser active for " .. fixture .. ": " .. tostring(parser_or_err))
    return
  end

  local root = parser_or_err:parse()[1]:root()
  if root:has_error() then
    fail("tree-sitter parse of " .. fixture .. " has ERROR/MISSING nodes:\n" .. root:sexpr())
    return
  end

  pass("tree-sitter parser active on " .. fixture .. " with no ERROR nodes")
end

-- Check: textDocument/hover at a given 0-based line/character returns
-- markdown content containing expect_substring.
local function check_hover(fixture, line, character, expect_substring)
  local path = vim.fn.fnamemodify(fixture, ":p")
  vim.cmd.edit({ path, bang = true })

  local attached = vim.wait(3000, function()
    local clients = vim.lsp.get_clients({ name = "edifact_ls", bufnr = 0 })
    return #clients > 0 and clients[1].initialized == true
  end, 50)
  if not attached then
    fail("edifact_ls LSP client did not attach to " .. path .. " within timeout")
    return
  end

  local params = {
    textDocument = vim.lsp.util.make_text_document_params(0),
    position = { line = line, character = character },
  }
  local results = vim.lsp.buf_request_sync(0, "textDocument/hover", params, 3000)

  local found = false
  for _, res in pairs(results or {}) do
    local hover = res.result
    if hover and hover.contents and hover.contents.value
        and hover.contents.value:find(expect_substring, 1, true) then
      found = true
    end
  end

  if not found then
    fail("hover at " .. path .. ":" .. line .. ":" .. character ..
      " did not include a message containing " .. vim.inspect(expect_substring) ..
      "; got: " .. vim.inspect(results))
    return
  end

  pass("hover at " .. path .. ":" .. line .. ":" .. character ..
    " includes a message containing " .. vim.inspect(expect_substring))
end

-- Check: textDocument/codeAction at a given 0-based line/character returns
-- a quickfix action whose title contains expect_title_substring; applying
-- that action's edit (exactly as a real client would) is expected to
-- resolve the problem, verified by re-checking diagnostics on the buffer
-- afterward. want_lines, if given, additionally asserts the exact
-- resulting buffer content -- needed to catch a fix leaving behind
-- unwanted artifacts (e.g. a blank line) that no diagnostic would flag.
local function check_code_action(fixture, line, character, expect_title_substring, want_lines)
  local path = vim.fn.fnamemodify(fixture, ":p")
  vim.cmd.edit({ path, bang = true })

  local attached = vim.wait(3000, function()
    local clients = vim.lsp.get_clients({ name = "edifact_ls", bufnr = 0 })
    return #clients > 0 and clients[1].initialized == true
  end, 50)
  if not attached then
    fail("edifact_ls LSP client did not attach to " .. path .. " within timeout")
    return
  end

  local params = {
    textDocument = vim.lsp.util.make_text_document_params(0),
    range = {
      start = { line = line, character = character },
      ["end"] = { line = line, character = character + 1 },
    },
    context = { diagnostics = {} },
  }
  local results = vim.lsp.buf_request_sync(0, "textDocument/codeAction", params, 3000)

  local action = nil
  for _, res in pairs(results or {}) do
    for _, a in ipairs(res.result or {}) do
      if a.title and a.title:find(expect_title_substring, 1, true) then
        action = a
      end
    end
  end

  if not action then
    fail("codeAction at " .. path .. ":" .. line .. ":" .. character ..
      " did not include an action titled like " .. vim.inspect(expect_title_substring) ..
      "; got: " .. vim.inspect(results))
    return
  end

  vim.lsp.util.apply_workspace_edit(action.edit, "utf-16")
  vim.wait(500)

  local errors = vim.tbl_filter(function(d) return d.severity == vim.diagnostic.severity.ERROR end, vim.diagnostic.get(0))
  if #errors > 0 then
    fail("after applying code action at " .. path .. ", still have error diagnostics: " ..
      vim.inspect(vim.tbl_map(function(d) return d.message end, errors)))
    return
  end

  if want_lines then
    local lines = vim.api.nvim_buf_get_lines(0, 0, -1, false)
    if #lines ~= #want_lines then
      fail("after applying code action at " .. path .. ", buffer has " .. #lines ..
        " lines, want " .. #want_lines .. ": " .. vim.inspect(lines))
      return
    end
    for i, w in ipairs(want_lines) do
      if lines[i] ~= w then
        fail("after applying code action at " .. path .. ", line " .. i .. " = " ..
          vim.inspect(lines[i]) .. ", want " .. vim.inspect(w))
        return
      end
    end
  end

  pass("codeAction at " .. path .. ":" .. line .. ":" .. character ..
    " (" .. expect_title_substring .. ") applies cleanly")
end

check_lsp_attaches()
check_diagnostic("testdata/syntax-error.edi", "invalid segment tag", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/envelope-error.edi", "missing UNZ", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/lint-warning.edi", "reserved", vim.diagnostic.severity.WARN)
check_diagnostic("testdata/lint-info.edi", "version 4", vim.diagnostic.severity.INFO)
check_diagnostic("testdata/iftmcs-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/content-violation.edi", "function code qualifier", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/desadv-violation.edi", "maximum of 5", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/iftsta-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/invrpt-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/delfor-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/orders-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/ordrsp-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/invoic-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/iftmin-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/pricat-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/aperak-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/cuscar-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/orders-d99b-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/ordrsp-d99b-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/invoic-d99b-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/desadv-d99b-violation.edi", "maximum of 5", vim.diagnostic.severity.ERROR)
-- Spot-checks of edifact-ls-13gu's bulk D.20A batch (182 new message
-- types) in a real nvim session -- not one per type (182 would be
-- redundant with the generic unit tests), but representative samples
-- including the two real bugs the batch surfaced and fixed:
check_diagnostic("testdata/qality-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR) -- exercises the wrapped-segment-name parsing fix
check_no_error_diagnostic("testdata/cusdec-conformant.edi") -- exercises the adjacent-same-tag (double UNS) matcher fix directly
check_diagnostic("testdata/genral-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/mscons-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_diagnostic("testdata/coarri-violation.edi", "maximum of 1", vim.diagnostic.severity.ERROR)
check_no_error_diagnostic("testdata/baplie-conformant.edi")
check_hover("testdata/minimal.edi", 0, 1, "Interchange header")
check_hover("testdata/minimal.edi", 2, 1, "Beginning of message")
-- Regression: a segment whose data is soft-wrapped across lines (a literal
-- embedded newline, not just one between segments) used to hang the server
-- entirely (see internal/edifact/lexer.go's next()) -- if that regressed,
-- this hover would simply never come back within its wait timeout.
check_hover("testdata/embedded-newline.edi", 4, 1, "Name and address")
-- edifact-ls-pcm0: tier-3 hover, message-specific segment-group context.
-- TOD is the leading segment of IFTMCS's real segment group 2 and isn't
-- in the tier-1 tag table, so this also exercises the "group context
-- alone, no tier-1 description" fallback.
check_hover("testdata/iftmcs-group-context.edi", 3, 1, "Part of segment group 2")
-- edifact-ls-x3pb: quick-fix code actions for the two mechanically-fixable
-- diagnostic kinds.
-- want_lines asserts the whole UNA line is gone (not just emptied) --
-- regression coverage for the trailing-newline fix in edifact-ls-x3pb.
check_code_action("testdata/lint-info.edi", 0, 0, "Remove redundant UNA", {
  "UNB+UNOA:1+SENDER:ZZ+RECEIVER:ZZ+201001:1200+1'",
  "UNH+1+ORDERS:D:96A:UN'",
  "BGM+220'",
  "UNT+3+1'",
  "UNZ+1+1'",
})
-- Cursor on the wrong value itself ("99" in "UNT+99"), not on the UNT tag --
-- regression coverage for the range-matching fix in edifact-ls-x3pb (the
-- action used to be reachable only from the tag's exact first byte).
check_code_action("testdata/envelope-count-mismatch.edi", 3, 4, "Fix UNT segment count")
check_formatting()
check_minify()
check_treesitter()
check_treesitter("testdata/functional-group.edi")

if failed then
  vim.cmd("cquit 1")
else
  -- Force-quit: checks like check_formatting intentionally leave buffers
  -- modified (we never save), so a plain :qa would refuse with E37.
  vim.cmd("qa!")
end
