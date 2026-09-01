package edifact

import "testing"

func TestSegmentElementsRegistered(t *testing.T) {
	for _, tag := range []string{"BGM", "DTM", "CTA"} {
		if _, ok := segmentElementSchemas[tag]; !ok {
			t.Errorf("%s has no registered element schema", tag)
		}
	}
}

func TestDTMConformantMessagePasses(t *testing.T) {
	ic, errs := Parse("UNH+1'DTM+137:20100101:102'UNT+3+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateSegmentContent(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a conformant DTM: %v", got)
	}
}

func TestDTMMissingMandatoryQualifier(t *testing.T) {
	// The qualifier component (2005) is missing -- only the date text is
	// given, e.g. ":20100101:102" instead of "137:20100101:102".
	ic, errs := Parse("UNH+1'DTM+:20100101:102'UNT+3+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}

	got := ValidateSegmentContent(ic)
	if len(got) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(got), got)
	}
	if !containsMessage(got, `"DTM"`) || !containsMessage(got, "component 1") || !containsMessage(got, "function code qualifier") {
		t.Errorf("diagnostic = %v, want an error naming DTM's missing qualifier component", got[0])
	}
}

func TestDTMMissingEntirely(t *testing.T) {
	ic, errs := Parse("UNH+1'DTM'UNT+2+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}

	got := ValidateSegmentContent(ic)
	if len(got) != 1 || !containsMessage(got, "missing its mandatory element 1") {
		t.Fatalf("got %v, want one error about the missing mandatory Date/time/period element", got)
	}
}

func TestBGMMinimalNeverFlags(t *testing.T) {
	// BGM with no elements at all -- nothing is mandatory, so nothing
	// should be flagged.
	ic, errs := Parse("UNH+1'BGM'UNT+2+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateSegmentContent(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal BGM: %v", got)
	}
}

func TestCTAMinimalNeverFlags(t *testing.T) {
	ic, errs := Parse("UNH+1'CTA'UNT+2+1'")
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateSegmentContent(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal CTA: %v", got)
	}
}
