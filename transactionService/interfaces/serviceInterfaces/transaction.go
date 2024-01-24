package serviceInterface

import (
	"net/http"
)

type TransactionService interface {
	AcceptTransaction(res http.ResponseWriter, req *http.Request)
}
