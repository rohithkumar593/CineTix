package repositoryInterface

import (
	"cine-tickets/models"
)

type ReservationRepository interface {
	StoreIntoTableHoldTix(*models.ReserveTix) error
	GetInformationByUserId(*models.UserInfo) (*models.ReserveTix, error) // takes user model
}
