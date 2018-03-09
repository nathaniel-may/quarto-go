package tests

import ("../quarto"
		"testing"
		"reflect"
		"fmt")

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

func TestInProgEqualityOneTurn(t *testing.T) {
	board := *quarto.NewBoard("test")
	board, error := board.TakeTurn(quarto.Pieces.WTQF,2, 3, quarto.Pieces.BSQF)
	if error != nil {
		t.Errorf(error.Error())
	}

	squares := make(map[quarto.Square]quarto.Piece)
	squares[quarto.Square{2,3}] = quarto.Pieces.WTQF
	inProg := quarto.InProgBoard("test", &squares, quarto.Pieces.BSQF)

	if !reflect.DeepEqual(board, inProg) {
		fmt.Println(board)
		fmt.Println(inProg)
		t.Fail()
	}
}