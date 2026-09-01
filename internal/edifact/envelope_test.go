package edifact

import (
	"os"
	"strings"
	"testing"
)

func validate(t *testing.T, src string) ErrorList {
	t.Helper()
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected syntax errors parsing %q: %v", src, errs)
	}
	return ValidateEnvelopes(ic)
}

func TestValidateEnvelopesValidInterchange(t *testing.T) {
	data, err := os.ReadFile("../../testdata/minimal.edi")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	errs := validate(t, string(data))
	if len(errs) != 0 {
		t.Fatalf("unexpected envelope errors: %v", errs)
	}
}

func TestValidateEnvelopesMissingUNZ(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "missing UNZ") {
		t.Fatalf("expected a missing-UNZ error, got: %v", errs)
	}
}

func TestValidateEnvelopesMissingUNT(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "missing its UNT") {
		t.Fatalf("expected a missing-UNT error, got: %v", errs)
	}
}

func TestValidateEnvelopesMismatchedInterchangeControlCount(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+2+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "interchange control count") {
		t.Fatalf("expected an interchange-control-count error, got: %v", errs)
	}
}

func TestValidateEnvelopesMismatchedInterchangeControlReference(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+999'"
	errs := validate(t, src)
	if !containsMessage(errs, "control reference") {
		t.Fatalf("expected a control-reference-mismatch error, got: %v", errs)
	}
}

func TestValidateEnvelopesMismatchedMessageSegmentCount(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+99+1'UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "UNT segment count") {
		t.Fatalf("expected a segment-count-mismatch error, got: %v", errs)
	}
}

func TestValidateEnvelopesMismatchedMessageReference(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+999'UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "UNT message reference") {
		t.Fatalf("expected a message-reference-mismatch error, got: %v", errs)
	}
}

func TestValidateEnvelopesMultipleMessages(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNH+2+ORDERS:D:96A:UN'BGM+221'UNT+3+2'" +
		"UNZ+2+1'"
	errs := validate(t, src)
	if len(errs) != 0 {
		t.Fatalf("unexpected envelope errors: %v", errs)
	}
}

func containsMessage(errs ErrorList, substr string) bool {
	for _, e := range errs {
		if strings.Contains(e.Message, substr) {
			return true
		}
	}
	return false
}
