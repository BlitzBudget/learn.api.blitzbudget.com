package bucket

import (
	"bytes"
	"fmt"
	"learn-api-blitzbudget-com/service/config"
	"learn-api-blitzbudget-com/service/errors"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func StoreDataInBucket(fileData []byte, s3Client s3iface.S3API) error {
	fmt.Println("Uploading file to S3 bucket...")

	fileURL := fmt.Sprintf("%sindex.json", config.SKPrefix)

	// Upload the JSON data to S3 bucket
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(fileURL),
		Body:   aws.ReadSeekCloser(bytes.NewReader(fileData)),
	})
	if err != nil {
		log.Println("Failed to upload JSON to S3:", err)
		return errors.ErrUnableToStoreFileInS3
	}

	fmt.Println("JSON file uploaded to S3 successfully!")
	return err
}

func FormatName(input string) string {
	// Convert the input string to lowercase.
	input = strings.ToLower(input)

	// Define a regular expression to match any character that is not alphanumeric or hyphen.
	regex := regexp.MustCompile("[^a-z0-9]+")

	// Use the regular expression to replace all matched characters with an empty string.
	result := regex.ReplaceAllString(input, "-")

	// Remove leading and trailing hyphens.
	result = strings.Trim(result, "-")

	return result
}
