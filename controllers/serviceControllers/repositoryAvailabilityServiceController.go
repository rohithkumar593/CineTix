package serviceController

import (
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
	"cine-tickets/repository"
	"cine-tickets/services"
)

func AvailabilityServiceController() serviceInterface.AvailabilityService {
	availabilityService := services.AvailabilityService{}
	availabilityRepo := &repository.AvailabilityRepository{}
	availabilityService.AvailabilityRepository = availabilityRepo
	return &availabilityService
}
