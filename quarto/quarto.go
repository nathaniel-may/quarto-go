package quarto

import (
	"errors"
	"strconv"
	"fmt"
	"reflect"
)

type Square struct {
	H int
	V int
}

type line struct {
	line string
	attr Attribute
}

type board struct {
	id string
	squares map[Square]Piece
	pieces map[Piece]bool
	lines map[line]int
	active Piece
}

type Quarto interface {
	IsWon() bool
	TakeTurn(toPlace Piece, h int, v int, active Piece) (Quarto, error)
	Validate() (bool, error)
	GetActive() Piece
	String() string
	setSquares(squares map[Square]Piece)
	setPieces(pieces map[Piece]bool)
	setLines(lines map[line]int)
	setActive(piece Piece)

}

func (board *board) IsWon() bool{
	for line, count := range board.lines{
		//TODO DELETE THIS PRINT
		fmt.Println("line: ", line.line, " attr: ", line.attr.AttrStr(), " count: ", count)
		if count >= 4 {
			return true
		}
	}
	fmt.Println("****************************************")
	return false
}

func (board board) TakeTurn(toPlace Piece, h int, v int, active Piece) (Quarto, error) {
	if board.pieces[toPlace] == true {
		return NilBoard(), errors.New("the placed piece is already on the board")
	}
	if board.pieces[active] == true {
		return NilBoard(), errors.New("the active piece is already on the board")
	}
	if _, ok := board.squares[Square {h,v}]; ok {
		return NilBoard(), errors.New("the square is already occupied")
	}
	if reflect.DeepEqual(board.GetActive(), NilPiece()) && reflect.DeepEqual(board.GetActive(), toPlace) {
		//TODO: DELETE THESE PRINT LINES
		if reflect.DeepEqual(board.GetActive(), NilPiece()) {
			fmt.Println("active isn't nil")
		}
		fmt.Println("board.GetActive(): ", board.GetActive())
		fmt.Println("toPlace: ", toPlace)
		return NilBoard(), errors.New("must place the active piece")
	}

	board.squares[Square {h,v}] = toPlace
	board.pieces[toPlace] = true
	board.setActive(active)
	for _, line := range getLines(toPlace, h, v) {
		board.lines[line]++
	}
	_, error := board.Validate()
	if error != nil {
		return NilBoard(), error
	}

	return &board, nil
}

func (board *board) Validate() (bool, error) {
	if NilPiece() == board.GetActive() && len(board.squares) != 0 && !board.IsWon() {
		return false, errors.New("active piece required")
	}
	if board.GetActive() == NilPiece() && len(board.pieces) != 0 && !board.IsWon(){
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

func (board *board) GetActive() Piece {
	return board.active
}

func (board *board) String() string{
	s := "|"
	h := 1
	v := 1
	for h <= 4 {
		for v <= 4 {
			if piece, ok := board.squares[Square{h, v}]; ok {
				s += piece.String() + "|"
			} else {
				s += NilPiece().String() + "|"
			}
			if v == 4{
				s += "\n"
			}
			v++
		}
		if h != 4 && v != 4 {
			s += "|"
		}
		h++
		v = 1
	}

	return s
}

func NewBoard(id string) Quarto {
	return &board {id, make(map[Square]Piece), make(map[Piece]bool), make(map[line]int), NilPiece()}
}

func NilBoard() Quarto {
	return NewBoard("nil")
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

	lines = append(lines, line{hline, piece.GetColor()})
	lines = append(lines, line{hline, piece.GetHeight()})
	lines = append(lines, line{hline, piece.GetShape()})
	lines = append(lines, line{hline, piece.GetTop()})

	lines = append(lines, line{vline, piece.GetColor()})
	lines = append(lines, line{vline, piece.GetHeight()})
	lines = append(lines, line{vline, piece.GetShape()})
	lines = append(lines, line{vline, piece.GetTop()})

	if dline != "" {
		lines = append(lines, line{dline, piece.GetColor()})
		lines = append(lines, line{dline, piece.GetHeight()})
		lines = append(lines, line{dline, piece.GetShape()})
		lines = append(lines, line{dline, piece.GetTop()})
	}

	return lines
}

type color string
type height string
type shape string
type top string

type Attribute interface{
	AttrStr() string
}

func (color color) AttrStr() string {
	return string(color)
}

func (height height) AttrStr() string {
	return string(height)
}

func (shape shape) AttrStr() string {
	return string(shape)
}

func (top top) AttrStr() string {
	return string(top)
}


const (
	WHITE color = "WHITE"
	BLACK color = "WHITE"
	TALL height = "TALL"
	SHORT height = "SHORT"
	ROUND shape = "ROUND"
	SQUARE shape = "SQUARE"
	FLAT top = "FLAT"
	HOLE top = "HOLE"
)

type piece struct {
	color color
	height height
	shape shape
	top top
}

func NewPiece(color color, height height, shape shape, top top) Piece{
	return &piece {color, height, shape, top}
}

func NilPiece() Piece {
	return &piece {}
}

type Piece interface {
	GetColor() Attribute
	GetHeight() Attribute
	GetShape() Attribute
	GetTop() Attribute
	String() string
}

func (piece *piece) GetColor() Attribute {
	return &piece.color
}

func (piece *piece) GetHeight() Attribute {
	return &piece.height
}

func (piece *piece) GetShape() Attribute {
	return &piece.shape
}

func (piece *piece) GetTop() Attribute {
	return &piece.top
}

func (piece *piece) String() string{
	var s string
	if piece.GetColor().AttrStr() == WHITE.AttrStr() {
		s += "W"
	} else if piece.GetColor().AttrStr() == BLACK.AttrStr() {
		s += "B"
	} else {
		s += " "
	}
	if piece.GetHeight().AttrStr() == TALL.AttrStr() {
		s += "T"
	} else if piece.GetHeight().AttrStr() == SHORT.AttrStr() {
		s += "S"
	} else {
		s += " "
	}
	if piece.GetShape().AttrStr() == ROUND.AttrStr() {
		s += "R"
	} else if piece.GetShape().AttrStr() == SQUARE.AttrStr() {
		s += "Q"
	} else {
		s += " "
	}
	if piece.GetTop().AttrStr() == FLAT.AttrStr() {
		s += "F"
	} else if piece.GetTop().AttrStr() == HOLE.AttrStr() {
		s += "H"
	} else {
		s += " "
	}

	return s
}