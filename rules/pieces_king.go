package rules

func (this piece) getKingThreatsFrom(from square) (targets []square) {
	for _, offset := range kingMoveOffsets {
		if target := from.Offset(offset); target.IsValidSquare() {
			targets = append(targets, target)
		}
	}
	return targets
}

func (this piece) calculateKingMovesFrom(square square, board board) (moves []move) {
	for _, offset := range kingMoveOffsets {
		target := square.Offset(offset)
		if !target.IsValidSquare() {
			continue
		}
		if board.GetPieceAt(target).Player() == this.Player() {
			continue
		}
		if board.IsUnderThreat(target, this.Player().Other()) {
			continue
		}
		moves = append(moves, move{Piece: this, From: square, To: target})
	}
	return moves
}
