package lspserver

import (
	"fmt"
	"log"

	"github.com/malteehrlen/edifact-ls/internal/edifact"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// CommandMinify is the workspace/executeCommand ID that collapses a
// document to single-line "wire" EDIFACT. The command expects exactly one
// argument: the document's URI (a string).
const CommandMinify = "edifact-ls.minify"

func (st *state) workspaceExecuteCommand(context *glsp.Context, params *protocol.ExecuteCommandParams) (any, error) {
	switch params.Command {
	case CommandMinify:
		return nil, st.minify(context, params)
	default:
		return nil, fmt.Errorf("unknown command: %q", params.Command)
	}
}

func (st *state) minify(context *glsp.Context, params *protocol.ExecuteCommandParams) error {
	uri, err := commandDocumentURI(params)
	if err != nil {
		return err
	}

	st.docsMu.Lock()
	text, ok := st.documents[uri]
	st.docsMu.Unlock()
	if !ok {
		return fmt.Errorf("document not open: %s", uri)
	}

	ic, errs := edifact.Parse(text)
	if errs.HasErrors() {
		// Consistent with textDocument/formatting: don't touch a document
		// we can't fully make sense of.
		return nil
	}

	minified := edifact.Render(ic, false)
	if minified == text {
		return nil
	}

	edit := protocol.WorkspaceEdit{
		Changes: map[protocol.DocumentUri][]protocol.TextEdit{
			uri: {wholeDocumentReplace(text, minified)},
		},
	}

	// context.Call blocks for a response on the same connection this
	// handler is itself being invoked from -- glsp/jsonrpc2 dispatch
	// incoming requests inline on their single read-loop goroutine, so a
	// synchronous nested call here would deadlock that read loop against
	// itself (it can never read the applyEdit response while stuck waiting
	// for this handler to return). Firing it from a separate goroutine
	// avoids that; workspace/executeCommand's own response doesn't need to
	// wait for the edit to actually land, so failures are only logged.
	go func() {
		var response protocol.ApplyWorkspaceEditResponse
		context.Call(string(protocol.ServerWorkspaceApplyEdit), protocol.ApplyWorkspaceEditParams{Edit: edit}, &response)
		if !response.Applied {
			reason := "unknown reason"
			if response.FailureReason != nil {
				reason = *response.FailureReason
			}
			log.Printf("edifact-ls: client did not apply the %s edit for %s: %s", CommandMinify, uri, reason)
		}
	}()

	return nil
}

func commandDocumentURI(params *protocol.ExecuteCommandParams) (protocol.DocumentUri, error) {
	if len(params.Arguments) == 0 {
		return "", fmt.Errorf("command %q requires a document URI argument", params.Command)
	}
	uri, ok := params.Arguments[0].(string)
	if !ok || uri == "" {
		return "", fmt.Errorf("command %q's first argument must be a document URI string", params.Command)
	}
	return uri, nil
}
