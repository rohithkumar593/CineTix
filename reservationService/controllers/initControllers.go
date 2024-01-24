package controllers

import (
	serviceController "cine-tickets/controllers/serviceControllers"
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
)

type ServiceControllers struct {
	ReservationService serviceInterface.ReservationService
}

func InitServiceControllers() *ServiceControllers {
	services := ServiceControllers{}
	services.ReservationService = serviceController.ReservationServiceController()
	return &services
}
