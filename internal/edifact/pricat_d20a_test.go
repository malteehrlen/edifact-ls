package edifact

import "testing"

func TestPRICATD20ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "PRICAT", Version: "D", Release: "20A", Agency: "UN"}]; !ok {
		t.Fatal("PRICAT D.20A schema is not registered")
	}
}

func TestPRICATD20AMinimalConformantMessage(t *testing.T) {
	src := "UNH+1+PRICAT:D:20A:UN'BGM+9'DTM+137:20100101:102'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant PRICAT message: %v", got)
	}
}

func TestPRICATD20AMissingMandatoryDTM(t *testing.T) {
	src := "UNH+1+PRICAT:D:20A:UN'BGM+9'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"DTM"`) {
		t.Fatalf("got %v, want an error about missing mandatory DTM", got)
	}
}

func TestPRICATD20AExceededRepeat(t *testing.T) {
	// BGM itself is capped at 1 occurrence.
	src := "UNH+1+PRICAT:D:20A:UN'BGM+9'BGM+9'DTM+137:20100101:102'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"BGM"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about BGM exceeding its max repeat of 1", got)
	}
}
