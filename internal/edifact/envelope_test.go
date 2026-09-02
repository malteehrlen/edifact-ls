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

func TestValidateEnvelopesFunctionalGroupsValid(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNH+2+ORDERS:D:96A:UN'BGM+221'UNT+3+2'" +
		"UNE+2+G1'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	if len(errs) != 0 {
		t.Fatalf("unexpected envelope errors: %v", errs)
	}
}

func TestValidateEnvelopesUNZCountsGroupsNotMessages(t *testing.T) {
	// One group containing 2 messages: UNZ's count must be 1 (the number
	// of groups), not 2 (the number of messages).
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNH+2+ORDERS:D:96A:UN'BGM+221'UNT+3+2'" +
		"UNE+2+G1'" +
		"UNZ+2+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "number of functional groups") {
		t.Fatalf("expected a group-count error, got: %v", errs)
	}
}

func TestValidateEnvelopesFunctionalGroupMismatchedMessageCount(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNE+99+G1'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "UNE number of messages") {
		t.Fatalf("expected a UNE message-count error, got: %v", errs)
	}
}

func TestValidateEnvelopesFunctionalGroupMismatchedReference(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNE+1+WRONG'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "UNE functional group reference") {
		t.Fatalf("expected a UNE reference-mismatch error, got: %v", errs)
	}
}

func TestValidateEnvelopesFunctionalGroupMissingUNE(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "missing its UNE trailer") {
		t.Fatalf("expected a missing-UNE error, got: %v", errs)
	}
}

func TestValidateEnvelopesUNEWithoutUNG(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNE+1+G1'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "UNE functional group trailer found without a preceding UNG") {
		t.Fatalf("expected a stray-UNE error, got: %v", errs)
	}
}

func TestValidateEnvelopesMixedGroupedAndUngroupedMessages(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNE+1+G1'" +
		"UNH+2+ORDERS:D:96A:UN'BGM+221'UNT+3+2'" +
		"UNZ+2+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "mixes messages inside functional groups") {
		t.Fatalf("expected a mixed-grouping error, got: %v", errs)
	}
}

func TestValidateEnvelopesUNSValidValues(t *testing.T) {
	for _, val := range []string{"D", "S"} {
		src := "UNB+UNOA:1+S+R+201001:1200+1'" +
			"UNH+1+ORDERS:D:96A:UN'BGM+220'UNS+" + val + "'UNT+4+1'" +
			"UNZ+1+1'"
		errs := validate(t, src)
		if len(errs) != 0 {
			t.Fatalf("UNS+%s: unexpected envelope errors: %v", val, errs)
		}
	}
}

func TestValidateEnvelopesUNSInvalidValue(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNS+X'UNT+4+1'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	if !containsMessage(errs, "UNS section control value") {
		t.Fatalf("expected a UNS-value error, got: %v", errs)
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

// errorContaining returns the first Error whose Message contains substr,
// or nil.
func errorContaining(errs ErrorList, substr string) *Error {
	for _, e := range errs {
		if strings.Contains(e.Message, substr) {
			return e
		}
	}
	return nil
}

// applyFix splices f into src, the same way a real client would apply the
// WorkspaceEdit derived from it.
func applyFix(src string, f *Fix) string {
	return src[:f.Pos.Offset] + f.NewText + src[f.Pos.Offset+len(f.OldText):]
}

// assertFixResolves applies e's Fix to src and asserts re-validating the
// result no longer contains wantGoneSubstr -- i.e. the fix actually fixes
// the problem, not just that some Fix is present.
func assertFixResolves(t *testing.T, src string, e *Error, wantGoneSubstr string) {
	t.Helper()
	if e == nil {
		t.Fatalf("no diagnostic found containing %q", wantGoneSubstr)
	}
	if e.Fix == nil {
		t.Fatalf("diagnostic %q has no Fix", e.Message)
	}
	fixed := applyFix(src, e.Fix)
	errs := validate(t, fixed)
	if containsMessage(errs, wantGoneSubstr) {
		t.Fatalf("applying Fix did not resolve the diagnostic; still present after fix: %v (fixed source: %q)", errs, fixed)
	}
}

func TestValidateEnvelopesMismatchedInterchangeControlCountHasFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+2+1'"
	errs := validate(t, src)
	e := errorContaining(errs, "interchange control count")
	if e.Code != "envelope-count-mismatch" {
		t.Errorf("Code = %q, want %q", e.Code, "envelope-count-mismatch")
	}
	assertFixResolves(t, src, e, "interchange control count")
}

func TestValidateEnvelopesMismatchedInterchangeControlReferenceHasFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+999'"
	errs := validate(t, src)
	e := errorContaining(errs, "control reference")
	if e.Code != "envelope-reference-mismatch" {
		t.Errorf("Code = %q, want %q", e.Code, "envelope-reference-mismatch")
	}
	assertFixResolves(t, src, e, "control reference")
}

func TestValidateEnvelopesMismatchedMessageSegmentCountHasFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+99+1'UNZ+1+1'"
	errs := validate(t, src)
	e := errorContaining(errs, "UNT segment count")
	if e.Code != "envelope-count-mismatch" {
		t.Errorf("Code = %q, want %q", e.Code, "envelope-count-mismatch")
	}
	assertFixResolves(t, src, e, "UNT segment count")
}

func TestValidateEnvelopesMismatchedMessageReferenceHasFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+999'UNZ+1+1'"
	errs := validate(t, src)
	e := errorContaining(errs, "UNT message reference")
	if e.Code != "envelope-reference-mismatch" {
		t.Errorf("Code = %q, want %q", e.Code, "envelope-reference-mismatch")
	}
	assertFixResolves(t, src, e, "UNT message reference")
}

func TestValidateEnvelopesFunctionalGroupMismatchedMessageCountHasFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNE+99+G1'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	e := errorContaining(errs, "UNE number of messages")
	if e.Code != "envelope-count-mismatch" {
		t.Errorf("Code = %q, want %q", e.Code, "envelope-count-mismatch")
	}
	assertFixResolves(t, src, e, "UNE number of messages")
}

func TestValidateEnvelopesFunctionalGroupMismatchedReferenceHasFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'" +
		"UNE+1+G999'" +
		"UNZ+1+1'"
	errs := validate(t, src)
	e := errorContaining(errs, "functional group reference")
	if e.Code != "envelope-reference-mismatch" {
		t.Errorf("Code = %q, want %q", e.Code, "envelope-reference-mismatch")
	}
	assertFixResolves(t, src, e, "functional group reference")
}

func TestValidateEnvelopesMissingUNTHasNoFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNZ+1+1'"
	errs := validate(t, src)
	e := errorContaining(errs, "missing its UNT")
	if e.Fix != nil {
		t.Errorf("Fix = %+v, want nil -- there's nothing to safely insert", e.Fix)
	}
	if e.Code != "" {
		t.Errorf("Code = %q, want empty for a non-fixable diagnostic", e.Code)
	}
}
