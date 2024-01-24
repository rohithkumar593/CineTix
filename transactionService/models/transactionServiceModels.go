package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type Transaction struct {
	TransactionId string `json:"transaction_id" validate:"required"`
}

type StateTransaction struct {
	State string `json:"state"`
}

func (order Transaction) Validate() error {
	validate := validator.New()
	err := validate.Struct(order)
	return err
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
