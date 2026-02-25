package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/api"
	"github.com/nexus-planet/nexus-planet-api/internal/auth"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
	"github.com/nexus-planet/nexus-planet-api/internal/user"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cfg := config.Load()

	db, err := db.ConnectContext(ctx, db.Postgres, cfg.DataSourceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
	}

	// api v1
	repo := user.NewRepository(db)
	svc := user.NewService(repo)
	handler := user.NewHandler(svc)

	server := api.NewServer(&cfg)
	server.MountMiddlewares()
	server.MountRoutes("/api", func(r chi.Router) {
		userRouter(r, handler)
	})
	server.StartServer()
}

func userRouter(r chi.Router, handler *user.Handler) chi.Router {
	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.JwtToken))
			r.Use(jwtauth.Authenticator(auth.JwtToken))

			r.Post("/", handler.CreateUser)
			r.Get("/{id}", handler.FindOne)
			r.Get("/", handler.FindAll)
		})
	})

	return r
}
