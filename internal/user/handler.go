package user

import "net/http"

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) FindOneUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeactivateUser(w http.ResponseWriter, r *http.Request) {

}
