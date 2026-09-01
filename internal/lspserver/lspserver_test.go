package lspserver

import (
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/sourcegraph/jsonrpc2"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// pipeStream glues together the two independent os.Pipe halves that make up
// one direction of client<->server traffic each, so it can be used as the
// client side's io.ReadWriteCloser in jsonrpc2.NewBufferedStream.
type pipeStream struct {
	r *os.File
	w *os.File
}

func (p pipeStream) Read(b []byte) (int, error)  { return p.r.Read(b) }
func (p pipeStream) Write(b []byte) (int, error) { return p.w.Write(b) }
func (p pipeStream) Close() error {
	err := p.r.Close()
	if werr := p.w.Close(); err == nil {
		err = werr
	}
	return err
}

// TestHandshake drives the server over a real stdio-shaped transport (two
// os.Pipe pairs standing in for stdin/stdout) through initialize ->
// initialized -> shutdown -> exit, exactly as a real LSP client would, and
// asserts the responses and resulting process exit code.
func TestHandshake(t *testing.T) {
	serverIn, clientOut, err := os.Pipe() // client writes here, server reads as stdin
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	clientIn, serverOut, err := os.Pipe() // server writes here, client reads as stdout
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}

	origStdin, origStdout := os.Stdin, os.Stdout
	os.Stdin, os.Stdout = serverIn, serverOut
	t.Cleanup(func() {
		os.Stdin, os.Stdout = origStdin, origStdout
	})

	srv := New()
	serverDone := make(chan error, 1)
	go func() { serverDone <- srv.RunStdio() }()

	ctx := context.Background()
	clientStream := jsonrpc2.NewBufferedStream(pipeStream{r: clientIn, w: clientOut}, jsonrpc2.VSCodeObjectCodec{})
	noopHandler := jsonrpc2.HandlerWithError(func(context.Context, *jsonrpc2.Conn, *jsonrpc2.Request) (any, error) {
		return nil, nil
	})
	client := jsonrpc2.NewConn(ctx, clientStream, noopHandler)
	t.Cleanup(func() { client.Close() })

	var initResult protocol.InitializeResult
	if err := client.Call(ctx, "initialize", protocol.InitializeParams{}, &initResult); err != nil {
		t.Fatalf("initialize call: %v", err)
	}
	if initResult.ServerInfo == nil || initResult.ServerInfo.Name != Name {
		t.Fatalf("unexpected serverInfo: %+v", initResult.ServerInfo)
	}
	if hp, ok := initResult.Capabilities.HoverProvider.(bool); !ok || !hp {
		t.Fatalf("Capabilities.HoverProvider = %+v, want true", initResult.Capabilities.HoverProvider)
	}

	if err := client.Notify(ctx, "initialized", protocol.InitializedParams{}); err != nil {
		t.Fatalf("initialized notify: %v", err)
	}

	// An unhandled method should come back as a proper JSON-RPC error, not
	// crash the connection. (textDocument/hover used to serve as this
	// example; it's a real handled method now -- see hover_test.go.)
	var unused any
	err = client.Call(ctx, "textDocument/definition", struct{}{}, &unused)
	rpcErr, ok := err.(*jsonrpc2.Error)
	if !ok || rpcErr.Code != jsonrpc2.CodeMethodNotFound {
		t.Fatalf("expected method-not-found error, got: %v", err)
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

	if got, want := srv.ExitCode(), 0; got != want {
		t.Fatalf("ExitCode() = %d, want %d (shutdown preceded exit)", got, want)
	}
}

var _ io.ReadWriteCloser = pipeStream{}
