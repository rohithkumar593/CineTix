package queryLayer

const (
	ReservationCheckQuery = `SELECT seats_booked from holdTix where theatre_id =$1 and movie_id=$2 and $3=ANY(seats_booked)`
	BookReservationQuery  = `INSERT INTO holdTix(movie_id,theatre_id,booking_date,seats_booked,show_time,user_id,status) Values($1,$2,$3,$4,$5,$6,$7)`
)

// func CheckIfReservationExists(db interfaces.DatabaseAccessObject, args []any) bool {

// 	row := db.QueryRow(ReservationCheckQuery, args)
// 	var seat models.SeatBooked
// 	fmt.Println(row, "row reservation")
// 	err := row.Scan(&seat.SeatNumber)
// 	if err != nil {
// 		fmt.Println("Reservation does not exist", err)
// 		return true
// 	}
// 	fmt.Println("Reservation exists", seat.SeatNumber)
// 	return false
// }

// func MakeReservation(reserveTix *models.ReserveTix) (bool, error) {

// 	result, err := db.QueryDB(BookReservationQuery, []any{reserveTix.MovieId, reserveTix.TheatreId, reserveTix.BookingDate, pq.Array(reserveTix.SeatsBooked), reserveTix.ShowTime, reserveTix.UserId, reserveTix.Status})
// 	fmt.Println("Reservation response", result, err)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
