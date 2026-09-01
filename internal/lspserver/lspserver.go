// Package lspserver wires the edifact-ls language server's capabilities into
// the glsp protocol handler and transport.
package lspserver

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	glspserver "github.com/tliron/glsp/server"
)

// Name is the server name reported to clients during initialize.
const Name = "edifact-ls"

// Version is the server version reported to clients during initialize.
var Version = "0.0.1"

// Server is the edifact-ls language server.
type Server struct {
	glsp  *glspserver.Server
	state *state
}

// state tracks session-lifecycle facts the handler funcs need across calls.
type state struct {
	shutdownReceived bool
}

// New builds a Server with no language features enabled yet beyond the base
// initialize/initialized/shutdown/exit lifecycle.
func New() *Server {
	st := &state{}

	handler := &protocol.Handler{
		Initialize:  st.initialize,
		Initialized: st.initialized,
		Shutdown:    st.shutdown,
		Exit:        st.exit,
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

func (st *state) initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	return protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{},
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    Name,
			Version: &Version,
		},
	}, nil
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
