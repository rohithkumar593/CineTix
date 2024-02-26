package customErrors

var errorMessage map[string]string = map[string]string{
	"DBError":          "Unable to connect to db",
	"FindSeatsError":   "Issue while checking availability, No seats might be available",
	"GenericError":     "InternalServerError",
	"FindTheatreError": "Issue while finding theatres in given location",
}

type CustomError struct {
	Msg string
}

func (custom CustomError) Error() string {
	return errorMessage[custom.Msg]
}
