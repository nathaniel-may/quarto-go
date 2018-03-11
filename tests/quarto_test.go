package tests

import ("quarto-go/quarto"
		"testing"
		"reflect"
)

type quartoPieces struct {
	NIL quarto.Piece
	WTQF, WTQH, WTRF, WTRH quarto.Piece
	WSQF, WSQH, WSRF, WSRH quarto.Piece
	BTQF, BTQH, BTRF, BTRH quarto.Piece
	BSQF, BSQH, BSRF, BSRH quarto.Piece
}

var pieces = quartoPieces {
	quarto.NilPiece(),
	quarto.NewPiece(quarto.WHITE, quarto.TALL, quarto.SQUARE, quarto.FLAT),
	quarto.NewPiece(quarto.WHITE, quarto.TALL, quarto.SQUARE, quarto.HOLE),
	quarto.NewPiece(quarto.WHITE, quarto.TALL, quarto.ROUND, quarto.FLAT),
	quarto.NewPiece(quarto.WHITE, quarto.TALL, quarto.ROUND, quarto.HOLE),
	quarto.NewPiece(quarto.WHITE, quarto.SHORT, quarto.SQUARE, quarto.FLAT),
	quarto.NewPiece(quarto.WHITE, quarto.SHORT, quarto.SQUARE, quarto.HOLE),
	quarto.NewPiece(quarto.WHITE, quarto.SHORT, quarto.ROUND, quarto.FLAT),
	quarto.NewPiece(quarto.WHITE, quarto.SHORT, quarto.ROUND, quarto.HOLE),
	quarto.NewPiece(quarto.BLACK, quarto.TALL, quarto.SQUARE, quarto.FLAT),
	quarto.NewPiece(quarto.BLACK, quarto.TALL, quarto.SQUARE, quarto.HOLE),
	quarto.NewPiece(quarto.BLACK, quarto.TALL, quarto.ROUND, quarto.FLAT),
	quarto.NewPiece(quarto.BLACK, quarto.TALL, quarto.ROUND, quarto.HOLE),
	quarto.NewPiece(quarto.BLACK, quarto.SHORT, quarto.SQUARE, quarto.FLAT),
	quarto.NewPiece(quarto.BLACK, quarto.SHORT, quarto.SQUARE, quarto.HOLE),
	quarto.NewPiece(quarto.BLACK, quarto.SHORT, quarto.ROUND, quarto.FLAT),
	quarto.NewPiece(quarto.BLACK, quarto.SHORT, quarto.ROUND, quarto.HOLE),
}

func TestNewBoardActivePieceIsNilPiece(t *testing.T) {
	if !reflect.DeepEqual(quarto.NewBoard("test").GetActive(), quarto.NilPiece()) {
		t.Fail()
	}
}

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
	error := board.TakeTurn(pieces.WTQF,2, 3, pieces.BSQF)
	if error != nil {
		t.Errorf(error.Error())
	}

	squares := make(map[quarto.Square]quarto.Piece)
	squares[quarto.Square{2,3}] = pieces.WTQF
	inProg := quarto.InProgBoard("test", squares, pieces.BSQF)

	if !reflect.DeepEqual(board, inProg) {
		t.Fail()
	}
}

func TestWinningMove(t *testing.T) {
	board := quarto.NewBoard("test")
	error := board.TakeTurn(pieces.WTQF, 1, 1, pieces.WTQH)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WTQH, 1, 2, pieces.WTRF)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WTRF, 1, 3, pieces.WSQH)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WSQH, 1, 4, pieces.NIL)
	if error != nil {
		t.Errorf(error.Error())
	}

	if !board.IsWon() {
		t.Fail()
	}
}

func TestD1WinningMove(t *testing.T) {
	board := quarto.NewBoard("test")
	error := board.TakeTurn(pieces.WTQF, 1, 1, pieces.WTQH)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WTQH, 2, 2, pieces.WTRF)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WTRF, 3, 3, pieces.WSQH)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WSQH, 4, 4, pieces.NIL)
	if error != nil {
		t.Errorf(error.Error())
	}

	if !board.IsWon() {
		t.Fail()
	}
}

func TestD2WinningMove(t *testing.T) {
	board := quarto.NewBoard("test")
	error := board.TakeTurn(pieces.WTQF, 1, 4, pieces.WTQH)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WTQH, 2, 3, pieces.WTRF)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WTRF, 3, 2, pieces.WSQH)
	if error != nil {
		t.Errorf(error.Error())
	}
	error = board.TakeTurn(pieces.WSQH, 4, 1, pieces.NIL)
	if error != nil {
		t.Errorf(error.Error())
	}

	if !board.IsWon() {
		t.Fail()
	}
}

func TestRejectNilPieceIfNotWon(t *testing.T) {
	board := quarto.NewBoard("test")
	error := board.TakeTurn(pieces.WTQF, 1, 1, pieces.NIL)
	if error != nil {
		t.Fail()
	}
}