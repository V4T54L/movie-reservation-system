package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/V4T54L/movie-reservation-system/internals/middlewares"
	"github.com/V4T54L/movie-reservation-system/internals/schemas"
	"github.com/V4T54L/movie-reservation-system/internals/store"
	"github.com/V4T54L/movie-reservation-system/internals/utils"
	"github.com/go-chi/chi/v5"
)

func addMovieRoutes(r chi.Router) {
	r.Get("/", getMovies)
	r.Group(func(sr chi.Router) {
		sr.Use(middlewares.AuthMiddleware, middlewares.AdminOnlyMiddleware)
		sr.Post("/", addMovie)
	})
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	movies, err := store.GetPostgresStore().GetMovieDetails(ctx)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, movies)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	movie := schemas.AddMovie{}
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "error parsing the body : "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	err := store.GetPostgresStore().AddMovie(ctx, movie)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MessageResponse(w, http.StatusOK, "movie created successfully")
}
