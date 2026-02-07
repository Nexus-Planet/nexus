package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
)

type Server struct {
	port int
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		port: cfg.ServerPort,
	}
}

func (s *Server) StartServer(r chi.Router) {
	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s] %s %d", method, route, len(middlewares))
		return nil
	})

	fmt.Printf("Server listening on port %d...\n", s.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
	if err != nil {
		log.Fatalf("ERROR:server failed to start %v\n", err)
	}
}
