package edifact

import "testing"

func TestORDRSPD99BRegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "ORDRSP", Version: "D", Release: "99B", Agency: "UN"}]; !ok {
		t.Fatal("ORDRSP D.99B schema is not registered")
	}
}

func TestORDRSPD99BMinimalConformantMessage(t *testing.T) {
	src := "UNH+1+ORDRSP:D:99B:UN'BGM+231'DTM+137:20100101:102'UNS+D'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant ORDRSP D.99B message: %v", got)
	}
}

func TestORDRSPD99BMissingMandatoryDTM(t *testing.T) {
	src := "UNH+1+ORDRSP:D:99B:UN'BGM+231'UNS+D'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"DTM"`) {
		t.Fatalf("got %v, want an error about missing mandatory DTM", got)
	}
}

func TestORDRSPD99BExceededRepeat(t *testing.T) {
	src := "UNH+1+ORDRSP:D:99B:UN'BGM+231'BGM+231'DTM+137:20100101:102'UNS+D'UNT+5+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"BGM"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about BGM exceeding its max repeat of 1", got)
	}
}
