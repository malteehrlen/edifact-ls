package edifact

import "testing"

func TestIFTSTAD20ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "IFTSTA", Version: "D", Release: "20A", Agency: "UN"}]; !ok {
		t.Fatal("IFTSTA D.20A schema is not registered")
	}
}

func TestIFTSTAD20AMinimalConformantMessage(t *testing.T) {
	src := "UNH+1+IFTSTA:D:20A:UN'BGM+23'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant IFTSTA message: %v", got)
	}
}

func TestIFTSTAD20AMissingMandatoryBGM(t *testing.T) {
	src := "UNH+1+IFTSTA:D:20A:UN'DTM+132'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"BGM"`) {
		t.Fatalf("got %v, want an error about missing mandatory BGM", got)
	}
}

func TestIFTSTAD20AExceededRepeat(t *testing.T) {
	// TSR is conditional but capped at 1 occurrence.
	src := "UNH+1+IFTSTA:D:20A:UN'BGM+23'TSR'TSR'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"TSR"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about TSR exceeding its max repeat of 1", got)
	}
}
