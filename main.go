package main

import (
	"cine-tickets/controllers"
	"cine-tickets/middlewares"
	routes "cine-tickets/routers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()
	middlewares.InitializeMiddlewares(router)
	log.Println("Set Middlewares : Done")
	serviceController := controllers.InitController()
	log.Println("Set ServiceController : Done")
	routes.InitializeRoutes(router, serviceController)

	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Issue while listening to server at port 5000")
	}
}
