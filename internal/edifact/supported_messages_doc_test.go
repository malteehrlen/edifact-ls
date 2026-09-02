package edifact

import (
	"os"
	"testing"
)

// TestSupportedMessagesDocIsUpToDate keeps docs/SUPPORTED_MESSAGES.md
// structurally unable to drift from the registry it's generated from:
// this compares the checked-in file against RenderSupportedMessagesDoc()
// directly, the same function `make docs` uses to write it, so the two
// can never silently disagree.
func TestSupportedMessagesDocIsUpToDate(t *testing.T) {
	want := RenderSupportedMessagesDoc()

	got, err := os.ReadFile("../../docs/SUPPORTED_MESSAGES.md")
	if err != nil {
		t.Fatalf("reading docs/SUPPORTED_MESSAGES.md: %v (run `make docs` to generate it)", err)
	}

	if string(got) != want {
		t.Fatalf("docs/SUPPORTED_MESSAGES.md is out of date -- run `make docs` to regenerate it and commit the result")
	}
}
