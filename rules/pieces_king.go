package rules

func (this piece) getKingCoverageFrom(from square) (targets []square) {
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
		moves = append(moves, move{
			Piece: this,

			From: from,
			To:   target,

			Captured:   targetPiece,
			CapturedOn: target,
		})
	}
	if board.OO(this.Player()) {
		moves = append(moves, move{
			Piece:   this,
			From:    from,
			To:      IntSquare(from.Int() + 2),
			Castles: true,
		})
	}
	if board.OOO(this.Player()) {
		moves = append(moves, move{
			Piece:   this,
			From:    from,
			To:      IntSquare(from.Int() - 2),
			Castles: true,
		})
	}
	return moves
}
