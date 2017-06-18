package rules

import "unicode"

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

func (this piece) String() string {
	return string(this)
}

func (this piece) GetThreatsFrom(from square) []square {
	switch {
	case this.IsKing():
		return this.getKingThreatsFrom(from)
	default:
		return nil
	}
}

func (this piece) CalculateMovesFrom(square square, board board) (moves []move) {
	// TODO: if the player's king is currently in check and a move can't do anything to prevent check, that move is invalid
	// TODO: if the player's king is currently in check and a move can remove check by blocking or capturing the aggressor, that move is valid
	// TODO: if executing a move would cause discovered check, that move is invalid
	switch {
	case this.IsKing():
		return this.calculateKingMovesFrom(square, board)
	case this.IsKnight():
		return this.calculateKnightMovesFrom(square, board)
	case this.IsPawn():
		return this.calculatePawnMovesFrom(square, board)
	default:
		return nil
	}
}
func (this piece) IsPawn() bool   { return this == WhitePawn || this == BlackPawn }
func (this piece) IsKing() bool   { return this == WhiteKing || this == BlackKing }
func (this piece) IsKnight() bool { return this == WhiteKnight || this == BlackKnight }

func (this piece) Player() player {
	if this == Void {
		return Neither
	}
	if unicode.IsUpper(rune(this[0])) {
		return White
	}
	return Black
}
