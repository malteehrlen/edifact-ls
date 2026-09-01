package lspserver

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sourcegraph/jsonrpc2"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestOffsetToLSPPosition(t *testing.T) {
	text := "UNH+1'\nBGM+220'\n"
	cases := []struct {
		offset   int
		wantLine protocol.UInteger
		wantChar protocol.UInteger
	}{
		{0, 0, 0},
		{6, 0, 6},  // right before the first '\n'
		{7, 1, 0},  // right after the first '\n', start of "BGM..."
		{10, 1, 3}, // "BGM" -> 3 chars in
	}
	for _, c := range cases {
		got := offsetToLSPPosition(text, c.offset)
		if got.Line != c.wantLine || got.Character != c.wantChar {
			t.Errorf("offsetToLSPPosition(%d) = %+v, want {Line:%d Character:%d}", c.offset, got, c.wantLine, c.wantChar)
		}
	}
}

func TestDiagnosticsForTextValidDocument(t *testing.T) {
	data, err := os.ReadFile("../../testdata/minimal.edi")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	diags := diagnosticsForText(string(data))
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics for valid document: %+v", diags)
	}
}

func TestDiagnosticsForTextSyntaxError(t *testing.T) {
	// A complete, otherwise-valid envelope with one malformed segment, so
	// envelope validation doesn't also fire (that's covered separately by
	// TestDiagnosticsForTextEnvelopeError) and this isolates the syntax error.
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'1BC+garbage'UNT+3+1'UNZ+1+1'"
	diags := diagnosticsForText(src)
	if len(diags) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %+v", len(diags), diags)
	}
	d := diags[0]
	if !strings.Contains(d.Message, "invalid segment tag") {
		t.Errorf("message = %q, want it to mention an invalid segment tag", d.Message)
	}
	if d.Severity == nil || *d.Severity != protocol.DiagnosticSeverityError {
		t.Errorf("severity = %v, want DiagnosticSeverityError", d.Severity)
	}
	if d.Range.Start.Line != 0 {
		t.Errorf("range start line = %d, want 0", d.Range.Start.Line)
	}
}

func TestDiagnosticsForTextEnvelopeError(t *testing.T) {
	src := "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'BGM+220'UNT+3+1'"
	diags := diagnosticsForText(src)
	if len(diags) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "missing UNZ") {
		t.Errorf("message = %q, want it to mention missing UNZ", diags[0].Message)
	}
}

// capturingHandler records every notification/request the test client
// receives from the server (e.g. textDocument/publishDiagnostics) onto a
// channel so tests can assert on them.
type capturingHandler struct {
	received chan *jsonrpc2.Request
}

func newCapturingHandler() *capturingHandler {
	return &capturingHandler{received: make(chan *jsonrpc2.Request, 16)}
}

func (h *capturingHandler) jsonrpc2Handler() jsonrpc2.Handler {
	return jsonrpc2.HandlerWithError(func(_ context.Context, _ *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error) {
		h.received <- req
		return nil, nil
	})
}

func (h *capturingHandler) waitForPublishDiagnostics(t *testing.T, timeout time.Duration) protocol.PublishDiagnosticsParams {
	t.Helper()
	deadline := time.After(timeout)
	for {
		select {
		case req := <-h.received:
			if req.Method != string(protocol.ServerTextDocumentPublishDiagnostics) {
				continue
			}
			var params protocol.PublishDiagnosticsParams
			if err := json.Unmarshal(*req.Params, &params); err != nil {
				t.Fatalf("unmarshaling publishDiagnostics params: %v", err)
			}
			return params
		case <-deadline:
			t.Fatal("timed out waiting for textDocument/publishDiagnostics")
			return protocol.PublishDiagnosticsParams{}
		}
	}
}

// TestDiagnosticsPublishedOnOpenAndChange drives the server over the same
// stdio-shaped in-memory transport as TestHandshake, opens a known-bad
// document, and asserts the exact diagnostics published -- then edits the
// document to valid content and asserts the diagnostics are replaced with
// an empty set (not accumulated).
func TestDiagnosticsPublishedOnOpenAndChange(t *testing.T) {
	serverIn, clientOut, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	clientIn, serverOut, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}

	origStdin, origStdout := os.Stdin, os.Stdout
	os.Stdin, os.Stdout = serverIn, serverOut
	t.Cleanup(func() { os.Stdin, os.Stdout = origStdin, origStdout })

	srv := New()
	serverDone := make(chan error, 1)
	go func() { serverDone <- srv.RunStdio() }()

	ctx := context.Background()
	clientStream := jsonrpc2.NewBufferedStream(pipeStream{r: clientIn, w: clientOut}, jsonrpc2.VSCodeObjectCodec{})
	capture := newCapturingHandler()
	client := jsonrpc2.NewConn(ctx, clientStream, capture.jsonrpc2Handler())
	t.Cleanup(func() { client.Close() })

	var initResult protocol.InitializeResult
	if err := client.Call(ctx, "initialize", protocol.InitializeParams{}, &initResult); err != nil {
		t.Fatalf("initialize call: %v", err)
	}
	if err := client.Notify(ctx, "initialized", protocol.InitializedParams{}); err != nil {
		t.Fatalf("initialized notify: %v", err)
	}

	const uri = "file:///test.edi"
	const badText = "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'1BC+garbage'UNT+3+1'UNZ+1+1'"

	if err := client.Notify(ctx, "textDocument/didOpen", protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{URI: uri, LanguageID: "edifact", Version: 1, Text: badText},
	}); err != nil {
		t.Fatalf("didOpen notify: %v", err)
	}

	params := capture.waitForPublishDiagnostics(t, 2*time.Second)
	if params.URI != uri {
		t.Errorf("published diagnostics URI = %q, want %q", params.URI, uri)
	}
	if len(params.Diagnostics) != 1 {
		t.Fatalf("got %d diagnostics on open, want 1: %+v", len(params.Diagnostics), params.Diagnostics)
	}
	if !strings.Contains(params.Diagnostics[0].Message, "invalid segment tag") {
		t.Errorf("diagnostic message = %q, want it to mention an invalid segment tag", params.Diagnostics[0].Message)
	}

	// Now "fix" the document via didChange (full-sync) and expect the
	// diagnostics to be replaced with an empty set, not appended to.
	const fixedText = "UNB+UNOA:1+S+R+201001:1200+1'UNH+1+ORDERS:D:96A:UN'UNT+2+1'UNZ+1+1'"
	if err := client.Notify(ctx, "textDocument/didChange", protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{URI: uri},
			Version:                2,
		},
		ContentChanges: []any{
			protocol.TextDocumentContentChangeEventWhole{Text: fixedText},
		},
	}); err != nil {
		t.Fatalf("didChange notify: %v", err)
	}

	params = capture.waitForPublishDiagnostics(t, 2*time.Second)
	if len(params.Diagnostics) != 0 {
		t.Fatalf("got %d diagnostics after fixing the document, want 0 (stale diagnostic left behind): %+v", len(params.Diagnostics), params.Diagnostics)
	}

	var shutdownResult any
	if err := client.Call(ctx, "shutdown", nil, &shutdownResult); err != nil {
		t.Fatalf("shutdown call: %v", err)
	}
	if err := client.Notify(ctx, "exit", nil); err != nil {
		t.Fatalf("exit notify: %v", err)
	}

	select {
	case err := <-serverDone:
		if err != nil {
			t.Fatalf("RunStdio returned error: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for server to disconnect after exit")
	}
}
