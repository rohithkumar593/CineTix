package repositoryInterface

import (
	"cine-tickets/models"
	"cine-tickets/responses"
)

type AvailabilityRepository interface {
	GetSeatByTheatreId(*models.GetSeatByTheatre) *responses.ResponseFormat
	GetTheatresAndMoviesByLocation(*models.Theatre) *responses.ResponseFormat
}
