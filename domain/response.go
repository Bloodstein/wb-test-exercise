package domain

type ErrorResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Result   string    `json:"result"`
	Response *struct{} `json:"response"`
}
