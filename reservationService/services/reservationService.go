package services

import (
	"cine-tickets/configs"
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

	transaction_id, err := reservation.ReservationRepo.StoreIntoTableHoldTix(reserveTix)
	if err != nil {
		log.Println(err)
		return
	}
	//generate payment url with call back
	callBackUrl := "http://localhost" + configs.AppConfig.Server.Host + configs.AppConfig.Server.Port + "/transaction"
	url, err := reservation.GeneratePaymentUrlWithTransaction(transaction_id, callBackUrl)
	if err != nil {
		log.Println(err)
		return
	}

	bytesString, err := json.Marshal(url)
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
	bookingInfo := displayInfo.ReservationRepo.GetInformationByUserId(user)

	bytesString, err := json.Marshal(bookingInfo)
	if err != nil {
		log.Println(err, "error while formatting response")
		return
	}
	res.Write(bytesString)

}

func (reservation *ReservationService) GeneratePaymentUrlWithTransaction(transactionId string, callBackUrl string) (url *models.PaymentUrl, err error) {
	var payment models.PaymentUrl
	//make post request to payment gateway with transactionId
	payment.Url = "paymentUrl.com/paytome"

	return &payment, nil
}
