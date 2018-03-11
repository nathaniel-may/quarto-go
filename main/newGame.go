package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Lambda function handler
func NewGameHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//TODO use SNS to call the dal lambda function to create a new game

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	//TODO return json representation of the game (new package)
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(NewGameHandler)
}