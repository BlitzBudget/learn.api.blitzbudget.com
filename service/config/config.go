package config

import "os"

var (
	S3Bucket      = os.Getenv("S3_BUCKET_NAME")
	DynamoDBTable = os.Getenv("DYNAMODB_TABLE_NAME")
)
