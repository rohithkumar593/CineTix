package services

import (
	repositoryInterface "cine-tickets/interfaces/repositoryInterfaces"
	FormattersIO "cine-tickets/models/inputFormat"
	"cine-tickets/utils"
	"fmt"
	"io"
	"net/http"
)

type AvailabilityService struct {
	AvailabilityRepository repositoryInterface.AvailabilityRepository
}

func (availabilityService *AvailabilityService) DisplaySeats(res http.ResponseWriter, req *http.Request) {

	data, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error while reading request body", err) // wrap res with 422 result
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}
	seat, err := FormattersIO.AvailabilityServiceSeatsFormatIO(data)
	if err != nil {
		fmt.Println("Error while formatting input body", err) // wrap res with 422 result
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return

	}
	seats := availabilityService.AvailabilityRepository.GetSeatByTheatreId(seat)
	utils.GetResponseFormatter(req).WithOkResult(seats)

}

func (availabilityService *AvailabilityService) DisplayMoviesAndTheatres(res http.ResponseWriter, req *http.Request) {

	// format input

	data, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error while reading request body", err)
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}

	theatre, err := FormattersIO.AvailabilityServiceTheatresFormatIO(data)
	if err != nil {
		fmt.Println("Error while formatting input", err)
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}
	theatres := availabilityService.AvailabilityRepository.GetTheatresAndMoviesByLocation(theatre)
	utils.GetResponseFormatter(req).WithOkResult(theatres)
}
