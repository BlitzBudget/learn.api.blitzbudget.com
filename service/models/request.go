package models

type Request struct {
	URL *string `validate:"required" json:"url"`
}
