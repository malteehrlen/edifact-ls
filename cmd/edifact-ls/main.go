// Command edifact-ls is a language server for UN/EDIFACT files, speaking LSP
// over stdio.
package main

import (
	"fmt"
	"os"

	"github.com/malteehrlen/edifact-ls/internal/lspserver"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-version", "-v":
			fmt.Println(lspserver.Version)
			return
		case "--help", "-help", "-h":
			fmt.Printf("%s speaks LSP over stdio; run it from an editor, not a terminal.\n", lspserver.Name)
			fmt.Println("Usage: edifact-ls [--version]")
			return
		}
	}

	srv := lspserver.New()

	if err := srv.RunStdio(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(srv.ExitCode())
}
