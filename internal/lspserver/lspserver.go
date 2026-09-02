// Package lspserver wires the edifact-ls language server's capabilities into
// the glsp protocol handler and transport.
package lspserver

import (
	"sync"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	glspserver "github.com/tliron/glsp/server"
)

// Name is the server name reported to clients during initialize.
const Name = "edifact-ls"

// Version is the server version reported to clients during initialize. Set
// at build time via -ldflags "-X .../internal/lspserver.Version=vX.Y.Z"
// (see Makefile); "dev" when built without that flag.
var Version = "dev"

// Server is the edifact-ls language server.
type Server struct {
	glsp  *glspserver.Server
	state *state
}

// state tracks session-lifecycle facts the handler funcs need across calls.
type state struct {
	shutdownReceived bool

	docsMu    sync.Mutex
	documents map[protocol.DocumentUri]string
}

// New builds a Server with no language features enabled yet beyond the base
// initialize/initialized/shutdown/exit lifecycle.
func New() *Server {
	st := &state{documents: map[protocol.DocumentUri]string{}}

	handler := &protocol.Handler{
		Initialized:             st.initialized,
		Shutdown:                st.shutdown,
		Exit:                    st.exit,
		TextDocumentDidOpen:     st.textDocumentDidOpen,
		TextDocumentDidChange:   st.textDocumentDidChange,
		TextDocumentDidClose:    st.textDocumentDidClose,
		TextDocumentFormatting:  st.textDocumentFormatting,
		TextDocumentHover:       st.textDocumentHover,
		TextDocumentCodeAction:  st.textDocumentCodeAction,
		WorkspaceExecuteCommand: st.workspaceExecuteCommand,
	}
	// Deriving capabilities from the handler itself (rather than
	// hand-listing them) means CreateServerCapabilities always reflects
	// whichever handler funcs are actually wired above, so a new feature
	// can't be added here without also being advertised to clients.
	handler.Initialize = func(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
		capabilities := handler.CreateServerCapabilities()
		// We track whole-document text (see setDocument), not incremental
		// ranges, so require full-document sync rather than the library's
		// incremental default.
		if sync, ok := capabilities.TextDocumentSync.(*protocol.TextDocumentSyncOptions); ok {
			full := protocol.TextDocumentSyncKindFull
			sync.Change = &full
		}
		// CreateServerCapabilities sets ExecuteCommandProvider whenever
		// WorkspaceExecuteCommand is wired, but (unlike the other
		// capabilities it derives) leaves Commands empty -- the LSP spec
		// requires it to list the supported command IDs.
		if capabilities.ExecuteCommandProvider != nil {
			capabilities.ExecuteCommandProvider.Commands = []string{CommandMinify}
		}
		return protocol.InitializeResult{
			Capabilities: capabilities,
			ServerInfo: &protocol.InitializeResultServerInfo{
				Name:    Name,
				Version: &Version,
			},
		}, nil
	}

	return &Server{
		glsp:  glspserver.NewServer(handler, Name, false),
		state: st,
	}
}

// RunStdio serves the language server over stdin/stdout until the client
// disconnects (normally after sending an "exit" notification).
func (s *Server) RunStdio() error {
	return s.glsp.RunStdio()
}

// ExitCode reports the process exit code to use once RunStdio has returned,
// per the LSP spec: 0 if the client sent "shutdown" before disconnecting,
// 1 otherwise.
func (s *Server) ExitCode() int {
	if s.state.shutdownReceived {
		return 0
	}
	return 1
}

func (st *state) initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func (st *state) shutdown(context *glsp.Context) error {
	st.shutdownReceived = true
	return nil
}

func (st *state) exit(context *glsp.Context) error {
	return nil
}
