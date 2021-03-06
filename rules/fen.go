package rules

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

const startingPositionFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Forsyth–Edwards Notation
// https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
type FEN struct {
	buffer *bytes.Buffer

	squares []piece
	toMove  player

	whiteOO  bool
	whiteOOO bool
	blackOO  bool
	blackOOO bool

	enPassantTargetSquare int
	fullMoveCount         int
	halfMoveCount         int
}

/**************************************************************************/

func ParseFEN(raw string) (this *FEN, err error) {
	this = &FEN{}
	fields := strings.Split(raw, " ")
	this.parseSquares(fields[0])
	this.parsePlayerToMove(fields)
	this.parseCastlingOpportunities(fields)
	// TODO: this.parseEnPassantTargetSquare(fields[3])
	// TODO: this.parseHalfMoveCount
	// TODO: this.parseFullMoveCount
	return this, nil
}
func (this *FEN) parseSquares(fenBoard string) {
	ranks := strings.Split(fenBoard, "/")
	this.squares = make([]piece, 64)
	for r, rank := range ranks {
		square := 64 - ((r + 1) * 8)
		for _, c := range rank {
			if unicode.IsDigit(c) {
				square += int(c - '0')
			} else {
				this.squares[square] = piece(string(c))
				square++
			}
		}
	}
}

func (this *FEN) parsePlayerToMove(fields []string) {
	this.toMove = White
	if len(fields) <= 1 {
		return
	}
	player := fields[1]
	if player == "b" {
		this.toMove = Black
	}
}
func (this *FEN) parseCastlingOpportunities(fields []string) {
	if len(fields) <= 2 {
		return
	}
	castle := fields[2]
	for _, c := range castle {
		switch c {
		case 'K':
			this.whiteOO = true
		case 'k':
			this.blackOO = true
		case 'Q':
			this.whiteOOO = true
		case 'q':
			this.blackOOO = true
		}
	}

}
func (this *FEN) parseHalfMoveCount(count string) (err error) {
	this.halfMoveCount, err = strconv.Atoi(count)
	return err
}

func (this *FEN) parseFullMoveCount(count string) (err error) {
	this.fullMoveCount, err = strconv.Atoi(count)
	return err
}

/**************************************************************************/

func PrepareFEN(squares map[square]piece, game *Game) *FEN {
	return &FEN{
		buffer:                new(bytes.Buffer),
		squares:               copyMapToArray(squares),
		toMove:                game.PlayerToMove(),
		whiteOO:               game.wOO,
		blackOO:               game.bOO,
		whiteOOO:              game.wOOO,
		blackOOO:              game.bOOO,
		enPassantTargetSquare: 0, // TODO
		halfMoveCount:         0, // TODO
		fullMoveCount:         1, // TODO
	}
}

func copyMapToArray(squares map[square]piece) []piece {
	pieces := make([]piece, len(squares))
	for square, piece := range squares {
		pieces[square.Int()] = piece
	}
	return pieces
}

func (this *FEN) String() string {
	if this.buffer.Len() == 0 {
		this.recordPiecePlacement()
		this.recordGameMetadata()
	}
	return this.buffer.String()
}

func (this *FEN) recordPiecePlacement() {
	for rank := 7; rank >= 0; rank-- {
		this.recordSquaresInRank(rank)

		if rank > 0 {
			this.buffer.WriteString("/")
		}
	}
}
func (this *FEN) recordSquaresInRank(rank int) {
	voids := 0
	for file := 0; file < 8; file++ {
		piece := this.squares[(rank*8)+file]
		if piece == Void {
			voids++
		} else if voids > 0 {
			this.buffer.WriteString(strconv.Itoa(voids))
			voids = 0
		}
		this.buffer.WriteString(string(piece))
	}
	if voids > 0 {
		this.buffer.WriteString(strconv.Itoa(voids))
	}
}
func (this *FEN) recordActiveColor() {
	if this.toMove == White {
		this.buffer.WriteString("w")
	} else {
		this.buffer.WriteString("b")
	}
}
func (this *FEN) recordCastlingOpportunities() {
	initial := this.buffer.Len()

	if this.whiteOO {
		this.buffer.WriteString("K")
	}
	if this.whiteOOO {
		this.buffer.WriteString("Q")
	}
	if this.blackOO {
		this.buffer.WriteString("k")
	}
	if this.blackOOO {
		this.buffer.WriteString("q")
	}

	if this.buffer.Len() == initial {
		this.buffer.WriteString("-")
	}
}
func (this *FEN) recordGameMetadata() {
	this.space()
	this.recordActiveColor()
	this.space()
	this.recordCastlingOpportunities()
	this.space()
	this.buffer.WriteString("-") // TODO: En-passant target square
	this.space()
	this.buffer.WriteString("0") // TODO: Half-move count
	this.space()
	this.buffer.WriteString("1") // TODO: Full-move count
}
func (this *FEN) space() {
	this.buffer.WriteString(" ")
}
