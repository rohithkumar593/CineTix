package models

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type SeatBooked struct {
	SeatNumber int
}

type ReserveTix struct {
	UserId           int           `json:"user_id" gorm:"column:user_id" validate:"required"`
	MovieId          int           `json:"movie_id" gorm:"column:movie_id" validate:"required"`
	TheatreId        int           `json:"theatre_id"  gorm:"column:theatre_id" validate:"required"`
	BookingDate      string        `json:"booking_date"  gorm:"column:booking_date" validate:"timestampWithoutZone"`
	SeatsBooked      pq.Int64Array `json:"seats_booked" gorm:"type:integer[]" gorm:"column:seats_booked" validate:"required" `
	ShowTime         string        `json:"show_time"  gorm:"column:show_time" validate:"len=5"`
	Status           int           `json:"status"  gorm:"column:status"`
	BookingTimestamp time.Time     `json:"booking_timestamp" gorm:"autoCreateTime:true"  gorm:"column:booking_timestamp" `
}

type UserInfo struct {
	UserId int `json:"user_id" gorm:"column:user_id" validate:"required"`
}

func (user UserInfo) Validate() error {
	validate := validator.New()
	err := validate.Struct(user)
	return err
}

func (ticket ReserveTix) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("timestampWithoutZone", timeZoneWithoutTimestamp)
	if err != nil {
		log.Fatal("Error while registering validation", err)
		return err
	}
	err = validate.Struct(ticket)
	return err
}

func (ReserveTix) TableName() string {
	return "holdtix"
}

func timeZoneWithoutTimestamp(fl validator.FieldLevel) bool {
	layout := "2006-01-02 15:04:05"
	data := fl.Field().String()
	if _, err := time.Parse(layout, data); err != nil {
		log.Println(err, "error on timezone")
		return false
	}
	return true
}
