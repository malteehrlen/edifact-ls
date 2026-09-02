// Command edifact-ls is a language server for UN/EDIFACT files, speaking LSP
// over stdio.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
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
			fmt.Println("Usage: edifact-ls [--version] [--help] | check <file> | schemas")
			return
		case "check":
			if len(os.Args) < 3 {
				fmt.Fprintln(os.Stderr, "usage: edifact-ls check <file>")
				os.Exit(2)
			}
			os.Exit(runCheck(os.Stdout, os.Args[2]))
		case "schemas":
			runSchemas(os.Stdout)
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

// runCheck parses and validates the EDIFACT file at path exactly as the
// LSP server would, printing every diagnostic to w (one per line, as
// "line:col: severity: message"). It returns 1 if any error-severity
// diagnostic was found (or the file couldn't be read), 0 otherwise -- for
// scripted/CI use independent of the editor.
func runCheck(w io.Writer, path string) int {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(w, err)
		return 1
	}

	_, errs := edifact.Validate(string(data))
	for _, e := range errs {
		fmt.Fprintln(w, e)
	}
	if errs.HasErrors() {
		return 1
	}
	return 0
}

// runSchemas prints every message specification structural validation is
// available for, one per line as "TYPE VERSION:RELEASE:AGENCY SOURCE" --
// the same edifact.ListRegisteredSchemas() data the generated
// docs/SUPPORTED_MESSAGES.md is built from (see tools/gendocs), so the two
// can never disagree about what's actually registered.
func runSchemas(w io.Writer) {
	for _, info := range edifact.ListRegisteredSchemas() {
		fmt.Fprintf(w, "%s %s:%s:%s %s\n", info.ID.Type, info.ID.Version, info.ID.Release, info.ID.Agency, info.Source)
	}
}
