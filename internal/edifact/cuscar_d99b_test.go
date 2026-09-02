package edifact

import "testing"

func TestCUSCARD99BRegistered(t *testing.T) {
	if _, ok := schemaRegistry[MessageID{Type: "CUSCAR", Version: "D", Release: "99B", Agency: "UN"}]; !ok {
		t.Fatal("CUSCAR D.99B schema is not registered")
	}
}

func TestCUSCARD99BMinimalConformantMessage(t *testing.T) {
	// BGM is CUSCAR's only unconditionally-mandatory top-level node.
	// Segment group 14 (leading GID) is mandatory too, but only *within*
	// segment group 8 (leading RFF), itself only within segment group 7
	// (leading CNI) -- both conditional, so a bare BGM with none of that
	// entered at all is a genuinely clean, minimal message.
	src := "UNH+1+CUSCAR:D:99B:UN'BGM+785'UNT+2+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	if got := ValidateMessageSchemas(ic); len(got) != 0 {
		t.Fatalf("unexpected diagnostics for a minimal conformant CUSCAR message: %v", got)
	}
}

func TestCUSCARD99BMissingMandatoryGID(t *testing.T) {
	// Entering segment group 7 (CNI) then segment group 8 (RFF) makes
	// segment group 8's mandatory child, segment group 14 (leading GID),
	// apply -- omitting GID here should be flagged, unlike the bare-BGM
	// case above where SG7/SG8 are never entered at all.
	src := "UNH+1+CUSCAR:D:99B:UN'BGM+785'CNI+1'RFF+BM:X'UNT+4+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, "missing mandatory") || !containsMessage(got, `"GID"`) {
		t.Fatalf("got %v, want an error about the missing mandatory GID group", got)
	}
}

func TestCUSCARD99BExceededRepeat(t *testing.T) {
	// BGM itself is capped at 1 occurrence.
	src := "UNH+1+CUSCAR:D:99B:UN'BGM+785'BGM+785'UNT+3+1'"
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	got := ValidateMessageSchemas(ic)
	if len(got) != 1 || !containsMessage(got, `"BGM"`) || !containsMessage(got, "maximum of 1") {
		t.Fatalf("got %v, want an error about BGM exceeding its max repeat of 1", got)
	}
}
