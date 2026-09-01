package edifact

import "testing"

func TestIFTMIND20ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "IFTMIN", Version: "D", Release: "20A", Agency: "UN"}]; !ok {
		t.Fatal("IFTMIN D.20A schema is not registered")
	}
}

func TestIFTMIND20AMinimalConformantMessage(t *testing.T) {
	// BGM is mandatory, and so is segment group 12 (leading with NAD) --
	// unlike every other message in this batch, IFTMIN has a mandatory
	// *group*, not just mandatory top-level segments.
	src := "UNH+1+IFTMIN:D:20A:UN'BGM+610'NAD+1'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant IFTMIN message: %v", got)
	}
}

func TestIFTMIND20AMissingMandatoryBGM(t *testing.T) {
	src := "UNH+1+IFTMIN:D:20A:UN'NAD+1'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"BGM"`) {
		t.Fatalf("got %v, want an error about missing mandatory BGM", got)
	}
}

func TestIFTMIND20AExceededRepeat(t *testing.T) {
	// CTA is conditional but capped at 1 occurrence. Order matters here --
	// CTA must appear where the schema expects it (right after BGM, before
	// segment group 12's mandatory NAD), since the matcher is
	// order-sensitive.
	src := "UNH+1+IFTMIN:D:20A:UN'BGM+610'CTA'CTA'NAD+1'UNT+5+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"CTA"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about CTA exceeding its max repeat of 1", got)
	}
}
