package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", addExtraRoutes)
	r.Route("/auth", addAuthRoutes)
	r.Route("/users", addUserRoutes)
	r.Route("/movies", addMovieRoutes)
	r.Route("/shows", addShowtimeRoutes)

	return r
}
