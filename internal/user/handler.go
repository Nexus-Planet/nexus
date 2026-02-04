package user

import (
	"encoding/json"
	"net/http"

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
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	user, err := h.svc.CreateUser(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	api.JSON(w, http.StatusCreated, user)
}

func (h *Handler) FindOneUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.svc.FindOneUser(r.Context(), id)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
	}
	api.JSON(w, http.StatusOK, user)
}

func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.FindAllUsers(r.Context())
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
	}

	api.JSON(w, http.StatusOK, users)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeactivateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
