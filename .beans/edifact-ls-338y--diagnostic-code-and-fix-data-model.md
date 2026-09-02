---
# edifact-ls-338y
title: Diagnostic Code + Fix data model
status: completed
type: feature
priority: normal
created_at: 2026-09-02T11:38:49Z
updated_at: 2026-09-02T11:42:35Z
parent: edifact-ls-x3pb
---

# Description

Lay the data-model groundwork the code-action handler (edifact-ls-n3yd)
will consume. No LSP-visible behavior change in this story.

Add to `edifact.Error`:

```go
type Error struct {
    Pos      Position
    Severity Severity
    Message  string
    Code     string // stable id, e.g. "redundant-una", "envelope-count-mismatch"
    Fix      *Fix   // nil if not mechanically fixable
}

type Fix struct {
    Title   string   // human-readable, e.g. "Remove redundant UNA service string advice"
    Pos     Position // start of the exact span to replace
    OldText string   // sanity-checked against current buffer text before applying
    NewText string   // "" for a deletion
}
```

Populate `Code`+`Fix` at the two existing call sites that already know
the fix:

- `lint.go`'s redundant-`UNA` diagnostic: `ic.UNA.Pos`/`ic.UNA.Raw` are
  already available there.
- `envelope.go`'s six `UNT`/`UNE`/`UNZ` count-and-reference comparisons:
  each already computes the "correct" value it's comparing against.
  These currently only hold the component's string value, not its
  `Pos`/`Raw` -- switch to `Element0(i).Components[0]` (or equivalent)
  at each of the six call sites to get an exact replaceable span.

Thread `Code` into `internal/lspserver/diagnostics.go`'s existing
translation loop, setting `protocol.Diagnostic.Code` wherever
`edifact.Error.Code` is non-empty. This gives editors/clients a stable,
filterable id independent of message wording even before code actions
exist.

Every other diagnostic source (`ValidateMessageSchemas`,
`ValidateSegmentContent`, `Parse`, the reserved-tag warning in `Lint`)
gets no `Code`/`Fix` -- there's nothing to derive.

# Acceptance Criteria

[x] `Error.Code` and `Error.Fix` (with `Fix.{Title,Pos,OldText,NewText}`)
added to `internal/edifact`
[x] Redundant-UNA diagnostic carries `Code: "redundant-una"` and a
correct `Fix` deleting exactly the `UNA` segment's raw text
[x] All six envelope count/reference mismatches carry a `Code` (e.g.
`"envelope-count-mismatch"` / `"envelope-reference-mismatch"`) and a
correct `Fix` replacing exactly the wrong component with the value
already computed as correct
[x] `internal/lspserver/diagnostics.go` sets `protocol.Diagnostic.Code`
from `edifact.Error.Code` wherever present
[x] Unit tests assert `Fix` is populated with the right
`Pos`/`OldText`/`NewText` for the redundant-UNA case and each of the
six envelope-mismatch cases, and is nil for every non-fixable
diagnostic kind
[x] Full suite (`make test`) and e2e harness (`make test-e2e`) still
green -- no LSP-facing behavior change expected from this story

## Summary of Changes

internal/edifact/edifact.go: added `Error.Code`/`Error.Fix` and a `Fix`
type (`Title`, `Pos`, `OldText`, `NewText`), plus `ErrorList.AddFixable`
alongside the existing `Add` (kept unchanged, still used by the 31
non-fixable call sites -- no signature break).

internal/edifact/lint.go: the redundant-UNA diagnostic now uses
AddFixable with Code "redundant-una" and a Fix deleting exactly
`ic.UNA.Raw` (the 9-byte advice including the terminator byte, which is
one of the 6 delimiter characters UNA packs rather than a separately
appended terminator).

internal/edifact/envelope.go: added two small helpers --
`componentRaw` (returns a component's still-escaped Raw text, or "" if
absent) and `fixReplaceComponent` (builds a Fix replacing one
component's exact span, or nil if the component doesn't exist -- an
absent component is an insertion, not a replacement, and this story
intentionally doesn't attempt that). All six count/reference mismatch
call sites (UNZ count, UNZ reference, UNE count, UNE reference, UNT
count, UNT reference) now use AddFixable with Code
"envelope-count-mismatch" or "envelope-reference-mismatch" and a Fix
built from these helpers. Reference fixes copy the source component's
Raw text verbatim (rather than its unescaped Value) so no re-escaping
decision has to be made when splicing it into another segment.

internal/lspserver/diagnostics.go: threads `Error.Code` into
`protocol.Diagnostic.Code` via a small `diagnosticCode` helper (nil
when Code is empty).

18 new unit tests across lint_test.go and envelope_test.go. The six
envelope-mismatch tests go further than asserting Fix's fields are
non-nil: `assertFixResolves` actually applies the Fix to the source
text and re-validates the result, confirming the fix genuinely
resolves the diagnostic rather than just existing. Two tests
(reserved-tag warning, missing-UNT) confirm non-fixable diagnostics
correctly have Code=="" and Fix==nil. Full suite (`make test`) and
e2e harness (`make test-e2e`, all 35 checks) still green -- as
expected, no LSP-facing behavior changed in this story.
