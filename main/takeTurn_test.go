package main

import (
	"testing"
	"github.com/aws/aws-lambda-go/events"
)

func TestTakeTurnHandler(t *testing.T) {
	//TODO "goify" with input/expect pairs
	var response events.APIGatewayProxyResponse
	response, err := NewGameHandler(events.APIGatewayProxyRequest{}) //TODO add boardId,piece,h,v,active to body
	if err != nil {
		t.Errorf(err.Error())
	}

	//TODO STUB
	if response.Body != "not ok" {
		t.Fail()
	}
}