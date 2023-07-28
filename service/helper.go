package service

import (
	"encoding/json"
	"learn-api-blitzbudget-com/service/models"
	"log"
)

func ParserRequest(body *string) (*models.Request, error) {
	log.Println("Unmarshalling request")
	// Read the SK value from the request JSON
	var reqPayload models.Request
	err := json.Unmarshal([]byte(*body), &reqPayload)
	if err != nil {
		log.Println("Error parsing request JSON:", err)
		return nil, err
	}

	log.Println("Request parsed successfully: ", reqPayload)
	return &reqPayload, nil
}
