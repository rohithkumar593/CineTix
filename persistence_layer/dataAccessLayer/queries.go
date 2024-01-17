package queryLayer

const (
	ReservationCheckQuery = `SELECT seats_booked from holdTix where theatre_id =$1 and movie_id=$2 and $3=ANY(seats_booked)`
	BookReservationQuery  = `INSERT INTO holdTix(movie_id,theatre_id,booking_date,seats_booked,show_time,user_id,status) Values($1,$2,$3,$4,$5,$6,$7)`
	HoldTix               = ` 	CREATE EXTENSION btree_gist;
								CREATE EXTENSION IF NOT EXISTS intarray;
								CREATE TABLE HOLDTIX (
								tix_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
								movie_id INT references movies ON DELETE NO ACTION NOT NULL,
								theatre_id INT references theatres ON DELETE NO ACTION NOT NULL,
								booking_date timestamp NOT NULL,
								seats_booked int[] NOT NULL,
								show_time varchar(5) NOT NULL,
								user_id INT NOT NULL,
								status SMALLINT NOT NULL,
								booking_timestamp timestamp not null,
								EXCLUDE USING GIST(booking_date with =, movie_id with =, theatre_id with =, show_time with =, seats_booked with &&)
								);
							`
	Theatres 			 = `CREATE TABLE theatres (
									theatre_id INT     PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
									name      VARCHAR(50) NOT NULL,
									shows     text[]            NOT NULL,
									location  VARCHAR(50) NOT NULL,
									movie INT REFERENCES Movies ON DELETE CASCADE
							);`
	Movies 				 = `
								CREATE TABLE movies (
								
								movie_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
								
								name VARCHAR ( 50 ) NOT NULL,
								
								languages text[] NOT NULL,
								
								release_date TIMESTAMP NOT NULL
								);`
	Seats 				 = `
							CREATE TABLE seats (
							
							seat_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
							
							seat_type SMALLINT DEFAULT 1,
							
							seat_price INT NOT NULL,
							
							theatre_id INT REFERENCES theatres ON DELETE CASCADE
							
							);`
)
