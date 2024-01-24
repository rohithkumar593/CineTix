package utils

import (
	"cine-tickets/responses"
)

func RepositoryResponseLayer(response any, error error) *responses.ResponseFormat {
	data := new(responses.ResponseFormat)
	data.Body = response
	data.Error = error
	return data
}
