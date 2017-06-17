package rules

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestKingMovesFixture(t *testing.T) {
	gunit.Run(new(KingMovesFixture), t)
}

type KingMovesFixture struct {
	*gunit.Fixture
	*LegalPieceMovesFixture
}

func (this *KingMovesFixture) Setup() {
	this.LegalPieceMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

const (
	whiteKingAwayFromAnyEdge     = "8/1k6/8/p7/P7/8/1K6/8 w - - 0 1"
	whiteKingOnBottomEdge        = "8/1k6/8/p7/P7/8/8/1K6 w - - 0 1"
	blackKingOnTopEdge           = "1k6/8/8/p7/P7/8/8/1K6 b - - 0 1"
	blackKingOnLeftEdge          = "8/k7/8/p7/P7/8/8/1K6 b - - 0 1"
	blackKingOnRightEdge         = "8/7k/8/p7/P7/8/8/1K6 b - - 0 1"
	whiteKingInBottomLeftCorner  = "k7/8/8/p7/P7/8/8/K7 w - - 0 1"
	whiteKingInBottomRightCorner = "k7/8/8/p7/P7/8/8/7K w - - 0 1"
	blackKingInTopLeftCorner     = "k7/8/8/p7/P7/8/8/7K b - - 0 1"
	blackKingInTopRightCorner    = "7k/8/8/p7/P7/8/8/7K b - - 0 1"

	whiteKingSurroundedByFriendlyUnits         = "k7/8/8/8/1BQR4/1NKN4/1PPP4/8 w - - 0 1"
	whiteKingSurroundedByUnprotectedEnemyUnits = "k7/8/8/8/8/8/nn6/Kn6 w - - 0 1"
	kingCannotApproachOtherKing                = "kq6/8/K7/8/8/8/8/8 w - - 0 1"
	whiteKingSurroundedByThreatenedSquares     = "2r1r3/8/8/r7/3K4/r7/8/8 w - - 0 1"     // TODO
	whiteKingSurroundedByProtectedEnemyUnits   = "8/8/8/2qqq3/2qKq3/2qqq3/8/8 w - - 0 1" // TODO
)

func (this *KingMovesFixture) TestAwayFromAnyEdge() {
	this.assertLegalPieceMoves(
		whiteKingAwayFromAnyEdge, "b2", WhiteKing,
		"a1", "a2", "a3",
		"b1" /***/, "b3",
		"c1", "c2", "c3")
}
func (this *KingMovesFixture) TestOnBottomEdge() {
	this.assertLegalPieceMoves(whiteKingOnBottomEdge, "b1", WhiteKing,
		"a2", "b2", "c2",
		"a1" /***/, "c1")
}
func (this *KingMovesFixture) TestOnTopEdge() {
	this.assertLegalPieceMoves(blackKingOnTopEdge, "b8", BlackKing,
		"a8" /***/, "c8",
		"a7", "b7", "c7")
}
func (this *KingMovesFixture) TestOnLeftEdge() {
	this.assertLegalPieceMoves(blackKingOnLeftEdge, "a7", BlackKing,
		"a8", "b8",
		/***/ "b7",
		"a6", "b6")
}
func (this *KingMovesFixture) TestOnRightEdge() {
	this.assertLegalPieceMoves(blackKingOnRightEdge, "h7", BlackKing,
		"g8", "h8",
		"g7",
		"g6", "h6")
}
func (this *KingMovesFixture) TestInBottomLeftCorner() {
	this.assertLegalPieceMoves(whiteKingInBottomLeftCorner, "a1", WhiteKing,
		"a2", "b2",
		/***/ "b1")
}
func (this *KingMovesFixture) TestInBottomRightCorner() {
	this.assertLegalPieceMoves(whiteKingInBottomRightCorner, "h1", WhiteKing,
		"g2", "h2",
		"g1" /***/)
}
func (this *KingMovesFixture) TestInTopLeftCorner() {
	this.assertLegalPieceMoves(blackKingInTopLeftCorner, "a8", BlackKing,
		/***/ "b8",
		"a7", "b7")
}
func (this *KingMovesFixture) TestInTopRightCorner() {
	this.assertLegalPieceMoves(blackKingInTopRightCorner, "h8", BlackKing,
		"g8", /***/
		"g7", "h7")
}
func (this *KingMovesFixture) TestBlockedByFriendlyUnits() {
	this.assertLegalPieceMoves(whiteKingSurroundedByFriendlyUnits, "c3", WhiteKing)
}
func (this *KingMovesFixture) TestSurroundedByUnprotectedEnemyUnits() {
	this.assertLegalPieceMoves(whiteKingSurroundedByUnprotectedEnemyUnits, "a1", WhiteKing,
		"a2", "b2",
		/***/ "b1")
}

func (this *KingMovesFixture) TestKingsRepelEachOther() {
	this.assertLegalPieceMoves(kingCannotApproachOtherKing, "a8", BlackKing)
}

// TODO: can't move into check
// TODO: can't capture into check
