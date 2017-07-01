package rules

import "log"

type Game struct {
	squares map[square]piece
	player  player

	enPassantTarget square
}

func NewGame() *Game {
	game := new(Game)
	game.initialize(make(map[square]piece, 64))
	game.Reset()
	return game
}

func (this *Game) initialize(state map[square]piece) {
	this.squares = state
}

func (this *Game) Reset() {
	this.player = White
	this.enPassantTarget = IntSquare(-1)
	this.MustLoadFEN(startingPositionFEN)
}

func (this *Game) MustLoadFEN(raw string) {
	if err := this.LoadFEN(raw); err != nil {
		log.Panicf("Could not load fen [%s] because of err:", err)
	}
}
func (this *Game) LoadFEN(raw string) error {
	fen, err := ParseFEN(raw)
	if err != nil {
		return err
	}
	squares := map[square]piece{}
	for s, piece := range fen.squares {
		squares[IntSquare(s)] = piece
	}
	this.initialize(squares)
	this.player = fen.toMove
	return nil
}
func (this *Game) ExportFEN() *FEN {
	return PrepareFEN(this.squares, this)
}

func (this *Game) PlayerToMove() player {
	return this.player
}

func (this *Game) GetEnPassantTarget() square {
	return this.enPassantTarget
}

func (this *Game) IsOver() bool {
	return this.IsInCheckmate(White) || this.IsInCheckmate(Black)
}

func (this *Game) IsInCheckmate(player player) bool {
	return this.IsInCheck(player) && len(this.GetLegalMoves(player)) == 0
}
func (this *Game) IsInCheck(player player) bool {
	kingSquare := this.findKing(player)
	return this.SquareIsCoveredBy(kingSquare, player.Other())
}
func (this *Game) findKing(player player) square {
	for square, piece := range this.squares {
		if piece.IsKing() && piece.Player() == player {
			return square
		}
	}
	return IntSquare(-1)
}

func (this *Game) Attempt(moveSAN string) move {
	legalMoves := this.GetLegalMoves(this.PlayerToMove())
	for _, move := range legalMoves {
		if move.SAN() == moveSAN {
			this.Execute(move)
			return move
		}
	}
	return move{}
}

func (this *Game) Execute(move move) {
	this.squares[move.To], this.squares[move.From] = this.squares[move.From], Void

	if move.Promotion != Void {
		this.squares[move.To] = move.Promotion
	} else if move.EnPassant {
		this.squares[Square(move.To.File()+move.From.Rank())] = Void
	}

	this.enPassantTarget = calculateEnPassantTarget(move)
	this.player = this.player.Other()
}

func (this *Game) TakeBack(move move) {
	this.squares[move.From], this.squares[move.To] = this.squares[move.To], move.Captured

	if move.Promotion != Void {
		this.squares[move.From] = move.Piece
	} else if move.EnPassant {
		this.squares[move.To] = Void
		this.squares[Square(move.To.File()+move.From.Rank())] = move.Captured
		this.enPassantTarget = move.To
	}

	this.player = this.player.Other()
}

func (this *Game) GetPieceAt(square square) piece {
	return this.squares[square]
}

func (this *Game) SquareIsCoveredBy(subject square, aggressor player) bool {
	for square, piece := range this.squares {
		if piece.Player() == aggressor {
			for _, covered := range piece.GetCoverageForPieceOn(square, this) {
				if covered == subject {
					return true
				}
			}
		}
	}
	return false
}

func (this *Game) copyGame() *Game {
	game := new(Game)
	game.initialize(this.copySquares())
	// TODO: other game state???
	return game
}

func (this *Game) copySquares() map[square]piece {
	other := make(map[square]piece, 64)
	for square, piece := range this.squares {
		other[square] = piece
	}
	return other
}

func (this *Game) GetLegalMoves(player player) (moves []move) {
	imagination := this.copyGame()

	for square, piece := range this.squares {
		if piece.Player() == player {
			for _, move := range piece.CalculateMovesFrom(square, this) {
				imagination.Execute(move)
				if !imagination.IsInCheck(player) {
					if imagination.IsInCheckmate(player.Other()) {
						move.Checkmate = true
					} else if imagination.IsInCheck(player.Other()) {
						move.Check = true
					}
					moves = append(moves, move)
				}
				imagination.TakeBack(move)
			}
		}
	}
	return moves
}
