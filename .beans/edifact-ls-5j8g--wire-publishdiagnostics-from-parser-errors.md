---
# edifact-ls-5j8g
title: Wire publishDiagnostics from parser errors
status: todo
type: feature
priority: normal
created_at: 2026-09-01T13:03:16Z
updated_at: 2026-09-01T13:03:16Z
parent: edifact-ls-z9te
---

# Description
On document open/change, run the parser core over the document text and
translate its structured errors into LSP `Diagnostic` objects, published via
`textDocument/publishDiagnostics`.

# Acceptance Criteria
- [ ] Diagnostics published on open and on every change (debounced if needed)
- [ ] Byte/line/column positions from the parser correctly map to LSP
      `Range` (UTF-16 code unit columns per spec)
- [ ] Previously-published diagnostics are cleared/replaced (not accumulated)
      on each new publish
- [ ] Unit test driving the server with a known-bad document and asserting
      the exact diagnostics published
