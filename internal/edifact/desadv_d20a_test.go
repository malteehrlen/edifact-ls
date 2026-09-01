package edifact

import "testing"

func TestDESADVD20ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "DESADV", Version: "D", Release: "20A", Agency: "UN"}]; !ok {
		t.Fatal("DESADV D.20A schema is not registered")
	}
}

func TestDESADVD20AMinimalConformantMessage(t *testing.T) {
	src := "UNH+1+DESADV:D:20A:UN'BGM+351'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant DESADV message: %v", got)
	}
}

func TestDESADVD20AMissingMandatoryBGM(t *testing.T) {
	src := "UNH+1+DESADV:D:20A:UN'DTM+11'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"BGM"`) {
		t.Fatalf("got %v, want an error about missing mandatory BGM", got)
	}
}

func TestDESADVD20AExceededRepeat(t *testing.T) {
	// ALI is conditional but capped at 5 occurrences.
	src := "UNH+1+DESADV:D:20A:UN'BGM+351'ALI'ALI'ALI'ALI'ALI'ALI'UNT+8+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"ALI"`) || !containsMessage(got, "maximum of 5") {
		t.Fatalf("got %v, want an error about ALI exceeding its max repeat of 5", got)
	}
}
