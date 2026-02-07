package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
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
	err := sonic.ConfigDefault.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	user, err := h.svc.CreateSession(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	api.JSON(w, http.StatusCreated, user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var c Credentials
	err := sonic.ConfigDefault.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if c.Email == "" || c.Password == "" {
		http.Error(w, "Missing email or password", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)

	ctx := context.WithValue(r.Context(), TokenKey, token)

	respToken, err := h.svc.Login(ctx, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	api.JSON(w, http.StatusOK, respToken)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
