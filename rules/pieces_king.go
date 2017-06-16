package rules

func (this Piece) calculateKingMovesFrom(square Square, board map[Square]Piece) (moves []Move) {
	for _, offset := range kingMoveOffsets {
		target := square.Offset(offset)
		if !target.IsValidSquare() {
			continue
		}
		if board[target].Player() == this.Player() {
			continue
		}
		//if board[target].ThreatenedBy(this.Player().Other()) {
		//	continue
		//}
		moves = append(moves, Move{Piece: this, From: square, To: target})
	}
	return moves
}
