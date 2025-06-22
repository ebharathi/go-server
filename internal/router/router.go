package router

import (
	"net/http"

	"server/internal/handler"
	"server/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	router := chi.NewRouter()

	//public routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.RequestLogger)
		r.Get("/", handler.HomeHandler)
		r.Post("/users", handler.CreateUser)
		r.Post("/login", handler.LoginUser)
	})

	//protected routes
	router.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Use(middleware.RequestLogger)
		protected.Get("/me", handler.GetMe)
	})

	return router
}
