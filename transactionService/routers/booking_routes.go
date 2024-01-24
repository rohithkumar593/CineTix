package routes

import (
	"cine-tickets/controllers"

	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(router *chi.Mux, serviceController *controllers.ServiceControllers) {
	router.Post("/transaction", serviceController.TransactionService.AcceptTransaction)
}
