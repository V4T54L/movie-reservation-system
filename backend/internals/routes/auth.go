package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/V4T54L/movie-reservation-system/internals/schemas"
	"github.com/V4T54L/movie-reservation-system/internals/store"
	"github.com/V4T54L/movie-reservation-system/internals/utils"
	"github.com/go-chi/chi/v5"
)

func addAuthRoutes(r chi.Router) {
	r.Post("/login", login)
	r.Post("/signup", signup)
}

func login(w http.ResponseWriter, r *http.Request) {
	creds := schemas.UserLogin{}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "error parsing the body : "+err.Error())
		return
	}

	creds.EncodedPass = utils.Hash(utils.DecodePassword(creds.EncodedPass))

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	token, err := store.GetPostgresStore().UserLogin(ctx, creds)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"access_token": token,
	})
}

func signup(w http.ResponseWriter, r *http.Request) {
	creds := schemas.UserSignup{}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "error parsing the body : "+err.Error())
		return
	}

	creds.EncodedPass = utils.Hash(utils.DecodePassword(creds.EncodedPass))

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	err := store.GetPostgresStore().UserSignup(ctx, creds)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MessageResponse(w, http.StatusOK, "user created")
}
