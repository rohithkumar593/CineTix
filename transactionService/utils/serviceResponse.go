package utils

import (
	"cine-tickets/responses"
	"net/http"
)

func RepositoryResponseLayer(response any, err error) *responses.Response {
	data := new(responses.Response)
	data.Body = response
	if err != nil {
		data.Error = err.Error()
	}
	return data
}

func GetResponseFormatter(r *http.Request) *responses.ResponseFormat {
	responseFormat := responses.ResponseFormat{
		Request: r,
	}
	return &responseFormat
}
