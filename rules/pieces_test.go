package rules

import (
	"sort"

	"github.com/mdwhatcott/gofish/console"
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

func (this *LegalPieceMovesFixture) assertAllLegalMoves(actualMoves []move, expectedMovesSAN ...string) {
	this.So(actualMoves, should.HaveLength, len(expectedMovesSAN))
	actualMovesSan := []string{}
	for _, move := range actualMoves {
		this.So(move.To.String(), should.NotResemble, move.From.String())
		actualMovesSan = append(actualMovesSan, move.String())
	}
	sort.Strings(actualMovesSan)
	sort.Strings(expectedMovesSAN)
	this.So(actualMovesSan, should.Resemble, expectedMovesSAN)
	if this.Failed() {
		picture := console.NewBoard()
		picture.Setup(this.game.ExportFEN())
		this.Println(console.NewCoordinateBoard(picture.String()))
	}
}

func (this *LegalPieceMovesFixture) assertLegalPieceMoves(
	position string, from string, piece piece, expectedPieceMovesSAN ...string) {

	if len(expectedPieceMovesSAN) == 0 {
		expectedPieceMovesSAN = []string{}
	}

	this.game.MustLoadFEN(position)

	moves := this.game.GetLegalMoves(piece.Player())
	actualPieceMoves := filterMovesByPiece(moves, piece)
	this.So(actualPieceMoves, should.HaveLength, len(expectedPieceMovesSAN))

	actualTargets := []string{}
	for _, move := range actualPieceMoves {
		this.So(move.Piece, should.Equal, piece)
		this.So(move.From.String(), should.Equal, from)
		this.So(move.To.String(), should.NotResemble, move.From.String())
		actualTargets = append(actualTargets, move.String())
	}
	sort.Strings(actualTargets)
	sort.Strings(expectedPieceMovesSAN)
	this.So(actualTargets, should.Resemble, expectedPieceMovesSAN)
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
