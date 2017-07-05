package rules

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestKingMovesFixture(t *testing.T) {
	gunit.Run(new(KingMovesFixture), t)
}

type KingMovesFixture struct {
	*gunit.Fixture
	*LegalMovesFixture
}

func (this *KingMovesFixture) Setup() {
	this.LegalMovesFixture = NewLegalGameMovesFixture(this.Fixture)
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
	whiteKingSurroundedByThreatenedSquares     = "2r1r3/8/8/r7/3K4/r7/8/8 w - - 0 1"
	whiteKingSurroundedByProtectedEnemyUnits   = "8/8/8/2qqq3/2qKq3/2qqq3/8/8 w - - 0 1"
)

func (this *KingMovesFixture) TestAwayFromAnyEdge() {
	this.assertLegalPieceMoves(
		whiteKingAwayFromAnyEdge, "b2", WhiteKing,
		"Ka1", "Ka2", "Ka3",
		"Kb1" /****/, "Kb3",
		"Kc1", "Kc2", "Kc3")
	this.assertLegalPieceMoves(whiteKingOnBottomEdge, "b1", WhiteKing,
		"Ka2", "Kb2", "Kc2",
		"Ka1" /****/, "Kc1")
	this.assertLegalPieceMoves(blackKingOnTopEdge, "b8", BlackKing,
		"Ka8" /****/, "Kc8",
		"Ka7", "Kb7", "Kc7")
	this.assertLegalPieceMoves(blackKingOnLeftEdge, "a7", BlackKing,
		"Ka8", "Kb8",
		/****/ "Kb7",
		"Ka6", "Kb6")
	this.assertLegalPieceMoves(blackKingOnRightEdge, "h7", BlackKing,
		"Kg8", "Kh8",
		"Kg7",
		"Kg6", "Kh6")
	this.assertLegalPieceMoves(whiteKingInBottomLeftCorner, "a1", WhiteKing,
		"Ka2", "Kb2",
		/****/ "Kb1")
	this.assertLegalPieceMoves(whiteKingInBottomRightCorner, "h1", WhiteKing,
		"Kg2", "Kh2",
		"Kg1" /****/)
	this.assertLegalPieceMoves(blackKingInTopLeftCorner, "a8", BlackKing,
		/****/ "Kb8",
		"Ka7", "Kb7")
	this.assertLegalPieceMoves(blackKingInTopRightCorner, "h8", BlackKing,
		"Kg8", /***/
		"Kg7", "Kh7")
	this.assertLegalPieceMoves(whiteKingSurroundedByUnprotectedEnemyUnits, "a1", WhiteKing,
		"Kxa2", "Kxb2",
		/*****/ "Kxb1")
	this.assertLegalPieceMoves(whiteKingSurroundedByFriendlyUnits, "c3", WhiteKing)
}

func (this *KingMovesFixture) TestKingCannotEnterCheck() {
	this.assertLegalPieceMoves(kingCannotApproachOtherKing, "a8", BlackKing)
	this.assertLegalPieceMoves(whiteKingSurroundedByThreatenedSquares, "d4", WhiteKing)
	this.assertLegalPieceMoves(whiteKingSurroundedByProtectedEnemyUnits, "d4", WhiteKing)
}

func (this *KingMovesFixture) TestCaptureMovesAreMarkedAsSuch() {
	this.game.MustLoadFEN(whiteKingSurroundedByUnprotectedEnemyUnits)
	moves := filterMovesByPieceOnSquare(this.game.GetLegalMoves(White), WhiteKing, "a1")
	for _, move := range moves {
		this.So(move.Captured, should.Equal, BlackKnight)
		this.So(move.CapturedOn.String(), should.Equal, move.To.String())
	}
}

