package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/V4T54L/movie-reservation-system/internals/store"
	"github.com/go-chi/chi/v5"
)

func addExtraRoutes(r chi.Router) {
	r.Get("/", home)
	r.Get("/health", health)
}

func home(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "home/landing page"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func health(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["server"] = "healthy"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	if err := store.GetPostgresStore().Health(ctx); err != nil {
		resp["database"] = "error : " + err.Error()
	} else {
		resp["database"] = "healthy"
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
