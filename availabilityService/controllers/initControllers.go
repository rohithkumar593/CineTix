package controllers

import (
	serviceController "cine-tickets/controllers/serviceControllers"
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
)

type ServiceControllers struct {
	AvailabilityService serviceInterface.AvailabilityService
}

func InitServiceControllers() *ServiceControllers {
	services := ServiceControllers{}
	services.AvailabilityService = serviceController.AvailabilityServiceController()
	return &services
}
