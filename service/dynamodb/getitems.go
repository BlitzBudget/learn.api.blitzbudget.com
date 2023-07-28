package dynamodb

import (
	"learn-api-blitzbudget-com/service/config"
	"learn-api-blitzbudget-com/service/models"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func GetItems(dbClient dynamodbiface.DynamoDBAPI, req *models.Request) (*dynamodb.QueryOutput, error) {
	// Construct the DynamoDB query parameters
	params := &dynamodb.QueryInput{
		TableName: aws.String(config.DynamoDBTable), // Replace with your DynamoDB table name
		KeyConditions: map[string]*dynamodb.Condition{
			"PK": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(config.S3Bucket),
					},
				},
			},
			"SK": {
				ComparisonOperator: aws.String("BEGINS_WITH"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(*req.URL),
					},
				},
			},
		},
	}

	// Perform the DynamoDB query
	result, err := dbClient.Query(params)
	if err != nil {
		log.Println("Error querying DynamoDB:", err)
		return nil, nil
	}

	return result, nil
}
