package rules

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type LegalPieceMovesFixture struct {
	*gunit.Fixture

	game *Game
}

func NewLegalGameMovesFixture(fixture *gunit.Fixture) *LegalPieceMovesFixture {
	return &LegalPieceMovesFixture{Fixture: fixture, game: NewGame()}
}

func (this *LegalPieceMovesFixture) assertLegalPieceMoves(position string, piece Piece, from string, targets []string) {
	this.game.MustLoadFEN(position)
	moves := filterMovesByPiece(this.game.CalculateAvailableMoves(), piece)
	this.So(moves, should.HaveLength, len(targets))
	for _, to := range targets {
		this.So(moves, should.Contain, Move{Piece: piece, From: ParseSquare(from), To: ParseSquare(to)})
	}
}

func filterMovesByPiece(moves []Move, piece Piece) (filtered []Move) {
	for _, move := range moves {
		if move.Piece == piece {
			filtered = append(filtered, move)
		}
	}
	return filtered
}
