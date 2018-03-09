package quarto

import (
	"errors"
	"strconv"
)

type Square struct {
	H int
	V int
}

type line struct {
	line string
	attr string
}

type board struct {
	id string
	squares map[Square]Piece
	pieces map[Piece]bool
	lines map[line]int
	active Piece
}

type Quarto interface {
	Nil() Quarto
	IsWon() bool
	TakeTurn(toPlace Piece, h int, v int, active Piece) (Quarto, error)
	Validate() (bool, error)
	setSquares(squares map[Square]Piece)
	setPieces(pieces map[Piece]bool)
	setLines(lines map[line]int)
	setActive(piece Piece)

}

func (board *board) Nil() Quarto {
	return NewBoard("nil")
}

func (board *board) IsWon() bool{
	for _, count := range board.lines{
		if count >= 4 {
			return true
		}
	}
	return false
}

func (board board) TakeTurn(toPlace Piece, h int, v int, active Piece) (Quarto, error) {
	if board.pieces[toPlace] == true {
		return board.Nil(), errors.New("the placed piece is already on the board")
	}
	if board.pieces[active] == true {
		return board.Nil(), errors.New("the active piece is already on the board")
	}
	if _, ok := board.squares[Square {h,v}]; ok {
		return board.Nil(), errors.New("the square is already occupied")
	}
	if board.active != Pieces.ZERO && board.active != toPlace {
		return board.Nil(), errors.New("must place the active piece")
	}

	board.squares[Square {h,v}] = toPlace
	board.pieces[toPlace] = true
	board.active = active
	for _, line := range getLines(toPlace, h, v) {
		board.lines[line]++
	}
	_, error := board.Validate()
	if error != nil {
		return board.Nil(), error
	}

	return &board, nil
}

func (board *board) Validate() (bool, error) {
	if Pieces.ZERO == board.active && len(board.squares) != 0 && !board.IsWon() {
		return false, errors.New("active piece required")
	}
	if board.active == Pieces.ZERO && len(board.pieces) != 0 && !board.IsWon(){
		return false, errors.New("cannot set ZERO piece as active unless the game is won")
	}
	//TODO other probs
	return true, nil
}

func (board *board) setSquares(squares map[Square]Piece) {
	board.squares = squares
}

func (board *board) setPieces(pieces map[Piece]bool) {
	board.pieces = pieces
}

func (board *board) setLines(lines map[line]int) {
	board.lines = lines
}

func (board *board) setActive(piece Piece) {
	board.active = piece
}

func NewBoard(id string) Quarto {
	return &board {id, make(map[Square]Piece), make(map[Piece]bool), make(map[line]int), Pieces.ZERO}
}

func InProgBoard(id string, squares map[Square]Piece, active Piece) Quarto {
	board := NewBoard(id)
	board.setSquares(squares)
	pieces := make(map[Piece]bool)
	lines := make(map[line]int)
	for square, piece := range squares {
		pieces[piece]=true
		for _, line := range getLines(piece, square.H, square.V) {
			lines[line]++
		}
	}
	board.setPieces(pieces)
	board.setLines(lines)
	board.setActive(active)
	board.Validate()
	return board
}

func getLines(piece Piece, h int, v int) []line{
	var lines []line
	hline := "h" + strconv.Itoa(h)
	vline := "v" + strconv.Itoa(v)
	var dline string

	if h == v {
		dline = "d1"
	} else if h == 1 && v == 4 {
		dline = "d2"
	} else if h == 2 && v == 3 {
		dline = "d2"
	} else if h == 3 && v == 2 {
		dline = "d2"
	} else if h == 4 && v == 1 {
		dline = "d2"
	}

	lines = append(lines, line{hline, piece.getColor()})
	lines = append(lines, line{hline, piece.getHeight()})
	lines = append(lines, line{hline, piece.getShape()})
	lines = append(lines, line{hline, piece.getTop()})

	lines = append(lines, line{vline, piece.getColor()})
	lines = append(lines, line{vline, piece.getHeight()})
	lines = append(lines, line{vline, piece.getShape()})
	lines = append(lines, line{vline, piece.getTop()})

	if dline != "" {
		lines = append(lines, line{dline, piece.getColor()})
		lines = append(lines, line{dline, piece.getHeight()})
		lines = append(lines, line{dline, piece.getShape()})
		lines = append(lines, line{dline, piece.getTop()})
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