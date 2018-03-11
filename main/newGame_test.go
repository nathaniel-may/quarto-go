package main

import (
	"testing"
	"github.com/aws/aws-lambda-go/events"
)

func TestNewGameHandler(t *testing.T) {
	//TODO "goify" with input/expect pairs
	var response events.APIGatewayProxyResponse
	response, err := NewGameHandler(events.APIGatewayProxyRequest{})
	if err != nil {
		t.Errorf(err.Error())
	}

	//TODO STUB
	if response.Body != "exactly what is expected" {
		t.Fail()
	}
}