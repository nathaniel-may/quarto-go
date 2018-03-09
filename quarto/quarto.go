package quarto

import (
	"errors"
	"strconv"
)

type Square struct {
	H int
	V int
}

type board struct {
	id string
	squares map[Square]Piece
	pieces map[Piece]bool
	lines map[string]int
	active Piece
}

type Quarto interface {
	isWon() bool
	TakeTurn(toPlace piece, h int, v int, active piece) Quarto
	Nil()
}

func (board *board) Nil() *board {
	return NewBoard("nil")
}

func (board *board) isWon() bool{
	//TODO
	return false
}

func (board board) TakeTurn(toPlace Piece, h int, v int, active Piece) (board, error) {
	if board.pieces[toPlace] == true {
		return *board.Nil(), errors.New("the placed piece is already on the board")
	}
	if board.pieces[active] == true {
		return *board.Nil(), errors.New("the active piece is already on the board")
	}
	if _, ok := board.squares[Square {h,v}]; ok {
		return *board.Nil(), errors.New("the square is already occupied")
	}
	if board.active != Pieces.ZERO && board.active != toPlace {
		return *board.Nil(), errors.New("must place the active piece")
	}

	board.squares[Square {h,v}] = toPlace
	board.pieces[toPlace] = true
	board.active = active
	for _, line := range getLines(h, v) {
		board.lines[line]++
	}

	return board, nil
}

func (board *board) Validate() (bool, error) {
	if Pieces.ZERO == board.active && len(board.squares) != 0 && !board.isWon() {
		return false, errors.New("active piece required")
	}
	//TODO other probs
	return true, nil
}

func (board *board) setSquares(squares *map[Square]Piece) {
	board.squares = *squares
}

func NewBoard(id string) *board {
	return &board {id, make(map[Square]Piece), make(map[Piece]bool), make(map[string]int), Pieces.ZERO}
}

func InProgBoard(id string, squares *map[Square]Piece, active Piece) board {
	board := NewBoard(id)
	board.setSquares(squares)
	pieces := make(map[Piece]bool)
	lines := make(map[string]int)
	for square, piece := range *squares {
		pieces[piece]=true
		for _, line := range getLines(square.H, square.V) {
			lines[line]++
		}
	}
	board.pieces = pieces
	board.lines = lines
	board.active = active
	board.Validate()
	return *board
}

func getLines(h int, v int) []string{
	var lines []string
	lines = append(lines, "h" + strconv.Itoa(h))
	lines = append(lines, "v" + strconv.Itoa(v))

	if h == v {
		lines = append(lines, "d" + strconv.Itoa(h))
	} else if h == 1 && v == 4 {
		lines = append(lines, "d2")
	} else if h == 2 && v == 3 {
		lines = append(lines, "d2")
	} else if h == 3 && v == 2 {
		lines = append(lines, "d2")
	} else if h == 4 && v == 1 {
		lines = append(lines, "d2")
	}

	return lines
}


type piece struct {
	color string
	height string
	shape string
	top string
}

type Piece interface {
	getColor() string
	getHeight() string
	getShape() string
	getTop() string
}

func (piece *piece) getColor() string{
	return piece.color
}

func (piece *piece) getHeight() string{
	return piece.height
}

func (piece *piece) getShape() string{
	return piece.shape
}

func (piece *piece) getTop() string{
	return piece.top
}

type pieces struct {
	ZERO Piece
	WTQF, WTQH, WTRF, WTRH Piece
	WSQF, WSQH, WSRF, WSRH Piece
	BTQF, BTQH, BTRF, BTRH Piece
	BSQF, BSQH, BSRF, BSRH Piece
}

var Pieces = pieces {
		&piece {"zero","zero","zero","zero"},
		&piece {"white","tall","square","flat"},
		&piece {"white","tall","square","hole"},
		&piece {"white","tall","round","flat"},
		&piece {"white","tall","round","hole"},
		&piece {"white","short","square","flat"},
		&piece {"white","short","square","hole"},
		&piece {"white","short","round","flat"},
		&piece {"white","short","round","hole"},
		&piece {"black","tall","square","flat"},
		&piece {"black","tall","square","hole"},
		&piece {"black","tall","round","flat"},
		&piece {"black","tall","round","hole"},
		&piece {"black","short","square","flat"},
		&piece {"black","short","square","hole"},
		&piece {"black","short","round","flat"},
		&piece {"black","short","round","hole"},
}