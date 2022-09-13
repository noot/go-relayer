package rpc

import (
	"fmt"
	"net/http"

	"github.com/AthanorLabs/go-relayer/relayer"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"

	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("rpc")

// Server represents the JSON-RPC server
type Server struct {
	s    *rpc.Server
	port uint16
}

// Config ...
type Config struct {
	Port    uint16
	Relayer *relayer.Relayer
}

// NewServer ...
func NewServer(cfg *Config) (*Server, error) {
	s := rpc.NewServer()
	s.RegisterCodec(NewCodec(), "application/json")

	rs, err := NewRelayerService(cfg.Relayer)
	if err != nil {
		return nil, err
	}

	if err := s.RegisterService(rs, "relayer"); err != nil {
		return nil, err
	}

	return &Server{
		s:    s,
		port: cfg.Port,
	}, nil
}

// Start starts the JSON-RPC server.
func (s *Server) Start() <-chan error {
	errCh := make(chan error)

	go func() {
		r := mux.NewRouter()
		r.Handle("/", s.s)

		headersOk := handlers.AllowedHeaders([]string{"content-type", "username", "password"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		originsOk := handlers.AllowedOrigins([]string{"*"})

		log.Infof("starting RPC server on http://localhost:%d", s.port)

		if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), handlers.CORS(headersOk, methodsOk, originsOk)(r)); err != nil { //nolint:lll
			log.Errorf("failed to start http RPC server: %s", err)
			errCh <- err
		}
	}()

	return errCh
}
