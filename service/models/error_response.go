package models

type ErrorHttpResponse struct {
	Message *string `validate:"required" json:"message"`
	Code    *int    `validate:"required" json:"error_code"`
}
