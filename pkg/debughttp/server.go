// Package debughttp exposes development-only runtime diagnostics on a
// dedicated HTTP listener.
package debughttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"time"
)

const shutdownTimeout = 2 * time.Second

// Config controls the optional debug HTTP listener.
type Config struct {
	Enabled  bool
	ListenOn string
}

// Server is a managed component suitable for merged-service runners.
type Server struct {
	config Config
}

// New creates a debug HTTP component.
func New(config Config) *Server {
	return &Server{config: config}
}

// Name identifies the component in merged-service errors.
func (s *Server) Name() string {
	return "debug-http"
}

// Run serves runtime diagnostics until ctx is canceled.
func (s *Server) Run(ctx context.Context) error {
	if !s.config.Enabled {
		return nil
	}

	listener, err := net.Listen("tcp", s.config.ListenOn)
	if err != nil {
		return fmt.Errorf("listen debug HTTP on %s: %w", s.config.ListenOn, err)
	}

	server := &http.Server{
		Handler:           NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}
	serveErr := make(chan error, 1)
	go func() {
		err := server.Serve(listener)
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		serveErr <- err
	}()

	select {
	case err := <-serveErr:
		if err != nil {
			return fmt.Errorf("serve debug HTTP: %w", err)
		}
		return nil
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if shutdownErr := server.Shutdown(shutdownCtx); shutdownErr != nil {
			// Shutdown does not close connections whose handlers ignore request
			// cancellation. Parent cancellation is still a normal component exit,
			// so force-close after the grace period and drain Serve below.
			if closeErr := server.Close(); closeErr != nil && !errors.Is(closeErr, http.ErrServerClosed) {
				return fmt.Errorf("force close debug HTTP after shutdown error %v: %w", shutdownErr, closeErr)
			}
		}
		if err := <-serveErr; err != nil {
			return fmt.Errorf("serve debug HTTP: %w", err)
		}
		return nil
	}
}

// NewHandler returns the private pprof handler used by this component. The
// component never serves http.DefaultServeMux.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return mux
}
