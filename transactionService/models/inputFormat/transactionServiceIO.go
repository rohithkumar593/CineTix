package FormattersIO

import (
	"cine-tickets/models"
	"encoding/json"
	"log"
)

func TransactionIDFormatIO(data []byte) (*models.Transaction, error) {
	var order models.Transaction
	err := json.Unmarshal(data, &order)
	if err != nil {
		log.Println("Error while constructing model", err)
		return nil, err
	} else {
		err = order.Validate()
		if err != nil {
			log.Println("Error while validating model", err)
			return nil, err
		} else {
			return &order, nil
		}
	}
}
