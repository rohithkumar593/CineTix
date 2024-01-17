package FormattersIO

import (
	"cine-tickets/models"
	"encoding/json"
	"log"
)

func AvailabilityServiceSeatsFormatIO(data []byte) (*models.GetSeatByTheatre, error) {
	var seat models.GetSeatByTheatre
	err := json.Unmarshal(data, &seat)
	if err != nil {
		log.Println("Error while constructing model", err)
		return nil, err
	} else {
		err = seat.Validate()
		if err != nil {
			log.Println("Error while validating model", err)
			return nil, err
		} else {
			return &seat, nil
		}
	}
}

func AvailabilityServiceTheatresFormatIO(data []byte) (*models.Theatre, error) {
	var theatre models.Theatre
	err := json.Unmarshal(data, &theatre)
	if err != nil {
		log.Println("Error while constructing model", err)
		return nil, err
	} else {
		err = theatre.Validate()
		if err != nil {
			log.Println("Error while validating model", err)
			return nil, err
		} else {
			return &theatre, nil
		}
	}
}
