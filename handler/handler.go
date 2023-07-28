package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-template/service"
	httpErrors "golang-template/service/errors"
	"golang-template/service/models"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jsonReq, _ := json.Marshal(request)
	fmt.Printf("Processing request data for request %v.\n", string(jsonReq))

	//service.SaveRequest(&request.Body)
	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}

	err := service.SaveRequest(&request.Body)
	if err != nil {
		errorAsBytes, httpStatusCode := fetchErrorMessage(err)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: httpStatusCode, Headers: header}, nil
	}

	message := "The password has been changed successfully"
	httpResponse := models.HttpResponse{
		Message: &message,
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
