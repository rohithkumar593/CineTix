package FormattersIO

import (
	"cine-tickets/models"
	"encoding/json"
	"log"
)

func ReservationServiceIOFormatBookTix(data []byte) (*models.ReserveTix, error) {
	var ticket *models.ReserveTix = new(models.ReserveTix)
	err := json.Unmarshal(data, ticket)
	if err != nil {
		log.Println("Error while constructing model", err)
		return nil, err
	} else {
		err = ticket.Validate()
		if err != nil {
			log.Println("Error while validating model", err)
			return nil, err
		} else {
			return ticket, nil
		}
	}
}

func ReservationServiceIOFormatUserInfo(data []byte) (*models.UserInfo, error) {
	var user models.UserInfo
	err := json.Unmarshal(data, &user)
	if err != nil {
		log.Println("Error while constructing model", err)
		return nil, err
	} else {
		err = user.Validate()
		if err != nil {
			log.Println("Error while validating model", err)
			return nil, err
		} else {
			return &user, nil
		}
	}
}
