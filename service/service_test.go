package service

import (
	"errors"
	"learn-api-blitzbudget-com/service/config"
	"learn-api-blitzbudget-com/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDynamoDBAPI is a mock implementation of dynamodbiface.DynamoDBAPI for testing
type MockDynamoDBAPI struct {
	mock.Mock
	dynamodbiface.DynamoDBAPI
}

// Query is a method in the mock implementation that returns predefined values
func (m *MockDynamoDBAPI) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.QueryOutput), args.Error(1)
}

func TestFetchResultsSuccess(t *testing.T) {
	// Create a mock DynamoDBAPI with successful query response
	dbClient := new(MockDynamoDBAPI)
	expectedResult := []models.DBItem{
		{
			PK:           "12345",
			SK:           "file",
			Author:       "John Doe",
			Category:     "Technology",
			CreationDate: "2023-07-28",
			File:         "example.pdf",
			Name:         "Sample Document",
			Tags:         "tag1,tag2",
		},
		{
			PK:           "67890",
			SK:           "file",
			Author:       "Jane Smith",
			Category:     "Science",
			CreationDate: "2023-07-27",
			File:         "sample.docx",
			Name:         "Scientific Report",
			Tags:         "research,report",
		},
	}

	queryOutput := &dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{},
	}
	// Convert the expectedResult slice into the format expected by QueryOutput
	itemMap := map[string]*dynamodb.AttributeValue{
		config.PK:           {S: &expectedResult[0].PK},
		config.SK:           {S: &expectedResult[0].SK},
		config.Author:       {S: &expectedResult[0].Author},
		config.Category:     {S: &expectedResult[0].Category},
		config.CreationDate: {S: &expectedResult[0].CreationDate},
		config.File:         {S: &expectedResult[0].File},
		config.Name:         {S: &expectedResult[0].Name},
		config.Tags:         {S: &expectedResult[0].Tags},
	}
	queryOutput.Items = append(queryOutput.Items, itemMap)
	itemMap2 := map[string]*dynamodb.AttributeValue{
		config.PK:           {S: &expectedResult[1].PK},
		config.SK:           {S: &expectedResult[1].SK},
		config.Author:       {S: &expectedResult[1].Author},
		config.Category:     {S: &expectedResult[1].Category},
		config.CreationDate: {S: &expectedResult[1].CreationDate},
		config.File:         {S: &expectedResult[1].File},
		config.Name:         {S: &expectedResult[1].Name},
		config.Tags:         {S: &expectedResult[1].Tags},
	}
	queryOutput.Items = append(queryOutput.Items, itemMap2)

	dbClient.On("Query", mock.AnythingOfType("*dynamodb.QueryInput")).Return(queryOutput, nil)

	body := `{"url": "https://example.com"}`

	// Call the FetchResults function with the provided mock DynamoDBAPI
	result, err := FetchResults(dbClient, &body)

	// Assert the result matches the expected output
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, *result)

	// Assert that the expected methods on the mock were called
	dbClient.AssertExpectations(t)
}

func TestFetchResultsError(t *testing.T) {
	// Create a mock DynamoDBAPI with an error response
	dbClient := new(MockDynamoDBAPI)
	someError := errors.New("some error")
	dbClient.On("Query", mock.AnythingOfType("*dynamodb.QueryInput")).Return(&dynamodb.QueryOutput{}, someError)

	body := `{"url": "https://example.com"}`

	// Call the FetchResults function with the provided mock DynamoDBAPI
	result, err := FetchResults(dbClient, &body)

	// Assert the result is nil due to error
	assert.Nil(t, result)

	// Assert the error matches the expected error
	assert.Equal(t, someError, err)

	// Assert that the expected methods on the mock were called
	dbClient.AssertExpectations(t)
}
