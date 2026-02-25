package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nexus-planet/nexus/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	port   int
	root   chi.Router
	logger *zap.Logger
}

func NewServer(cfg *config.Config) *Server {
	router := chi.NewRouter()

	l, _ := zap.NewProduction()

	return &Server{
		port:   cfg.ServerPort,
		root:   router,
		logger: l,
	}
}

func (s *Server) StartServer() {
	defer s.logger.Sync()

	fmt.Printf(color.GreenString("Server listening on port %s...\n"), color.YellowString(strconv.Itoa(s.port)))

	chi.Walk(s.root, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s%s%s %s %s %s\n", color.HiGreenString("["), color.GreenString(method), color.HiGreenString("]"), color.CyanString(route), color.YellowString("middlewares"), color.BlueString(strconv.Itoa(len(middlewares))))
		return nil
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.root)
	if err != nil {
		log.Fatalf("ERROR:server failed to start %v\n", err)
	}
}

func (s *Server) MountRoutes(prefix string, mount func(r chi.Router)) {
	apiRouter := chi.NewRouter()
	mount(apiRouter)
	s.root.Mount(prefix, apiRouter)
}

func (s *Server) MountMiddlewares() {
	s.root.Use(middleware.Recoverer)
	s.root.Use(s.Logger)
}

func (s *Server) Logger(next http.Handler) http.Handler {
	return s.ZapLogger(next)
}

func (s *Server) ZapLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wr := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(w, r)

		reqLogger := s.logger.With(
			zap.String("protocol", r.Proto),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("id", middleware.GetReqID(r.Context())),
			zap.String("host", r.Host),
			zap.String("user_agent", r.UserAgent()),
			zap.Int("status", wr.Status()),
			zap.Int("size", wr.BytesWritten()),
		)
		reqLogger.Info("Request Info")
	})
}
