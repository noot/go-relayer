package rpc

import (
	"context"
	"errors"
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

	if err = rpcServer.RegisterService(relayService, "relayer"); err != nil {
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
func (s *Server) HttpURL() string {
	return fmt.Sprintf("http://%s", s.httpServer.Addr)
}

// Start starts the JSON-RPC server.
func (s *Server) Start() error {
	log.Infof("Starting RPC server on %s", s.HttpURL())
	serverErr := make(chan error, 1)
	go func() {
		// Serve never returns nil. Even if it stops because Shutdown() was called, it
		// will return http.ErrServerClosed.
		serverErr <- s.httpServer.Serve(s.listener)
	}()

	select {
	case <-s.ctx.Done():
		// We are passing Shutdown a context that is already closed, which means it will
		// shut down immediately. If you pass a live context, it will stop listening for
		// new connections, but try to finish serving existing request while the context
		// is valid.
		if err := s.httpServer.Shutdown(s.ctx); err != nil {
			log.Warnf("RPC server shutdown errored: %s", err)
		} else {
			log.Infof("RPC server shut down")
		}
		// We shut down because the context was cancelled, so that's the error to return
		return s.ctx.Err()
	case err := <-serverErr:
		if errors.Is(err, http.ErrServerClosed) {
			log.Infof("RPC server shut down")
		} else {
			log.Errorf("RPC server failed: %s", err)
		}
		return err
	}
}

// Stop shuts down the http server. If the context the server was created with is still
// valid, it will be a graceful shutdown where existing connections are serviced until
// finished. If the context is cancelled, the shutdown will be immediate.
func (s *Server) Stop() error {
	return s.httpServer.Shutdown(s.ctx)
}
