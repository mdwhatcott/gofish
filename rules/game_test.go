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
	this.So(this.game.IsInCheck(White), should.BeFalse)
	this.So(this.game.IsInCheck(Black), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, White)
	this.So(this.game.ExportFEN().String(), should.StartWith, startingPositionFEN)
}

func (this *GameFixture) TestGameConditionsAfterFirstPawnMoveByWhite() {
	move := move{From: Square("a2"), To: Square("a3"), Piece: WhitePawn}
	this.game.Execute(move)

	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, Black)
	this.So(this.game.ExportFEN().String(), should.StartWith, positionAfter1A3)
}
func (this *GameFixture) TestGameConditionsAfterTakingBackTheFirstMove() {
	move := move{From: Square("a2"), To: Square("a3"), Piece: WhitePawn}
	this.game.Execute(move)
	this.game.TakeBack(move)
	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, White)
	this.So(this.game.ExportFEN().String(), should.StartWith, startingPositionFEN)
}

func (this *GameFixture) TestTakeBackPromotion() {
	this.game.MustLoadFEN("8/8/8/8/8/7k/K3p3/8")
	move := move{Piece: BlackPawn, From: Square("e2"), To: Square("e1"), Promotion: BlackQueen}
	this.game.Execute(move)
	this.game.TakeBack(move)
	this.So(this.game.IsInCheck(White), should.BeFalse)
}

func (this *GameFixture) TestLoadFEN() {
	const kingsOnBackRanks = "4k3/8/8/8/8/8/8/4K3"
	err := this.game.LoadFEN(kingsOnBackRanks)
	this.So(err, should.BeNil)
	this.So(this.game.ExportFEN().String(), should.StartWith, kingsOnBackRanks)
}
func (this *GameFixture) TestCheckMate() {
	this.assertInCheckMate("7k/5KQ1/8/8/8/8/8/8", Black)
}
func (this *GameFixture) assertInCheckMate(position string, player player) {
	this.game.MustLoadFEN(position)
	this.So(this.game.IsInCheckmate(player), should.BeTrue)
	this.So(this.game.IsOver(), should.BeTrue)
}
func (this *GameFixture) TestCheckByPawnAggressor() {
	this.assertInCheck("3k4/4P3/8/8/8/8/8/4K3", Black)
	this.assertInCheck("3k4/2P5/8/8/8/8/8/4K3", Black)
	this.assertNotInCheck("2k5/2P5/8/8/8/8/8/4K3", Black)

	this.assertInCheck("2k5/8/8/8/8/8/3p4/4K3", White)
	this.assertInCheck("2k5/8/8/8/8/8/5p2/4K3", White)
	this.assertNotInCheck("2k5/8/8/8/8/8/4p3/4K3", White)
}
func (this *GameFixture) TestCheckByKnightAggressor() {
	this.assertInCheck("2k5/8/3N4/8/8/8/8/4K3", Black)
	this.assertInCheck("2k5/4N3/8/8/8/8/8/4K3", Black)
	this.assertNotInCheck("3k4/4N3/8/8/8/8/8/4K3", Black)

	this.assertInCheck("2K5/8/3n4/8/8/8/8/4k3", White)
	this.assertInCheck("2K5/4n3/8/8/8/8/8/4k3", White)
	this.assertNotInCheck("3K4/4n3/8/8/8/8/8/4k3", White)
}
func (this *GameFixture) TestCheckByRookAggressor() {
	this.assertInCheck("3k4/8/8/8/8/8/3R4/3K4", Black)
	this.assertInCheck("8/8/8/8/8/8/3r3K/3k4", White)
	this.assertNotInCheck("8/8/8/8/8/8/3R2pk/3K4", Black)
	this.assertNotInCheck("8/8/7K/8/8/8/3r4/3k4", White)
}
func (this *GameFixture) TestCheckByBishopAggressor() {
	this.assertInCheck("7k/8/8/8/P7/8/8/B2K4", Black)
	this.assertInCheck("b2k4/8/8/7p/8/8/8/7K", White)
	this.assertNotInCheck("8/8/8/8/P7/8/k7/B2K4", Black)
	this.assertNotInCheck("b2k4/K7/8/7p/8/8/8/8", White)
}
func (this *GameFixture) TestCheckByQueenAggressor() {
	this.assertInCheck("3k4/3q4/8/8/8/8/8/3K4", White)
	this.assertInCheck("3k4/8/8/7q/8/8/8/3K4", White)
	this.assertInCheck("3k4/8/8/8/8/8/8/2qK4", White)
	this.assertNotInCheck("3k4/3q4/8/8/8/8/8/4K3", White)

	this.assertInCheck("3K4/3Q4/8/8/8/8/8/3k4", Black)
	this.assertInCheck("3K4/8/8/7Q/8/8/8/3k4", Black)
	this.assertInCheck("3K4/8/8/8/8/8/8/2Qk4", Black)
	this.assertNotInCheck("3K4/3Q4/8/8/8/8/8/4k3", Black)
}
func (this *GameFixture) assertInCheck(position string, color player) {
	this.game.MustLoadFEN(position)
	this.So(this.game.IsInCheck(color), should.BeTrue)
}
func (this *GameFixture) assertNotInCheck(position string, color player) {
	this.game.MustLoadFEN(position)
	this.So(this.game.IsInCheck(color), should.BeFalse)
}

const positionAfter1A3 = "rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR"
