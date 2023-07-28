package service

import (
	"encoding/json"
	"learn-api-blitzbudget-com/service/models"
	"log"
)

func ParserRequest(body *string) (*models.Request, error) {
	// Read the SK value from the request JSON
	var reqPayload *models.Request
	err := json.Unmarshal([]byte(*body), reqPayload)
	if err != nil {
		log.Println("Error parsing request JSON:", err)
		return nil, nil
	}

	return reqPayload, nil
}
