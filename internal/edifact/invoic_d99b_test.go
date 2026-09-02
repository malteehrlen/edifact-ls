package edifact

import "testing"

func TestINVOICD99BRegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "INVOIC", Version: "D", Release: "99B", Agency: "UN"}]; !ok {
		t.Fatal("INVOIC D.99B schema is not registered")
	}
}

func TestINVOICD99BMinimalConformantMessage(t *testing.T) {
	// BGM, DTM, and UNS are mandatory top-level segments; segment group
	// 50 (leading MOA) is a mandatory group -- same shape as INVOIC D.20A.
	src := "UNH+1+INVOIC:D:99B:UN'BGM+380'DTM+137:20100101:102'UNS+D'MOA+1'UNT+5+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant INVOIC D.99B message: %v", got)
	}
}

func TestINVOICD99BMissingMandatoryDTM(t *testing.T) {
	src := "UNH+1+INVOIC:D:99B:UN'BGM+380'UNS+D'MOA+1'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"DTM"`) {
		t.Fatalf("got %v, want an error about missing mandatory DTM", got)
	}
}

func TestINVOICD99BExceededRepeat(t *testing.T) {
	// PAI is conditional but capped at 1 occurrence. Order matters here --
	// PAI must appear where the schema expects it (right after DTM, before
	// UNS and segment group 50), since the matcher is order-sensitive.
	src := "UNH+1+INVOIC:D:99B:UN'BGM+380'DTM+137:20100101:102'PAI'PAI'UNS+D'MOA+1'UNT+7+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"PAI"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about PAI exceeding its max repeat of 1", got)
	}
}
