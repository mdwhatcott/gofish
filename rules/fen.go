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

	squares []Piece
	toMove  Player

	whiteCanCastleKingside  bool
	whiteCanCastleQueenside bool
	blackCanCastleKingside  bool
	blackCanCastleQueenside bool

	enPassantTargetSquare int
	fullMoveCount         int
	halfMoveCount         int
}

/**************************************************************************/

func ParseFEN(raw string) (this *FEN, err error) {
	this = &FEN{}
	fields := strings.Split(raw, " ")
	this.parseSquares(fields[0])
	this.parsePlayerToMove(fields[1])
	this.parseCastlingOpportunities(fields[2])
	// TODO: this.parseEnPassantTargetSquare(fields[3])
	err = this.parseHalfMoveCount(fields[4])
	if err != nil {
		return nil, err
	}
	err = this.parseFullMoveCount(fields[5])
	if err != nil {
		return nil, err
	}
	return this, nil
}
func (this *FEN) parseSquares(fenBoard string) {
	ranks := strings.Split(fenBoard, "/")
	this.squares = make([]Piece, 64)
	for r, rank := range ranks {
		square := 64 - ((r + 1) * 8)
		for _, c := range rank {
			if unicode.IsDigit(c) {
				square += int(c - '0')
			} else {
				this.squares[square] = Piece(string(c))
				square++
			}
		}
	}
}

func (this *FEN) parsePlayerToMove(player string) {
	if player == "w" {
		this.toMove = White
	} else {
		this.toMove = Black
	}
}
func (this *FEN) parseCastlingOpportunities(castle string) {
	for _, c := range castle {
		switch c {
		case 'K':
			this.whiteCanCastleKingside = true
		case 'k':
			this.blackCanCastleKingside = true
		case 'Q':
			this.whiteCanCastleQueenside = true
		case 'q':
			this.blackCanCastleQueenside = true
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

func PrepareFEN(squares map[int]Piece, game *Game) *FEN {
	return &FEN{
		buffer:                  new(bytes.Buffer),
		squares:                 copyMapToArray(squares),
		toMove:                  game.PlayerToMove(),
		whiteCanCastleKingside:  game.CanCastleKingside(White),
		whiteCanCastleQueenside: game.CanCastleQueenside(White),
		blackCanCastleKingside:  game.CanCastleKingside(Black),
		blackCanCastleQueenside: game.CanCastleQueenside(Black),
		enPassantTargetSquare:   0, // TODO
		fullMoveCount:           game.FullMoveCount(),
		halfMoveCount:           game.HalfMoveCount(),
	}
}

func copyMapToArray(squares map[int]Piece) (pieces []Piece) {
	for x := 0; x < 64; x++ {
		pieces = append(pieces, squares[x])
	}
	return pieces
}

func (this *FEN) String() string {
	this.recordPiecePlacement()
	this.space()
	this.recordGameMetadata()
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

	if this.whiteCanCastleKingside {
		this.buffer.WriteString("K")
	}
	if this.whiteCanCastleQueenside {
		this.buffer.WriteString("Q")
	}
	if this.blackCanCastleKingside {
		this.buffer.WriteString("k")
	}
	if this.blackCanCastleQueenside {
		this.buffer.WriteString("q")
	}

	if this.buffer.Len() == initial {
		this.buffer.WriteString("-")
	}
}
func (this *FEN) recordGameMetadata() {
	this.recordActiveColor()
	this.space()
	this.recordCastlingOpportunities()
	this.space()
	this.buffer.WriteString("-") // TODO: En-passant target square
	this.space()
	this.buffer.WriteString(strconv.Itoa(this.halfMoveCount))
	this.space()
	this.buffer.WriteString(strconv.Itoa(this.fullMoveCount))
}
func (this *FEN) space() {
	this.buffer.WriteString(" ")
}
