package rules

import "log"

type Game struct {
	squares map[square]piece
	player  player

	fullMoveCount int
	halfMoveCount int

	whiteCanCastleKingside  bool
	blackCanCastleKingside  bool
	whiteCanCastleQueenside bool
	blackCanCastleQueenside bool
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
	this.LoadFEN(startingPositionFEN)
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
	for s, piece := range fen.squares {
		this.squares[IntSquare(s)] = piece
	}
	this.player = fen.toMove
	this.blackCanCastleQueenside = fen.blackCanCastleQueenside
	this.blackCanCastleKingside = fen.blackCanCastleKingside
	this.whiteCanCastleQueenside = fen.whiteCanCastleQueenside
	this.whiteCanCastleKingside = fen.whiteCanCastleKingside
	this.fullMoveCount = fen.fullMoveCount
	this.halfMoveCount = fen.halfMoveCount
	return nil
}
func (this *Game) ExportFEN() string {
	return PrepareFEN(this.squares, this).String()
}

func (this *Game) PlayerToMove() player {
	return this.player
}

func (this *Game) IsOver() bool {
	return false
}

func (this *Game) FullMoveCount() int {
	return this.fullMoveCount
}

func (this *Game) HalfMoveCount() int {
	return this.halfMoveCount
}

func (this *Game) CanCastleKingside(player player) bool {
	if player == White {
		return this.whiteCanCastleKingside
	} else {
		return this.blackCanCastleKingside
	}
}

func (this *Game) CanCastleQueenside(player player) bool {
	if player == White {
		return this.whiteCanCastleQueenside
	} else {
		return this.blackCanCastleQueenside
	}
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

func (this *Game) Execute(move move) {
	this.squares[move.To], this.squares[move.From] = this.squares[move.From], Void
	this.player = this.player.Other()
}

func (this *Game) TakeBack(move move) {
	this.squares[move.From], this.squares[move.To] = this.squares[move.To], move.Captured
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
					moves = append(moves, move) // TODO: if this move puts the other player in check, mark the move as such
				}
				imagination.TakeBack(move)
			}
		}
	}
	return moves
}
