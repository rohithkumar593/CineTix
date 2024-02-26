package models

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

func timeZoneWithoutTimestamp(fl validator.FieldLevel) bool {
	layout := "2006-01-02"
	data := fl.Field().String()
	if _, err := time.Parse(layout, data); err != nil {
		log.Println(err, "error on timezone")
		return false
	}
	return true
}


