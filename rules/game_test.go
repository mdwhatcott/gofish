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
	*LegalMovesFixture
}

func (this *GameFixture) Setup() {
	this.LegalMovesFixture = NewLegalGameMovesFixture(this.Fixture)
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

const positionAfter1A3 = "rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR b KQkq - 0 1"
