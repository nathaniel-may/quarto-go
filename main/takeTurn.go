package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"errors"
)

// Lambda function handler
func TakeTurnHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//TODO use SNS to call the dal lambda function to take the turn once validated

	// body must take boardId, piece toPlace, h and v of placement and active piece
	// note, currently no security here for turns on other boards or play out of turn
	if len(request.Body) < 4 {
		return events.APIGatewayProxyResponse{}, errors.New("bad turn request")
	}

	//TODO return ok or not ok
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(TakeTurnHandler)
}