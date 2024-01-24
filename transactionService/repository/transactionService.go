package repository

import (
	"cine-tickets/configs"
	"cine-tickets/models"
	dataAccessLayer "cine-tickets/persistence_layer"
	"cine-tickets/responses"
	"cine-tickets/utils"
)

type TransactionRepository struct {
}

func (transaction *TransactionRepository) UpdateTransaction(order *models.Transaction) *responses.ResponseFormat {
	var statusInformation models.StateTransaction
	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		return utils.RepositoryResponseLayer(nil, err)
	}
	postgresClient := dataAccessLayer.GetPostgresClient()
	status := postgresClient.Model(&models.ReserveTix{}).Where("transaction_id=?", order.TransactionId).Update("status", configs.AppConfig.Booking.Confirmed)
	if status.Error != nil {
		return utils.RepositoryResponseLayer("Internal Server Error", status.Error)
	}
	if status.RowsAffected == 1 {
		statusInformation.State = "Confirmed"
	}
	return utils.RepositoryResponseLayer(statusInformation, status.Error)
}
