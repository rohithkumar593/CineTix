package repositoryInterface

import (
	"cine-tickets/models"
)

type TransactionRepository interface {
	UpdateTransaction(*models.Transaction) (models.StateTransaction, error)
}
