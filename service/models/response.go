package models

type HttpResponse struct {
	Message *string `validate:"required" json:"message"`
}
