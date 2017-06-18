package rules

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestGameFixture(t *testing.T) {
	gunit.Run(new(GameFixture), t)
}

type GameFixture struct {
	*gunit.Fixture
	*LegalPieceMovesFixture
}

func (this *GameFixture) Setup() {
	this.LegalPieceMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

func (this *GameFixture) TestStartingGameConditions() {
	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, White)
	this.So(this.game.FullMoveCount(), should.Equal, 1)
	this.So(this.game.HalfMoveCount(), should.Equal, 0)
	this.So(this.game.CanCastleKingside(White), should.BeTrue)
	this.So(this.game.CanCastleKingside(Black), should.BeTrue)
	this.So(this.game.CanCastleQueenside(White), should.BeTrue)
	this.So(this.game.CanCastleQueenside(Black), should.BeTrue)
	this.So(this.game.ExportFEN(), should.Equal, startingPositionFEN)
}

func (this *GameFixture) TestGameConditionsAfterFirstPawnMoveByWhite() {
	move := move{From: Square("a2"), To: Square("a3"), Piece: WhitePawn}
	this.game.Execute(move)

	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, Black)
	this.So(this.game.FullMoveCount(), should.Equal, 1)
	this.So(this.game.HalfMoveCount(), should.Equal, 0) // pawn move
	this.So(this.game.ExportFEN(), should.Equal, positionAfter1A3)
}
func (this *GameFixture) TestGameConditionsAfterTakingBackTheFirstMove() {
	move := move{From: Square("a2"), To: Square("a3"), Piece: WhitePawn}
	this.game.Execute(move)
	this.game.TakeBack(move)
	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, White)
	this.So(this.game.FullMoveCount(), should.Equal, 1)
	this.So(this.game.HalfMoveCount(), should.Equal, 0) // pawn move
	this.So(this.game.ExportFEN(), should.Equal, startingPositionFEN)
}

func (this *GameFixture) TestLoadFEN() { // TODO: test many more pieces and scenarios
	const kingsOnBackRanks = "4k3/8/8/8/8/8/8/4K3 w - - 0 1"
	err := this.game.LoadFEN(kingsOnBackRanks)
	this.So(err, should.BeNil)
	this.So(this.game.ExportFEN(), should.Equal, kingsOnBackRanks)
}
func (this *GameFixture) TestLegalFirstMovesForWhite() {
	moves := this.game.GetLegalMoves(this.game.PlayerToMove())
	this.assertAllLegalMoves(moves,
		"a3", "a4",
		"b3", "b4",
		"c3", "c4",
		"d3", "d4",
		"e3", "e4",
		"f3", "f4",
		"g3", "g4",
		"h3", "h4",
		"Na3", "Nc3",
		"Nf3", "Nh3",
	)
}

func (this *GameFixture) TestLegalFirstMovesForBlack() {
	this.game.Execute(move{From: Square("a2"), To: Square("a3"), Piece: WhitePawn})
	moves := this.game.GetLegalMoves(this.game.PlayerToMove())
	this.assertAllLegalMoves(moves,
		"a6", "a5",
		"b6", "b5",
		"c6", "c5",
		"d6", "d5",
		"e6", "e5",
		"f6", "f5",
		"g6", "g5",
		"h6", "h5",
		"Na6", "Nc6",
		"Nf6", "Nh6",
	)
}

const positionAfter1A3 = "rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR b KQkq - 0 1"
