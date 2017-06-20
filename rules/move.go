package rules

import (
	"bytes"
	"strings"
)

type move struct {
	Piece   piece
	From    square
	To      square
	NotFrom []square

	Promotion piece

	Captured   piece
	CapturedOn square

	Castles bool // When true, the Piece will be the king. Which side is inferred by the direction moved by the king.

	Check     bool
	Checkmate bool
}

func (this move) String() string {
	return this.SAN()
}

func (this move) SAN() string {
	buffer := new(bytes.Buffer)

	if this.Castles {
		this.determineCastleSide(buffer)
	} else if this.Piece.IsPawn() {
		this.pawnMove(buffer)
	} else {
		this.pieceMove(buffer)
	}

	this.accountForCheckAndMate(buffer)
	return buffer.String()
}
func (this move) determineCastleSide(buffer *bytes.Buffer) {
	buffer.WriteString("O-O")
	if this.To.File() == "c" {
		buffer.WriteString("-O")
	}
}
func (this move) pawnMove(buffer *bytes.Buffer) {
	if this.Captured != Void {
		buffer.WriteString(this.From.File())
		buffer.WriteString(takes)
	}

	buffer.WriteString(this.To.String())

	if this.Promotion != Void {
		buffer.WriteString(promotes)
		buffer.WriteString(string(this.Promotion))
	}
}
func (this move) pieceMove(buffer *bytes.Buffer) {
	buffer.WriteString(strings.ToUpper(string(this.Piece)))

	if this.isAmbiguous() {
		this.disambiguate(buffer)
	}

	if this.Captured != Void {
		buffer.WriteString(takes)
	}
	buffer.WriteString(this.To.String())
}
func (this move) isAmbiguous() bool {
	return len(this.NotFrom) > 0
}
func (this move) disambiguate(buffer *bytes.Buffer) {
	if this.ambiguitiesOccupyUniqueFiles() {
		buffer.WriteString(this.From.File())
	} else if this.ambiguitiesOccupyUniqueRanks() {
		buffer.WriteString(this.From.Rank())
	} else {
		buffer.WriteString(this.From.String())
	}
}
func (this move) ambiguitiesOccupyUniqueFiles() bool {
	return append(squares(this.NotFrom), this.From).filesAreUnique()
}
func (this move) ambiguitiesOccupyUniqueRanks() bool {
	return append(squares(this.NotFrom), this.From).ranksAreUnique()
}

func (this move) accountForCheckAndMate(buffer *bytes.Buffer) {
	if this.Check {
		buffer.WriteString(check)
	} else if this.Checkmate {
		buffer.WriteString(checkmate)
	}
}

const (
	takes     = "x"
	promotes  = "="
	check     = "+"
	checkmate = "#"
)
