package lspserver

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
	"github.com/sourcegraph/jsonrpc2"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestWorkspaceExecuteCommandUnknownCommand(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	_, err := st.workspaceExecuteCommand(nil, &protocol.ExecuteCommandParams{Command: "not-a-real-command"})
	if err == nil {
		t.Fatal("expected an error for an unknown command")
	}
}

func TestMinifyMissingArguments(t *testing.T) {
	st := &state{documents: map[protocol.DocumentUri]string{}}
	err := st.minify(nil, &protocol.ExecuteCommandParams{Command: CommandMinify})
	if err == nil {
		t.Fatal("expected an error when no document URI argument is given")
	}
}

func TestMinifyRoundTripsToFormat(t *testing.T) {
	data, err := os.ReadFile("../../testdata/minimal.edi")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}

	ic, errs := edifact.Parse(string(data))
	if errs.HasErrors() {
		t.Fatalf("unexpected parse errors: %v", errs)
	}
	minified := edifact.Render(ic, false)
	if strings.Contains(minified, "\n") {
		t.Fatalf("minified output contains a newline: %q", minified)
	}

	reformatted, errs := edifact.Parse(minified)
	if errs.HasErrors() {
		t.Fatalf("re-parsing minified output failed: %v", errs)
	}
	// format(minify(x)) should land back on the same text as format(x), and
	// since the fixture is already in format's own style, that's x itself.
	if got := edifact.Render(reformatted, true); got != string(data) {
		t.Fatalf("format(minify(x)) = %q, want %q", got, string(data))
	}
}

// TestMinifyAppliesEditViaWorkspaceApplyEdit drives the server over the
// real stdio transport, invokes the edifact-ls.minify command, and asserts
// the server sends a workspace/applyEdit request containing the expected
// single-line replacement -- the same integration-level guarantee
// TestDiagnosticsPublishedOnOpenAndChange gives the diagnostics path.
func TestMinifyAppliesEditViaWorkspaceApplyEdit(t *testing.T) {
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

	applyEditReceived := make(chan protocol.ApplyWorkspaceEditParams, 1)
	handler := jsonrpc2.HandlerWithError(func(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error) {
		if req.Method == string(protocol.ServerWorkspaceApplyEdit) {
			var params protocol.ApplyWorkspaceEditParams
			if err := json.Unmarshal(*req.Params, &params); err != nil {
				t.Errorf("unmarshaling applyEdit params: %v", err)
				return nil, err
			}
			applyEditReceived <- params
			return protocol.ApplyWorkspaceEditResponse{Applied: true}, nil
		}
		return nil, nil
	})
	client := jsonrpc2.NewConn(ctx, clientStream, handler)
	t.Cleanup(func() { client.Close() })

	var initResult protocol.InitializeResult
	if err := client.Call(ctx, "initialize", protocol.InitializeParams{}, &initResult); err != nil {
		t.Fatalf("initialize call: %v", err)
	}
	if initResult.Capabilities.ExecuteCommandProvider == nil {
		t.Fatal("server did not advertise ExecuteCommandProvider")
	}
	if got := initResult.Capabilities.ExecuteCommandProvider.Commands; len(got) != 1 || got[0] != CommandMinify {
		t.Fatalf("advertised commands = %v, want [%s]", got, CommandMinify)
	}
	if err := client.Notify(ctx, "initialized", protocol.InitializedParams{}); err != nil {
		t.Fatalf("initialized notify: %v", err)
	}

	const uri = "file:///minify.edi"
	const text = "UNH+1'\nBGM+220'\n"
	if err := client.Notify(ctx, "textDocument/didOpen", protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{URI: uri, LanguageID: "edifact", Version: 1, Text: text},
	}); err != nil {
		t.Fatalf("didOpen notify: %v", err)
	}

	var result any
	if err := client.Call(ctx, "workspace/executeCommand", protocol.ExecuteCommandParams{
		Command:   CommandMinify,
		Arguments: []any{uri},
	}, &result); err != nil {
		t.Fatalf("executeCommand call: %v", err)
	}

	select {
	case params := <-applyEditReceived:
		edits, ok := params.Edit.Changes[uri]
		if !ok || len(edits) != 1 {
			t.Fatalf("applyEdit changes for %s = %+v, want exactly 1 edit", uri, params.Edit.Changes)
		}
		if want := "UNH+1'BGM+220'"; edits[0].NewText != want {
			t.Errorf("minify NewText = %q, want %q", edits[0].NewText, want)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for workspace/applyEdit")
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
