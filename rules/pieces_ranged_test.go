package rules

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestRangedPieceMoveFixture(t *testing.T) {
	gunit.Run(new(RangedPieceMoveFixture), t)
}

type RangedPieceMoveFixture struct {
	*gunit.Fixture
	*LegalMovesFixture
}

func (this *RangedPieceMoveFixture) Setup() {
	this.LegalMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

const (
	unhinderedRook               = "5K1k/8/7p/8/8/8/8/R7 w - - 0 1"
	rookWithCaptureOpportunities = "5K1k/8/7p/b7/8/8/8/R2b4 w - - 0 1"
	rookWithLimitedMovement      = "5K1k/8/7p/B7/8/8/8/R2B4 w - - 0 1"
	bishopWithMaximumRange       = "k7/8/p7/8/3B4/8/8/7K w - - 0 1"
	// TODO: rookWithNoMovement
	// TODO: bishopWithLimitedRange
	// TODO: bishopWithNoMovement
	// TODO: queenWithMaximumRange
	// TODO: queenWithLimitedMovement
	// TODO: queenWithNoMovement
)

func (this *RangedPieceMoveFixture) TestRook() {
	this.assertLegalPieceMoves(unhinderedRook, "a1", WhiteRook,
		"Ra2", "Ra3", "Ra4", "Ra5", "Ra6", "Ra7", "Ra8",
		"Rb1", "Rc1", "Rd1", "Re1", "Rf1", "Rg1", "Rh1")

	this.assertLegalPieceMoves(rookWithCaptureOpportunities, "a1", WhiteRook,
		"Ra2", "Ra3", "Ra4", "Rxa5",
		"Rb1", "Rc1", "Rxd1")

	this.assertLegalPieceMoves(rookWithLimitedMovement, "a1", WhiteRook,
		"Ra2", "Ra3", "Ra4",
		"Rb1", "Rc1")
}

func (this *RangedPieceMoveFixture) TestBishop() {
	this.assertLegalPieceMoves(bishopWithMaximumRange, "d4", WhiteBishop,
		"Ba1", "Bb2", "Bc3", "Be5", "Bf6", "Bg7", "Bh8",
		"Ba7", "Bb6", "Bc5", "Be3", "Bf2", "Bg1")
}
