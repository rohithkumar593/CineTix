package services

import (
	"cine-tickets/models"
	FormattersIO "cine-tickets/models/inputFormat"
	"cine-tickets/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ReservationService struct {
	ReservationRepo *repository.ReservationRepository
}

func (reservation *ReservationService) BookTix(res http.ResponseWriter, req *http.Request) {
	var reserveTix *models.ReserveTix
	log.Println("started reservation")
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err, "error while reading input body")
		return
	}
	reserveTix, err = FormattersIO.ReservationServiceIOFormatBookTix(body)
	if err != nil {
		log.Fatal(err, "error while formatting input body")
		return
	}

	err = reservation.ReservationRepo.StoreIntoTableHoldTix(reserveTix)
	if err != nil {
		log.Println(err)
		return
	}
	bytesString, err := json.Marshal(reserveTix)
	if err != nil {
		log.Println(err, "error while formatting response")
		return
	}
	res.Write(bytesString)
}

func (displayInfo *ReservationService) DisplayBooking(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err, "error while reading input body")
		return
	}

	user, err := FormattersIO.ReservationServiceIOFormatUserInfo(body)
	if err != nil {
		log.Fatal(err, "error while formatting input body")
		return
	}
	bookingInfo, err := displayInfo.ReservationRepo.GetInformationByUserId(user)
	if err != nil {
		log.Println(err)
		return
	}
	bytesString, err := json.Marshal(bookingInfo)
	if err != nil {
		log.Println(err, "error while formatting response")
		return
	}
	res.Write(bytesString)

}
