package repositoryInterface

import (
	"cine-tickets/models"
)

type AvailabilityRepository interface {
	GetSeatByTheatreId(*models.GetSeatByTheatre) ([]models.Seat, error)
	GetTheatresAndMoviesByLocation(*models.Theatre) ([]models.Theatre, error)
}
