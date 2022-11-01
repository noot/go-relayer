package rpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/athanorlabs/go-relayer/relayer"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"

	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("rpc")

// Server represents the JSON-RPC server
type Server struct {
	ctx        context.Context
	listener   net.Listener
	httpServer *http.Server
}

// Config is the RPC server configuration passed to NewServer
type Config struct {
	Ctx     context.Context
	Address string // "IP:port" to listen on
	Relayer *relayer.Relayer
}

// NewServer creates but does not start the relayer RPC service. Pre-start configuration
// that can fail, including reserving the server port, is handled here.
func NewServer(cfg *Config) (*Server, error) {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(NewCodec(), "application/json")

	relayService, err := NewRelayerService(cfg.Relayer)
	if err != nil {
		return nil, err
	}

	if err := rpcServer.RegisterService(relayService, "relayer"); err != nil {
		return nil, err
	}

	addr := cfg.Address
	if addr == "" {
		// default to restricted localhost access with a randomized port
		addr = "127.0.0.1:0"
	}

	router := mux.NewRouter()
	router.Handle("/", rpcServer)

	headersOk := handlers.AllowedHeaders([]string{"content-type", "username", "password"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	lc := net.ListenConfig{}
	ln, err := lc.Listen(cfg.Ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		ctx:      cfg.Ctx,
		listener: ln,
		httpServer: &http.Server{
			Addr:              ln.Addr().String(),
			ReadHeaderTimeout: time.Second,
			Handler:           handlers.CORS(headersOk, methodsOk, originsOk)(router),
			BaseContext: func(listener net.Listener) context.Context {
				return cfg.Ctx
			},
		},
	}, nil
}

// HttpURL returns the URL used for HTTP requests
func (s *Server) HttpURL() string { //nolint:revive
	return fmt.Sprintf("http://%s", s.httpServer.Addr)
}

// Start starts the JSON-RPC server.
func (s *Server) Start() <-chan error {
	errCh := make(chan error, 1)

	log.Infof("Starting RPC server on %s", s.HttpURL())

	go func() {
		err := s.httpServer.Serve(s.listener) // Serve never returns nil
		log.Errorf("RPC server failed: %s", err)
		errCh <- err
	}()

	return errCh
}
