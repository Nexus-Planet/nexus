package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"
)

func NewRouter(pattern string) chi.Router {
	r := chi.NewRouter()

	if pattern == "" {

		return r.Route("/api", func(r chi.Router) {})
	}
	return r.Route(fmt.Sprintf("/api%s", pattern), func(r chi.Router) {})

}
