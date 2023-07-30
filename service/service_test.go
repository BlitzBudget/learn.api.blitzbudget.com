package service

import (
	"errors"
	"learn-api-blitzbudget-com/service/config"
	"learn-api-blitzbudget-com/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
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

type mockS3Client struct {
	s3iface.S3API
	putObjectFunc func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func (m *mockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return m.putObjectFunc(input)
}

func TestFetchResultsSuccess(t *testing.T) {
	// Create a mock DynamoDBAPI with successful query response
	dbClient := new(MockDynamoDBAPI)
	// Create a mock S3 client
	s3Client := &mockS3Client{
		putObjectFunc: func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			// Verify the input parameters
			assert.Equal(t, config.S3Bucket, aws.StringValue(input.Bucket))
			assert.Equal(t, "content/index.json", aws.StringValue(input.Key))

			// Return a successful response
			return &s3.PutObjectOutput{}, nil
		},
	}
	expectedResult := []models.DBItem{
		{
			PK:           "12345",
			SK:           "file",
			Author:       "John Doe",
			Category:     "Technology",
			CreationDate: "2023-07-28",
			Name:         "Sample Document",
			Tags:         "tag1,tag2",
		},
		{
			PK:           "67890",
			SK:           "file",
			Author:       "Jane Smith",
			Category:     "Science",
			CreationDate: "2023-07-27",
			Name:         "Scientific Report",
			Tags:         "research,report",
		},
	}

	count := int64(1)
	queryOutput := &dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{},
		Count: &count,
	}
	// Convert the expectedResult slice into the format expected by QueryOutput
	itemMap := map[string]*dynamodb.AttributeValue{
		config.PK:           {S: &expectedResult[0].PK},
		config.SK:           {S: &expectedResult[0].SK},
		config.Author:       {S: &expectedResult[0].Author},
		config.Category:     {S: &expectedResult[0].Category},
		config.CreationDate: {S: &expectedResult[0].CreationDate},
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
		config.Name:         {S: &expectedResult[1].Name},
		config.Tags:         {S: &expectedResult[1].Tags},
	}
	queryOutput.Items = append(queryOutput.Items, itemMap2)

	dbClient.On("Query", mock.AnythingOfType("*dynamodb.QueryInput")).Return(queryOutput, nil)

	body := `{"url": "https://example.com"}`

	// Call the FetchResults function with the provided mock DynamoDBAPI
	result, err := FetchResults(dbClient, s3Client, &body)

	// Assert the result matches the expected output
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, *result)

	// Assert that the expected methods on the mock were called
	dbClient.AssertExpectations(t)
}

func TestFetchResultsError(t *testing.T) {
	// Create a mock DynamoDBAPI with an error response
	dbClient := new(MockDynamoDBAPI)
	s3Client := new(mockS3Client)
	someError := errors.New("some error")
	dbClient.On("Query", mock.AnythingOfType("*dynamodb.QueryInput")).Return(&dynamodb.QueryOutput{}, someError)

	body := `{"url": "https://example.com"}`

	// Call the FetchResults function with the provided mock DynamoDBAPI
	result, err := FetchResults(dbClient, s3Client, &body)

	// Assert the result is nil due to error
	assert.Nil(t, result)

	// Assert the error matches the expected error
	assert.Equal(t, someError, err)

	// Assert that the expected methods on the mock were called
	dbClient.AssertExpectations(t)
}
