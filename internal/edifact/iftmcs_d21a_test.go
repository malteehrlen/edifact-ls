package edifact

import "testing"

// These exercise the real, registered IFTMCS D.21A schema (iftmcs_d21a.go)
// end to end, as opposed to schema_registry_test.go's hand-built
// stand-ins. Every one of IFTMCS's 41 segment groups is itself
// conditional, and each group's sole mandatory child is always its own
// leading (group-detecting) segment -- so a "missing mandatory group" or
// "missing mandatory child within a started group" violation can never
// actually occur against this schema: recognizing a group's presence at
// all requires its leading, mandatory child to already be there. BGM is
// the only mandatory node anywhere in the schema that isn't a group's own
// leading child, so it's the one real "missing mandatory" scenario;
// exceeding a repeat cap (e.g. CTA's) is the other kind of real
// violation. The engine's general handling of missing-mandatory-group and
// missing-mandatory-nested-child is already covered generically by
// schema_test.go's hand-built schemas.

func TestIFTMCSD21ARegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "IFTMCS", Version: "D", Release: "21A", Agency: "UN"}]; !ok {
		t.Fatal("IFTMCS D.21A schema is not registered")
	}
}

func TestIFTMCSD21AMinimalConformantMessage(t *testing.T) {
	src := "UNH+1+IFTMCS:D:21A:UN'BGM+30'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant IFTMCS message: %v", got)
	}
}

func TestIFTMCSD21AMissingMandatoryBGM(t *testing.T) {
	src := "UNH+1+IFTMCS:D:21A:UN'CTA+IC'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(got), got)
	}
	if got[0].Severity != SeverityError || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"BGM"`) {
		t.Errorf("diagnostic = %v, want an error about missing mandatory BGM", got[0])
	}
}

func TestIFTMCSD21AExceededRepeat(t *testing.T) {
	// CTA is conditional but capped at 1 occurrence.
	src := "UNH+1+IFTMCS:D:21A:UN'BGM+30'CTA+IC'CTA+IC'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(got), got)
	}
	if got[0].Severity != SeverityError || !containsMessage(got, `"CTA"`) || !containsMessage(got, "maximum of 1") {
		t.Errorf("diagnostic = %v, want an error about CTA exceeding its max repeat of 1", got[0])
	}
}

// TestIFTMCSD21AViolationCitesMessageContext confirms every structural
// violation names which message type/version it was validated against
// and which UNH declared it -- important once an interchange has more
// than one message, so it's clear which UNH a violation is scoped to.
func TestIFTMCSD21AViolationCitesMessageContext(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'\n" +
		"UNH+1+IFTMCS:D:21A:UN'\n" +
		"BGM+30'\n" +
		"CTA+IC'\n" +
		"CTA+IC'\n" +
		"UNT+5+1'\n" +
		"UNZ+1+1'\n"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}

	got := ValidateMessageSchemas(ic)
	if len(got) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %v", len(got), got)
	}
	want := "as specified by message type IFTMCS D:21A:UN on line 2"
	if !containsMessage(got, want) {
		t.Errorf("message = %q, want it to contain %q (UNH is on line 2)", got[0].Message, want)
	}
}
