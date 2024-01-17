package serviceController

import (
	serviceInterface "cine-tickets/interfaces/serviceInterfaces"
	"cine-tickets/repository"
	"cine-tickets/services"
)

func TransactionServiceController() serviceInterface.TransactionService {
	transactionService := services.TransactionService{}
	transactionRepository := repository.TransactionRepository{}
	transactionService.TransactionRepo = &transactionRepository
	return &transactionService
}
