package serviceInterface

import (
	"net/http"
)

type AvailabilityService interface {
	DisplaySeats(res http.ResponseWriter, req *http.Request)
	DisplayMoviesAndTheatres(res http.ResponseWriter, req *http.Request)
}
