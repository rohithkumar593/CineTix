package services

import (
	FormattersIO "cine-tickets/models/inputFormat"
	"cine-tickets/repository"
	"cine-tickets/utils"
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
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}

	order, err := FormattersIO.TransactionIDFormatIO(body)
	if err != nil {
		log.Fatal(err, "error while formatting input body")
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}

	stateOfTransaction := transaction.TransactionRepo.UpdateTransaction(order)

	if err != nil {
		log.Println(err)
		utils.GetResponseFormatter(req).WithBadRequest(err.Error())
		return
	}

	utils.GetResponseFormatter(req).WithOkResult(stateOfTransaction)

	// incase of any error ask to rollback payment for user, by sending refund initiated status

}
