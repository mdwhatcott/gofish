package rules

func (this piece) calculateRangedPieceMovesFrom(from square, board board, offsetLines [][]square) (moves []move) {
	for _, line := range offsetLines {
		for _, offset := range line {
			if target := from.Offset(offset); !target.IsValidSquare() {
				break
			} else if board.GetPieceAt(target).Player() == this.Player() {
				break
			} else if targetPiece := board.GetPieceAt(target); targetPiece.Player() == Neither {
				moves = append(moves, move{Piece: this, From: from, To: target})
			} else if targetPiece.Player() == this.Player().Other() {
				moves = append(moves, move{
					Piece: this, From: from, To: target,
					Captured: targetPiece, CapturedOn: target,
				})
				break
			}
		}
	}
	return moves
}
