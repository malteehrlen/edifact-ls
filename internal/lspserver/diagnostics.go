package lspserver

import (
	"unicode/utf16"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (st *state) textDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	st.setDocument(params.TextDocument.URI, params.TextDocument.Text)
	st.publishDiagnostics(context, params.TextDocument.URI)
	return nil
}

func (st *state) textDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	// We advertise TextDocumentSyncKindFull, so the client sends the whole
	// document text on every change; take the last change event as the new
	// full content.
	for _, raw := range params.ContentChanges {
		switch change := raw.(type) {
		case protocol.TextDocumentContentChangeEventWhole:
			st.setDocument(params.TextDocument.URI, change.Text)
		case protocol.TextDocumentContentChangeEvent:
			st.setDocument(params.TextDocument.URI, change.Text)
		}
	}
	st.publishDiagnostics(context, params.TextDocument.URI)
	return nil
}

func (st *state) textDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	st.docsMu.Lock()
	delete(st.documents, params.TextDocument.URI)
	st.docsMu.Unlock()

	// Clear any diagnostics we'd published for a document that's no longer open.
	context.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         params.TextDocument.URI,
		Diagnostics: []protocol.Diagnostic{},
	})
	return nil
}

func (st *state) setDocument(uri protocol.DocumentUri, text string) {
	st.docsMu.Lock()
	st.documents[uri] = text
	st.docsMu.Unlock()
}

func (st *state) publishDiagnostics(context *glsp.Context, uri protocol.DocumentUri) {
	st.docsMu.Lock()
	text, ok := st.documents[uri]
	st.docsMu.Unlock()
	if !ok {
		return
	}

	// Always computed fresh from the current text, and always sent as a
	// complete replacement set (LSP publishDiagnostics semantics: the
	// client replaces whatever it had for this URI), so there's nothing to
	// accumulate or explicitly clear between publishes.
	diagnostics := diagnosticsForText(text)
	context.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         uri,
		Diagnostics: diagnostics,
	})
}

// diagnosticsForText parses, envelope-validates, and lints text and
// translates the resulting structured errors into LSP diagnostics.
func diagnosticsForText(text string) []protocol.Diagnostic {
	ic, errs := edifact.Parse(text)
	errs = append(errs, edifact.ValidateEnvelopes(ic)...)
	errs = append(errs, edifact.Lint(ic)...)

	source := Name
	diagnostics := make([]protocol.Diagnostic, 0, len(errs))
	for _, e := range errs {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:    errorRange(text, e.Pos),
			Severity: diagnosticSeverity(e.Severity),
			Source:   &source,
			Message:  e.Message,
		})
	}
	return diagnostics
}

func diagnosticSeverity(s edifact.Severity) *protocol.DiagnosticSeverity {
	var sev protocol.DiagnosticSeverity
	switch s {
	case edifact.SeverityWarning:
		sev = protocol.DiagnosticSeverityWarning
	case edifact.SeverityInfo:
		sev = protocol.DiagnosticSeverityInformation
	default:
		sev = protocol.DiagnosticSeverityError
	}
	return &sev
}

// errorRange builds a (non-zero-width, where possible) LSP Range for a byte
// offset in text, so clients render a visible squiggle rather than a
// zero-width marker.
func errorRange(text string, pos edifact.Position) protocol.Range {
	start := offsetToLSPPosition(text, pos.Offset)
	end := offsetToLSPPosition(text, pos.Offset+1)
	if end.Line == start.Line && end.Character == start.Character {
		// pos.Offset was at (or past) EOF; widen backwards instead so the
		// range still isn't zero-width when there's a preceding character.
		if pos.Offset > 0 {
			start = offsetToLSPPosition(text, pos.Offset-1)
		}
	}
	return protocol.Range{Start: start, End: end}
}

// offsetToLSPPosition converts a 0-based byte offset into text to an LSP
// Position: a 0-based line number and a UTF-16 code-unit character offset
// within that line, per the LSP spec.
func offsetToLSPPosition(text string, offset int) protocol.Position {
	if offset < 0 {
		offset = 0
	}
	if offset > len(text) {
		offset = len(text)
	}

	line := 0
	lineStart := 0
	for i := 0; i < offset; i++ {
		if text[i] == '\n' {
			line++
			lineStart = i + 1
		}
	}

	units := utf16.Encode([]rune(text[lineStart:offset]))
	return protocol.Position{Line: protocol.UInteger(line), Character: protocol.UInteger(len(units))}
}
