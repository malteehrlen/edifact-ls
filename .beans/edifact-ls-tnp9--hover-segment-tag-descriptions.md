---
# edifact-ls-tnp9
title: 'Hover: segment-tag descriptions'
status: completed
type: epic
priority: normal
created_at: 2026-09-01T17:53:47Z
updated_at: 2026-09-01T19:45:30Z
parent: edifact-ls-gdt6
---

# Goal

Hover support, tier 1: hovering a segment tag (e.g. `BGM`) shows its
name and a one-line description (e.g. "Beginning of message -- a
segment for beginning, class identification of a message and its
number"), sourced from the UN Segment Directory (UNSD). This is
independent of message type -- it applies to any `.edi` file, not just
IFTMCS -- and doesn't need any of the UN message-directory sourcing the
structural spec-validation epic (edifact-ls-3uzr) requires. It's the
starting point of the hover strategy: ship this first, since it's small,
standalone, and immediately useful regardless of what happens with the
deeper tiers below.

# Strategy across all three hover epics

1. **This epic** (segment-tag hover) -- start here. Small, static lookup
data + a straightforward `textDocument/hover` handler reusing the AST
positions diagnostics already walk. No dependency on the message-spec
epics.
2. edifact-ls-6xaz (data-element/coded-value hover) -- deferred,
blocked by this epic (reuses its hover wiring). Needs deeper UN
directory data (element structure + code lists), likely shareable
with the content-validation epic (edifact-ls-9ger).
3. edifact-ls-pcm0 (message-context hover) -- deferred, blocked by
both this epic and the structural spec-validation epic
(edifact-ls-3uzr), since it needs to know which segment group a
given segment occurrence belongs to.

Only this epic is meant to be picked up now; 2 and 3 are intentionally
placeholders until their dependencies land.

# Acceptance Criteria

[x] `initialize` advertises `hoverProvider: true`
[x] `textDocument/hover` returns markdown (name + one-line description)
for a recognized segment tag under the cursor
[x] Unrecognized tags return no hover content (nil), not a placeholder
[x] e2e check: hovering a known segment tag in a fixture returns content
containing the expected description

## Summary of Changes

All three stories completed: a static segment-tag description table
(20 tags: all 8 recognized service segments plus 12 business
segments), a `textDocument/hover` handler matching whole segment spans
(not just the 3 tag characters, so hovering anywhere on a segment's
line resolves it) against that table, and end-to-end e2e coverage
confirming it works in a real nvim + LSP session for both a service
segment (UNB) and a business segment (BGM). `hoverProvider: true` is
derived automatically from the handler wiring, same pattern already
used for every other capability.

## Retro

- The "derive capabilities from which handler funcs are wired" pattern
  already established for formatting/diagnostics/commands paid off
  again here: adding hover was one line in the handler struct, and
  `hoverProvider: true` followed for free, tested explicitly rather
  than just trusted.
- Found and fixed a real, if narrow, side effect: TestHandshake had
  been using `textDocument/hover` as its canonical example of an
  *unhandled* LSP method. Implementing hover for real quietly turned
  that into a false assumption the test would have kept asserting
  wrongly forever if not caught -- swapped it for
  `textDocument/definition`, which is still genuinely unhandled.
  Worth remembering: a test that names a specific "this doesn't exist
  yet" method is a landmine for whatever epic eventually implements
  that method.
- The "match the whole segment span, not just the tag" design choice
  (deliberately broader than what the epic's own title literally says)
  was a judgment call made now, in a way that should stay compatible
  with tier 2 (element-level hover) later -- tier 2 can simply take
  priority for element sub-spans once it exists, since a segment's
  span still contains its elements' spans. Worth confirming this
  actually holds when tier 2 is picked up, rather than assuming it.
- No real surprises otherwise -- this epic was exactly as small and
  self-contained as scoped, unlike the structural-validation epic
  before it. No adjustments needed before starting whatever's next.
