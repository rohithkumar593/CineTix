package responses

import (
	"context"
	"net/http"
)

type ResponseKey string
type Response struct {
	Body       interface{} `json:"body"`
	Error      string      `json:"message,omitempty"`
	StatusCode int         `json:"status_code"`
}

type ResponseFormat struct {
	Request *http.Request
}

func (r *ResponseFormat) setResponse(bag Response) {
	*r.Request = *r.Request.WithContext(context.WithValue(r.Request.Context(), ResponseKey("response"), bag))
}

func (response *ResponseFormat) WithOkResult(data *Response) {
	data.StatusCode = http.StatusOK
	response.setResponse(*data)

}

func (response *ResponseFormat) WithUnprocessableEntity(data string) {
	res := Response{
		StatusCode: http.StatusUnprocessableEntity,
		Body:       "Internal Server Error",
		Error:      data,
	}
	response.setResponse(res)

}

func (response *ResponseFormat) WithBadRequest(data string) {
	res := Response{
		StatusCode: http.StatusBadRequest,
		Body:       "Internal Server Error",
		Error:      data,
	}
	response.setResponse(res)

}
