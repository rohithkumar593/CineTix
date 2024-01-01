package routes

import (
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"

	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(router *chi.Mux, serviceController serviceInterface.ReservationService) {
	router.Post("/bookTix", serviceController.BookTix)
	router.Post("/displayBooking", serviceController.DisplayBooking)
}
