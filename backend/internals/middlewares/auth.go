package middlewares

import (
	"context"
	"net/http"

	"github.com/V4T54L/movie-reservation-system/internals/schemas"
	"github.com/V4T54L/movie-reservation-system/internals/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Auth")
		if token == "" {
			utils.ErrorResponse(rw, http.StatusUnauthorized, "token not found")
			return
		}
		user, valid := utils.GetUserFromToken(token)
		if !valid {
			utils.ErrorResponse(rw, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), schemas.CurrentUser{}, user)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(schemas.CurrentUser{}).(schemas.UserToken)
		if user.Role != "admin" {
			utils.ErrorResponse(rw, http.StatusUnauthorized, "endpoint access is restricted")
			return
		}
		next.ServeHTTP(rw, r)
	})
}
