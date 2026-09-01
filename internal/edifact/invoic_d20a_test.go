package edifact

import "testing"

func TestINVOICD20ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "INVOIC", Version: "D", Release: "20A", Agency: "UN"}]; !ok {
		t.Fatal("INVOIC D.20A schema is not registered")
	}
}

func TestINVOICD20AMinimalConformantMessage(t *testing.T) {
	// BGM, DTM, and UNS are mandatory top-level segments; segment group 52
	// (leading with MOA) is a mandatory *group* -- INVOIC is the first
	// message in this batch with a mandatory group at the top level, not
	// just mandatory segments.
	src := "UNH+1+INVOIC:D:20A:UN'BGM+380'DTM+137:20100101:102'UNS+D'MOA+1'UNT+5+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant INVOIC message: %v", got)
	}
}

func TestINVOICD20AMissingMandatoryDTM(t *testing.T) {
	src := "UNH+1+INVOIC:D:20A:UN'BGM+380'UNS+D'MOA+1'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"DTM"`) {
		t.Fatalf("got %v, want an error about missing mandatory DTM", got)
	}
}

func TestINVOICD20AExceededRepeat(t *testing.T) {
	// PAI is conditional but capped at 1 occurrence. Order matters here --
	// PAI must appear where the schema expects it (right after DTM, before
	// UNS and segment group 52), since the matcher is order-sensitive.
	src := "UNH+1+INVOIC:D:20A:UN'BGM+380'DTM+137:20100101:102'PAI'PAI'UNS+D'MOA+1'UNT+8+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"PAI"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about PAI exceeding its max repeat of 1", got)
	}
}
