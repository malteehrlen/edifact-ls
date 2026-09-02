package edifact

import "testing"

func TestORDERSD99BRegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "ORDERS", Version: "D", Release: "99B", Agency: "UN"}]; !ok {
		t.Fatal("ORDERS D.99B schema is not registered")
	}
}

func TestORDERSD99BMinimalConformantMessage(t *testing.T) {
	// BGM, DTM, and UNS are all mandatory at the top level -- same shape
	// as ORDERS D.20A.
	src := "UNH+1+ORDERS:D:99B:UN'BGM+220'DTM+137:20100101:102'UNS+D'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant ORDERS D.99B message: %v", got)
	}
}

func TestORDERSD99BMissingMandatoryDTM(t *testing.T) {
	src := "UNH+1+ORDERS:D:99B:UN'BGM+220'UNS+D'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"DTM"`) {
		t.Fatalf("got %v, want an error about missing mandatory DTM", got)
	}
}

func TestORDERSD99BExceededRepeat(t *testing.T) {
	// BGM itself is capped at 1 occurrence.
	src := "UNH+1+ORDERS:D:99B:UN'BGM+220'BGM+220'DTM+137:20100101:102'UNS+D'UNT+5+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"BGM"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about BGM exceeding its max repeat of 1", got)
	}
}

// TestORDERSBothReleasesIndependentlyCorrect confirms registering a
// second release of ORDERS (D.99B) didn't disturb the existing D.20A
// registration -- the exact scenario this whole epic is about: multiple
// releases of the same "popular" message type coexisting correctly.
func TestORDERSBothReleasesIndependentlyCorrect(t *testing.T) {
	d20a := "UNH+1+ORDERS:D:20A:UN'BGM+220'DTM+137:20100101:102'UNS+D'UNT+4+1'"
	ic, errs := Parse(d20a)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics validating ORDERS D.20A: %v", got)
	}

	d99b := "UNH+1+ORDERS:D:99B:UN'BGM+220'DTM+137:20100101:102'UNS+D'UNT+4+1'"
	ic, errs = Parse(d99b)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics validating ORDERS D.99B: %v", got)
	}
}
