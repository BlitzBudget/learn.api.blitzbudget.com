package service

import (
	dynamoDBService "learn-api-blitzbudget-com/service/dynamodb"
	"learn-api-blitzbudget-com/service/models"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func FetchResults(dbClient dynamodbiface.DynamoDBAPI, body *string) (*[]models.DBItem, error) {
	req, err := ParserRequest(body)
	if err != nil {
		log.Println("Error in parsing request", err)
		return nil, err
	}

	res, err := dynamoDBService.GetItems(dbClient, req)
	if err != nil {
		log.Println("Error in fetching results from DynamoDB", err)
		return nil, err
	}

	response, err := dynamoDBService.ParseQueryOutput(res)
	if err != nil {
		log.Println("Error in parsing query output", err)
		return nil, err
	}

	return response, nil
}
