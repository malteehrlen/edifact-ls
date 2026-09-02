---
# edifact-ls-5wdy
title: CUSCAR D.99B schema data
status: completed
type: epic
priority: normal
created_at: 2026-09-02T08:27:17Z
updated_at: 2026-09-02T08:30:52Z
parent: edifact-ls-gdt6
---

# Goal

Add structural schema validation for CUSCAR (Customs cargo report
message), requested directly by the user against a real file they hit
the embedded-newline lexer bug on (test.edi, declaring
`CUSCAR:D:96B:UN`).

# Release: D.99B, not the originally-linked D.00A or the file's own D.96B

The user originally linked
https://service.unece.org/trade/untdid/d00a/trmd/cuscar_c.htm (release
D.00A). Their own file declares release D.96B. Checked before
implementing, per this project's established habit of verifying a
release is actually sourceable before committing to it:

- D.96B itself is not archived by the Wayback Machine under any path
  checked.
- D.96A *is* archived, but the snapshot is a stub -- literally
  placeholder text ("There is some standard text here"), not a real
  segment table (confirmed by reading it directly).
- D.99B and D.01B are both archived with real, full segment tables
  (~730 lines each, real "Pos Tag Name" header at line ~598/596).

Asked the user directly which to use; they picked **D.99B**. Registers
as MessageID{Type: "CUSCAR", Version: "D", Release: "99B", Agency:
"UN"} -- this will not structurally validate the user's own D.96B file
(different exact tuple, per this project's registry design), but will
produce the "recognized type, here's what's registered" info diagnostic
for it rather than silence, and gives real validation to any D.99B
CUSCAR message.

Source: https://service.unece.org/trade/untdid/d99b/trmd/cuscar_c.htm
(403s directly via Cloudflare, same as every other service.unece.org
page hit this project); archived copy used instead:
http://web.archive.org/web/20220126231724/https://service.unece.org/trade/untdid/d99b/trmd/cuscar_c.htm

# Approach

Same as every other message-type story (edifact-ls-7uhx and
edifact-ls-oton's stories): parse the source's exact rail-art column
positions mechanically (not by eye), verify the tree balances, generate
the Go SchemaNode literal from the verified tree, register, unit test
against real data, e2e-verify a real violation in nvim.

# Acceptance Criteria

[x] CUSCAR D.99B's real branching diagram transcribed accurately from
the cited source, verified to balance before transcription
[x] Registered for the exact tuple (CUSCAR, D, 99B, UN)
[x] Unit tests: a conformant CUSCAR message passes with no structural
violations; at least one fixture produces a real violation the
actual fetched structure supports
[x] e2e check: opening a fixture with a structural CUSCAR violation
shows the diagnostic in nvim
[x] Source URL(s) cited in the schema data's source comment, including
the Cloudflare/Wayback caveat and the D.96A-stub/D.99B-chosen note

## Summary of Changes

CUSCAR D.99B registered (17 segment groups, max depth 4), real-data
tested, e2e-verified. Confirmed the actual outcome for the user's own
D:96B file: correctly produces an info diagnostic naming what's
registered rather than a false match or silence.

## Retro

- The pre-implementation check paid off directly: the URL the user
  linked (D.00A) and their file's own declared release (D.96B) didn't
  match each other, and D.96B turned out not to be archived at all --
  checking this *before* implementing, and asking rather than guessing
  which nearby release to use, avoided shipping a schema that either
  silently didn't apply to the user's real file or (worse) quietly
  validated it against the wrong release's rules.
- A second, unplanned discovery mid-verification: the first archived
  D.96A snapshot looked promising (present, 200 OK) but turned out to
  be a stub page with placeholder text, not real content. Worth the
  reminder this establishes: "archived" and "usable" aren't the same
  thing -- always read a small sample of fetched content before
  trusting it, not just check that a fetch succeeded.
- The reusable extraction script (generalized across every message-type
  epic so far) hit its first real format mismatch on this older-vintage
  page: 4-digit position numbers and an actual "change indicator"
  column neither of the two page generations used previously exercised.
  Fixed by generalizing the script's parsing rather than hand-patching
  this one file's data -- keeping it a genuinely reusable tool as the
  corpus of sourced pages grows, not something that silently degrades
  into one-off hacks per message type.
- Caught my own wrong assumption before it became a bad test: an
  initial reading of the rendered tree assumed segment group 14's
  mandatory flag applied unconditionally at message level; building the
  "minimal conformant message" test against that assumption failed
  immediately, which is exactly the value of testing against the real
  generated schema rather than trusting a summary of it.
- This was a much smaller, single-message unit of work than the
  batch epics before it, and the ceremony (a whole epic + retro) is
  proportionally heavier than the work -- worth considering, next time
  a single ad-hoc message type is requested outside a larger batch,
  whether a plain story under an existing umbrella epic (or no epic at
  all) fits better than spinning up a new one.
