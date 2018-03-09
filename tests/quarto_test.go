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
	board := quarto.NewBoard("test")
	board, error := board.TakeTurn(quarto.Pieces.WTQF,2, 3, quarto.Pieces.BSQF)
	if error != nil {
		t.Errorf(error.Error())
	}

	squares := make(map[quarto.Square]quarto.Piece)
	squares[quarto.Square{2,3}] = quarto.Pieces.WTQF
	inProg := quarto.InProgBoard("test", squares, quarto.Pieces.BSQF)

	if !reflect.DeepEqual(board, inProg) {
		fmt.Println(board)
		fmt.Println(inProg)
		t.Fail()
	}
}

func TestWinningMove(t *testing.T) {
	board := quarto.NewBoard("test")
	board, _ = board.TakeTurn(quarto.Pieces.WTQF, 1, 1, quarto.Pieces.WTQH)
	board, _ = board.TakeTurn(quarto.Pieces.WTQH, 1, 2, quarto.Pieces.WTRF)
	board, _ = board.TakeTurn(quarto.Pieces.WTRF, 1, 3, quarto.Pieces.WSQH)
	board, _ = board.TakeTurn(quarto.Pieces.WSQH, 1, 4, quarto.Pieces.ZERO)

	if !board.IsWon() {
		t.Fail()
	}
}

func TestD1WinningMove(t *testing.T) {
	board := quarto.NewBoard("test")
	board, _ = board.TakeTurn(quarto.Pieces.WTQF, 1, 1, quarto.Pieces.WTQH)
	board, _ = board.TakeTurn(quarto.Pieces.WTQH, 2, 2, quarto.Pieces.WTRF)
	board, _ = board.TakeTurn(quarto.Pieces.WTRF, 3, 3, quarto.Pieces.WSQH)
	board, _ = board.TakeTurn(quarto.Pieces.WSQH, 4, 4, quarto.Pieces.ZERO)

	if !board.IsWon() {
		t.Fail()
	}
}

func TestD2WinningMove(t *testing.T) {
	board := quarto.NewBoard("test")
	board, _ = board.TakeTurn(quarto.Pieces.WTQF, 1, 4, quarto.Pieces.WTQH)
	board, _ = board.TakeTurn(quarto.Pieces.WTQH, 2, 3, quarto.Pieces.WTRF)
	board, _ = board.TakeTurn(quarto.Pieces.WTRF, 3, 2, quarto.Pieces.WSQH)
	board, _ = board.TakeTurn(quarto.Pieces.WSQH, 4, 1, quarto.Pieces.ZERO)

	if !board.IsWon() {
		t.Fail()
	}
}

func TestRejectZeroPieceIfNotWon(t *testing.T) {
	board := quarto.NewBoard("test")
	board, error := board.TakeTurn(quarto.Pieces.WTQF, 1, 1, quarto.Pieces.ZERO)
	if error == nil {
		t.Fail()
	}
}

//TODO this test fails
//func TestBoardDoesntChangeAfterTakeTurn(t *testing.T) {
//	board := quarto.NewBoard("test")
//	boardClone := quarto.NewBoard("test")
//	newBoard, error := board.TakeTurn(quarto.Pieces.WTQF, 1, 4, quarto.Pieces.WTQH)
//
//	if error != nil {
//		t.Fail()
//	}
//
//	if !reflect.DeepEqual(boardClone, board) {
//		fmt.Println(boardClone)
//		fmt.Println(board)
//		t.Fail()
//	}
//
//	if reflect.DeepEqual(newBoard, board) {
//		t.Fail()
//	}
//}

//TODO test that underlying board is not modified during take turn