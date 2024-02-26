package services

import (
	"cine-tickets/configs"
	"cine-tickets/models"
	FormattersIO "cine-tickets/models/inputFormat"
	"cine-tickets/repository"
	"cine-tickets/utils"
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
		log.Println(err, "error while reading input body")
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}
	reserveTix, err = FormattersIO.ReservationServiceIOFormatBookTix(body)
	if err != nil {
		log.Println(err, "error while formatting input body")
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}

	transactionResponse := reservation.ReservationRepo.StoreIntoTableHoldTix(reserveTix)
	if transactionResponse.Error != "" {
		log.Println(err)
		utils.GetResponseFormatter(req).WithUnprocessableEntity(transactionResponse.Error)
		return
	}
	//generate payment url with call back
	callBackUrl := "http://localhost" + configs.AppConfig.Server.Host + configs.AppConfig.Server.Port + "/transaction"
	url, err := reservation.GeneratePaymentUrlWithTransaction(transactionResponse.Body, callBackUrl)
	if err != nil {
		log.Println(err, url)
		utils.GetResponseFormatter(req).WithBadRequest(err.Error())
		return
	}

	utils.GetResponseFormatter(req).WithOkResult(utils.RepositoryResponseLayer(transactionResponse, nil))
}

func (displayInfo *ReservationService) DisplayBooking(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err, "error while reading input body")
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}

	user, err := FormattersIO.ReservationServiceIOFormatUserInfo(body)
	if err != nil {
		log.Fatal(err, "error while formatting input body")
		utils.GetResponseFormatter(req).WithUnprocessableEntity(err.Error())
		return
	}
	bookingInfo := displayInfo.ReservationRepo.GetInformationByUserId(user)
	utils.GetResponseFormatter(req).WithOkResult(bookingInfo)

}

func (reservation *ReservationService) GeneratePaymentUrlWithTransaction(transaction any, callBackUrl string) (url *models.PaymentUrl, err error) {
	var payment models.PaymentUrl
	//make post request to payment gateway with transactionId
	payment.Url = "paymentUrl.com/paytome"

	return &payment, nil
}
