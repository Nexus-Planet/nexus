package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
)

func StartServer(r chi.Router) {
	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s] %s %d", method, route, len(middlewares))
		return nil
	})

	if config.CustomServerPort == 0 {
		fmt.Printf("Server listening on port %d...\n", config.ServerPort)
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), r)
		if err != nil {
			log.Fatalf("ERROR:server failed to start %v\n", err)
		}
	} else {
		fmt.Printf("Server listening on port %d...\n", config.CustomServerPort)
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.CustomServerPort), r)
		if err != nil {
			log.Fatalf("ERROR:server failed to start %v\n", err)
		}
	}
}
