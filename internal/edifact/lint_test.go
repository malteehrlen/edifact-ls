package edifact

import "testing"

func lint(t *testing.T, src string) ErrorList {
	t.Helper()
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected syntax errors parsing %q: %v", src, errs)
	}
	return Lint(ic)
}

func TestLintCleanInterchangeHasNoWarnings(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 0 {
		t.Fatalf("unexpected lint diagnostics: %v", errs)
	}
}

func TestLintWarnsOnReservedUNPrefix(t *testing.T) {
	// UNX is not a recognized service segment, but starts with "UN".
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'UNX+1'UNT+4+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 1 {
		t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityWarning {
		t.Errorf("severity = %v, want warning", errs[0].Severity)
	}
	if !containsMessage(errs, `"UNX"`) || !containsMessage(errs, "reserved") {
		t.Errorf("message = %q, want it to mention UNX and the reserved prefix", errs[0].Message)
	}
}

func TestLintDoesNotWarnOnRecognizedServiceSegments(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'" +
		"UNG+INVOIC+15623+23457+201001:1200+G1+UN+96A:1'" +
		"UNH+1+ORDERS:D:96A:UN'BGM+220'UNS+D'UNT+4+1'" +
		"UNE+1+G1'" +
		"UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 0 {
		t.Fatalf("unexpected lint diagnostics for known service segments: %v", errs)
	}
}

func TestLintInfoOnRedundantDefaultUNA(t *testing.T) {
	src := "UNA:+.? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 1 {
		t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityInfo {
		t.Errorf("severity = %v, want info", errs[0].Severity)
	}
	if !containsMessage(errs, "default delimiters") {
		t.Errorf("message = %q, want it to mention default delimiters", errs[0].Message)
	}
}

func TestLintNoInfoOnCustomUNA(t *testing.T) {
	src := "UNA^*,\\#!UNB^UNOA:1^S^R^201001:1200^1!" // custom delimiters, not default
	errs := lint(t, src)
	if len(errs) != 0 {
		t.Fatalf("unexpected lint diagnostics for a non-default UNA: %v", errs)
	}
}
