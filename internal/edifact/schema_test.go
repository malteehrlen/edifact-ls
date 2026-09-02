package edifact

import "testing"

// seg builds a minimal Segment for schema tests -- only Tag and a position
// (used to tell diagnostics apart by which segment they point at) matter
// here; schema validation never looks at Elements.
func seg(tag string, line int) Segment {
	return Segment{Tag: tag, Pos: Position{Line: line, Column: 1}}
}

var testEnd = Position{Line: 999, Column: 1}

func TestValidateSchemaCleanPass(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "A", Mandatory: true, MaxRepeat: 1},
		{Segment: "B", Mandatory: false, MaxRepeat: 3},
	}}
	segments := []Segment{seg("A", 1), seg("B", 2), seg("B", 3)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 0 {
		t.Fatalf("unexpected diagnostics for a conformant sequence: %v", errs)
	}
}

func TestValidateSchemaMissingMandatorySegment(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "A", Mandatory: true, MaxRepeat: 1},
		{Segment: "B", Mandatory: true, MaxRepeat: 1},
	}}
	segments := []Segment{seg("A", 1)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityError {
		t.Errorf("severity = %v, want error", errs[0].Severity)
	}
	if !containsMessage(errs, "missing mandatory") || !containsMessage(errs, `"B"`) {
		t.Errorf("message = %q, want it to mention missing mandatory segment B", errs[0].Message)
	}
	if errs[0].Pos != testEnd {
		t.Errorf("pos = %+v, want the end-of-message fallback %+v (stream ran out)", errs[0].Pos, testEnd)
	}
}

func TestValidateSchemaMissingMandatoryGroup(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "A", Mandatory: true, MaxRepeat: 1},
		{
			Group: []SchemaNode{
				{Segment: "X", Mandatory: true, MaxRepeat: 1},
				{Segment: "Y", Mandatory: false, MaxRepeat: 9},
			},
			Mandatory: true,
			MaxRepeat: 1,
		},
	}}
	segments := []Segment{seg("A", 1)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if !containsMessage(errs, "missing mandatory") || !containsMessage(errs, "segment group starting with \"X\"") {
		t.Errorf("message = %q, want it to mention the missing mandatory group", errs[0].Message)
	}
}

func TestValidateSchemaExceededRepeat(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "A", Mandatory: true, MaxRepeat: 1},
		{Segment: "B", Mandatory: false, MaxRepeat: 2},
	}}
	segments := []Segment{seg("A", 1), seg("B", 2), seg("B", 3), seg("B", 4)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityError {
		t.Errorf("severity = %v, want error", errs[0].Severity)
	}
	if !containsMessage(errs, `"B"`) || !containsMessage(errs, "maximum of 2") {
		t.Errorf("message = %q, want it to mention B exceeding its max repeat of 2", errs[0].Message)
	}
	if errs[0].Pos.Line != 4 {
		t.Errorf("pos.Line = %d, want 4 (the 3rd, excess B)", errs[0].Pos.Line)
	}
}

func TestValidateSchemaWrongOrder(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "A", Mandatory: true, MaxRepeat: 1},
		{Segment: "B", Mandatory: true, MaxRepeat: 1},
	}}
	// B appears where A was expected, and A shows up only afterwards --
	// out of sequence rather than missing outright.
	segments := []Segment{seg("B", 1), seg("A", 2)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 2 {
		t.Fatalf("got %d diagnostics, want 2: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityError || !containsMessage(errs[:1], `"A"`) {
		t.Errorf("first diagnostic = %v, want an error about missing mandatory A", errs[0])
	}
	if errs[1].Severity != SeverityWarning || !containsMessage(errs[1:], "not expected here") {
		t.Errorf("second diagnostic = %v, want a warning about the leftover, out-of-place A", errs[1])
	}
}

// TestValidateSchemaNestedGroups exercises a group containing a nested,
// repeating child group -- the multi-level nesting real message types like
// IFTMCS rely on.
func TestValidateSchemaNestedGroups(t *testing.T) {
	inner := SchemaNode{
		Group: []SchemaNode{
			{Segment: "C", Mandatory: true, MaxRepeat: 1},
			{Segment: "D", Mandatory: false, MaxRepeat: 9},
		},
		Mandatory: false,
		MaxRepeat: 3,
	}
	outer := SchemaNode{
		Group: []SchemaNode{
			{Segment: "B", Mandatory: true, MaxRepeat: 1},
			inner,
		},
		Mandatory: true,
		MaxRepeat: 1,
	}
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "A", Mandatory: true, MaxRepeat: 1},
		outer,
	}}

	// A, then one Outer occurrence: B, then three Inner occurrences
	// (C+D+D, C alone, C alone).
	segments := []Segment{
		seg("A", 1), seg("B", 2),
		seg("C", 3), seg("D", 4), seg("D", 5),
		seg("C", 6),
		seg("C", 7),
	}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 0 {
		t.Fatalf("unexpected diagnostics for a conformant nested sequence: %v", errs)
	}
}

