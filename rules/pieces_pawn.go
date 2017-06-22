package rules

func (this piece) getPawnCoverageFrom(from square, board board) (covered []square) {
	var captures = whitePawnCaptureOffsets
	if this.Player() == Black {
		captures = blackPawnCaptureOffsets
	}
	for _, offset := range captures {
		covered = append(covered, from.Offset(offset))
	}
	return covered
}

func (this piece) calculatePawnMovesFrom(from square, board board) (moves []move) {
	var advancement []square
	var captures []square
	var promotions []piece
	if this.Player() == White {
		captures = whitePawnCaptureOffsets
		promotions = whitePawnPromotions
		if from.Rank() == "2" {
			advancement = whitePawnInitialAdvancementOffsets
		} else {
			advancement = whitePawnAdvancementOffsets
		}
	} else {
		captures = blackPawnCaptureOffsets
		promotions = blackPawnPromotions
		if from.Rank() == "7" {
			advancement = blackPawnInitialAdvancementOffsets
		} else {
			advancement = blackPawnAdvancementOffsets
		}
	}

	for _, offset := range advancement {
		target := from.Offset(offset)
		if board.GetPieceAt(target) == Void {
			if rank := target.Rank(); rank == "8" || rank == "1" {
				for _, promotion := range promotions {
					moves = append(moves, move{Piece: this, From: from, To: target, Promotion: promotion})
				}
			} else {
				moves = append(moves, move{Piece: this, From: from, To: target})
			}
		}
	}

	for _, offset := range captures {
		targetSquare := from.Offset(offset)
		targetPiece := board.GetPieceAt(targetSquare)
		if targetPiece.Player() == this.Player().Other() {
			if rank := targetSquare.Rank(); rank == "8" || rank == "1" {
				for _, promotion := range promotions {
					moves = append(moves, move{
						Piece:      this,
						From:       from,
						To:         targetSquare,
						Captured:   targetPiece,
						CapturedOn: targetSquare,
						Promotion:  promotion,
					})
				}
			} else {
				moves = append(moves, move{
					Piece:      this,
					From:       from,
					To:         targetSquare,
					Captured:   targetPiece,
					CapturedOn: targetSquare,
				})
			}
		}
	}
	return moves
}
