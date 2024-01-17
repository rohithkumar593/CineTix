package main

import (
	"cine-tickets/configs"
	"cine-tickets/controllers"
	"cine-tickets/middlewares"
	routes "cine-tickets/routers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.InitAppConfig()
	if err != nil {
		log.Fatal("unable to initialize app config", err)
	}
	log.Println("Set App Configs : Done")
	router := chi.NewRouter()
	middlewares.InitializeMiddlewares(router)
	log.Println("Set Middlewares : Done")
	serviceController := controllers.InitServiceControllers()
	log.Println("Set ServiceController : Done")
	routes.InitializeRoutes(router, serviceController)

	server := http.Server{
		Addr:    configs.AppConfig.Server.Host + configs.AppConfig.Server.Port,
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Issue while listening to server at port 5000")
	}
}
