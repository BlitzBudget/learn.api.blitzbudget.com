package models

type HttpResponse struct {
	Items *[]DBItem `json:"items"`
}
