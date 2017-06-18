package rules

import (
	"sort"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"github.com/mdwhatcott/gofish/console"
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
	if len(expectedPotentialTargetSquares) == 0 {
		expectedPotentialTargetSquares = []string{}
	}

	this.game.MustLoadFEN(position)

	moves := filterMovesByPiece(this.game.GetLegalMoves(piece.Player()), piece)

	this.So(moves, should.HaveLength, len(expectedPotentialTargetSquares))

	actualTargets := []string{}
	for _, move := range moves {
		this.So(move.Piece, should.Equal, piece)
		this.So(move.From.String(), should.Equal, from)
		this.So(move.To.String(), should.NotResemble, move.From.String())
		actualTargets = append(actualTargets, move.To.String())
	}
	sort.Strings(actualTargets)
	sort.Strings(expectedPotentialTargetSquares)
	this.So(actualTargets, should.Resemble, expectedPotentialTargetSquares)
	if this.Failed() {
		picture := console.NewBoard()
		picture.Setup(position)
		this.Println(console.NewCoordinateBoard(picture.String()))
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
