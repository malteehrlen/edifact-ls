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
