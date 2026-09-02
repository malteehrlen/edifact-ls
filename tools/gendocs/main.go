// Command gendocs prints docs/SUPPORTED_MESSAGES.md's content to stdout,
// generated from the schemas actually registered in internal/edifact.
// Run via `make docs`, not directly -- see that target for where the
// output is written. Not part of the edifact-ls binary; dev tooling only.
package main

import (
	"fmt"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
)

func main() {
	fmt.Print(edifact.RenderSupportedMessagesDoc())
}
