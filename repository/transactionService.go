package repository

import (
	"cine-tickets/configs"
	"cine-tickets/models"
	dataAccessLayer "cine-tickets/persistence_layer"
	"log"
)

type TransactionRepository struct {
}

func (transaction *TransactionRepository) UpdateTransaction(order *models.Transaction) (*models.StateTransaction, error) {
	var statusInformation models.StateTransaction
	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		log.Fatal("Error while acquiring postgres db", err)
		return &statusInformation, err
	}
	postgresClient := dataAccessLayer.GetPostgresClient()
	status := postgresClient.Model(&models.ReserveTix{}).Where("transaction_id=?", order.TransactionId).Update("status", configs.AppConfig.Booking.Confirmed)
	if status.Error != nil {
		return &statusInformation, status.Error
	}
	if status.RowsAffected == 1 {
		statusInformation.State = "Confirmed"
	}
	return &statusInformation, nil
}
