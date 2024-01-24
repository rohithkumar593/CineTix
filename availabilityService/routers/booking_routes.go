package routes

import (
	"cine-tickets/controllers"

	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(router *chi.Mux, serviceController *controllers.ServiceControllers) {
	router.Post("/displaySeats", serviceController.AvailabilityService.DisplaySeats)
	router.Post("/displayMoviesAndTheatres", serviceController.AvailabilityService.DisplayMoviesAndTheatres)

}
