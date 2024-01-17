package services

import (
	repositoryInterface "cine-tickets/interfaces/repositoryInterfaces"
	FormattersIO "cine-tickets/models/inputFormat"
	"encoding/json"
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
		fmt.Println("Error while reading request body", err)
	}
	seat, err := FormattersIO.AvailabilityServiceSeatsFormatIO(data)
	if err != nil {
		fmt.Println("Error while formatting input body", err)
	}
	seats, err := availabilityService.AvailabilityRepository.GetSeatByTheatreId(seat)
	if err != nil {
		fmt.Println("Error while getting seats by theatre id", err)
	}
	byteArray, err := json.Marshal(seats)
	if err != nil {
		fmt.Println(err, "error while formatting response")
		return
	}
	res.Write(byteArray)
}

func (availabilityService *AvailabilityService) DisplayMoviesAndTheatres(res http.ResponseWriter, req *http.Request) {

	// format input

	data, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error while reading request body", err)
	}

	theatre, err := FormattersIO.AvailabilityServiceTheatresFormatIO(data)
	if err != nil {
		fmt.Println("Error while formatting input", err)
	}
	theatres, err := availabilityService.AvailabilityRepository.GetTheatresAndMoviesByLocation(theatre)
	if err != nil {
		fmt.Println("Error while getting theaters and movies by location", err)
		return
	}
	byteArray, err := json.Marshal(theatres)
	if err != nil {
		fmt.Println(err, "error while formatting response")
		return
	}
	res.Write(byteArray)
}

