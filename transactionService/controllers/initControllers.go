package controllers

import (
	serviceController "cine-tickets/controllers/serviceControllers"
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
)

type ServiceControllers struct {
	TransactionService serviceInterface.TransactionService
}

func InitServiceControllers() *ServiceControllers {
	services := ServiceControllers{}
	services.TransactionService = serviceController.TransactionServiceController()
	return &services
}
