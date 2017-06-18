package rules

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestPawnMovesFixture(t *testing.T) {
	gunit.Run(new(PawnMovesFixture), t)
}

type PawnMovesFixture struct {
	*gunit.Fixture
	*LegalPieceMovesFixture
}

func (this *PawnMovesFixture) Setup() {
	this.LegalPieceMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

const (
	pawnAdvancement             = "7k/8/8/7p/P7/8/8/K7 "
	whitePawnAdvancement        = pawnAdvancement + "w - - 0 1"
	blackPawnAdvancement        = pawnAdvancement + "b - - 0 1"
	blockedPawn                 = "7K/8/8/p7/P7/8/8/K7 "
	whiteBlockedPawn            = blockedPawn + "w - - 0 1"
	blackBlockedPawn            = blockedPawn + "b - - 0 1"
	initialPawn                 = "k7/p7/8/8/8/8/P7/K7 "
	whiteInitialPawnAdvancement = initialPawn + "w - - 0 1"
	blackInitialPawnAdvancement = initialPawn + "b - - 0 1"
	singlePawnCapture           = "k7/8/8/4p3/3P4/8/8/K7 "
	whiteSinglePawnCapture      = singlePawnCapture + "w - - 0 1"
	blackSinglePawnCapture      = singlePawnCapture + "b - - 0 1"
	doublePawnCaptures          = "k7/8/8/n1n3p1/1P3N1N/8/8/K7 "
	whiteDoublePawnCapture      = doublePawnCaptures + "w - - 0 1"
	blackDoublePawnCapture      = doublePawnCaptures + "b - - 0 1"
)

func (this *PawnMovesFixture) Test() {
	this.assertLegalPieceMoves(whitePawnAdvancement, "a4", WhitePawn, "a5")
	this.assertLegalPieceMoves(blackPawnAdvancement, "h5", BlackPawn, "h4")
	this.assertLegalPieceMoves(whiteInitialPawnAdvancement, "a2", WhitePawn, "a3", "a4")
	this.assertLegalPieceMoves(blackInitialPawnAdvancement, "a7", BlackPawn, "a6", "a5")
	this.assertLegalPieceMoves(whiteBlockedPawn, "a4", WhitePawn)
	this.assertLegalPieceMoves(blackBlockedPawn, "a5", BlackPawn)
	this.assertLegalPieceMoves(whiteSinglePawnCapture, "d4", WhitePawn, "d5", "dxe5")
	this.assertLegalPieceMoves(blackSinglePawnCapture, "e5", BlackPawn, "e4", "exd4")
	this.assertLegalPieceMoves(whiteDoublePawnCapture, "b4", WhitePawn, "bxa5", "b5", "bxc5")
	this.assertLegalPieceMoves(blackDoublePawnCapture, "g5", BlackPawn, "gxf4", "g4", "gxh4")
}
