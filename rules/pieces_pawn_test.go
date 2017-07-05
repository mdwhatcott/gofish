package rules

import (
	"testing"

	"github.com/smartystreets/assertions/should"
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
	pawnAdvancement                  = "7k/8/8/7p/P7/8/8/K7"
	blockedPawn                      = "7K/8/8/p7/P7/8/8/K7"
	initialPawn                      = "k7/p7/8/8/8/8/P7/K7"
	singlePawnCapture                = "k7/8/8/4p3/3P4/8/8/K7"
	doublePawnCaptures               = "k7/8/8/n1n3p1/1P3N1N/8/8/K7"
	whitePawnAdvancesToPromote       = "8/5P2/8/8/8/7k/8/1K6"
	blackPawnAdvancesToPromote       = "8/8/K7/8/8/7k/4p3/8"
	whitePawnCapturesAndPromotes     = "2br4/3P4/7k/K7/8/8/8/8"
	blackPawnCapturesAndPromotes     = "8/8/7k/K7/8/8/4p3/4RN2"
	pawnAdvancesToCheckEnemyKing     = "7k/8/6P1/6P1/8/8/8/1K6"
	pawnCapturesToCheckEnemyKing     = "7k/6p1/5P2/8/8/8/8/1K6"
	pawnPromotesToCheckEnemyKing     = "7k/5P2/8/8/8/8/8/1K6"
	pawnAdvancesToCheckmateEnemyKing = "6Nk/5R1P/5PP1/6P1/8/8/2B5/1K6"
	pawnPromotesToCheckmateEnemyKing = "7k/5P1p/8/8/8/8/8/1K6"
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
	this.assertLegalPieceMoves(pawnAdvancesToCheckmateEnemyKing, "g6", WhitePawn, "g7#")
	this.assertLegalPieceMoves(pawnPromotesToCheckmateEnemyKing, "f7", WhitePawn, "f8=Q#", "f8=R+", "f8=B", "f8=N")
}

func (this *GameFixture) TestTakeBackPromotion() {
	fen := "8/8/8/8/8/7k/K3p3/8"
	this.game.MustLoadFEN(fen)
	move := move{Piece: BlackPawn, From: Square("e2"), To: Square("e1"), Promotion: BlackQueen}
	this.game.Execute(move)
	this.game.TakeBack(move)
	this.So(this.game.IsInCheck(White), should.BeFalse)
	this.So(this.game.ExportFEN().String(), should.StartWith, fen)
}

func (this *PawnMovesFixture) TestEnPassantIsLegalMove() {
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: []string{"e4", "e6", "e5", "d5"},
		FocusOnPiece:        WhitePawn,
		FromSquare:          "e5",
		ExpectedMovesSAN:    []string{"exd6"},
		ExpectedPositionFEN: "rnbqkbnr/ppp2ppp/4p3/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1", // w KQkq d6 0 3", // TODO
	})

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: []string{"e4", "e6", "e5", "f5"},
		FocusOnPiece:        WhitePawn,
		FromSquare:          "e5",
		ExpectedMovesSAN:    []string{"exf6"},
		ExpectedPositionFEN: "rnbqkbnr/pppp2pp/4p3/4Pp2/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1", // b KQkq f6 0 3", // TODO
	})

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: []string{"e4", "c5", "c3", "c4", "d4"},
		FocusOnPiece:        BlackPawn,
		FromSquare:          "c4",
		ExpectedMovesSAN:    []string{"cxd3"},
		ExpectedPositionFEN: "rnbqkbnr/pp1ppppp/8/8/2pPP3/2P5/PP3PPP/RNBQKBNR b KQkq - 0 1", // KQkq d3 0 3", TODO
	})

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: []string{"e4", "c5", "c3", "c4", "b4"},
		FocusOnPiece:        BlackPawn,
		FromSquare:          "c4",
		ExpectedMovesSAN:    []string{"cxb3"},
		ExpectedPositionFEN: "rnbqkbnr/pp1ppppp/8/8/1Pp1P3/2P5/P2P1PPP/RNBQKBNR b KQkq - 0 1", // KQkq d3 0 3", TODO
	})
}

func (this *PawnMovesFixture) TestEnPassantIsForfeitedIfNotPerformedImmediately() {
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: []string{
			"e4", "c5",
			"c3", "c4",
			"b4", "h6", // With this move by black, the privilege to capture en passant is forfeited.
			"h3",
		},
		FocusOnPiece:     BlackPawn,
		FromSquare:       "c4",
		ExpectedMovesSAN: nil,
	})
}

func (this *PawnMovesFixture) TestEnPassantMoveMechanics_CapturedPieceRemoved() {
	this.Play(LegalMovesSetup{PreparatoryMovesSAN: []string{
		"e4", "e6",
		"e5", "d5",
		"exd6",
	}})
	this.assertPosition("rnbqkbnr/ppp2ppp/3Pp3/8/8/8/PPPP1PPP/RNBQKBNR")

	this.Play(LegalMovesSetup{PreparatoryMovesSAN: []string{
		"e4", "e6",
		"e5", "f5",
		"exf6",
	}})
	this.assertPosition("rnbqkbnr/pppp2pp/4pP2/8/8/8/PPPP1PPP/RNBQKBNR")

	this.Play(LegalMovesSetup{PreparatoryMovesSAN: []string{
		"e3", "e5",
		"Nc3", "e4",
		"d4", "exd3",
	}})
	this.assertPosition("rnbqkbnr/pppp1ppp/8/8/8/2NpP3/PPP2PPP/R1BQKBNR")

	this.Play(LegalMovesSetup{PreparatoryMovesSAN: []string{
		"e3", "e5",
		"Nc3", "e4",
		"f4", "exf3",
	}})
	this.assertPosition("rnbqkbnr/pppp1ppp/8/8/8/2N1Pp2/PPPP2PP/R1BQKBNR")
}

func (this *PawnMovesFixture) TestWhitePawnEnPassantTakeBack() {
	this.Play(LegalMovesSetup{PreparatoryMovesSAN: []string{
		"e4", "e6",
		"e5", "d5",
	}})
	beforeEnPassant := this.game.ExportFEN().String()
	enPassant := this.game.Attempt("exd6")

	this.game.TakeBack(enPassant)

	this.assertPosition(beforeEnPassant)
	this.So(this.game.GetEnPassantTarget().String(), should.Equal, "d6")
}

func (this *PawnMovesFixture) TestBlackPawnEnPassantTakeBack() {
	this.Play(LegalMovesSetup{PreparatoryMovesSAN: []string{
		"e3", "e5",
		"Nc3", "e4",
		"d4",
	}})
	beforeEnPassant := this.game.ExportFEN().String()
	enPassant := this.game.Attempt("exd3")

	this.game.TakeBack(enPassant)

	this.assertPosition(beforeEnPassant)
	this.So(this.game.GetEnPassantTarget().String(), should.Equal, "d3")
}
