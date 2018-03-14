package main

import (
	"testing"
	"github.com/aws/aws-lambda-go/events"
)

func TestLoadGameHandler(t *testing.T) {
	//TODO "goify" with input/expect pairs
	var response events.APIGatewayProxyResponse
	response, err := NewGameHandler(events.APIGatewayProxyRequest{}) //TODO add boardId to body
	if err != nil {
		t.Errorf(err.Error())
	}

	//TODO STUB
	if response.Body != "a json representation of the correct board" {
		t.Fail()
	}
}