package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"tools/internals/config"
)

const (
	defaultReadTimeout     = 10 * time.Second
	defaultWriteTimeout    = 10 * time.Second
	defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	sv     *http.Server
	notify chan error
}

func NewServer(server http.Handler, config *config.Configuration) *Server {
	httpServer := &http.Server{
		Handler:        server,
		ReadTimeout:    defaultReadTimeout,
		WriteTimeout:   defaultWriteTimeout,
		MaxHeaderBytes: 1 << 20,
		Addr:           fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
	}
	s := &Server{
		sv:     httpServer,
		notify: make(chan error, 1),
	}
	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.sv.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
	defer cancel()

	return s.sv.Shutdown(ctx)
}
