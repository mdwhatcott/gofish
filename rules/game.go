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
	game := &Game{squares: make(map[square]piece, 64)}
	game.Reset()
	return game
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

/* board interface: *******************************************************/

func (this *Game) Execute(move move) {
	this.squares[move.To], this.squares[move.From] = this.squares[move.From], Void
	this.player = this.player.Other()
}

func (this *Game) TakeBack(move move) {
	this.squares[move.From], this.squares[move.To] = this.squares[move.To], Void
	this.player = this.player.Other()
}

func (this *Game) GetPieceAt(square square) piece {
	return this.squares[square]
}

func (this *Game) IsUnderThreat(subject square, aggressor player) bool {
	for square, piece := range this.squares {
		if piece.Player() == aggressor {
			for _, covered := range piece.GetThreatsFrom(square) {
				if covered == subject {
					return true
				}
			}
		}
	}
	return false
}

func (this *Game) GetLegalMoves(player player) (moves []move) {
	for square, piece := range this.squares {
		if piece.Player() == player {
			for _, move := range piece.CalculateMovesFrom(square, this) {
				moves = append(moves, move)
			}
		}
	}
	return moves
}
