package lspserver

import (
	"strings"
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

// diagnosticsForText runs edifact.Validate over text and translates the
// resulting structured errors into LSP diagnostics.
func diagnosticsForText(text string) []protocol.Diagnostic {
	_, errs := edifact.Validate(text)

	source := Name
	diagnostics := make([]protocol.Diagnostic, 0, len(errs))
	for _, e := range errs {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:    errorRange(text, e.Pos),
			Severity: diagnosticSeverity(e.Severity),
			Code:     diagnosticCode(e.Code),
			Source:   &source,
			Message:  e.Message,
		})
	}
	return diagnostics
}

// diagnosticCode wraps a non-empty edifact.Error.Code for the protocol's
// code field, or returns nil when there's no stable code to report.
func diagnosticCode(code string) *protocol.IntegerOrString {
	if code == "" {
		return nil
	}
	return &protocol.IntegerOrString{Value: code}
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

// lspPositionToOffset converts an LSP Position (0-based line, UTF-16
// code-unit character offset within that line) to a 0-based byte offset
// into text -- the reverse of offsetToLSPPosition. A line/character past
// the end of the text clamps to len(text).
func lspPositionToOffset(text string, pos protocol.Position) int {
	lineStart := 0
	for line := 0; line < int(pos.Line); line++ {
		idx := strings.IndexByte(text[lineStart:], '\n')
		if idx < 0 {
			return len(text)
		}
		lineStart += idx + 1
	}

	lineEnd := len(text)
	if idx := strings.IndexByte(text[lineStart:], '\n'); idx >= 0 {
		lineEnd = lineStart + idx
	}

	units := utf16.Encode([]rune(text[lineStart:lineEnd]))
	if int(pos.Character) >= len(units) {
		return lineEnd
	}
	return lineStart + len(string(utf16.Decode(units[:pos.Character])))
}
