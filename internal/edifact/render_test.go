package edifact

import (
	"os"
	"strings"
	"testing"
)

// interchangeDataEqual reports whether a and b carry the same logical data:
// same delimiters, same segment tags, and same component values (ignoring
// source position and any escaping-style differences in Raw text).
func interchangeDataEqual(t *testing.T, a, b *Interchange) bool {
	t.Helper()
	if a.Delimiters != b.Delimiters {
		t.Logf("delimiters differ: %+v vs %+v", a.Delimiters, b.Delimiters)
		return false
	}
	if len(a.Segments) != len(b.Segments) {
		t.Logf("segment count differs: %d vs %d", len(a.Segments), len(b.Segments))
		return false
	}
	for i := range a.Segments {
		sa, sb := a.Segments[i], b.Segments[i]
		if sa.Tag != sb.Tag {
			t.Logf("segment %d tag differs: %q vs %q", i, sa.Tag, sb.Tag)
			return false
		}
		if len(sa.TagControlNumbers) != len(sb.TagControlNumbers) {
			t.Logf("segment %d (%s) tag control numbers differ: %v vs %v", i, sa.Tag, sa.TagControlNumbers, sb.TagControlNumbers)
			return false
		}
		for k := range sa.TagControlNumbers {
			if sa.TagControlNumbers[k] != sb.TagControlNumbers[k] {
				t.Logf("segment %d (%s) tag control number %d differs: %q vs %q", i, sa.Tag, k, sa.TagControlNumbers[k], sb.TagControlNumbers[k])
				return false
			}
		}
		if len(sa.Elements) != len(sb.Elements) {
			t.Logf("segment %d (%s) element count differs: %d vs %d", i, sa.Tag, len(sa.Elements), len(sb.Elements))
			return false
		}
		for j := range sa.Elements {
			ea, eb := sa.Elements[j], sb.Elements[j]
			if len(ea.Components) != len(eb.Components) {
				t.Logf("segment %d element %d component count differs: %d vs %d", i, j, len(ea.Components), len(eb.Components))
				return false
			}
			for k := range ea.Components {
				va := ea.Components[k].Value(a.Delimiters)
				vb := eb.Components[k].Value(b.Delimiters)
				if va != vb {
					t.Logf("segment %d element %d component %d differs: %q vs %q", i, j, k, va, vb)
					return false
				}
			}
		}
	}
	return true
}

func mustParseClean(t *testing.T, src string) *Interchange {
	t.Helper()
	ic, errs := Parse(src)
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors for %q: %v", src, errs)
	}
	return ic
}

func TestRenderMultilineMinimalInterchange(t *testing.T) {
	data, err := os.ReadFile("../../testdata/minimal.edi")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	ic := mustParseClean(t, string(data))

	out := Render(ic, true)
	reparsed := mustParseClean(t, out)
	if !interchangeDataEqual(t, ic, reparsed) {
		t.Fatalf("Render(multiline) round-trip lost data; output:\n%s", out)
	}

	// The fixture is already in this exact style, so formatting it should
	// be a no-op.
	if out != string(data) {
		t.Errorf("Render(multiline) of an already-formatted fixture changed it:\ngot:\n%s\nwant:\n%s", out, string(data))
	}
}

func TestRenderPreservesTagControlNumbers(t *testing.T) {
	src := "GDS:1:2+data'"
	ic := mustParseClean(t, src)
	out := Render(ic, true)
	want := "GDS:1:2+data'\n"
	if out != want {
		t.Fatalf("Render(multiline) = %q, want %q", out, want)
	}

	reparsed := mustParseClean(t, out)
	if !interchangeDataEqual(t, ic, reparsed) {
		t.Fatalf("Render round-trip lost tag control numbers; output:\n%s", out)
	}
}

func TestRenderMultilineCompositeElements(t *testing.T) {
	src := "UNH+1+ORDERS:D:96A:UN'DTM+137:20100101:102'"
	ic := mustParseClean(t, src)
	out := Render(ic, true)
	want := "UNH+1+ORDERS:D:96A:UN'\nDTM+137:20100101:102'\n"
	if out != want {
		t.Fatalf("Render(multiline) = %q, want %q", out, want)
	}
}

func TestRenderIsIdempotent(t *testing.T) {
	data, err := os.ReadFile("../../testdata/minimal.edi")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	ic := mustParseClean(t, string(data))
	once := Render(ic, true)

	ic2 := mustParseClean(t, once)
	twice := Render(ic2, true)

	if once != twice {
		t.Fatalf("Render is not idempotent:\nonce:\n%s\ntwice:\n%s", once, twice)
	}
}

func TestRenderPreservesCustomUNADelimiters(t *testing.T) {
	src := "UNA^*,\\#!UNH*1*ORDERS^D^96A^UN!BGM*220!"
	ic := mustParseClean(t, src)
	out := Render(ic, true)

	reparsed, errs := Parse(out)
	if errs.HasErrors() {
		t.Fatalf("re-parsing rendered output failed: %v\noutput:\n%s", errs, out)
	}
	if reparsed.UNA == nil || reparsed.UNA.Raw != ic.UNA.Raw {
		t.Fatalf("UNA not preserved: got %+v, want Raw %q", reparsed.UNA, ic.UNA.Raw)
	}
	if !interchangeDataEqual(t, ic, reparsed) {
		t.Fatalf("Render round-trip lost data with custom delimiters; output:\n%s", out)
	}
}

// TestRenderRoundTripsAllCleanFixtures sweeps every checked-in testdata
// fixture that parses without syntax errors and asserts
// parse(Render(parse(x), multiline)) carries the same data as parse(x), in
// both the multiline and wire render modes. Fixtures with deliberate syntax
// errors (e.g. syntax-error.edi) are skipped: parser error-recovery
// discards the malformed segment's content, so re-rendering it is expected
// to lose data -- that's exactly why textDocument/formatting itself refuses
// to touch unparseable documents (see formatting.go).
func TestRenderRoundTripsAllCleanFixtures(t *testing.T) {
	entries, err := os.ReadDir("../../testdata")
	if err != nil {
		t.Fatalf("reading testdata: %v", err)
	}

	tested := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := "../../testdata/" + entry.Name()
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("reading %s: %v", path, err)
		}

		ic, errs := Parse(string(data))
		if errs.HasErrors() {
			t.Logf("skipping %s: has syntax errors by design", entry.Name())
			continue
		}
		tested++

		for _, multiline := range []bool{true, false} {
			out := Render(ic, multiline)
			reparsed, errs := Parse(out)
			if errs.HasErrors() {
				t.Fatalf("%s: re-parsing Render(multiline=%v) output failed: %v\noutput:\n%s", entry.Name(), multiline, errs, out)
			}
			if !interchangeDataEqual(t, ic, reparsed) {
				t.Fatalf("%s: Render(multiline=%v) round-trip lost data; output:\n%s", entry.Name(), multiline, out)
			}
		}
	}

	if tested == 0 {
		t.Fatal("no clean fixtures found under testdata/ to round-trip test")
	}
}

func TestRenderWireFormatIsSingleLine(t *testing.T) {
	data, err := os.ReadFile("../../testdata/minimal.edi")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	ic := mustParseClean(t, string(data))

	wire := Render(ic, false)
	if strings.Contains(wire, "\n") {
		t.Fatalf("Render(wire) contains a newline: %q", wire)
	}

	reparsed := mustParseClean(t, wire)
	if !interchangeDataEqual(t, ic, reparsed) {
		t.Fatalf("Render(wire) round-trip lost data; output:\n%s", wire)
	}
}
