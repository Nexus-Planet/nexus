package auth

import (
	"encoding/json"
	"net/http"

	"github.com/nexus-planet/nexus-planet-api/internal/api"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var c Credentials
	api.DecodeJSON(w, r, "", &c)

	user, err := h.svc.CreateUser(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	api.JSON(w, http.StatusCreated, user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var c Credentials
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if c.Email == "" || c.Password == "" {
		http.Error(w, "Missing email or password", http.StatusBadRequest)
		return
	}

	token, err := h.svc.Login(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	api.JSON(w, http.StatusOK, token)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
