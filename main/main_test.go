package main

import (
	"testing"
	"github.com/aws/aws-lambda-go/events"
	"reflect"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: "Paul"},
			expect:  "Hello Paul",
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     ErrNameNotProvided,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		if reflect.TypeOf(test.err) != reflect.TypeOf(err){
			t.Fail()
		}
		if !reflect.DeepEqual(test.expect, response.Body){
			t.Fail()
		}
	}
}