package dynamodb

import (
	"learn-api-blitzbudget-com/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestParseQueryOutput(t *testing.T) {
	// Define a sample DynamoDB QueryOutput with items
	sampleQueryOutput := &dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{
			{
				"pk": {
					S: stringPtr("learn.blitzbudget.com"),
				},
				"sk": {
					S: stringPtr("content/coding/backend/serverless/golang/golang-fundamentals/chapter-1-introduction-to-golang.json"),
				},
				"Author": {
					S: stringPtr("Nagarjun Nagesh"),
				},
				"Category": {
					S: stringPtr("coding/backend/serverless/golang/golang-fundamentals"),
				},
				"creation_date": {
					S: stringPtr("2023-07-24T18:08:01Z"),
				},
				"File": {
					S: stringPtr("content/coding/backend/serverless/golang/golang-fundamentals/chapter-1-introduction-to-golang.json"),
				},
				"Name": {
					S: stringPtr("Chapter 1: Introduction to Golang"),
				},
				"Tags": {
					S: stringPtr("coding, backend, serverless, golang, golang-fundamentals"),
				},
			},
		},
	}

	// Call the ParseQueryOutput function with the sample query output
	result, err := ParseQueryOutput(sampleQueryOutput)

	// Assert that there is no error and the result is not nil
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the result contains the expected number of items
	assert.Len(t, *result, 1)

	// Assert that the parsed item has the correct values
	expectedItem := models.DBItem{
		PK:           "learn.blitzbudget.com",
		SK:           "content/coding/backend/serverless/golang/golang-fundamentals/chapter-1-introduction-to-golang.json",
		Author:       "Nagarjun Nagesh",
		Category:     "coding/backend/serverless/golang/golang-fundamentals",
		CreationDate: "2023-07-24T18:08:01Z",
		File:         "content/coding/backend/serverless/golang/golang-fundamentals/chapter-1-introduction-to-golang.json",
		Name:         "Chapter 1: Introduction to Golang",
		Tags:         "coding, backend, serverless, golang, golang-fundamentals",
	}
	assert.Equal(t, expectedItem, (*result)[0])

	// Test the case where the QueryOutput is nil (empty result)
	emptyResult, err := ParseQueryOutput(nil)
	assert.NoError(t, err)
	assert.NotNil(t, emptyResult)
	assert.Len(t, *emptyResult, 0)
}

func TestParseRequestEmptyData(t *testing.T) {
	// Test the case where the QueryOutput has no items
	emptyItemsResult, err := ParseQueryOutput(&dynamodb.QueryOutput{})
	assert.NoError(t, err)
	assert.NotNil(t, emptyItemsResult)
	assert.Len(t, *emptyItemsResult, 0)
}

func stringPtr(s string) *string {
	return &s
}
