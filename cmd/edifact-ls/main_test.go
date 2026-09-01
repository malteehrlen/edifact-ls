package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunCheckCleanFileReturnsZero(t *testing.T) {
	var buf bytes.Buffer
	code := runCheck(&buf, "../../testdata/minimal.edi")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; output: %s", code, buf.String())
	}
	if buf.Len() != 0 {
		t.Errorf("expected no output for a clean file, got: %s", buf.String())
	}
}

func TestRunCheckViolationReturnsOne(t *testing.T) {
	var buf bytes.Buffer
	code := runCheck(&buf, "../../testdata/iftmcs-violation.edi")
	if code != 1 {
		t.Fatalf("exit code = %d, want 1; output: %s", code, buf.String())
	}
	if !strings.Contains(buf.String(), "maximum of 1") {
		t.Errorf("output = %q, want it to mention the CTA repeat violation", buf.String())
	}
}

func TestRunCheckMissingFileReturnsOne(t *testing.T) {
	var buf bytes.Buffer
	code := runCheck(&buf, "../../testdata/does-not-exist.edi")
	if code != 1 {
		t.Fatalf("exit code = %d, want 1 for a missing file", code)
	}
}
