package repository

import (
	"cine-tickets/models"
	dataAccessLayer "cine-tickets/persistence_layer"
	"log"
)

type ReservationRepository struct {
}

// first param : try -> model which model you are going to use

func (reservationRepo *ReservationRepository) StoreIntoTableHoldTix(ticket *models.ReserveTix) error {

	// Step 1. Get Postgres Gorm
	// Step 2. Insert Model into Postgres
	// Step 3 . Return Response

	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		log.Fatal("Error while acquiring postgres db", err)
		return err
	}

	postgresClient := dataAccessLayer.GetPostgresClient()

	insertedID := postgresClient.Create(ticket)
	if insertedID.Error != nil {
		return insertedID.Error
	}

	return nil

}

// first param : table name -> try using this
func (reservationRepo *ReservationRepository) GetInformationByUserId(userInfo *models.UserInfo) (*models.ReserveTix, error) {

	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		log.Fatal("Error while acquiring postgres db", err)
		return nil, err
	}
	postgresClient := dataAccessLayer.GetPostgresClient()
	var reserveTix models.ReserveTix
	postgresClient.Where("user_id=?", userInfo.UserId).First(&reserveTix).Scan(&reserveTix)
	return &reserveTix, nil
	// Step 1. Get Postgres Gorm
	// Step 2. Find booking by user id
	// Step 3 . Return Response

}
