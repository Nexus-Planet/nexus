package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/auth"
)

func V1Router(handler *auth.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.JwtToken))
			r.Use(jwtauth.Authenticator(auth.JwtToken))

			r.Post("/login", handler.Login)
		})

		r.Group(func(r chi.Router) {
			r.Post("/signup", handler.SignUp)
			r.Post("/login", handler.Login)
			r.Post("/logout", handler.Logout)
		})
	})
	return r
}
