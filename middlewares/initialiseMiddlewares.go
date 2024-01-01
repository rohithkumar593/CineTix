package middlewares

import (
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

		func() {
			w.Header().Set("Content-Type", "application/json")
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
