package edifact

import "testing"

func TestORDRSPD20ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "ORDRSP", Version: "D", Release: "20A", Agency: "UN"}]; !ok {
		t.Fatal("ORDRSP D.20A schema is not registered")
	}
}

func TestORDRSPD20AMinimalConformantMessage(t *testing.T) {
	// BGM, DTM, and UNS are all mandatory at the top level.
	src := "UNH+1+ORDRSP:D:20A:UN'BGM+231'DTM+137:20100101:102'UNS+D'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant ORDRSP message: %v", got)
	}
}

func TestORDRSPD20AMissingMandatoryDTM(t *testing.T) {
	src := "UNH+1+ORDRSP:D:20A:UN'BGM+231'UNS+D'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"DTM"`) {
		t.Fatalf("got %v, want an error about missing mandatory DTM", got)
	}
}

func TestORDRSPD20AExceededRepeat(t *testing.T) {
	// PAI is conditional but capped at 1 occurrence. Order matters here --
	// PAI must appear where the schema expects it (right after DTM, before
	// UNS), since the matcher is order-sensitive.
	src := "UNH+1+ORDRSP:D:20A:UN'BGM+231'DTM+137:20100101:102'PAI'PAI'UNS+D'UNT+7+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"PAI"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about PAI exceeding its max repeat of 1", got)
	}
}
