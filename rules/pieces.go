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

func (this piece) GetCoverageForPieceOn(from square, board board) []square {
	switch {
	case this.IsKing():
		return this.getKingCoverageFrom(from)
	case this.IsPawn():
		return this.getPawnCoverageFrom(from, board)
	case this.IsKnight():
		return this.getKnightCoverageFrom(from)
	case this.IsRook():
		return this.getRangedPieceCoverageFrom(from, board, rookMoveOffsetLines)
	case this.IsBishop():
		return this.getRangedPieceCoverageFrom(from, board, bishopMoveOffsetLines)
	case this.IsQueen():
		return this.getRangedPieceCoverageFrom(from, board, queenMoveOffsetLines)
	default:
		return nil
	}
}

func (this piece) CalculateMovesFrom(square square, board board) (moves []move) {
	// TODO: if the player's king is currently in check and a move can't do anything to prevent check, that move is invalid
	// TODO: if executing a move would cause discovered check, that move is invalid
	switch {
	case this.IsKing():
		return this.calculateKingMovesFrom(square, board)
	case this.IsKnight():
		return this.calculateKnightMovesFrom(square, board)
	case this.IsPawn():
		return this.calculatePawnMovesFrom(square, board)
	case this.IsBishop():
		return this.calculateRangedPieceMovesFrom(square, board, bishopMoveOffsetLines)
	case this.IsRook():
		return this.calculateRangedPieceMovesFrom(square, board, rookMoveOffsetLines)
	case this.IsQueen():
		return this.calculateRangedPieceMovesFrom(square, board, queenMoveOffsetLines)
	}
	return nil
}
func (this piece) IsPawn() bool   { return this == WhitePawn || this == BlackPawn }
func (this piece) IsKing() bool   { return this == WhiteKing || this == BlackKing }
func (this piece) IsKnight() bool { return this == WhiteKnight || this == BlackKnight }
func (this piece) IsBishop() bool { return this == WhiteBishop || this == BlackBishop }
func (this piece) IsRook() bool   { return this == WhiteRook || this == BlackRook }
func (this piece) IsQueen() bool  { return this == WhiteQueen || this == BlackQueen }

func (this piece) Player() player {
	if this == Void {
		return Neither
	}
	if unicode.IsUpper(rune(this[0])) {
		return White
	}
	return Black
}
