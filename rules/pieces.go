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

func (this Piece) CalculateMovesFrom(square Square, board map[Square]Piece) (moves []Move) {
	if this.IsKing() {
		return this.calculateKingMovesFrom(square, board)
	}
	return nil
}

func (this Piece) IsKing() bool {
	return this == WhiteKing || this == BlackKing
}

func (this Piece) Player() Player {
	if this == Void {
		return Neither
	}
	if strings.ToUpper(string(this)) == string(this) {
		return White
	}
	return Black
}
