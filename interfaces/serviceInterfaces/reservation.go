package serviceInterface

import (
	"net/http"
)

type ReservationService interface {
	BookTix(res http.ResponseWriter, req *http.Request)
	DisplayBooking(res http.ResponseWriter, req *http.Request)
}
