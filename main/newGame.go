package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"errors"
)

// Lambda function handler
func NewGameHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//TODO use SNS to call the dal lambda function to create a new game

	// only empty body accepted
	if len(request.Body) > 0 {
		return events.APIGatewayProxyResponse{}, errors.New("bad request")
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