package rules

import "strings"

type piece string

const (
	Void        piece = ""
	WhitePawn   piece = "P"
	WhiteKnight piece = "N"
	WhiteBishop piece = "B"
	WhiteRook   piece = "R"
	WhiteQueen  piece = "Q"
	WhiteKing   piece = "K"
	BlackPawn   piece = "p"
	BlackKnight piece = "n"
	BlackBishop piece = "b"
	BlackRook   piece = "r"
	BlackQueen  piece = "q"
	BlackKing   piece = "k"
)

func (this piece) GetThreatsFrom(from square) []square {
	switch {
	case this.IsKing():
		return this.getKingThreatsFrom(from)
	default:
		return nil
	}
}

func (this piece) CalculateMovesFrom(square square, board board) (moves []move) {
	switch {
	case this.IsKing():
		return this.calculateKingMovesFrom(square, board)
	case this.IsKnight():
		return this.calculateKnightMovesFrom(square, board)
	default:
		return nil
	}
}

func (this piece) IsKing() bool {
	return this == WhiteKing || this == BlackKing
}

func (this piece) IsKnight() bool {
	return this == WhiteKnight || this == BlackKnight
}

func (this piece) Player() player {
	if this == Void {
		return Neither
	}
	if strings.ToUpper(string(this)) == string(this) {
		return White
	}
	return Black
}
