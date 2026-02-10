package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	port   int
	root   *chi.Mux
	routes *chi.Mux
	logger *string
}

func NewServer(routes *chi.Mux, cfg *config.Config) *Server {
	router := chi.NewRouter()
	logger := "zap"

	return &Server{
		port:   cfg.ServerPort,
		root:   router,
		logger: &logger,
		routes: routes,
	}
}

func (s *Server) StartServer() {
	s.MountRoutes()

	fmt.Printf(color.GreenString("Server listening on port %s...\n"), color.YellowString(strconv.Itoa(s.port)))

	chi.Walk(s.root, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s%s%s %s %s\n", color.HiGreenString("["), color.GreenString(method), color.HiGreenString("]"), color.CyanString(route), color.BlueString(strconv.Itoa(len(middlewares))))
		return nil
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.root)
	if err != nil {
		log.Fatalf("ERROR:server failed to start %v\n", err)
	}
}

func (s *Server) MountRoutes() {
	s.root.Mount("/api", s.routes)
}

func (s *Server) MountMiddlewares() {
	s.root.Use(middleware.Recoverer)
	s.root.Use(s.Logger)
}

func (s *Server) Logger(next http.Handler) http.Handler {
	emptyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	if s.logger == nil {
		return emptyHandler
	}
	switch *s.logger {
	case "zap":
		return s.ZapLogger(next)
	default:
		fmt.Fprintf(os.Stderr, "ERROR:%s logger not supported", *s.logger)
	}

	return emptyHandler
}

func (s *Server) DefaultLogger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

func (s *Server) ZapLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		var payload interface{}
		err = sonic.Unmarshal(body, payload)
		if err != nil {
			payload = string(body)
		}

		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("Request Info",
			zap.String("method", r.Method),
			zap.String("route", r.URL.Path),
			zap.String("host", r.Host),
			zap.String("user_agent", r.UserAgent()),
			zap.String("protocol", r.Proto),
			zap.String("response_status", r.Response.Status),
			zap.Int("response_status_code", r.Response.StatusCode),
			zap.Any("request_body", payload),
		)
		next.ServeHTTP(w, r)
	})
}
