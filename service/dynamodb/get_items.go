package dynamodb

import (
	"fmt"
	"learn-api-blitzbudget-com/service/config"
	"learn-api-blitzbudget-com/service/models"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func GetItems(dbClient dynamodbiface.DynamoDBAPI, req *models.Request) (*dynamodb.QueryOutput, error) {
	log.Println("fetching the DynamoDB items")

	skPrefix := fmt.Sprintf("%s%s",config.SKPrefix,  *req.URL)
	// Construct the DynamoDB query parameters
	params := &dynamodb.QueryInput{
		TableName: aws.String(config.DynamoDBTable), // Replace with your DynamoDB table name
		KeyConditions: map[string]*dynamodb.Condition{
			config.PK: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(config.S3Bucket),
					},
				},
			},
			config.SK: {
				ComparisonOperator: aws.String("BEGINS_WITH"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(skPrefix),
					},
				},
			},
		},
		ProjectionExpression: &config.ProjectionExpression,
		ScanIndexForward:     &config.ScanIndexForward,
	}

	// Perform the DynamoDB query
	result, err := dbClient.Query(params)
	if err != nil {
		log.Println("Error querying DynamoDB:", err)
		return nil, err
	}

	log.Println("Successfully queried the DynamoDB table")
	return result, nil
}
