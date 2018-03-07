package tests

import "../quarto"
import "testing"

//func TestNewBoard(t *testing.T) {

//}

func TestBoardEquality(t *testing.T) {
	if quarto.NewBoard("test") != quarto.NewBoard("test") {
		t.Fail()
	}
}