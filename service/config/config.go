package config

import "os"

var ScanIndexForward = false
var (
	PK                   = "pk"
	SK                   = "sk"
	Author               = "Author"
	Category             = "Category"
	CreationDate         = "creation_date"
	Name                 = "BlogName"
	Tags                 = "Tags"
	ProjectionExpression = "pk, sk, Author, Category, creation_date, BlogName, Tags"
	SKPrefix             = "content/"
)

var (
	S3Bucket      = os.Getenv("S3_BUCKET_NAME")
	DynamoDBTable = os.Getenv("DYNAMODB_TABLE_NAME")
)
