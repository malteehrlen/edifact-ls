package edifact

import "testing"

// buildSegment parses src as a single segment for content-validation
// tests -- simpler than hand-building Segment/Element/Component structs,
// and exercises the real parser/AST at the same time.
func buildSegment(t *testing.T, src string) Segment {
	t.Helper()
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors for %q: %v", src, errs)
	}
	if len(ic.Segments) != 1 {
		t.Fatalf("parsed %d segments from %q, want 1", len(ic.Segments), src)
	}
	return ic.Segments[0]
}

func TestValidateSegmentElementsCleanPass(t *testing.T) {
	schema := SegmentElementSchema{Elements: []ElementSchema{
		{Name: "A", Mandatory: true, Components: []ComponentSchema{{Name: "A", Mandatory: true}}},
		{Name: "B", Mandatory: false, Components: []ComponentSchema{{Name: "B", Mandatory: false}}},
	}}
	seg := buildSegment(t, "TST+1+2'")

	errs := ValidateSegmentElements(schema, seg)
	if len(errs) != 0 {
		t.Fatalf("unexpected diagnostics for a conformant segment: %v", errs)
	}
}

func TestValidateSegmentElementsMissingMandatorySimpleElement(t *testing.T) {
	schema := SegmentElementSchema{Elements: []ElementSchema{
		{Name: "A", Mandatory: true, Components: []ComponentSchema{{Name: "A", Mandatory: true}}},
	}}
	seg := buildSegment(t, "TST'") // no elements at all

	errs := ValidateSegmentElements(schema, seg)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityError || !containsMessage(errs, "missing its mandatory element 1") || !containsMessage(errs, "(A)") {
		t.Errorf("diagnostic = %v, want an error about the missing mandatory element A", errs[0])
	}
}

func TestValidateSegmentElementsMissingMandatoryCompositeElement(t *testing.T) {
	schema := SegmentElementSchema{Elements: []ElementSchema{
		{Name: "Composite", Mandatory: true, Components: []ComponentSchema{
			{Name: "C1", Mandatory: true},
			{Name: "C2", Mandatory: false},
		}},
	}}
	seg := buildSegment(t, "TST'") // the whole composite element is absent

	errs := ValidateSegmentElements(schema, seg)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if !containsMessage(errs, "missing its mandatory element 1") || !containsMessage(errs, "(Composite)") {
		t.Errorf("diagnostic = %v, want an error about the missing mandatory composite element", errs[0])
	}
}

func TestValidateSegmentElementsMissingMandatoryComponent(t *testing.T) {
	schema := SegmentElementSchema{Elements: []ElementSchema{
		{Name: "Composite", Mandatory: true, Components: []ComponentSchema{
			{Name: "Qualifier", Mandatory: true},
			{Name: "Text", Mandatory: false},
		}},
	}}
	// The composite is present (via its first component) but the
	// mandatory Qualifier component itself is missing -- e.g. ":something"
	// instead of "qualifier:something".
	seg := buildSegment(t, "TST+:something'")

	errs := ValidateSegmentElements(schema, seg)
	if len(errs) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(errs), errs)
	}
	if !containsMessage(errs, "element 1 (Composite)") || !containsMessage(errs, "component 1 (Qualifier)") {
		t.Errorf("diagnostic = %v, want an error naming the missing mandatory Qualifier component", errs[0])
	}
}

// TestValidateSegmentElementsAllConditionalNeverFlags mirrors BGM/CTA's
// real shape (every element and component conditional) -- confirms the
// validator never false-positives when nothing is actually mandatory,
// even for a segment with no elements at all.
func TestValidateSegmentElementsAllConditionalNeverFlags(t *testing.T) {
	schema := SegmentElementSchema{Elements: []ElementSchema{
		{Name: "A", Mandatory: false, Components: []ComponentSchema{{Name: "A", Mandatory: false}}},
		{Name: "B", Mandatory: false, Components: []ComponentSchema{
			{Name: "B1", Mandatory: false},
			{Name: "B2", Mandatory: false},
		}},
	}}
	seg := buildSegment(t, "TST'")

	if errs := ValidateSegmentElements(schema, seg); len(errs) != 0 {
		t.Fatalf("unexpected diagnostics for an all-conditional schema: %v", errs)
	}
}
