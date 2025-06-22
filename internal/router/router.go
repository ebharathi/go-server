package router

import (
	"net/http"

	"server/internal/handler"
	"server/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	root := chi.NewRouter()

	root.Route("/api/v1", func(router chi.Router) {

		// Public routes
		router.Group(func(r chi.Router) {
			r.Use(middleware.RequestLogger)
			r.Get("/", handler.HomeHandler)
			r.Post("/users", handler.CreateUser)
			r.Post("/login", handler.LoginUser)
			r.Get("/auth/google/callback", handler.GoogleCallback)
		})

		// Protected routes
		router.Group(func(protected chi.Router) {
			protected.Use(middleware.AuthMiddleware)
			protected.Use(middleware.RequestLogger)
			protected.Get("/me", handler.GetMe)
		})
	})

	return root
}
