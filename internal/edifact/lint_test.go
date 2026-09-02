package edifact

import (
	"strings"
	"testing"
)

func lint(t *testing.T, src string) ErrorList {
	t.Helper()
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected syntax errors parsing %q: %v", src, errs)
	}
	return Lint(ic, src)
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

func TestLintInfoOnRedundantUNAWithDotDecimal(t *testing.T) {
	// Dot decimal mark: the common convention since ISO 9735 version 4.
	src := "UNA:+.? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 1 {
		t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityInfo {
		t.Errorf("severity = %v, want info", errs[0].Severity)
	}
	if !containsMessage(errs, "safely omitted") || !containsMessage(errs, "version 4") {
		t.Errorf("message = %q, want it to mention it's safe to omit and cite version 4", errs[0].Message)
	}
}

func TestLintInfoOnRedundantUNAWithCommaDecimal(t *testing.T) {
	// Comma decimal mark: the ISO 9735 version 1-3 default. Component,
	// element, release, and terminator are still the (only structurally
	// significant) defaults, so this should still be flagged as redundant
	// -- just described as the version 1-3 convention, not version 4's.
	src := "UNA:+,? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 1 {
		t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityInfo {
		t.Errorf("severity = %v, want info", errs[0].Severity)
	}
	if !containsMessage(errs, "safely omitted") || !containsMessage(errs, "version 1-3") {
		t.Errorf("message = %q, want it to mention it's safe to omit and cite version 1-3", errs[0].Message)
	}
}

func TestLintInfoOnRedundantUNAWithUnusualDecimal(t *testing.T) {
	// An unusual decimal mark doesn't stop this from being functionally
	// redundant (component/element/release/terminator are still default);
	// the message should just describe it factually rather than claiming
	// it matches either named convention.
	src := "UNA:+#? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 1 {
		t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
	}
	if !containsMessage(errs, `"#"`) {
		t.Errorf("message = %q, want it to mention the actual decimal mark used", errs[0].Message)
	}
}

func TestLintNoInfoOnCustomUNA(t *testing.T) {
	src := "UNA^*,\\#!UNB^UNOA:1^S^R^201001:1200^1!" // custom delimiters, not default
	errs := lint(t, src)
	if len(errs) != 0 {
		t.Fatalf("unexpected lint diagnostics for a non-default UNA: %v", errs)
	}
}

func TestLintRedundantUNAHasFix(t *testing.T) {
	src := "UNA:+.? 'UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 1 {
		t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
	}
	e := errs[0]
	if e.Code != "redundant-una" {
		t.Errorf("Code = %q, want %q", e.Code, "redundant-una")
	}
	if e.Fix == nil {
		t.Fatalf("Fix is nil, want a fix that removes the UNA segment")
	}
	if e.Fix.OldText != "UNA:+.? '" {
		t.Errorf("Fix.OldText = %q, want the 9-byte UNA advice", e.Fix.OldText)
	}
	if e.Fix.NewText != "" {
		t.Errorf("Fix.NewText = %q, want empty (deletion)", e.Fix.NewText)
	}
	if e.Fix.Pos != (Position{Offset: 0, Line: 1, Column: 1}) {
		t.Errorf("Fix.Pos = %+v, want the start of the UNA segment", e.Fix.Pos)
	}
}

func TestLintRedundantUNAFixConsumesTrailingNewline(t *testing.T) {
	// Regression: a real interchange with each segment on its own line
	// (the common, human-readable convention -- see
	// skipInterSegmentWhitespace) used to have its Fix delete only the
	// 9-byte UNA advice, leaving a blank line behind. The fix must
	// consume the line terminator too, so applying it removes the whole
	// line.
	for _, tc := range []struct {
		name     string
		newline  string
		wantText string
	}{
		{"LF", "\n", "UNA:+.? '\n"},
		{"CRLF", "\r\n", "UNA:+.? '\r\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			src := "UNA:+.? '" + tc.newline + "UNB+UNOA:1+S+R+201001:1200+1'\nUNH+1+ORDERS:D:96A:UN'\nBGM+220'\nUNT+3+1'\nUNZ+1+1'"
			errs := lint(t, src)
			if len(errs) != 1 {
				t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
			}
			e := errs[0]
			if e.Fix.OldText != tc.wantText {
				t.Fatalf("Fix.OldText = %q, want %q (UNA plus its line terminator)", e.Fix.OldText, tc.wantText)
			}
			fixed := src[:e.Fix.Pos.Offset] + src[e.Fix.Pos.Offset+len(e.Fix.OldText):]
			if fixed[0] == '\n' || fixed[0] == '\r' {
				t.Fatalf("fixed text still starts with a blank line: %q", fixed)
			}
			if !strings.HasPrefix(fixed, "UNB+") {
				t.Fatalf("fixed text = %q, want it to start directly with UNB", fixed)
			}
		})
	}
}

func TestLintReservedTagWarningHasNoFix(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'UNX+1'UNT+4+1'UNZ+1+1'"
	errs := lint(t, src)
	if len(errs) != 1 {
		t.Fatalf("got %d lint diagnostics, want 1: %v", len(errs), errs)
	}
	if errs[0].Fix != nil {
		t.Errorf("Fix = %+v, want nil -- there's no safe rename to derive", errs[0].Fix)
	}
	if errs[0].Code != "" {
		t.Errorf("Code = %q, want empty for a non-fixable diagnostic", errs[0].Code)
	}
}
