package models

type DBItem struct {
	PK           string `json:"pk"`
	SK           string `json:"sk"`
	Author       string `json:"Author"`
	Category     string `json:"Category"`
	CreationDate string `json:"creation_date"`
	File         string `json:"File"`
	Name         string `json:"Name"`
	Tags         string `json:"Tags"`
}
