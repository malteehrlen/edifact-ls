package edifact

import (
	"os"
	"testing"
	"time"
)

func TestParseValidInterchange(t *testing.T) {
	data, err := os.ReadFile("../../testdata/minimal.edi")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	ic, errs := Parse(string(data))
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if ic.UNA != nil {
		t.Fatalf("expected no UNA advice in fixture, got %+v", ic.UNA)
	}

	wantTags := []string{"UNB", "UNH", "BGM", "DTM", "UNT", "UNZ"}
	if len(ic.Segments) != len(wantTags) {
		t.Fatalf("got %d segments, want %d: %+v", len(ic.Segments), len(wantTags), ic.Segments)
	}
	for i, tag := range wantTags {
		if ic.Segments[i].Tag != tag {
			t.Errorf("segment %d tag = %q, want %q", i, ic.Segments[i].Tag, tag)
		}
	}

	bgm := ic.Segments[2]
	if got := bgm.Component0(0, ic.Delimiters); got != "220" {
		t.Errorf("BGM element 0 = %q, want %q", got, "220")
	}
	if got := bgm.Component0(1, ic.Delimiters); got != "ORDER123" {
		t.Errorf("BGM element 1 = %q, want %q", got, "ORDER123")
	}

	dtm := ic.Segments[3]
	if len(dtm.Elements) != 1 || len(dtm.Elements[0].Components) != 3 {
		t.Fatalf("DTM elements = %+v, want 1 element with 3 components", dtm.Elements)
	}
	if got := dtm.Elements[0].Components[1].Value(ic.Delimiters); got != "20100101" {
		t.Errorf("DTM date component = %q, want %q", got, "20100101")
	}
}

func TestParseWithCustomUNADelimiters(t *testing.T) {
	// Component separator '^', element separator '*', terminator '!'.
	src := "UNA^*,\\#!UNH*1*ORDERS^D^96A^UN!BGM*220!"
	ic, errs := Parse(src)
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if ic.UNA == nil || ic.UNA.Raw != "UNA^*,\\#!" {
		t.Fatalf("UNA = %+v, want Raw %q", ic.UNA, "UNA^*,\\#!")
	}
	if len(ic.Segments) != 2 {
		t.Fatalf("got %d segments, want 2: %+v", len(ic.Segments), ic.Segments)
	}
	unh := ic.Segments[0]
	if unh.Tag != "UNH" {
		t.Fatalf("segment 0 tag = %q, want UNH", unh.Tag)
	}
	if len(unh.Elements) != 2 || len(unh.Elements[1].Components) != 4 {
		t.Fatalf("UNH elements = %+v, want element 1 with 4 components", unh.Elements)
	}
	if got := unh.Elements[1].Components[0].Value(ic.Delimiters); got != "ORDERS" {
		t.Errorf("message type component = %q, want ORDERS", got)
	}
}

func TestParseEmptyElementsAndComponents(t *testing.T) {
	// A segment's elements are each *introduced* by a separator (grammar:
	// `tag (SEP element)*`), so "TAG++X'" is two elements -- an empty one
	// (nothing between the two separators) and "X" -- not three: the
	// separator immediately after the tag doesn't itself close an element.
	ic, errs := Parse("TAG++X'")
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	seg := ic.Segments[0]
	if len(seg.Elements) != 2 {
		t.Fatalf("got %d elements, want 2: %+v", len(seg.Elements), seg.Elements)
	}
	wantValues := [][]string{{""}, {"X"}}
	for i, el := range seg.Elements {
		if len(el.Components) != len(wantValues[i]) {
			t.Fatalf("element %d components = %+v, want %v", i, el.Components, wantValues[i])
		}
		for j, c := range el.Components {
			if c.Value(ic.Delimiters) != wantValues[i][j] {
				t.Errorf("element %d component %d = %q, want %q", i, j, c.Value(ic.Delimiters), wantValues[i][j])
			}
		}
	}
}

func TestParseSegmentWithNoOrOneEmptyElement(t *testing.T) {
	ic, errs := Parse("TAG'TAG+'")
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if len(ic.Segments) != 2 {
		t.Fatalf("got %d segments, want 2: %+v", len(ic.Segments), ic.Segments)
	}
	if got := ic.Segments[0].Elements; len(got) != 0 {
		t.Errorf("bare TAG' elements = %+v, want none", got)
	}
	if got := ic.Segments[1].Elements; len(got) != 1 || got[0].Components[0].Raw != "" {
		t.Errorf("TAG+' elements = %+v, want one empty element", got)
	}
}

func TestParseMissingTerminatorAtEOF(t *testing.T) {
	ic, errs := Parse("UNH+1+ORDERS'BGM+220+ORDER123")
	if len(ic.Segments) != 2 {
		t.Fatalf("got %d segments, want 2: %+v", len(ic.Segments), ic.Segments)
	}
	if len(errs) != 1 {
		t.Fatalf("got %d errors, want 1: %v", len(errs), errs)
	}
	if errs[0].Severity != SeverityError {
		t.Errorf("severity = %v, want error", errs[0].Severity)
	}
	// The unterminated BGM segment should still have been captured with
	// whatever elements it had, not dropped.
	bgm := ic.Segments[1]
	if bgm.Tag != "BGM" || bgm.Component0(1, ic.Delimiters) != "ORDER123" {
		t.Errorf("unterminated segment = %+v, want BGM with second element ORDER123", bgm)
	}
}

