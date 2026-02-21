package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/api"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
	"github.com/nexus-planet/nexus-planet-api/internal/websocket"
)

func main() {
	cfg := config.Load()
	ws := websocket.NewWebSocket()
	svc := websocket.NewService(ws.Upgrader)
	handler := websocket.NewHandler(svc)

	server := api.NewServer(&cfg)
	server.MountMiddlewares()
	server.MountRoutes("/ws", func(r chi.Router) {
		websocketRouter(r, handler)
	})
	server.StartServer()

}

func websocketRouter(r chi.Router, handler *websocket.Handler) chi.Router {
	r.Get("/", handler.Handler)

	return r
}
