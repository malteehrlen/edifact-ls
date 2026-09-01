---
# edifact-ls-fo5r
title: Wire textDocument/hover into the LSP server
status: completed
type: feature
priority: normal
created_at: 2026-09-01T17:53:47Z
updated_at: 2026-09-01T19:43:41Z
parent: edifact-ls-tnp9
blocked_by:
    - edifact-ls-1cwj
---

# Description

Wire `textDocument/hover` into the LSP server. Given a cursor position,
find the segment under it (reusing the same AST position-tracking
`diagnosticsForText`/`Lint` already use), look up its tag in the
description table from "Segment description data", and return a
`Hover` response with markdown `MarkupContent` (name + description). No
match -> return nil, not an empty/placeholder hover.

# Acceptance Criteria

[x] `initialize` response advertises `hoverProvider: true`
[x] `textDocument/hover` handler resolves the segment at the given
position and returns its description as markdown
[x] Position resolution reuses existing AST/position infrastructure
rather than re-parsing or re-implementing offset math
[x] Unrecognized tag or out-of-range position returns nil cleanly (no
panic, no empty-string hover)

## Summary of Changes

internal/lspserver/lspserver.go: wired `TextDocumentHover:
st.textDocumentHover` into the handler struct. Capabilities are
derived from the handler (existing pattern), so `hoverProvider: true`
follows automatically -- locked in with an explicit assertion in
TestHandshake rather than just trusting the derivation.

internal/lspserver/hover.go (new): textDocumentHover converts the
LSP position to a byte offset, parses the document with
edifact.Parse (same as diagnostics), and finds which segment's
[Pos.Offset, EndPos.Offset) span contains it via segmentTagAt --
matching the whole segment's span rather than just its 3 tag
characters, so hovering anywhere on the segment's line resolves it
(deliberate: simpler, and doesn't preclude tier-2 element-level hover
later, which can just take priority for element sub-spans once it
exists). UNA is handled too, via its own Raw/Pos fields, since it
isn't a Segment. No match, or a syntactically valid but undescribed
tag, returns nil.

internal/lspserver/diagnostics.go: added lspPositionToOffset, the
reverse of the existing offsetToLSPPosition, placed right next to it.
Note on the AC about reusing position infrastructure rather than
reimplementing offset math: this one function *is* new, since nothing
previously needed to go from an LSP position back to a byte offset
(diagnostics only ever go the other direction) -- but it's the direct
mirror of the existing conversion, not a parallel/duplicate system, and
everything else (parsing, segment spans) reuses what already exists.

internal/lspserver/lspserver_test.go: TestHandshake's "unhandled
method" example switched from textDocument/hover (now genuinely
handled) to textDocument/definition (still isn't); added an assertion
that Capabilities.HoverProvider is true.

internal/lspserver/hover_test.go (new): 7 unit tests -- service
segment (UNH), business segment (BGM), hovering an element rather than
the tag itself still resolves, UNA specifically, an unrecognized-but-
syntactically-valid tag, a position between segments, and a
never-opened document. All direct handler-func tests (no transport),
matching formatting_test.go's existing style.

Full suite (`make test`) and e2e harness (`make test-e2e`, unaffected
by this story) green.
