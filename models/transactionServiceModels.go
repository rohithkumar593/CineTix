package models

import (
	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	TransactionId string `json:"transaction_id" validate:"required"`
}

type StateTransaction struct {
	State string `json:"state"`
}

func (order Transaction) Validate() error {
	validate := validator.New()
	err := validate.Struct(order)
	return err
}
