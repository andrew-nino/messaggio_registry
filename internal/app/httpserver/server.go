package server

import (
	"context"
	"net/http"
	"time"

	v1 "github.com/andrew-nino/messaggio/internal/controller/http/v1"
	service "github.com/andrew-nino/messaggio/internal/service/registry"

	"github.com/sirupsen/logrus"
)

type Registry interface {
}

type Server struct {
	log        *logrus.Logger
	handler    *v1.Handler
	httpServer *http.Server
	port       string
}

func New(log *logrus.Logger, services *service.ApplicationServices, port string) *Server {

	handler := v1.NewHandler(log, services, services)

	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        handler.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    12 * time.Second,
		WriteTimeout:   12 * time.Second,
	}

	return &Server{
		log:        log,
		handler:    handler,
		httpServer: httpServer,
		port:       port,
	}
}

// Configure with the necessary parameters and start the server.
func (s *Server) MustRun() {

	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.log.Fatalf("HTTP server failed to start: %v", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.log.Infof("HTTP server shutdown with error: %v", err)
		return err
	}
	return nil
}
