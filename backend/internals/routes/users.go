package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/V4T54L/movie-reservation-system/internals/middlewares"
	"github.com/V4T54L/movie-reservation-system/internals/schemas"
	"github.com/V4T54L/movie-reservation-system/internals/store"
	"github.com/V4T54L/movie-reservation-system/internals/utils"
	"github.com/go-chi/chi/v5"
)

func addUserRoutes(r chi.Router) {
	r.Group(func(sr chi.Router) {
		sr.Use(middlewares.AuthMiddleware)
		sr.Get("/me", myProfile)
	})
}

func myProfile(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(schemas.CurrentUser{}).(schemas.UserToken)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	userDetails, err := store.GetPostgresStore().GetUserDetails(ctx, currentUser.ID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, userDetails)
}
