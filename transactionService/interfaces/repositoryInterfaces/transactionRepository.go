package repositoryInterface

import (
	"cine-tickets/models"
	"cine-tickets/responses"
)

type TransactionRepository interface {
	UpdateTransaction(*models.Transaction) *responses.ResponseFormat
}
