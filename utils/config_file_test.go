package utils

import (
	"testing"
	"reflect"
	"strings"
)

func TestNewBoardActivePieceIsNilPiece(t *testing.T) {
	config := Load()
	connString := config.GetDBConnString("dev")
	if reflect.TypeOf(connString) != reflect.TypeOf("") {
		t.Fail()
	}
	if !strings.HasPrefix(connString, "mongodb") {
		t.Fail()
	}
}