package repositoryInterface

import (
	"cine-tickets/models"
	"cine-tickets/responses"
)

type AvailabilityRepository interface {
	GetSeatByTheatreId(*models.GetSeatByTheatre) *responses.Response
	GetTheatresAndMoviesByLocation(*models.Theatre) *responses.Response
}
