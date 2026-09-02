package edifact

import "testing"

func TestAPERAKD00BRegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "APERAK", Version: "D", Release: "00B", Agency: "UN"}]; !ok {
		t.Fatal("APERAK D.00B schema is not registered")
	}
}

func TestAPERAKD00BMinimalConformantMessage(t *testing.T) {
	src := "UNH+1+APERAK:D:00B:UN'BGM+311'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant APERAK D.00B message: %v", got)
	}
}

func TestAPERAKD00BMissingMandatoryBGM(t *testing.T) {
	src := "UNH+1+APERAK:D:00B:UN'DTM+137'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"BGM"`) {
		t.Fatalf("got %v, want an error about missing mandatory BGM", got)
	}
}

func TestAPERAKD00BExceededRepeat(t *testing.T) {
	// BGM itself is capped at 1 occurrence.
	src := "UNH+1+APERAK:D:00B:UN'BGM+311'BGM+311'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"BGM"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about BGM exceeding its max repeat of 1", got)
	}
}
