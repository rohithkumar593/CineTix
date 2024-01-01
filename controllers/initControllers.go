package controllers

import (
	serviceController "cine-tickets/controllers/serviceControllers"
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
)

func InitController() serviceInterface.ReservationService {

	return serviceController.ReservationServiceController()
}
