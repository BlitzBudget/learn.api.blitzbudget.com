package service

import (
	"learn-api-blitzbudget-com/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func TestParserRequestSuccess(t *testing.T) {
	body := `{"url": "https://example.com"}`
	expectedRequest := &models.Request{
		URL: aws.String("https://example.com"),
	}

	// Call the ParserRequest function with the test input
	result, err := ParserRequest(&body)

	// Assert the result matches the expected output
	assert.Nil(t, err)
	assert.Equal(t, expectedRequest, result)
}

func TestParserRequestError(t *testing.T) {
	body := `{"invalid_key": "https://example.com}`

	// Call the ParserRequest function with the test input
	result, err := ParserRequest(&body)

	// Assert the function returns an error
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
