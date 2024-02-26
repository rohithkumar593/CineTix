package repository

import (
	"cine-tickets/models"
	dataAccessLayer "cine-tickets/persistence_layer"
	"cine-tickets/responses"
	"cine-tickets/utils"
	"log"
)

type ReservationRepository struct {
}

func (reservationRepo *ReservationRepository) StoreIntoTableHoldTix(ticket *models.ReserveTix) *responses.Response {

	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		log.Fatal("Error while acquiring postgres db", err)
		return utils.RepositoryResponseLayer(nil, err)
	}

	postgresClient := dataAccessLayer.GetPostgresClient()

	insertedID := postgresClient.Create(ticket)
	if insertedID.Error != nil {
		return utils.RepositoryResponseLayer(nil, insertedID.Error)
	}
	return utils.RepositoryResponseLayer(ticket.TransactionId, nil)

}

func (reservationRepo *ReservationRepository) GetInformationByUserId(userInfo *models.UserInfo) *responses.Response {

	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		log.Println("Error while acquiring postgres db", err)
		return utils.RepositoryResponseLayer(nil, err)
	}
	postgresClient := dataAccessLayer.GetPostgresClient()
	var reserveTix models.ReserveTix
	data := postgresClient.Where("user_id=?", userInfo.UserId).First(&reserveTix).Scan(&reserveTix)
	if data.Error != nil {
		return utils.RepositoryResponseLayer(nil, data.Error)
	}
	return utils.RepositoryResponseLayer(reserveTix, nil)
}
