package rules

type Game struct {
	squares map[int]Piece
	player  Player

	fullMoveCount int
	halfMoveCount int

	whiteCanCastleKingside  bool // TODO: no need for fields, infer this from squares
	blackCanCastleKingside  bool
	whiteCanCastleQueenside bool
	blackCanCastleQueenside bool
}

func NewGame() *Game {
	game := &Game{squares: make(map[int]Piece, 64)}
	game.Reset()
	return game
}

func (this *Game) Reset() {
	this.LoadFEN(startingPositionFEN)
}

func (this *Game) LoadFEN(raw string) error {
	fen, err := ParseFEN(raw)
	if err != nil {
		return err
	}
	for s, piece := range fen.squares {
		this.squares[s] = piece
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

func (this *Game) PlayerToMove() Player {
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

func (this *Game) CanCastleKingside(player Player) bool {
	if player == White {
		return this.whiteCanCastleKingside
	} else {
		return this.blackCanCastleKingside
	}
}

func (this *Game) CanCastleQueenside(player Player) bool {
	if player == White {
		return this.whiteCanCastleQueenside
	} else {
		return this.blackCanCastleQueenside
	}
}

func (this *Game) DumpFEN() string {
	return PrepareFEN(this.squares, this).String()
}

func (this *Game) Move(move *Move) error {
	from := squareIndex(move.From)
	to := squareIndex(move.To)
	this.squares[to], this.squares[from] = this.squares[from], Void
	this.player = this.player.Alternate()
	return nil
}

func (this *Game) CalculateAvailableMoves() (moves []Move) {
	moves = []Move{}
	return moves
}