// TestValidateSchemaNestedGroupExceedsOwnRepeat confirms that a repeating
// group exceeding its own cap is reported once, at the group level --
// not misattributed to its first child's local cap (see the insideGroup
// exemption in matchSequence).
func TestValidateSchemaNestedGroupExceedsOwnRepeat(t *testing.T) {
	inner := SchemaNode{
		Group: []SchemaNode{
			{Segment: "C", Mandatory: true, MaxRepeat: 1},
		},
		Mandatory: false,
		MaxRepeat: 2,
	}
	schema := Schema{Nodes: []SchemaNode{inner}}
	// Three occurrences of the C-group where only two are allowed.
	segments := []Segment{seg("C", 1), seg("C", 2), seg("C", 3)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if !containsMessage(errs, "segment group starting with \"C\"") || !containsMessage(errs, "maximum of 2") {
		t.Errorf("message = %q, want it to mention the group (not the leaf) exceeding its max repeat of 2", errs[0].Message)
	}
}

func TestValidateSchemaNestedGroupMissingMandatoryChild(t *testing.T) {
	outer := SchemaNode{
		Group: []SchemaNode{
			{Segment: "B", Mandatory: true, MaxRepeat: 1},
			{Segment: "C", Mandatory: true, MaxRepeat: 1},
		},
		Mandatory: true,
		MaxRepeat: 1,
	}
	schema := Schema{Nodes: []SchemaNode{outer}}
	// The group starts (B present) but its mandatory child C is missing.
	segments := []Segment{seg("B", 1)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if !containsMessage(errs, "missing mandatory") || !containsMessage(errs, `"C"`) {
		t.Errorf("message = %q, want it to mention the group's missing mandatory child C", errs[0].Message)
	}
}

// TestValidateSchemaAdjacentSiblingsSharingLeadingTag is a regression
// test for a real bug found while bulk-sourcing edifact-ls-13gu: two
// different, independent schema nodes at the same level that happen to
// share a leading tag (e.g. UN/EDIFACT's real convention of two
// separate UNS occurrences marking the header/detail and
// detail/summary section boundaries) used to be misattributed --
// the first node's own cap would "steal" the second occurrence as its
// own overflow, then the second node would report a spurious
// missing-mandatory error, even though the message was fully
// conformant.
func TestValidateSchemaAdjacentSiblingsSharingLeadingTag(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
		{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // header/detail boundary
		{Segment: "UNS", Mandatory: true, MaxRepeat: 1}, // detail/summary boundary
	}}
	segments := []Segment{seg("BGM", 1), seg("UNS", 2), seg("UNS", 3)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 0 {
		t.Fatalf("unexpected diagnostics for two legitimately separate same-tag occurrences: %v", errs)
	}
}

// TestValidateSchemaAdjacentSiblingsSharingLeadingTagAcrossOptionalGap
// covers the same scenario but with a conditional sibling the message
// instance doesn't use sitting between the two same-tag nodes in the
// schema (as CONPVA's and CUSDEC's real schemas do) -- confirming the
// fix looks past every remaining sibling, not just the immediately
// next one.
func TestValidateSchemaAdjacentSiblingsSharingLeadingTagAcrossOptionalGap(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
		{Segment: "UNS", Mandatory: true, MaxRepeat: 1},
		{Segment: "TAX", Mandatory: false, MaxRepeat: 9}, // unused optional gap
		{Segment: "UNS", Mandatory: true, MaxRepeat: 1},
	}}
	segments := []Segment{seg("BGM", 1), seg("UNS", 2), seg("UNS", 3)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 0 {
		t.Fatalf("unexpected diagnostics with an unused optional sibling between the two UNS nodes: %v", errs)
	}
}

// TestValidateSchemaGenuineOverflowStillDetectedNearSameTagSiblings
// confirms the fix didn't overcorrect: a real overflow (more
// occurrences than any node, current or later, can account for) must
// still be reported.
func TestValidateSchemaGenuineOverflowStillDetectedNearSameTagSiblings(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
		{Segment: "UNS", Mandatory: true, MaxRepeat: 1},
		{Segment: "UNS", Mandatory: true, MaxRepeat: 1},
	}}
	// Three UNS occurrences, but only two schema slots exist for it.
	segments := []Segment{seg("BGM", 1), seg("UNS", 2), seg("UNS", 3), seg("UNS", 4)}

	errs := ValidateSchema(schema, segments, testEnd)
	if len(errs) != 1 || !containsMessage(errs, `"UNS"`) || !containsMessage(errs, "maximum of 1") {
		t.Fatalf("got %v, want exactly one error about UNS exceeding its max repeat of 1", errs)
	}
}

func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGroupPathAtTopLevelSegmentHasNoPath(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
	}}
	segments := []Segment{seg("BGM", 1)}
	if got := GroupPathAt(schema, segments, 0); got != nil {
		t.Fatalf("GroupPathAt = %v, want nil for a top-level segment", got)
	}
}

