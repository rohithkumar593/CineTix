package responses

type ResponseFormat struct {
	Body  interface{} `json:"body,omitempty"`
	Error error       `json:"message,omitempty"`
}
