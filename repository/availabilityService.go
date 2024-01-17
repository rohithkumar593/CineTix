package repository

import (
	"cine-tickets/configs"
	"cine-tickets/models"
	dataAccessLayer "cine-tickets/persistence_layer"
	"log"
)

type AvailabilityRepository struct {
}

func (availabilityRepo *AvailabilityRepository) GetSeatByTheatreId(seat *models.GetSeatByTheatre) ([]models.Seat, error) {
	var availableSeats []models.Seat
	var seatCheck []int64
	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		log.Fatal("Error while acquiring postgres db", err)
		return nil, err
	}
	postgresClient := dataAccessLayer.GetPostgresClient()
	bookedSeats := postgresClient.Table("holdtix").Select("unnest(seats_booked)").Where("booking_date=? and show_time=? and (localtimestamp-booking_timestamp < ? * INTERVAL '1 minutes' OR status=?)", seat.BookingDate, seat.ShowTime, configs.AppConfig.Booking.HoldTime, configs.AppConfig.Booking.Confirmed).Scan(&seatCheck)
	if bookedSeats.Error != nil {
		return nil, bookedSeats.Error
	}
	data := postgresClient.Model(&models.Seat{}).Where("theatre_id=? and seat_id not in ?", seat.TheatreId, seatCheck).Scan(&availableSeats)
	return availableSeats, data.Error
}

func (availabilityRepo *AvailabilityRepository) GetTheatresAndMoviesByLocation(theatre *models.Theatre) ([]models.Theatre, error) {
	var Theatres []models.Theatre
	err := dataAccessLayer.GetDbByName("postgres")
	if err != nil {
		log.Fatal("Error while acquiring postgres db", err)
		return nil, err
	}
	postgresClient := dataAccessLayer.GetPostgresClient()
	data := postgresClient.Model(&models.Theatre{}).Select("theatres.location, theatres.theatre_id,theatres.name,theatres.shows,theatres.movie,movies.name,movies.languages,movies.release_date").Joins("inner join movies on theatres.movie=movies.movie_id").Where("theatres.location=?", theatre.Location).Scan(&Theatres)
	return Theatres, data.Error
}