const (
	whiteKingWithCastlingOpportunities                      = "r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w KQkq - 0 1"
	blackKingWithCastlingOpportunities                      = "r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R b KQkq - 0 1"
	whiteKingWithoutCastlingRights                          = "r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w - - 0 1"
	blackKingWithoutCastlingRights                          = "r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R b - - 0 1"
	whiteKingCannotCastleBecauseTravelSquaresAreOccupied    = "r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R2QKB1R w KQkq - 0 1"
	blackKingCannotCastleBecauseTravelSquaresAreOccupied    = "r2qkb1r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w KQkq - 0 1"
	whiteKingCannotCastleBecauseLandingSquaresAreOccupied   = "r3k2r/pppppppp/8/8/8/8/PPPPPPPP/R1B1K1NR w KQkq - 0 1"
	blackKingCannotCastleBecauseLandingSquaresAreOccupied   = "r1b1k1nr/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w KQkq - 0 1"
	whiteKingCannotCastleQueensideBecauseRookIsBlocked      = "r3k2r/pppppppp/8/8/8/8/PPPPPPPP/RN2K2R w KQkq - 0 1"
	blackKingCannotCastleQueensideBecauseRookIsBlocked      = "rn2k2r/pppppppp/8/8/8/8/PPPPPPPP/R3K2R w KQkq - 0 1"
	whiteKingCannotCastleWhenInCheck                        = "r3k2r/ppppqppp/8/8/8/8/PPPP1PPP/R3K2R w KQkq - 0 1"
	blackKingCannotCastleWhenInCheck                        = "r3k2r/pppp1ppp/8/8/8/8/PPPPQPPP/R3K2R w KQkq - 0 1"
	whiteKingCannotCastleKingsideWhenTravelingThroughCheck  = "r3k2r/pppppqpp/8/8/8/8/PPPPP1PP/R3K2R w KQkq - 0 1"
	whiteKingCannotCastleQueensideWhenTravelingThroughCheck = "r3k2r/pppqpppp/8/8/8/8/PPP1PPPP/R3K2R w KQkq - 0 1"
	blackKingCannotCastleKingsideWhenTravelingThroughCheck  = "r3k2r/ppppp1pp/8/8/8/8/PPPPPQPP/R3K2R w KQkq - 0 1"
	blackKingCannotCastleQueensideWhenTravelingThroughCheck = "r3k2r/ppp1pppp/8/8/8/8/PPPQPPPP/R3K2R w KQkq - 0 1"
	whiteKingCannotCastleKingsideWhenLandingInCheck         = "r3k2r/ppppppqp/8/8/8/8/PPPPPP1P/R3K2R w KQkq - 0 1"
	whiteKingCannotCastleQueensideWhenLandingInCheck        = "r3k2r/ppqppppp/8/8/8/8/PP1PPPPP/R3K2R w KQkq - 0 1"
	blackKingCannotCastleKingsideWhenLandingInCheck         = "r3k2r/pppppp1p/8/8/8/8/PPPPPPQP/R3K2R w KQkq - 0 1"
	blackKingCannotCastleQueensideWhenLandingInCheck        = "r3k2r/pp1ppppp/8/8/8/8/PPQPPPPP/R3K2R w KQkq - 0 1"
	whiteKingCANCastleEvenWhenRookWouldPassThroughAttack    = "r3k2r/pqpppppp/8/8/8/8/P1PPPPPP/R3K2R w KQkq - 0 1"
	blackKingCANCastleEvenWhenRookWouldPassThroughAttack    = "r3k2r/p1pppppp/8/8/8/8/PQPPPPPP/R3K2R w KQkq - 0 1"
)

func (this *KingMovesFixture) TestCanCastle() {
	this.assertLegalPieceMoves(whiteKingWithCastlingOpportunities, "e1", WhiteKing, "Kf1", "Kd1", "O-O", "O-O-O")
	this.assertLegalPieceMoves(blackKingWithCastlingOpportunities, "e8", BlackKing, "Kf8", "Kd8", "O-O", "O-O-O")
}

func (this *KingMovesFixture) TestCannotCastleBecauseOfFENSettings() {
	this.assertLegalPieceMoves(whiteKingWithoutCastlingRights, "e1", WhiteKing, "Kf1", "Kd1")
	this.assertLegalPieceMoves(blackKingWithoutCastlingRights, "e8", BlackKing, "Kf8", "Kd8")
}

