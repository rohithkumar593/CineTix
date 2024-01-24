package repositoryInterface

import (
	"cine-tickets/models"
	"cine-tickets/responses"
)

type ReservationRepository interface {
	StoreIntoTableHoldTix(*models.ReserveTix) *responses.ResponseFormat
	GetInformationByUserId(*models.UserInfo) *responses.ResponseFormat // takes user model
}
