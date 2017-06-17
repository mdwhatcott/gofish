package rules

type board interface {
	//InterpretSANMove(san string) []move
	Execute(move move)
	TakeBack(move move)
	GetPieceAt(square) piece
	IsUnderThreat(square, player) bool
	GetAvailableMoves(player player) []move
}
