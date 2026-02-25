package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/nexus-planet/nexus/internal/api"
	"github.com/nexus-planet/nexus/internal/auth"
	"github.com/nexus-planet/nexus/internal/config"
	"github.com/nexus-planet/nexus/internal/db"
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
	repo := auth.NewRepository(db)
	svc := auth.NewService(repo)
	handler := auth.NewHandler(svc)

	// possible v2 ??

	server := api.NewServer(&cfg)
	server.MountMiddlewares()
	server.MountRoutes("/api", func(r chi.Router) {
		authRouter(r, handler)
	})
	server.StartServer()
}

func authRouter(r chi.Router, handler *auth.Handler) chi.Router {

	r.Route("/auth", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.JwtToken))
			r.Use(jwtauth.Authenticator(auth.JwtToken))

			r.Post("/login", handler.Login)
		})

		r.Group(func(r chi.Router) {
			r.Post("/signup", handler.SignUp)
			r.Post("/login", handler.Login)
			r.Post("/logout", handler.Logout)
		})
	})

	return r
}
