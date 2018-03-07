package tests

import "../quarto"
import "testing"
import "reflect"

func TestBoardEquality(t *testing.T) {
	if !reflect.DeepEqual(quarto.NewBoard("test"), quarto.NewBoard("test")) {
		t.Fail()
	}
}

func TestValidate(t *testing.T) {
	board := quarto.NewBoard("test")
	valid, error := board.Validate()
	if !valid {
		t.Errorf(error.Error())
	}
}