func TestGroupPathAtSingleGroup(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
		{Group: []SchemaNode{
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},
			{Segment: "CTA", Mandatory: false, MaxRepeat: 1},
		}, Mandatory: false, MaxRepeat: 1},
	}}
	segments := []Segment{seg("BGM", 1), seg("NAD", 2), seg("CTA", 3)}

	if got := GroupPathAt(schema, segments, 0); got != nil {
		t.Errorf("BGM: GroupPathAt = %v, want nil (top-level)", got)
	}
	if got := GroupPathAt(schema, segments, 1); !intsEqual(got, []int{1}) {
		t.Errorf("NAD: GroupPathAt = %v, want [1] (segment group 1)", got)
	}
	if got := GroupPathAt(schema, segments, 2); !intsEqual(got, []int{1}) {
		t.Errorf("CTA: GroupPathAt = %v, want [1]", got)
	}
}

func TestGroupPathAtNestedGroups(t *testing.T) {
	// Numbering must follow the same preorder sequence UN/EDIFACT itself
	// uses (matching the "Segment group N" comments this project's real
	// generated schema data already carries, e.g. iftmcs_d21a.go): group
	// 1 first, then group 2 nested inside it, then group 3 as a sibling
	// of group 1 -- not depth-first-complete-subtree-then-number.
	schema := Schema{Nodes: []SchemaNode{
		{Segment: "BGM", Mandatory: true, MaxRepeat: 1},
		{Group: []SchemaNode{ // segment group 1
			{Segment: "TOD", Mandatory: true, MaxRepeat: 1},
			{Group: []SchemaNode{ // segment group 2, nested in 1
				{Segment: "LOC", Mandatory: true, MaxRepeat: 1},
			}, Mandatory: false, MaxRepeat: 1},
		}, Mandatory: false, MaxRepeat: 1},
		{Group: []SchemaNode{ // segment group 3, a sibling of group 1
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},
		}, Mandatory: false, MaxRepeat: 1},
	}}
	segments := []Segment{seg("BGM", 1), seg("TOD", 2), seg("LOC", 3), seg("NAD", 4)}

	if got := GroupPathAt(schema, segments, 1); !intsEqual(got, []int{1}) {
		t.Errorf("TOD: GroupPathAt = %v, want [1]", got)
	}
	if got := GroupPathAt(schema, segments, 2); !intsEqual(got, []int{1, 2}) {
		t.Errorf("LOC: GroupPathAt = %v, want [1 2] (nested)", got)
	}
	if got := GroupPathAt(schema, segments, 3); !intsEqual(got, []int{3}) {
		t.Errorf("NAD: GroupPathAt = %v, want [3], not renumbered by group 2's existence", got)
	}
}

// TestGroupPathAtNumberingStableAcrossRepeats is a regression test for the
// exact bug numberGroups's doc comment warns about: a nested group's
// number must be the same for every occurrence of its repeating parent,
// not incremented per occurrence.
func TestGroupPathAtNumberingStableAcrossRepeats(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{
		{Group: []SchemaNode{ // segment group 1, repeats up to 3 times
			{Segment: "NAD", Mandatory: true, MaxRepeat: 1},
			{Group: []SchemaNode{ // segment group 2, nested in 1
				{Segment: "CTA", Mandatory: true, MaxRepeat: 1},
			}, Mandatory: false, MaxRepeat: 1},
		}, Mandatory: false, MaxRepeat: 3},
	}}
	segments := []Segment{
		seg("NAD", 1), seg("CTA", 2), // 1st occurrence of group 1
		seg("NAD", 3), seg("CTA", 4), // 2nd occurrence
		seg("NAD", 5), seg("CTA", 6), // 3rd occurrence
	}

	for i := 0; i < len(segments); i += 2 {
		if got := GroupPathAt(schema, segments, i); !intsEqual(got, []int{1}) {
			t.Errorf("NAD at index %d: GroupPathAt = %v, want [1] every time", i, got)
		}
		if got := GroupPathAt(schema, segments, i+1); !intsEqual(got, []int{1, 2}) {
			t.Errorf("CTA at index %d: GroupPathAt = %v, want [1 2] every time, not renumbered per repeat", i+1, got)
		}
	}
}

func TestGroupPathAtOutOfRangeIndex(t *testing.T) {
	schema := Schema{Nodes: []SchemaNode{{Segment: "BGM", Mandatory: true, MaxRepeat: 1}}}
	segments := []Segment{seg("BGM", 1)}
	if got := GroupPathAt(schema, segments, -1); got != nil {
		t.Errorf("negative index: GroupPathAt = %v, want nil", got)
	}
	if got := GroupPathAt(schema, segments, 5); got != nil {
		t.Errorf("out-of-range index: GroupPathAt = %v, want nil", got)
	}
}
