package controllers

import (
	serviceController "cine-tickets/controllers/serviceControllers"
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
)

type ServiceControllers struct {
	ReservationService  serviceInterface.ReservationService
	AvailabilityService serviceInterface.AvailabilityService
	TransactionService  serviceInterface.TransactionService
}

func InitServiceControllers() *ServiceControllers {
	services := ServiceControllers{}
	services.ReservationService = serviceController.ReservationServiceController()
	services.AvailabilityService = serviceController.AvailabilityServiceController()
	services.TransactionService = serviceController.TransactionServiceController()
	return &services
}
