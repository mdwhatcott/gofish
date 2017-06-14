package rules

type Game struct {
	squares map[int]Piece
}

func NewGame() *Game {
	return &Game{
		squares: startingPosition(),
	}
}

func (this *Game) PlayerToMove() Player {
	return White
}

func (this *Game) IsOver() bool {
	return false
}

func (this *Game) FullMoveCount() int {
	return 1
}

func (this *Game) HalfMoveCount() int {
	return 0
}

func (this *Game) CanCastleKingside(player Player) bool {
	return true
}

func (this *Game) FEN() string {
	return PrepareFEN(this.squares, this).String()
}

func (this *Game) CanCastleQueenside(player Player) bool {
	return true
}

func (this *Game) CalculateAvailableMoves() (moves []Move) {
	moves = []Move{}
	return moves
}
