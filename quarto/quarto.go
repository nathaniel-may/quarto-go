package quarto

import "errors"

type board struct {
	id string
	squares map[int]map[int]piece
	pieces map[piece]bool
	lines map[string]int
	active piece
}

type Quarto interface {
	isWon() bool
	TakeTurn(toPlace piece, h int, v int, active piece) Quarto
}

func (board *board) isWon() bool{
	//TODO
	return false
}

func (board *board) TakeTurn(toPlace piece, h int, v int, active piece) Quarto {
	//TODO
	return board
}

func (board *board) Validate() (bool, error) {
	if Pieces().ZERO == board.active && len(board.squares) != 0 && !board.isWon() {
		return false, errors.New("active piece required")
	}
	//TODO other probs
	return true, nil
}

func (board *board) setSquares(squares *map[int]map[int]piece) {
	board.squares = *squares
}

func NewBoard(id string) *board {
	return &board {id, map[int]map[int]piece {}, map[piece]bool {}, map[string]int {}, Pieces().ZERO}
}

func InProgBoard(id string, squares *map[int]map[int]piece, active piece) *board {
	board := NewBoard(id)
	board.setSquares(squares)
	//TODO use squares to set lines and pieces too
	board.Validate()
	return board
}


type piece struct {
	color string
	height string
	shape string
	top string
}

type pieces struct {
	ZERO piece
	WTQF, WTQH, WTRF, WTRH piece
	WSQF, WSQH, WSRF, WSRH piece
	BTQF, BTQH, BTRF, BTRH piece
	BSQF, BSQH, BSRF, BSRH piece
}

func Pieces() *pieces {
	return &p
}

var p = pieces{
piece {"zero","zero","zero","zero"},
piece {"white","tall","square","flat"},
piece {"white","tall","square","hole"},
piece {"white","tall","round","flat"},
piece {"white","tall","round","hole"},
piece {"white","short","square","flat"},
piece {"white","short","square","hole"},
piece {"white","short","round","flat"},
piece {"white","short","round","hole"},
piece {"black","tall","square","flat"},
piece {"black","tall","square","hole"},
piece {"black","tall","round","flat"},
piece {"black","tall","round","hole"},
piece {"black","short","square","flat"},
piece {"black","short","square","hole"},
piece {"black","short","round","flat"},
piece {"black","short","round","hole"},
}