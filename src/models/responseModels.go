package models

type ValidatorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status int     `json:"status"`
	Res    Message `json:"res"`
}
type SuccessResponse struct {
	Status int `json:"status"`
	Res    any `json:"res"`
}
