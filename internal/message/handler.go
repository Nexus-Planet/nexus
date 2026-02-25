package message

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"
	"github.com/nexus-planet/nexus/internal/api"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{}
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := sonic.ConfigDefault.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	h.svc.SendMessage(r.Context(), msg)

	api.JSON(w, http.StatusOK, msg)
}

func (h *Handler) FindOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	msg, err := h.svc.FindOne(r.Context(), id)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
	}

	api.JSON(w, http.StatusOK, msg)
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	msgs, err := h.svc.FindAll(r.Context())
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
	}

	api.JSON(w, http.StatusOK, msgs)
}

func (h *Handler) UpdateData(w http.ResponseWriter, r *http.Request) {
	var m UpdateMessage
	err := sonic.ConfigDefault.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	msg, err := h.svc.UpdateData(r.Context(), m)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}

	api.JSON(w, http.StatusOK, msg)
}

func (h *Handler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.svc.SoftDelete(r.Context(), id)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}
