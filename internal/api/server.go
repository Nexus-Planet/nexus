package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/color"
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
	fmt.Printf("Server listening on port %d...\n", s.port)

	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s%s%s %s %s\n", color.HiGreenString("["), color.GreenString(method), color.HiGreenString("]"), color.CyanString(route), color.BlueString(strconv.Itoa(len(middlewares))))
		return nil
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
	if err != nil {
		log.Fatalf("ERROR:server failed to start %v\n", err)
	}
}
