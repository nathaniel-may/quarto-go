package utils

import (
	"testing"
	"reflect"
	"strings"
)

func TestNewBoardActivePieceIsNilPiece(t *testing.T) {
	config := Load("dev")
	connString := config.GetDBConnString()
	db := config.GetDB()
	if reflect.TypeOf(connString) != reflect.TypeOf("") {
		t.Fail()
	}
	if !strings.HasPrefix(connString, "mongodb") {
		t.Fail()
	}
	if !strings.HasPrefix(db, "quarto") {
		t.Fail()
	}
}