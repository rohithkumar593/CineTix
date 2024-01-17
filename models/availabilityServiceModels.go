package models

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Seat struct {
	SeatId             int    `json:"seat_id"`
	SeatRepresentation string `json:"seat_representation" `
	SeatPrice          int    `json:"seat_price" `
	SeatType           int    `json:"seat_type"  `
	TheatreId          int    `json:"theatre_id"`
}

type GetSeatByTheatre struct {
	TheatreId   int    `json:"theatre_id" validate:"required"`
	BookingDate string `json:"booking_date"  gorm:"column:booking_date" validate:"timestampWithoutZone"`
	ShowTime    string `json:"show_time"  gorm:"column:show_time" validate:"len=5"`
}

type Theatre struct {
	Location    string `json:"location" validate:"required" gorm:"type:text"`
	TheatreId   int    `json:"theatre_id"`
	Name        string `json:"name" `
	Shows       string `json:"shows"`
	MovieId     int    `json:"movie" gorm:"column:movie" `
	MovieName   string `json:"movie_name" gorm:"column:name"`
	Languages   string `json:"languages"`
	ReleaseDate string `json:"release_date"`
}

func (theatre Theatre) Validate() error {
	validate := validator.New()
	err := validate.Struct(theatre)
	return err
}

func (seat Seat) Validate() error {
	validate := validator.New()
	err := validate.Struct(seat)
	return err
}

func (seatRequest GetSeatByTheatre) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("timestampWithoutZone", timeZoneWithoutTimestamp)
	if err != nil {
		log.Fatal("Error while registering validation", err)
		return err
	}
	err = validate.Struct(seatRequest)
	return err
}
