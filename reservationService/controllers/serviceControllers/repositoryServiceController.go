package serviceController

import (
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
	"cine-tickets/repository"
	"cine-tickets/services"
)

func ReservationServiceController() serviceInterface.ReservationService {
	reservationService := services.ReservationService{}
	reservationRepository := repository.ReservationRepository{}
	reservationService.ReservationRepo = &reservationRepository
	return &reservationService
}


