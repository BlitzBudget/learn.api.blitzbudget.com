package service

import (
	"encoding/json"
	"learn-api-blitzbudget-com/service/bucket"
	dynamoDBService "learn-api-blitzbudget-com/service/dynamodb"
	"learn-api-blitzbudget-com/service/errors"
	"learn-api-blitzbudget-com/service/models"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func FetchResults(dbClient dynamodbiface.DynamoDBAPI, s3Client s3iface.S3API, body *string) (*[]models.DBItem, error) {
	log.Println("Fetching results from DynamoDB")
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

	// Store it in a Put Object
	responseAsBytes, err := json.Marshal(response)
	if err != nil {
		log.Println("Error in marshalling response", err)
		return nil, errors.ErrMarshalingResponseFromDB
	}
	bucket.StoreDataInBucket(responseAsBytes, s3Client)
	log.Println("Results fetched from DynamoDB and the item count is: ", len(*response))
	return response, nil
}
