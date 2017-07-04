package rules

type board interface {
	Execute(move move)
	TakeBack(move move)
	GetPieceAt(square) piece
	GetLegalMoves(player player) []move
	GetEnPassantTarget() square
	OO(player) bool
	OOO(player) bool
}
