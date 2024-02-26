package middlewares

import (
	"cine-tickets/responses"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeMiddlewares(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(someHandler)
}

func someHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			responseStruct := r.Context().Value(responses.ResponseKey("response")).(responses.Response)
			WriteResponse(responseStruct, w)

		}()
		next.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}

func WriteResponse(response responses.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)

}
