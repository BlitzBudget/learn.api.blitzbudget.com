package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"learn-api-blitzbudget-com/service"
	"learn-api-blitzbudget-com/service/config"
	httpErrors "learn-api-blitzbudget-com/service/errors"
	"learn-api-blitzbudget-com/service/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jsonReq, _ := json.Marshal(request)
	fmt.Printf("Processing request data for request %v.\n", string(jsonReq))

	//service.FetchResults(&request.Body)
	header := map[string]string{
		"Access-Control-Allow-Origin":      "https://" + config.S3Bucket,
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dbClient := dynamodb.New(sess)
	res, err := service.FetchResults(dbClient, &request.Body)
	if err != nil {
		errorAsBytes, httpStatusCode := fetchErrorMessage(err)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: httpStatusCode, Headers: header}, nil
	}

	httpResponse := models.HttpResponse{
		Items: res,
	}
	responseAsBytes, _ := json.Marshal(httpResponse)
	return events.APIGatewayProxyResponse{Body: string(responseAsBytes), StatusCode: 200, Headers: header}, nil
}

func fetchErrorMessage(err error) ([]byte, int) {
	errorMessage := err.Error()
	errorCode := httpErrors.ExtractErrorCode(err)
	errorRespose := models.ErrorHttpResponse{
		Message: &errorMessage,
		Code:    errorCode,
	}
	errorAsBytes, _ := json.Marshal(errorRespose)
	httpStatusCode := 400

	// TODO : Need to handle error codes properly
	return errorAsBytes, httpStatusCode
}
