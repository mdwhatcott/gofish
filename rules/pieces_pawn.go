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
	return NewPawnMoveCalculator(this, from, board).calculateLegalMoves()
}

/**************************************************************************/

type PawnMoveCalculator struct {
	piece piece
	from  square
	board board

	advancement []square
	captures    []square
	promotions  []piece

	moves []move
}

func NewPawnMoveCalculator(piece piece, from square, board board) *PawnMoveCalculator {
	calculator := &PawnMoveCalculator{
		piece: piece,
		from:  from,
		board: board,
	}
	calculator.determinePossibilities()
	return calculator
}
func (this *PawnMoveCalculator) determinePossibilities() {
	if this.piece.Player() == White {
		this.determineWhitePawnPossibilities()
	} else {
		this.determineBlackPawnPossibilities()
	}
}
func (this *PawnMoveCalculator) determineWhitePawnPossibilities() {
	this.captures = whitePawnCaptureOffsets
	this.promotions = whitePawnPromotions
	if this.from.Rank() == "2" {
		this.advancement = whitePawnInitialAdvancementOffsets
	} else {
		this.advancement = whitePawnAdvancementOffsets
	}
}

func (this *PawnMoveCalculator) determineBlackPawnPossibilities() {
	this.captures = blackPawnCaptureOffsets
	this.promotions = blackPawnPromotions
	if this.from.Rank() == "7" {
		this.advancement = blackPawnInitialAdvancementOffsets
	} else {
		this.advancement = blackPawnAdvancementOffsets
	}
}

func (this *PawnMoveCalculator) calculateLegalMoves() []move {
	this.calculateAdvancements()
	this.calculateCaptures()
	return this.moves
}
func (this *PawnMoveCalculator) calculateAdvancements() {
	for _, offset := range this.advancement {
		if target := this.from.Offset(offset); this.canAdvanceTo(target) {
			this.calculateAdvancement(target)
		}
	}
}

func (this *PawnMoveCalculator) canAdvanceTo(target square) bool {
	return this.board.GetPieceAt(target) == Void
}
func (this *PawnMoveCalculator) calculateAdvancement(target square) {
	if this.canPromoteOnNextMove(target) {
		this.appendAdvancementPromotions(target)
	} else {
		this.appendAdvancement(target)
	}
}
func (this *PawnMoveCalculator) canPromoteOnNextMove(target square) bool {
	rank := target.Rank()
	return rank == "8" || rank == "1"
}

func (this *PawnMoveCalculator) appendAdvancementPromotions(target square) {
	for _, promotion := range this.promotions {
		this.moves = append(this.moves, move{
			Piece:     this.piece,
			From:      this.from,
			To:        target,
			Promotion: promotion,
		})
	}
}

func (this *PawnMoveCalculator) appendAdvancement(target square) {
	this.moves = append(this.moves, move{
		Piece: this.piece,
		From:  this.from,
		To:    target,
	})
}

func (this *PawnMoveCalculator) calculateCaptures() {
	for _, offset := range this.captures {
		targetSquare := this.from.Offset(offset)
		targetPiece := this.board.GetPieceAt(targetSquare)
		this.calculateCapture(targetSquare, targetPiece)
	}
}

func (this *PawnMoveCalculator) calculateCapture(targetSquare square, targetPiece piece) {
	isEnPassant := targetSquare == this.board.GetEnPassantTarget()
	if isEnPassant {
		if this.piece == WhitePawn {
			targetPiece = BlackPawn
		} else if this.piece == BlackPawn {
			targetPiece = WhitePawn
		}
	}

	if this.canCapture(targetPiece) {
		if this.canPromoteOnNextMove(targetSquare) {
			this.appendCapturingPromotions(targetSquare, targetPiece)
		} else {
			this.appendCapture(targetSquare, targetPiece, isEnPassant)
		}
	}
}

func (this *PawnMoveCalculator) canCapture(targetPiece piece) bool {
	return targetPiece.Player() == this.piece.Player().Other()
}

func (this *PawnMoveCalculator) appendCapturingPromotions(targetSquare square, targetPiece piece) {
	for _, promotion := range this.promotions {
		this.moves = append(this.moves, move{
			Piece:      this.piece,
			From:       this.from,
			To:         targetSquare,
			Captured:   targetPiece,
			CapturedOn: targetSquare,
			Promotion:  promotion,
		})
	}
}
func (this *PawnMoveCalculator) appendCapture(targetSquare square, targetPiece piece, isEnPassant bool) {
	this.moves = append(this.moves, move{
		Piece:      this.piece,
		From:       this.from,
		To:         targetSquare,
		Captured:   targetPiece,
		CapturedOn: targetSquare,
		EnPassant:  isEnPassant,
	})
}

/**************************************************************************/

func calculateEnPassantTarget(move move) square {
	if isBlackPawnDoubleAdvancement(move) {
		return IntSquare(move.To.Int() + 8)
	} else if isWhitePawnDoubleAdvancement(move) {
		return IntSquare(move.From.Int() + 8)
	} else {
		return IntSquare(-1)
	}
}
func isWhitePawnDoubleAdvancement(move move) bool {
	if move.Piece != WhitePawn {
		return false
	}
	if move.From.Rank() != "2" {
		return false
	}
	return move.To.Rank() == "4"
}
func isBlackPawnDoubleAdvancement(move move) bool {
	if move.Piece != BlackPawn {
		return false
	}
	if move.From.Rank() != "7" {
		return false
	}
	return move.To.Rank() == "5"
}
