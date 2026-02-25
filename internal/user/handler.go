package user

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/api"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var c UserCredentials
	err := sonic.ConfigDefault.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	user, err := h.svc.CreateUser(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	api.JSON(w, http.StatusCreated, user)
}

func (h *Handler) FindOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.svc.FindOne(r.Context(), id)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
	}
	api.JSON(w, http.StatusOK, user)
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.FindAll(r.Context())
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
	}

	api.JSON(w, http.StatusOK, users)
}

func (h *Handler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Deactivate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Reactivate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
