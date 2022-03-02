package domain

type ErrorResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}
