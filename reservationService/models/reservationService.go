package models

import (
	"cine-tickets/configs"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type SeatBooked struct {
	SeatNumber int
}
type PaymentUrl struct {
	Url string `json:"payment_url"`
}

type ReserveTix struct {
	TransactionId string        `json:"tix_id" default:"3"`
	UserId        int           `json:"user_id" gorm:"column:user_id" validate:"required"`
	MovieId       int           `json:"movie_id" gorm:"column:movie_id" validate:"required"`
	TheatreId     int           `json:"theatre_id"  gorm:"column:theatre_id" validate:"required"`
	BookingDate   string        `json:"booking_date"  gorm:"column:booking_date" validate:"timestampWithoutZone"`
	SeatsBooked   pq.Int64Array `json:"seats_booked" gorm:"type:integer[]" gorm:"column:seats_booked" validate:"required" `
	ShowTime      string        `json:"show_time"  gorm:"column:show_time" validate:"len=5"`
	Status        int           `json:"status"  gorm:"column:status"`
	CreatedAt     time.Time     `json:"booking_timestamp"  gorm:"column:booking_timestamp" `
}

type UserInfo struct {
	UserId int `json:"user_id" gorm:"column:user_id" validate:"required"`
}

func (ReserveTix) TableName() string {
	return "holdtix"
}

func (ticket *ReserveTix) BeforeCreate(tx *gorm.DB) (err error) {
	ticket.TransactionId = uuid.New().String()
	ticket.CreatedAt = time.Now().UTC()
	res := tx.Delete(&ReserveTix{}, "status = ? and localtimestamp-booking_timestamp > ? * INTERVAL '1 minutes'", configs.AppConfig.Booking.Initialised, configs.AppConfig.Booking.HoldTime)
	log.Println(res)
	return
}

func (user UserInfo) Validate() error {
	validate := validator.New()
	err := validate.Struct(user)
	return err
}

func (ticket *ReserveTix) Validate() error {
	validate := validator.New()

	err := validate.RegisterValidation("timestampWithoutZone", timeZoneWithoutTimestamp)
	if err != nil {
		log.Fatal("Error while registering validation", err)
		return err
	}
	err = validate.Struct(ticket)
	return err
}
