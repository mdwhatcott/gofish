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
	blockedPawn                      = "7K/8/8/p7/P7/8/8/K7 "
	initialPawn                      = "k7/p7/8/8/8/8/P7/K7 "
	singlePawnCapture                = "k7/8/8/4p3/3P4/8/8/K7 "
	doublePawnCaptures               = "k7/8/8/n1n3p1/1P3N1N/8/8/K7 "
	whitePawnAdvancesToPromote       = "8/5P2/8/8/8/7k/8/1K6"
	blackPawnAdvancesToPromote       = "8/8/K7/8/8/7k/4p3/8"
	whitePawnCapturesAndPromotes     = "2br4/3P4/7k/K7/8/8/8/8"
	blackPawnCapturesAndPromotes     = "8/8/7k/K7/8/8/4p3/4RN2"
	pawnAdvancesToCheckEnemyKing     = "7k/8/6P1/6P1/8/8/8/1K6"
	pawnCapturesToCheckEnemyKing     = "7k/6p1/5P2/8/8/8/8/1K6"
	pawnPromotesToCheckEnemyKing     = "7k/5P2/8/8/8/8/8/1K6"
	pawnPromotesToCheckmateEnemyKing = "7k/5P1p/8/8/8/8/8/1K6"
	// TODO: en-passant move mechanics
	// TODO: pawnCanCheckmateEnemyKing advancement
	// TODO: pawnCanCheckmateEnemyKing capture
	// TODO: pawnCanCheckmateEnemyKing promotion = Q
	// TODO: pawnCanCheckmateEnemyKing promotion = R
	// TODO: pawnCanCheckmateEnemyKing promotion = B
	// TODO: pawnCanCheckmateEnemyKing promotion = N
	// TODO: pawnCanCheckmateEnemyKingWithEnPassant

)

func (this *PawnMovesFixture) Test() {
	this.assertLegalPieceMoves(pawnAdvancement, "a4", WhitePawn, "a5")
	this.assertLegalPieceMoves(pawnAdvancement, "h5", BlackPawn, "h4")
	this.assertLegalPieceMoves(initialPawn, "a2", WhitePawn, "a3", "a4")
	this.assertLegalPieceMoves(initialPawn, "a7", BlackPawn, "a6", "a5")
	this.assertLegalPieceMoves(blockedPawn, "a4", WhitePawn)
	this.assertLegalPieceMoves(blockedPawn, "a5", BlackPawn)
	this.assertLegalPieceMoves(singlePawnCapture, "d4", WhitePawn, "d5", "dxe5")
	this.assertLegalPieceMoves(singlePawnCapture, "e5", BlackPawn, "e4", "exd4")
	this.assertLegalPieceMoves(doublePawnCaptures, "b4", WhitePawn, "bxa5", "b5", "bxc5")
	this.assertLegalPieceMoves(doublePawnCaptures, "g5", BlackPawn, "gxf4", "g4", "gxh4")
	this.assertLegalPieceMoves(pawnAdvancesToCheckEnemyKing, "g6", WhitePawn, "g7+")
	this.assertLegalPieceMoves(pawnCapturesToCheckEnemyKing, "f6", WhitePawn, "f7", "fxg7+")
	this.assertLegalPieceMoves(whitePawnAdvancesToPromote, "f7", WhitePawn, "f8=Q", "f8=R", "f8=B", "f8=N")
	this.assertLegalPieceMoves(blackPawnAdvancesToPromote, "e2", BlackPawn, "e1=Q", "e1=R", "e1=B", "e1=N")
	this.assertLegalPieceMoves(blackPawnCapturesAndPromotes, "e2", BlackPawn, "exf1=Q", "exf1=R", "exf1=B", "exf1=N")
	this.assertLegalPieceMoves(whitePawnCapturesAndPromotes, "d7", WhitePawn, "dxc8=Q", "dxc8=R", "dxc8=B", "dxc8=N")
	this.assertLegalPieceMoves(pawnPromotesToCheckEnemyKing, "f7", WhitePawn, "f8=Q+", "f8=R+", "f8=B", "f8=N")
	//this.assertLegalPieceMoves(pawnPromotesToCheckmateEnemyKing, "f7", WhitePawn, "f8=Q#", "f8=R+", "f8=B", "f8=N")
}
