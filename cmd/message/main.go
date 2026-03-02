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
	"github.com/nexus-planet/nexus/internal/message"
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
	repo := message.NewRepository(db)
	svc := message.NewService(repo)
	handler := message.NewHandler(svc)

	server := api.NewServer(&cfg)
	server.MountMiddlewares()
	server.MountRoutes("/api", func(r chi.Router) {
		messageRouter(r, handler)
	})
	server.StartServer()
}

func messageRouter(r chi.Router, handler *message.Handler) chi.Router {
	r.Route("/messages", func(r chi.Router) {
		r.Use(jwtauth.Verifier(auth.JwtToken))
		r.Use(jwtauth.Authenticator(auth.JwtToken))

		r.Post("/", handler.SendMessage)
		r.Get("/{id}", handler.FindOne)
		r.Get("/", handler.FindAll)
	})

	return r
}
