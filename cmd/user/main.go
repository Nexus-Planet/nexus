package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/auth"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
	"github.com/nexus-planet/nexus-planet-api/internal/routes"
	"github.com/nexus-planet/nexus-planet-api/internal/user"
)

func main() {

	ctx := context.Background()

	flag.IntVar(&config.CustomServerPort, "p", 0, "Changes the default port for server")
	flag.StringVar(&config.CustomDatabaseUrl, "du", "", "Changes the default database url")
	flag.String("default", "", "Use default options from environment variables of system\ni.e:\nDATABASE_URL=<url>\nJWT_SECRET=<secret>")
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	var err error
	var conn *pgx.Conn

	if config.CustomDatabaseUrl == "" {
		conn, err = pgx.Connect(ctx, config.DatabaseUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR:Unable to connect to database %v\n", err)
		}
	} else {
		conn, err = pgx.Connect(ctx, config.CustomDatabaseUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR:Unable to connect to database %v\n", err)
		}
	}
	defer conn.Close(ctx)

	r := routes.NewRouter("/v1")
	q := db.New(conn)
	repo := user.NewRepository(q)
	svc := user.NewService(repo)
	handler := user.NewHandler(svc)
	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.JwtToken))
			r.Use(jwtauth.Authenticator(auth.JwtToken))

			r.Post("", handler.CreateUser)
			r.Get("/{id}", handler.FindOneUser)
			r.Get("", handler.FindAllUsers)
		})
	})
}