func TestParseInvalidSegmentTagRecovers(t *testing.T) {
	// "1BC" is not a valid segment tag; the parser should record an error
	// but resync at the following terminator and keep parsing UNZ correctly.
	ic, errs := Parse("UNH+1'1BC+garbage+more'UNZ+1+1'")
	if len(errs) != 1 {
		t.Fatalf("got %d errors, want 1: %v", len(errs), errs)
	}
	wantTags := []string{"UNH", "1BC", "UNZ"}
	if len(ic.Segments) != len(wantTags) {
		t.Fatalf("got %d segments, want %d: %+v", len(ic.Segments), len(wantTags), ic.Segments)
	}
	for i, tag := range wantTags {
		if ic.Segments[i].Tag != tag {
			t.Errorf("segment %d tag = %q, want %q", i, ic.Segments[i].Tag, tag)
		}
	}
	// Recovery should have discarded the malformed segment's body.
	if len(ic.Segments[1].Elements) != 0 {
		t.Errorf("recovered segment elements = %+v, want none", ic.Segments[1].Elements)
	}
	// And parsing of the following, well-formed segment should be unaffected.
	if got := ic.Segments[2].Component0(0, ic.Delimiters); got != "1" {
		t.Errorf("UNZ element 0 = %q, want %q", got, "1")
	}
}

func TestParseErrorPositionIsAccurate(t *testing.T) {
	// Two well-formed segments, then EOF mid-third-segment on line 3.
	src := "UNH+1'\nBGM+220'\nUNT+incomplete"
	_, errs := Parse(src)
	if len(errs) != 1 {
		t.Fatalf("got %d errors, want 1: %v", len(errs), errs)
	}
	if errs[0].Pos.Line != 3 {
		t.Errorf("error line = %d, want 3 (pos: %+v)", errs[0].Pos.Line, errs[0].Pos)
	}
}

func TestParseSegmentTagWithExplicitControlNumbers(t *testing.T) {
	ic, errs := Parse("GDS:1:2+data'")
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	seg := ic.Segments[0]
	if seg.Tag != "GDS" {
		t.Errorf("Tag = %q, want %q", seg.Tag, "GDS")
	}
	if want := []string{"1", "2"}; !slicesEqual(seg.TagControlNumbers, want) {
		t.Errorf("TagControlNumbers = %v, want %v", seg.TagControlNumbers, want)
	}
	if got := seg.Component0(0, ic.Delimiters); got != "data" {
		t.Errorf("element 0 = %q, want %q (control numbers must not leak into Elements)", got, "data")
	}
}

func TestParseSegmentTagWithoutControlNumbersUnaffected(t *testing.T) {
	// Implicit representation (the common case, no control numbers) must
	// parse exactly as before.
	ic, errs := Parse("GDS+data'")
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	seg := ic.Segments[0]
	if len(seg.TagControlNumbers) != 0 {
		t.Errorf("TagControlNumbers = %v, want none", seg.TagControlNumbers)
	}
	if got := seg.Component0(0, ic.Delimiters); got != "data" {
		t.Errorf("element 0 = %q, want %q", got, "data")
	}
}

func TestParseSegmentTagWithControlNumberButNoElements(t *testing.T) {
	// "GDS:1'" -- control number present, but the segment has no elements
	// at all (terminator immediately follows).
	ic, errs := Parse("GDS:1'")
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	seg := ic.Segments[0]
	if want := []string{"1"}; !slicesEqual(seg.TagControlNumbers, want) {
		t.Errorf("TagControlNumbers = %v, want %v", seg.TagControlNumbers, want)
	}
	if len(seg.Elements) != 0 {
		t.Errorf("Elements = %+v, want none", seg.Elements)
	}
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestParseDoesNotHangOnEmbeddedNewline is a regression test for a real
// infinite loop found on user-supplied input: a segment whose data was
// soft-wrapped across multiple physical lines (a literal '\n' embedded in
// the data, not just between segments -- common in hand-wrapped or
// legacy-gateway-exported EDIFACT, e.g. a long NAD address). The lexer's
// data-run loop used to stop at a mid-data newline without consuming it,
// so parseSegmentBody's next call saw the same byte again forever. Run
// with an explicit deadline so a reintroduced hang fails this test
// quickly instead of stalling the whole suite (or CI) indefinitely.
func TestParseDoesNotHangOnEmbeddedNewline(t *testing.T) {
	inputs := []string{
		"UNH+1'NAD+N1++AB\nCD'UNT+2+1'",
		"UNH+1'NAD+N1++AB\r\nCD'UNT+2+1'",
		// Multiple embedded newlines in one segment, and one right at the
		// very end just before the terminator.
		"UNH+1'FTX+AAA+++A\nB\nC\n'UNT+2+1'",
	}
	for _, src := range inputs {
		done := make(chan struct{})
		go func() {
			defer close(done)
			Parse(src)
		}()
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			t.Fatalf("Parse(%q) did not return within 2s (likely an infinite loop)", src)
		}
	}
}

func TestParseDoesNotPanic(t *testing.T) {
	inputs := []string{
		"",
		"UNA",
		"UNA:+.? '",
		"'''",
		"++++",
		"::::",
		"?",
		"UNH+1?",
		string([]byte{0x00, 0x01, 0x02, '\''}),
		"a very long segment tag that is definitely not three letters+data'",
	}
	for _, src := range inputs {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parse(%q) panicked: %v", src, r)
				}
			}()
			Parse(src)
		}()
	}
}
