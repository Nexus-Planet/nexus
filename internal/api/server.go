package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
)

type Server struct {
	port   int
	router *chi.Mux
}

func NewServer(router *chi.Mux, cfg *config.Config) *Server {
	return &Server{
		port:   cfg.ServerPort,
		router: router,
	}
}

func (s *Server) StartServer() {
	s.MountRoutes()

	fmt.Printf("Server listening on port %d...\n", s.port)

	chi.Walk(s.router, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s%s%s %s %s\n", color.HiGreenString("["), color.GreenString(method), color.HiGreenString("]"), color.CyanString(route), color.BlueString(strconv.Itoa(len(middlewares))))
		return nil
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
	if err != nil {
		log.Fatalf("ERROR:server failed to start %v\n", err)
	}
}

func (s *Server) MountRoutes() {
	s.router.Mount("/api", s.router)
}

func (s *Server) MountMiddlewares() {
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Logger)
}