func (this *KingMovesFixture) TestCannotCastle() {
	this.assertLegalPieceMoves(whiteKingCannotCastleBecauseTravelSquaresAreOccupied, "e1", WhiteKing)
	this.assertLegalPieceMoves(blackKingCannotCastleBecauseTravelSquaresAreOccupied, "e8", BlackKing)
	this.assertLegalPieceMoves(whiteKingCannotCastleBecauseLandingSquaresAreOccupied, "e1", WhiteKing, "Kd1", "Kf1")
	this.assertLegalPieceMoves(blackKingCannotCastleBecauseLandingSquaresAreOccupied, "e8", BlackKing, "Kd8", "Kf8")
	this.assertLegalPieceMoves(whiteKingCannotCastleQueensideBecauseRookIsBlocked, "e1", WhiteKing, "Kd1", "Kf1", "O-O")
	this.assertLegalPieceMoves(blackKingCannotCastleQueensideBecauseRookIsBlocked, "e8", BlackKing, "Kd8", "Kf8", "O-O")
	this.assertLegalPieceMoves(whiteKingCannotCastleWhenInCheck, "e1", WhiteKing, "Kd1", "Kf1")
	this.assertLegalPieceMoves(blackKingCannotCastleWhenInCheck, "e8", BlackKing, "Kd8", "Kf8")
	this.assertLegalPieceMoves(whiteKingCannotCastleKingsideWhenTravelingThroughCheck, "e1", WhiteKing, "Kd1", "O-O-O")
	this.assertLegalPieceMoves(whiteKingCannotCastleQueensideWhenTravelingThroughCheck, "e1", WhiteKing, "Kf1", "O-O")
	this.assertLegalPieceMoves(blackKingCannotCastleKingsideWhenTravelingThroughCheck, "e8", BlackKing, "Kd8", "O-O-O")
	this.assertLegalPieceMoves(blackKingCannotCastleQueensideWhenTravelingThroughCheck, "e8", BlackKing, "Kf8", "O-O")
	this.assertLegalPieceMoves(whiteKingCannotCastleKingsideWhenLandingInCheck, "e1", WhiteKing, "Kd1", "Kf1", "O-O-O")
	this.assertLegalPieceMoves(whiteKingCannotCastleQueensideWhenLandingInCheck, "e1", WhiteKing, "Kd1", "Kf1", "O-O")
	this.assertLegalPieceMoves(blackKingCannotCastleKingsideWhenLandingInCheck, "e8", BlackKing, "Kd8", "Kf8", "O-O-O")
	this.assertLegalPieceMoves(blackKingCannotCastleQueensideWhenLandingInCheck, "e8", BlackKing, "Kd8", "Kf8", "O-O")
	this.assertLegalPieceMoves(whiteKingCANCastleEvenWhenRookWouldPassThroughAttack, "e1", WhiteKing,
		"Kf1", "Kd1", "O-O", "O-O-O")
	this.assertLegalPieceMoves(blackKingCANCastleEvenWhenRookWouldPassThroughAttack, "e8", BlackKing,
		"Kf8", "Kd8", "O-O", "O-O-O")
}

func (this *KingMovesFixture) TestCastlingIsLegalBeforeHavingMovedTheKing() {
	kingsIndianReadyToCastle := []string{
		"Nf3", "Nf6",
		"g3", "g6",
		"Bg2", "Bg7",
	}

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianReadyToCastle,
		FromSquare:          "e1",
		FocusOnPiece:        WhiteKing,
		ExpectedMovesSAN:    []string{"Kf1", "O-O"},
	})
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianReadyToCastle,
		FromSquare:          "e8",
		FocusOnPiece:        BlackKing,
		ExpectedMovesSAN:    []string{"Kf8", "O-O"},
	})
}

func (this *KingMovesFixture) TestCastlingKingsideIsNoLongerLegalAfterMovingTheKing() {
	kingsIndianReadyToCastle := []string{
		"Nf3", "Nf6",
		"g3", "g6",
		"Bg2", "Bg7",
		"Kf1", "Kf8",
		"Ke1", "Ke8",
	}

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianReadyToCastle,
		FromSquare:          "e1",
		FocusOnPiece:        WhiteKing,
		ExpectedMovesSAN:    []string{"Kf1"},
		ExpectedPositionFEN: "rnbqk2r/ppppppbp/5np1/8/8/5NP1/PPPPPPBP/RNBQK2R w - - 0 1",
	})
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianReadyToCastle,
		FromSquare:          "e8",
		FocusOnPiece:        BlackKing,
		ExpectedMovesSAN:    []string{"Kf8"},
		ExpectedPositionFEN: "rnbqk2r/ppppppbp/5np1/8/8/5NP1/PPPPPPBP/RNBQK2R w - - 0 1",
	})
}

func (this *KingMovesFixture) TestCastlingKingsideIsNoLongerLegalAfterMovingTheKingsideRook() {
	kingsIndianReadyToCastle := []string{
		"Nf3", "Nf6",
		"g3", "g6",
		"Bg2", "Bg7",
		"Rg1", "Rg8",
		"Rh1", "Rh8",
	}

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianReadyToCastle,
		FromSquare:          "e1",
		FocusOnPiece:        WhiteKing,
		ExpectedMovesSAN:    []string{"Kf1"},
		ExpectedPositionFEN: "rnbqk2r/ppppppbp/5np1/8/8/5NP1/PPPPPPBP/RNBQK2R w Qq - 0 1",
	})
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianReadyToCastle,
		FromSquare:          "e8",
		FocusOnPiece:        BlackKing,
		ExpectedMovesSAN:    []string{"Kf8"},
		ExpectedPositionFEN: "rnbqk2r/ppppppbp/5np1/8/8/5NP1/PPPPPPBP/RNBQK2R w Qq - 0 1",
	})
}

