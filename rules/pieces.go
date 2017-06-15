package rules

import "strings"

type Piece string

const (
	Void        Piece = ""
	WhitePawn   Piece = "P"
	WhiteKnight Piece = "N"
	WhiteBishop Piece = "B"
	WhiteRook   Piece = "R"
	WhiteQueen  Piece = "Q"
	WhiteKing   Piece = "K"
	BlackPawn   Piece = "p"
	BlackKnight Piece = "n"
	BlackBishop Piece = "b"
	BlackRook   Piece = "r"
	BlackQueen  Piece = "q"
	BlackKing   Piece = "k"
)

var kingMoveOffsets = []Square{
	{File: -1, Rank: 1}, {File: 0, Rank: 1}, {File: 1, Rank: 1},
	{File: -1, Rank: 0}, {File: 1, Rank: 0},
	{File: -1, Rank: -1}, {File: 0, Rank: -1}, {File: 1, Rank: -1},
}

func (this Piece) CalculateMovesFrom(square Square) (moves []Move) {
	if this.IsKing() {
		for _, offset := range kingMoveOffsets {
			if target := square.Offset(offset); target.IsValid() {
				moves = append(moves, Move{Piece: this, From: square, To: target})
			}
		}
	}
	return moves
}

func (this Piece) IsKing() bool {
	return this == WhiteKing || this == BlackKing
}

func (this Piece) Player() Player {
	if this == Void {
		return Neutral
	}
	if strings.ToUpper(string(this)) == string(this) {
		return White
	}
	return Black
}
