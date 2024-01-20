package services

import (
	FormattersIO "cine-tickets/models/inputFormat"
	"cine-tickets/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type TransactionService struct {
	TransactionRepo *repository.TransactionRepository
}

func (transaction *TransactionService) AcceptTransaction(res http.ResponseWriter, req *http.Request) {
	// format io
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err, "error while reading input body")
		return
	}

	order, err := FormattersIO.TransactionIDFormatIO(body)
	if err != nil {
		log.Fatal(err, "error while formatting input body")
		return
	}

	stateOfTransaction := transaction.TransactionRepo.UpdateTransaction(order)

	if err != nil {
		log.Println(err)
		return
	}

	bytesString, err := json.Marshal(stateOfTransaction)
	if err != nil {
		log.Println(err, "error while formatting response")
		return
	}
	res.Write(bytesString)

	// incase of any error ask to rollback payment for user, by sending refund initiated status

	

}
