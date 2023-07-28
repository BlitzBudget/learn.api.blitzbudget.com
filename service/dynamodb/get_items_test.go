package dynamodb

import (
	"errors"
	"learn-api-blitzbudget-com/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
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

func TestGetItemsSuccessfulQuery(t *testing.T) {
	// Create a mock DynamoDBAPI with successful query response
	dbClient := new(MockDynamoDBAPI)
	expectedResult := &dynamodb.QueryOutput{}
	dbClient.On("Query", mock.AnythingOfType("*dynamodb.QueryInput")).Return(expectedResult, nil)

	request := &models.Request{
		URL: aws.String("testURL"),
	}

	// Call the GetItems function with the provided mock DynamoDBAPI
	result, err := GetItems(dbClient, request)

	// Assert the result matches the expected output
	assert.Equal(t, expectedResult, result)

	// Assert there is no error
	assert.Nil(t, err)

	// Assert that the expected method on the mock was called
	dbClient.AssertExpectations(t)
}

func TestGetItemsErrorQueryingDynamoDB(t *testing.T) {
	// Create a mock DynamoDBAPI with an error response
	someError := errors.New("some-error")
	dbClient := new(MockDynamoDBAPI)
	expectedResult := &dynamodb.QueryOutput{}
	dbClient.On("Query", mock.AnythingOfType("*dynamodb.QueryInput")).Return(expectedResult, someError)

	request := &models.Request{
		URL: aws.String("testURL"),
	}

	// Call the GetItems function with the provided mock DynamoDBAPI
	result, err := GetItems(dbClient, request)

	// Assert the result is nil due to error
	assert.Nil(t, result)

	// Assert the error matches the expected error
	assert.Equal(t, someError, err)

	// Assert that the expected method on the mock was called
	dbClient.AssertExpectations(t)
}