func (this *KingMovesFixture) TestCastlingQueensideLegalBeforeMovingTheKing() {
	vacatedQueenside := []string{
		"Nc3", "Nc6",
		"e4", "e5",
		"d4", "d5",
		"Qf3", "Qf6",
		"Bf4", "Bf5",
	}

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: vacatedQueenside,
		FromSquare:          "e1",
		FocusOnPiece:        WhiteKing,
		ExpectedMovesSAN:    []string{"Kd1", "Kd2", "Ke2", "O-O-O"},
		ExpectedPositionFEN: "r3kbnr/ppp2ppp/2n2q2/3ppb2/3PPB2/2N2Q2/PPP2PPP/R3KBNR w KQkq - 0 1",
	})
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: vacatedQueenside,
		FromSquare:          "e8",
		FocusOnPiece:        BlackKing,
		ExpectedMovesSAN:    []string{"Kd8", "Kd7", "Ke7", "O-O-O"},
		ExpectedPositionFEN: "r3kbnr/ppp2ppp/2n2q2/3ppb2/3PPB2/2N2Q2/PPP2PPP/R3KBNR w KQkq - 0 1",
	})
}

func (this *KingMovesFixture) TestCastlingQueensideIsNoLongerLegalAfterMovingTheKing() {
	vacatedQueenside := []string{
		"Nc3", "Nc6",
		"e4", "e5",
		"d4", "d5",
		"Qf3", "Qf6",
		"Bf4", "Bf5",
		"Kd1", "Kd8",
		"Ke1", "Ke8",
	}

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: vacatedQueenside,
		FromSquare:          "e1",
		FocusOnPiece:        WhiteKing,
		ExpectedMovesSAN:    []string{"Kd1", "Kd2", "Ke2"},
		ExpectedPositionFEN: "r3kbnr/ppp2ppp/2n2q2/3ppb2/3PPB2/2N2Q2/PPP2PPP/R3KBNR w - - 0 1",
	})
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: vacatedQueenside,
		FromSquare:          "e8",
		FocusOnPiece:        BlackKing,
		ExpectedMovesSAN:    []string{"Kd8", "Kd7", "Ke7"},
		ExpectedPositionFEN: "r3kbnr/ppp2ppp/2n2q2/3ppb2/3PPB2/2N2Q2/PPP2PPP/R3KBNR w - - 0 1",
	})
}

func (this *KingMovesFixture) TestCastlingQueensideIsNoLongerLegalAfterMovingTheQueensideRook() {
	vacatedQueenside := []string{
		"Nc3", "Nc6",
		"e4", "e5",
		"d4", "d5",
		"Qf3", "Qf6",
		"Bf4", "Bf5",
		"Rb1", "Rb8",
		"Ra1", "Ra8",
	}

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: vacatedQueenside,
		FromSquare:          "e1",
		FocusOnPiece:        WhiteKing,
		ExpectedMovesSAN:    []string{"Kd1", "Kd2", "Ke2"},
		ExpectedPositionFEN: "r3kbnr/ppp2ppp/2n2q2/3ppb2/3PPB2/2N2Q2/PPP2PPP/R3KBNR w Kk - 0 1",
	})
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: vacatedQueenside,
		FromSquare:          "e8",
		FocusOnPiece:        BlackKing,
		ExpectedMovesSAN:    []string{"Kd8", "Kd7", "Ke7"},
		ExpectedPositionFEN: "r3kbnr/ppp2ppp/2n2q2/3ppb2/3PPB2/2N2Q2/PPP2PPP/R3KBNR w Kk - 0 1",
	})
}

func (this *KingMovesFixture) TestCastlingCanOnlyHappenOnce() {
	kingsIndianAfterWhiteCastlesKingside := []string{
		"Nf3", "Nf6",
		"g3", "g6",
		"Bg2", "Bg7",
		"O-O",
	}

	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianAfterWhiteCastlesKingside,
		ExpectedPositionFEN: "rnbqk2r/ppppppbp/5np1/8/8/5NP1/PPPPPPBP/RNBQ1RK1 b kq - 0 1", // Only black can still castle

		FocusOnPiece: WhiteBishop, // focus on a piece with not moves
		FromSquare:   "c1",
	})
	kingsIndianAfterWhiteAndBlackCastleKingside := []string{
		"Nf3", "Nf6",
		"g3", "g6",
		"Bg2", "Bg7",
		"O-O", "O-O",
	}
	this.PlayAndValidate(LegalMovesSetup{
		PreparatoryMovesSAN: kingsIndianAfterWhiteAndBlackCastleKingside,
		ExpectedPositionFEN: "rnbq1rk1/ppppppbp/5np1/8/8/5NP1/PPPPPPBP/RNBQ1RK1 w - - 0 1", // No one can castle anymore

		FocusOnPiece: BlackBishop, // focus on a piece with not moves
		FromSquare:   "c8",
	})
}

// TODO: Executing a castle queenside should abolish future right to castle
// TODO: Executing a castling move should move the king and the involved rook ([√] kingside and [ ] queenside)
// TODO: Taking back a castling move should restore the king and the involved rook
