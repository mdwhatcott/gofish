package rules

func (this piece) getKingThreatsFrom(from square) (targets []square) {
	for _, offset := range kingMoveOffsets {
		if target := from.Offset(offset); target.IsValidSquare() {
			targets = append(targets, target)
		}
	}
	return targets
}

func (this piece) calculateKingMovesFrom(from square, board board) (moves []move) {
	for _, offset := range kingMoveOffsets {
		target := from.Offset(offset)
		if !target.IsValidSquare() {
			continue
		}
		targetPiece := board.GetPieceAt(target)
		if targetPiece.Player() == this.Player() {
			continue
		}
		if board.IsUnderThreat(target, this.Player().Other()) {
			continue
		}
		moves = append(moves, move{
			Piece: this,

			From: from,
			To:   target,

			Captured:   targetPiece,
			CapturedOn: target,
		})
	}
	return moves
}
