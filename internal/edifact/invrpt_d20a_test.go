package edifact

import "testing"

func TestINVRPTD20ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "INVRPT", Version: "D", Release: "20A", Agency: "UN"}]; !ok {
		t.Fatal("INVRPT D.20A schema is not registered")
	}
}

func TestINVRPTD20AMinimalConformantMessage(t *testing.T) {
	// Unlike most messages here, INVRPT has *two* mandatory top-level
	// nodes: BGM and DTM (D.20A capped it C507's presence at exactly 10
	// occurrences, but at least one is mandatory).
	src := "UNH+1+INVRPT:D:20A:UN'BGM+23'DTM+137:20100101:102'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant INVRPT message: %v", got)
	}
}

func TestINVRPTD20AMissingMandatoryDTM(t *testing.T) {
	src := "UNH+1+INVRPT:D:20A:UN'BGM+23'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"DTM"`) {
		t.Fatalf("got %v, want an error about missing mandatory DTM", got)
	}
}

func TestINVRPTD20AExceededRepeat(t *testing.T) {
	// BGM itself is capped at 1 occurrence.
	src := "UNH+1+INVRPT:D:20A:UN'BGM+23'BGM+23'DTM+137:20100101:102'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"BGM"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about BGM exceeding its max repeat of 1", got)
	}
}
