package quarto

import (
	"errors"
	"reflect"
)

type Square struct {
	H int
	V int
}

type line string

const (
	H1 line = "H1"
	H2 line = "H2"
	H3 line = "H3"
	H4 line = "H4"
	V1 line = "V1"
	V2 line = "V2"
	V3 line = "V3"
	V4 line = "V4"
	D1 line = "D1"
	D2 line = "D2"
)

type linePair struct {
	line line
	attr Attribute
}

type board struct {
	id string
	squares map[Square]Piece
	pieces map[Piece]bool
	lines map[linePair]int
	active Piece
}

type Quarto interface {
	IsWon() bool
	TakeTurn(toPlace Piece, h int, v int, active Piece) error
	Validate() (bool, error)
	GetActive() Piece
	String() string
	setSquares(squares map[Square]Piece)
	setPieces(pieces map[Piece]bool)
	setLines(lines map[linePair]int)
	setActive(piece Piece)
}

func (board *board) IsWon() bool{
	for _, count := range board.lines{
		if count >= 4 {
			return true
		}
	}
	return false
}

func (board *board) TakeTurn(toPlace Piece, h int, v int, active Piece) error {
	if board.pieces[toPlace] == true {
		return errors.New("the placed piece is already on the board")
	}
	if board.pieces[active] == true {
		return errors.New("the active piece is already on the board")
	}
	if _, ok := board.squares[Square {h,v}]; ok {
		return errors.New("the square is already occupied")
	}
	if reflect.DeepEqual(board.GetActive(), NilPiece()) && reflect.DeepEqual(board.GetActive(), toPlace) {
		return errors.New("must place the active piece")
	}

	board.squares[Square {h,v}] = toPlace
	board.pieces[toPlace] = true
	board.setActive(active)
	for _, line := range getLines(toPlace, h, v) {
		board.lines[line]++
	}
	_, error := board.Validate()
	if error != nil {
		return error
	}

	return nil
}

func (board *board) Validate() (bool, error) {
	if NilPiece() == board.GetActive() && len(board.squares) != 0 && !board.IsWon() {
		return false, errors.New("active piece required")
	}
	if board.GetActive() == NilPiece() && len(board.pieces) != 0 && !board.IsWon(){
		return false, errors.New("cannot set nil piece as active unless the game is won")
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

func (board *board) setLines(lines map[linePair]int) {
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
	return &board {id, make(map[Square]Piece), make(map[Piece]bool), make(map[linePair]int), NilPiece()}
}

func NilBoard() Quarto {
	return NewBoard("nil")
}

func InProgBoard(id string, squares map[Square]Piece, active Piece) Quarto {
	board := NewBoard(id)
	board.setSquares(squares)
	pieces := make(map[Piece]bool)
	lines := make(map[linePair]int)
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

func getLines(piece Piece, h int, v int) []linePair {
	var lines []linePair
	var hline line
	var vline line
	var dline line

	switch h {
	case 1:
		hline = H1
	case 2:
		hline = H2
	case 3:
		hline = H3
	case 4:
		hline = H4
	default:
		hline = ""
	}

	switch v {
	case 1:
		vline = V1
	case 2:
		vline = V2
	case 3:
		vline = V3
	case 4:
		vline = V4
	default:
		vline = ""
	}

	if h == v {
		dline = D1
	} else if h == 1 && v == 4 {
		dline = D2
	} else if h == 2 && v == 3 {
		dline = D2
	} else if h == 3 && v == 2 {
		dline = D2
	} else if h == 4 && v == 1 {
		dline = D2
	}

	lines = append(lines, linePair{hline, piece.GetColor()})
	lines = append(lines, linePair{hline, piece.GetHeight()})
	lines = append(lines, linePair{hline, piece.GetShape()})
	lines = append(lines, linePair{hline, piece.GetTop()})

	lines = append(lines, linePair{vline, piece.GetColor()})
	lines = append(lines, linePair{vline, piece.GetHeight()})
	lines = append(lines, linePair{vline, piece.GetShape()})
	lines = append(lines, linePair{vline, piece.GetTop()})

	if dline != "" {
		lines = append(lines, linePair{dline, piece.GetColor()})
		lines = append(lines, linePair{dline, piece.GetHeight()})
		lines = append(lines, linePair{dline, piece.GetShape()})
		lines = append(lines, linePair{dline, piece.GetTop()})
	}

	return lines
}

type color string
type height string
type shape string
type top string

type Attribute interface{
	String() string
	attrStr() string
}

func (color color) attrStr() string {
	return color.String()
}

func (color color) String() string {
	return string(color)
}

func (height height) attrStr() string {
	return height.String()
}

func (height height) String() string {
	return string(height)
}

func (shape shape) attrStr() string {
	return shape.String()
}

func (shape shape) String() string {
	return string(shape)
}

func (top top) attrStr() string {
	return top.String()
}

func (top top) String() string {
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
	return piece.color
}

func (piece *piece) GetHeight() Attribute {
	return piece.height
}

func (piece *piece) GetShape() Attribute {
	return piece.shape
}

func (piece *piece) GetTop() Attribute {
	return piece.top
}

func (piece *piece) String() string{
	var s string
	if piece.GetColor() == WHITE {
		s += "W"
	} else if piece.GetColor() == BLACK {
		s += "B"
	} else {
		s += " "
	}
	if piece.GetHeight() == TALL {
		s += "T"
	} else if piece.GetHeight() == SHORT {
		s += "S"
	} else {
		s += " "
	}
	if piece.GetShape() == ROUND {
		s += "R"
	} else if piece.GetShape() == SQUARE {
		s += "Q"
	} else {
		s += " "
	}
	if piece.GetTop() == FLAT {
		s += "F"
	} else if piece.GetTop() == HOLE {
		s += "H"
	} else {
		s += " "
	}

	return s
}