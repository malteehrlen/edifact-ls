package edifact

import "testing"

func collectTokens(l *lexer) []token {
	var toks []token
	for {
		tok := l.next()
		toks = append(toks, tok)
		if tok.Kind == tokenEOF {
			return toks
		}
	}
}

func TestLexerDefaultDelimiters(t *testing.T) {
	var errs ErrorList
	l := newLexer("BGM+220+ORDER123+9'", DefaultDelimiters(), &errs)
	toks := collectTokens(l)

	want := []tokenKind{
		tokenData, tokenElementSep, tokenData, tokenElementSep,
		tokenData, tokenElementSep, tokenData, tokenSegmentTerminator, tokenEOF,
	}
	if len(toks) != len(want) {
		t.Fatalf("got %d tokens, want %d: %+v", len(toks), len(want), toks)
	}
	for i, k := range want {
		if toks[i].Kind != k {
			t.Errorf("token %d: kind = %v, want %v", i, toks[i].Kind, k)
		}
	}
	if toks[0].Raw != "BGM" {
		t.Errorf("tag token Raw = %q, want %q", toks[0].Raw, "BGM")
	}
	if len(errs) != 0 {
		t.Errorf("unexpected errors: %v", errs)
	}
}

func TestDetectUNACustomDelimiters(t *testing.T) {
	src := "UNA:+.? 'UNH+1'"
	delims, consumed, present := detectUNA(src)
	if !present {
		t.Fatal("expected UNA to be detected")
	}
	if consumed != 9 {
		t.Fatalf("consumed = %d, want 9", consumed)
	}
	want := Delimiters{Component: ':', Element: '+', Decimal: '.', Release: '?', Reserved: ' ', Terminator: '\''}
	if delims != want {
		t.Fatalf("delims = %+v, want %+v", delims, want)
	}
}

func TestDetectUNAWithNonDefaultCharacters(t *testing.T) {
	// A UNA that redefines every delimiter to something unusual: component
	// separator '^', element separator '*', decimal ',', release '\',
	// reserved '#', terminator '!'.
	src := "UNA^*,\\#!" + "UNH*1^A!"
	delims, consumed, present := detectUNA(src)
	if !present || consumed != 9 {
		t.Fatalf("detectUNA(%q) = %+v, %d, %v", src, delims, consumed, present)
	}
	want := Delimiters{Component: '^', Element: '*', Decimal: ',', Release: '\\', Reserved: '#', Terminator: '!'}
	if delims != want {
		t.Fatalf("delims = %+v, want %+v", delims, want)
	}

	var errs ErrorList
	l := newLexer(src, delims, &errs)
	l.skip(consumed)
	toks := collectTokens(l)
	wantKinds := []tokenKind{
		tokenData, tokenElementSep, tokenData, tokenComponentSep,
		tokenData, tokenSegmentTerminator, tokenEOF,
	}
	if len(toks) != len(wantKinds) {
		t.Fatalf("got %d tokens, want %d: %+v", len(toks), len(wantKinds), toks)
	}
	for i, k := range wantKinds {
		if toks[i].Kind != k {
			t.Errorf("token %d: kind = %v, want %v (%+v)", i, toks[i].Kind, k, toks[i])
		}
	}
}

func TestDetectUNAAbsentUsesDefaults(t *testing.T) {
	delims, consumed, present := detectUNA("UNB+UNOA:1+...")
	if present || consumed != 0 {
		t.Fatalf("expected no UNA detected, got consumed=%d present=%v", consumed, present)
	}
	if delims != DefaultDelimiters() {
		t.Fatalf("delims = %+v, want defaults", delims)
	}
}

func TestLexerEscapedDelimiterCharacters(t *testing.T) {
	var errs ErrorList
	// "12?+34" contains a release-escaped '+' that must NOT be treated as an
	// element separator; "56??78" contains an escaped release character
	// itself, which also must not split the run.
	l := newLexer("12?+34+56??78'", DefaultDelimiters(), &errs)
	toks := collectTokens(l)

	want := []tokenKind{tokenData, tokenElementSep, tokenData, tokenSegmentTerminator, tokenEOF}
	if len(toks) != len(want) {
		t.Fatalf("got %d tokens, want %d: %+v", len(toks), len(want), toks)
	}
	if toks[0].Raw != "12?+34" {
		t.Errorf("toks[0].Raw = %q, want %q", toks[0].Raw, "12?+34")
	}
	if toks[2].Raw != "56??78" {
		t.Errorf("toks[2].Raw = %q, want %q", toks[2].Raw, "56??78")
	}

	got := Unescape(toks[0].Raw, DefaultDelimiters().Release)
	if got != "12+34" {
		t.Errorf("Unescape(%q) = %q, want %q", toks[0].Raw, got, "12+34")
	}
	got2 := Unescape(toks[2].Raw, DefaultDelimiters().Release)
	if got2 != "56?78" {
		t.Errorf("Unescape(%q) = %q, want %q", toks[2].Raw, got2, "56?78")
	}
	if len(errs) != 0 {
		t.Errorf("unexpected errors: %v", errs)
	}
}

func TestLexerDanglingReleaseCharacterAtEOF(t *testing.T) {
	var errs ErrorList
	l := newLexer("ABC?", DefaultDelimiters(), &errs)
	toks := collectTokens(l)
	if len(toks) != 2 || toks[0].Kind != tokenData || toks[1].Kind != tokenEOF {
		t.Fatalf("unexpected tokens: %+v", toks)
	}
	if len(errs) != 1 {
		t.Fatalf("expected 1 error, got %d: %v", len(errs), errs)
	}
}

func TestLexerEmptyComponentsAndElements(t *testing.T) {
	// "A::B" -> data "A", componentSep, componentSep (empty component
	// between them), data "B" -- i.e. an empty component is represented by
	// two adjacent separators with no data token between them.
	var errs ErrorList
	l := newLexer("A::B++C'", DefaultDelimiters(), &errs)
	toks := collectTokens(l)

	want := []tokenKind{
		tokenData, tokenComponentSep, tokenComponentSep, tokenData,
		tokenElementSep, tokenElementSep, tokenData, tokenSegmentTerminator, tokenEOF,
	}
	if len(toks) != len(want) {
		t.Fatalf("got %d tokens, want %d: %+v", len(toks), len(want), toks)
	}
	for i, k := range want {
		if toks[i].Kind != k {
			t.Errorf("token %d: kind = %v, want %v", i, toks[i].Kind, k)
		}
	}
}

func TestLexerSkipsInterSegmentNewlines(t *testing.T) {
	var errs ErrorList
	l := newLexer("A'\r\nB'", DefaultDelimiters(), &errs)
	first := l.next() // "A"
	_ = first
	term := l.next() // '
	if term.Kind != tokenSegmentTerminator {
		t.Fatalf("expected terminator, got %v", term.Kind)
	}
	l.skipInterSegmentWhitespace()
	next := l.next()
	if next.Kind != tokenData || next.Raw != "B" {
		t.Fatalf("expected data 'B' after whitespace skip, got %+v", next)
	}
}
