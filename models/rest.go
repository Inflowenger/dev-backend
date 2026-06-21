package models


type Response struct{
	Data any `json:"data"`
	Error any `json:"error"`
}
type ErrorResponse struct{
	Message string `json:"message"`
	Code int `json:"code"`
}

