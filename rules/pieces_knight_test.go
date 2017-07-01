package rules

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestKnightMovesFixture(t *testing.T) {
	gunit.Run(new(KnightMovesFixture), t)
}

type KnightMovesFixture struct {
	*gunit.Fixture
	*LegalMovesFixture
}

func (this *KnightMovesFixture) Setup() {
	this.LegalMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

const (
	knight8TargetSquares = "7k/8/8/7p/7P/2N5/8/7K w - - 0 1"
	knight6TargetSquares = "7k/8/8/7p/7P/8/2N5/7K w - - 0 1"
	knight4TargetSquares = "7k/8/8/7p/7P/8/1N6/7K w - - 0 1"
	knight3TargetSquares = "7k/8/8/7p/7P/8/8/1N5K w - - 0 1"
	knight2TargetSquares = "7k/8/8/7p/7P/8/8/N6K w - - 0 1"

	knightCapturesOpposingUnits       = "7k/8/2p1p3/1p3p1p/3N3P/1p3p2/2p1p3/7K w - - 0 1"
	knightBlockedByFriendlyUnits      = "7k/8/2P1P3/1P3P1p/3N3P/1P3P2/2P1P3/7K w - - 0 1"
	knightCanCheckEnemyKingFromCorner = "7N/8/8/4k3/8/8/8/1K6 w - - 0 1"

	// TODO: knightCanCheckmateEnemyKing
)

func (this *KnightMovesFixture) Test() {
	this.assertLegalPieceMoves(knight8TargetSquares, "c3", WhiteKnight, "Na2", "Nb1", "Nd1", "Ne2", "Ne4", "Nd5", "Nb5", "Na4")
	this.assertLegalPieceMoves(knight6TargetSquares, "c2", WhiteKnight, "Na1", "Ne1", "Ne3", "Nd4", "Nb4", "Na3")
	this.assertLegalPieceMoves(knight4TargetSquares, "b2", WhiteKnight, "Na4", "Nc4", "Nd3", "Nd1")
	this.assertLegalPieceMoves(knight3TargetSquares, "b1", WhiteKnight, "Na3", "Nc3", "Nd2")
	this.assertLegalPieceMoves(knight2TargetSquares, "a1", WhiteKnight, "Nb3", "Nc2")
	this.assertLegalPieceMoves(knightBlockedByFriendlyUnits, "d4", WhiteKnight)
	this.assertLegalPieceMoves(knightCapturesOpposingUnits, "d4", WhiteKnight,
		"Nxb3", "Nxb5", "Nxc2", "Nxc6", "Nxe2", "Nxe6", "Nxf3", "Nxf5")
	this.assertLegalPieceMoves(knightCanCheckEnemyKingFromCorner, "h8", WhiteKnight, "Ng6+", "Nf7+")
}

func (this *KnightMovesFixture) TestCaptureMovesAreMarkedAsSuch() {
	this.game.MustLoadFEN(knightCapturesOpposingUnits)
	moves := filterMovesByPieceOnSquare(this.game.GetLegalMoves(White), WhiteKnight, "d4")
	for _, move := range moves {
		this.So(move.Captured, should.Equal, BlackPawn)
		this.So(move.CapturedOn.String(), should.Equal, move.To.String())
	}
}
