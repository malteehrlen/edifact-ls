// Command edifact-ls is a language server for UN/EDIFACT files, speaking LSP
// over stdio.
package main

import (
	"fmt"
	"os"

	"github.com/malteehrlen/edifact-ls/internal/lspserver"
)

func main() {
	srv := lspserver.New()

	if err := srv.RunStdio(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(srv.ExitCode())
}
