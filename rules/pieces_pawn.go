package rules

func (this piece) calculatePawnMovesFrom(from square, board board) (moves []move) {
	var advancement []square
	var captures []square
	if this.Player() == White {
		captures = whitePawnCaptureOffsets
		if from.Rank() == "2" {
			advancement = whitePawnInitialAdvancementOffsets
		} else {
			advancement = whitePawnAdvancementOffsets
		}
	} else {
		captures = blackPawnCaptureOffsets
		if from.Rank() == "7" {
			advancement = blackPawnInitialAdvancementOffsets
		} else {
			advancement = blackPawnAdvancementOffsets
		}
	}

	for _, offset := range advancement {
		target := from.Offset(offset)
		if board.GetPieceAt(target) == Void {
			moves = append(moves, move{Piece: this, From: from, To: target})
		}
	}

	for _, offset := range captures {
		targetSquare := from.Offset(offset)
		targetPiece := board.GetPieceAt(targetSquare)
		if targetPiece.Player() == this.Player().Other() {
			moves = append(moves, move{
				Piece:      this,
				From:       from,
				To:         targetSquare,
				Captured:   targetPiece,
				CapturedOn: targetSquare,
			})
		}
	}
	return moves
}
