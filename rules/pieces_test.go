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
	for _, target := range expectedPotentialTargetSquares {
		this.So(moves, should.Contain, move{Piece: piece, From: Square(from), To: Square(target)})
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
