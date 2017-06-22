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
	*LegalMovesFixture
}

func (this *PawnMovesFixture) Setup() {
	this.LegalMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

const (
	pawnAdvancement                  = "7k/8/8/7p/P7/8/8/K7 "
	whitePawnAdvancement             = pawnAdvancement + "w - - 0 1"
	blackPawnAdvancement             = pawnAdvancement + "b - - 0 1"
	blockedPawn                      = "7K/8/8/p7/P7/8/8/K7 "
	whiteBlockedPawn                 = blockedPawn + "w - - 0 1"
	blackBlockedPawn                 = blockedPawn + "b - - 0 1"
	initialPawn                      = "k7/p7/8/8/8/8/P7/K7 "
	whiteInitialPawnAdvancement      = initialPawn + "w - - 0 1"
	blackInitialPawnAdvancement      = initialPawn + "b - - 0 1"
	singlePawnCapture                = "k7/8/8/4p3/3P4/8/8/K7 "
	whiteSinglePawnCapture           = singlePawnCapture + "w - - 0 1"
	blackSinglePawnCapture           = singlePawnCapture + "b - - 0 1"
	doublePawnCaptures               = "k7/8/8/n1n3p1/1P3N1N/8/8/K7 "
	whiteDoublePawnCapture           = doublePawnCaptures + "w - - 0 1"
	blackDoublePawnCapture           = doublePawnCaptures + "b - - 0 1"
	whitePawnAdvancesToPromote       = "8/5P2/8/8/8/7k/8/1K6 w - - 0 1"
	blackPawnAdvancesToPromote       = "8/8/K7/8/8/7k/4p3/8 b - - 0 1"
	whitePawnCapturesAndPromotes     = "2br4/3P4/7k/K7/8/8/8/8 w - - 0 1"
	blackPawnCapturesAndPromotes     = "8/8/7k/K7/8/8/4p3/4RN2 b - - 0 1"
	pawnAdvancesToCheckEnemyKing     = "7k/8/6P1/6P1/8/8/8/1K6 w - - 0 1"
	pawnCapturesToCheckEnemyKing     = "7k/6p1/5P2/8/8/8/8/1K6 w - - 0 1"
	pawnPromotesToCheckEnemyKing     = "7k/5P2/8/8/8/8/8/1K6 w - - 0 1"
	pawnPromotesToCheckmateEnemyKing = "7k/5P1p/8/8/8/8/8/1K6 w - - 0 1"
	// TODO: pawnCanCheckmateEnemyKing (with advancement, capture, or promotion = Q or R or B or N)
	// TODO: en-passant

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
	this.assertLegalPieceMoves(pawnAdvancesToCheckEnemyKing, "g6", WhitePawn, "g7+")
	this.assertLegalPieceMoves(pawnCapturesToCheckEnemyKing, "f6", WhitePawn, "f7", "fxg7+")
	this.assertLegalPieceMoves(whitePawnAdvancesToPromote, "f7", WhitePawn, "f8=Q", "f8=R", "f8=B", "f8=N")
	this.assertLegalPieceMoves(blackPawnAdvancesToPromote, "e2", BlackPawn, "e1=Q", "e1=R", "e1=B", "e1=N")
	this.assertLegalPieceMoves(blackPawnCapturesAndPromotes, "e2", BlackPawn, "exf1=Q", "exf1=R", "exf1=B", "exf1=N")
	this.assertLegalPieceMoves(whitePawnCapturesAndPromotes, "d7", WhitePawn, "dxc8=Q", "dxc8=R", "dxc8=B", "dxc8=N")
	this.assertLegalPieceMoves(pawnPromotesToCheckEnemyKing, "f7", WhitePawn, "f8=Q+", "f8=R+", "f8=B", "f8=N")
	//this.assertLegalPieceMoves(pawnPromotesToCheckmateEnemyKing, "f7", WhitePawn, "f8=Q#", "f8=R+", "f8=B", "f8=N")
}
