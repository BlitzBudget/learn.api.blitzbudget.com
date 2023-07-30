package bucket

import (
	"learn-api-blitzbudget-com/service/config"
	"learn-api-blitzbudget-com/service/errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

type mockS3Client struct {
	s3iface.S3API
	putObjectFunc func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func (m *mockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return m.putObjectFunc(input)
}

func TestStoreDataInBucket_Success(t *testing.T) {
	// Create a sample request and file data
	fileData := []byte("Sample file data")

	// Create a mock S3 client
	mockClient := &mockS3Client{
		putObjectFunc: func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			// Verify the input parameters
			assert.Equal(t, config.S3Bucket, aws.StringValue(input.Bucket))
			assert.Equal(t, "content/index.json", aws.StringValue(input.Key))

			// Return a successful response
			return &s3.PutObjectOutput{}, nil
		},
	}

	// Call the function under test
	err := StoreDataInBucket(fileData, mockClient)

	// Verify the output
	assert.NoError(t, err)
}

func TestStoreDataInBucket_Error(t *testing.T) {
	// Create a sample request and file data
	fileData := []byte("Sample file data")

	// Create a mock S3 client
	mockClient := &mockS3Client{
		putObjectFunc: func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			// Simulate an error during upload
			return nil, errors.ErrUnableToStoreFileInS3
		},
	}

	// Call the function under test
	err := StoreDataInBucket(fileData, mockClient)

	// Verify the error
	assert.Error(t, err)
	assert.Equal(t, errors.ErrUnableToStoreFileInS3, err)
}

func TestFormatString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "#Chapter 8: Using SNS for Notifications - Publish and Subscribe to Events",
			expected: "chapter-8-using-sns-for-notifications-publish-and-subscribe-to-events",
		},
		{
			input:    "Testing Function to REMOVE Special Characters @#!!",
			expected: "testing-function-to-remove-special-characters",
		},
		{
			input:    "----TEST----", // Only special characters and hyphens
			expected: "test",
		},
		{
			input:    "    Space     ", // Only spaces
			expected: "space",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := FormatName(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
