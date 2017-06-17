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

func (this *LegalPieceMovesFixture) assertLegalPieceMoves(
	position string, from string, piece piece, expectedPotentialTargetSquares ...string) {

	this.game.MustLoadFEN(position)

	moves := filterMovesByPiece(this.game.GetAvailableMoves(piece.Player()), piece)

	this.So(moves, should.HaveLength, len(expectedPotentialTargetSquares))

	actualTargets := []string{}
	for _, move := range moves {
		this.So(move.Piece, should.Equal, piece)
		this.So(move.From.String(), should.Equal, from)
		this.So(move.To.String(), should.NotResemble, move.From.String())
		actualTargets = append(actualTargets, move.To.String())
	}
	for _, target := range expectedPotentialTargetSquares {
		this.So(actualTargets, should.Contain, target)
	}
}

func filterMovesByPiece(moves []move, piece piece) (filtered []move) {
	for _, move := range moves {
		if move.Piece == piece {
			filtered = append(filtered, move)
		}
	}
	return filtered
}
