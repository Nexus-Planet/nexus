package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/nexus-planet/nexus-planet-api/internal/auth"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
	"github.com/nexus-planet/nexus-planet-api/internal/routes"
	"github.com/nexus-planet/nexus-planet-api/internal/user"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ERROR:%v\n", err)
	}

	flag.IntVar(&config.CustomServerPort, "p", 0, "Changes the default port for server")
	flag.StringVar(&config.CustomDatabaseUrl, "du", "", "Changes the default database url")
	flag.Parse()

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

	// api v1
	r := routes.NewRouter("/v1")
	q := db.New(conn)
	authRepo := auth.NewRepository(q)
	userRepo := user.NewRepository(q)
	svc := auth.NewService(authRepo, userRepo)
	handler := auth.NewHandler(svc)

	r.Route("/auth", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.JwtToken))

			r.Post("/login", handler.Login)
		})

		r.Group(func(r chi.Router) {
			r.Post("/signup", handler.SignUp)
			r.Post("/login", handler.Login)
			r.Post("/logout", handler.Logout)
		})
	})

	// possible v2 ??

	if config.CustomServerPort == 0 {
		fmt.Printf("Server listening on port %d...\n", config.ServerPort)
		err = http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), r)
		if err != nil {
			log.Fatalf("ERROR:server failed to start %v\n", err)
		}
	} else {
		fmt.Printf("Server listening on port %d...\n", config.CustomServerPort)
		err = http.ListenAndServe(fmt.Sprintf(":%d", config.CustomServerPort), r)
		if err != nil {
			log.Fatalf("ERROR:server failed to start %v\n", err)
		}
	}
}
