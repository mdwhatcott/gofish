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
	*LegalPieceMovesFixture
}

func (this *KnightMovesFixture) Setup() {
	this.LegalPieceMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

const (
	knight8TargetSquares = "7k/8/8/7p/7P/2N5/8/7K w - - 0 1"
	knight6TargetSquares = "7k/8/8/7p/7P/8/2N5/7K w - - 0 1"
	knight4TargetSquares = "7k/8/8/7p/7P/8/1N6/7K w - - 0 1"
	knight3TargetSquares = "7k/8/8/7p/7P/8/8/1N5K w - - 0 1"
	knight2TargetSquares = "7k/8/8/7p/7P/8/8/N6K w - - 0 1"

	knightCapturesOpposingUnits  = "7k/8/2p1p3/1p3p1p/3N3P/1p3p2/2p1p3/7K w - - 0 1"
	knightBlockedByFriendlyUnits = "7k/8/2P1P3/1P3P1p/3N3P/1P3P2/2P1P3/7K w - - 0 1"

	// TODO: if the knight's king is in check and the knight can't do anything to prevent check, no knight moves are possible
	// TODO: if the knight's king king is in check and the knight can remove check by blocking or capturing the aggressor, those are the only knight moves available
	// TODO: if moving the knig	ht would cause discovered check, no knight moves are possible (requires ranged piece movement behavior).
)

func (this *KnightMovesFixture) Test() {
	this.assertLegalPieceMoves(knight8TargetSquares, "c3", WhiteKnight, "a2", "b1", "d1", "e2", "e4", "d5", "b5", "a4")
	this.assertLegalPieceMoves(knight6TargetSquares, "c2", WhiteKnight, "a1", "e1", "e3", "d4", "b4", "a3")
	this.assertLegalPieceMoves(knight4TargetSquares, "b2", WhiteKnight, "a4", "c4", "d3", "d1")
	this.assertLegalPieceMoves(knight3TargetSquares, "b1", WhiteKnight, "a3", "c3", "d2")
	this.assertLegalPieceMoves(knight2TargetSquares, "a1", WhiteKnight, "b3", "c2")
	this.assertLegalPieceMoves(knightBlockedByFriendlyUnits, "d4", WhiteKnight)
	this.assertLegalPieceMoves(knightCapturesOpposingUnits, "d4", WhiteKnight,
		"b3", "b5", "c2", "c6", "e2", "e6", "f3", "f5")
}

func (this *KnightMovesFixture) TestCaptureMovesAreMarkedAsSuch() {
	this.game.MustLoadFEN(knightCapturesOpposingUnits)
	moves := filterMovesByPiece(this.game.GetAvailableMoves(White), WhiteKnight)
	for _, move := range moves {
		this.So(move.Capture, should.Equal, BlackPawn)
		this.So(move.CaptureOn.String(), should.Equal, move.To.String())
	}
}
