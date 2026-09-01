package edifact

import "testing"

func TestValidateSegmentContentAppliesRegisteredSchema(t *testing.T) {
	RegisterSegmentElementSchema("TST", SegmentElementSchema{Elements: []ElementSchema{
		{Name: "A", Mandatory: true, Components: []ComponentSchema{{Name: "A", Mandatory: true}}},
	}})
	defer delete(segmentElementSchemas, "TST")

	ic, errs := Parse("UNH+1'TST'UNT+2+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}

	got := ValidateSegmentContent(ic)
	if len(got) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(got), got)
	}
	if !containsMessage(got, `"TST"`) || !containsMessage(got, "missing its mandatory element 1") {
		t.Errorf("diagnostic = %v, want an error about TST's missing mandatory element", got[0])
	}
}

func TestValidateSegmentContentUnregisteredTagIsSilent(t *testing.T) {
	ic, errs := Parse("UNH+1'BGM'UNT+2+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	// No schema registered for BGM in this test (real BGM data doesn't
	// exist yet -- see edifact-ls-arr3).
	if got := ValidateSegmentContent(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for an unregistered tag: %v", got)
	}
}

// TestValidateSegmentContentPlugAndPlaySecondTag registers two unrelated
// tags back to back and validates a message containing both, proving the
// registry/dispatch needs no per-tag code -- only a registration call --
// mirroring the same guardrail edifact-ls-ogqj established for the
// message-schema registry.
func TestValidateSegmentContentPlugAndPlaySecondTag(t *testing.T) {
	RegisterSegmentElementSchema("AAA", SegmentElementSchema{Elements: []ElementSchema{
		{Name: "X", Mandatory: true, Components: []ComponentSchema{{Name: "X", Mandatory: true}}},
	}})
	defer delete(segmentElementSchemas, "AAA")

	RegisterSegmentElementSchema("BBB", SegmentElementSchema{Elements: []ElementSchema{
		{Name: "Y", Mandatory: true, Components: []ComponentSchema{{Name: "Y", Mandatory: true}}},
	}})
	defer delete(segmentElementSchemas, "BBB")

	ic, errs := Parse("UNH+1'AAA'BBB+1'UNT+3+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}

	got := ValidateSegmentContent(ic)
	if len(got) != 1 || !containsMessage(got, `"AAA"`) {
		t.Fatalf("got %v, want exactly one violation, for AAA (BBB is conformant)", got)
	}
}
