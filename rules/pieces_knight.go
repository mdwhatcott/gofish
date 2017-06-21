package rules
func (this piece) getKnightCoverageFrom(from square) (covered []square) {
	for _, offset := range knightMoveOffsets {
		target := from.Offset(offset)
		covered = append(covered, target)
	}
	return covered
}


func (this piece) calculateKnightMovesFrom(square square, board board) (moves []move) {
	for _, offset := range knightMoveOffsets {
		target := square.Offset(offset)
		if !target.IsValidSquare() {
			continue
		}
		targetPiece := board.GetPieceAt(target)
		if targetPiece.Player() == this.Player() {
			continue
		}
		moves = append(moves, move{
			Piece:      this,
			From:       square,
			To:         target,
			Captured:   targetPiece,
			CapturedOn: target,
		})
	}
	return moves
}
