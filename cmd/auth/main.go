package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/api"
	"github.com/nexus-planet/nexus-planet-api/internal/auth"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
	"github.com/nexus-planet/nexus-planet-api/internal/routes"
)

func main() {
	ctx := context.Background()

	config.LoadArgs()

	db, err := db.Connect(ctx, db.Postgres)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:%v", err)
	}

	// api v1
	r := routes.NewRouter("/v1")
	repo := auth.NewRepository(db)
	svc := auth.NewService(repo)
	handler := auth.NewHandler(svc)

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

	// possible v2 ??

	api.StartServer(r)
}
