package websocket

import (
	"fmt"
	"net/http"
	"os"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := h.svc.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
		http.Error(w, "Failed to establish connection to the socket", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	go h.svc.StartHub()
	client := NewClient(conn, h.svc)
	h.svc.register <- client

	go client.reader()
	go client.writer()
}
