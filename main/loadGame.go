package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"errors"
)

// Lambda function handler
func LoadGameHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//TODO use SNS to call the dal lambda function to save the new state after validation

	// If no boardId is provided in the HTTP request body, throw an error
	if len(request.Body) < 1{
		return events.APIGatewayProxyResponse{}, errors.New("must request with boardId")
	}

	//TODO return json representation of the game (new package)
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(LoadGameHandler)
}