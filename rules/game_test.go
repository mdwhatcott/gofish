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

	game *Game
}

func (this *GameFixture) Setup() {
	this.game = NewGame()
}

func (this *GameFixture) TestStartingGameConditions() {
	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, White)
	this.So(this.game.FullMoveCount(), should.Equal, 1)
	this.So(this.game.HalfMoveCount(), should.Equal, 0)
	this.So(this.game.CanCastleKingside(White), should.BeTrue)
	this.So(this.game.CanCastleKingside(Black), should.BeTrue)
	this.So(this.game.CanCastleQueenside(White), should.BeTrue)
	this.So(this.game.CanCastleQueenside(Black), should.BeTrue)
	this.So(this.game.FEN(), should.Equal, startingPositionFEN)
}

func (this *GameFixture) TestGameConditionsAfterFirstPawnMove() {
	err := this.game.Move(&Move{From: "a2", To: "a3", Piece: WhitePawn})

	this.So(err, should.BeNil)
	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, Black)
	this.So(this.game.FullMoveCount(), should.Equal, 1)
	this.So(this.game.HalfMoveCount(), should.Equal, 0) // pawn move
	this.So(this.game.FEN(), should.Equal, positionAfter1A3)
}

func (this *GameFixture) TestLoadFEN() {
	const kingsOnBackRanks = "4k3/8/8/8/8/8/8/4K3 w - - 0 1"
	err := this.game.LoadFEN(kingsOnBackRanks)
	this.So(err, should.BeNil)
	this.So(this.game.FEN(), should.Equal, kingsOnBackRanks)
}

func (this *GameFixture) TestLegalFirstMoves() {
	this.assertFirstMoveSuccessful(&Move{From: "a2", To: "a3", Piece: WhitePawn}, positionAfter1A3)
	this.assertFirstMoveSuccessful(&Move{From: "a2", To: "a4", Piece: WhitePawn}, positionAfter1A4)

	this.assertFirstMoveSuccessful(&Move{From: "b2", To: "b3", Piece: WhitePawn}, positionAfter1B3)
	this.assertFirstMoveSuccessful(&Move{From: "b2", To: "b4", Piece: WhitePawn}, positionAfter1B4)

	this.assertFirstMoveSuccessful(&Move{From: "c2", To: "c3", Piece: WhitePawn}, positionAfter1C3)
	this.assertFirstMoveSuccessful(&Move{From: "c2", To: "c4", Piece: WhitePawn}, positionAfter1C4)

	this.assertFirstMoveSuccessful(&Move{From: "d2", To: "d3", Piece: WhitePawn}, positionAfter1D3)
	this.assertFirstMoveSuccessful(&Move{From: "d2", To: "d4", Piece: WhitePawn}, positionAfter1D4)

	this.assertFirstMoveSuccessful(&Move{From: "e2", To: "e3", Piece: WhitePawn}, positionAfter1E3)
	this.assertFirstMoveSuccessful(&Move{From: "e2", To: "e4", Piece: WhitePawn}, positionAfter1E4)

	this.assertFirstMoveSuccessful(&Move{From: "f2", To: "f3", Piece: WhitePawn}, positionAfter1F3)
	this.assertFirstMoveSuccessful(&Move{From: "f2", To: "f4", Piece: WhitePawn}, positionAfter1F4)

	this.assertFirstMoveSuccessful(&Move{From: "g2", To: "g3", Piece: WhitePawn}, positionAfter1G3)
	this.assertFirstMoveSuccessful(&Move{From: "g2", To: "g4", Piece: WhitePawn}, positionAfter1G4)

	this.assertFirstMoveSuccessful(&Move{From: "h2", To: "h3", Piece: WhitePawn}, positionAfter1H3)
	this.assertFirstMoveSuccessful(&Move{From: "h2", To: "h4", Piece: WhitePawn}, positionAfter1H4)
}
func (this *GameFixture) assertFirstMoveSuccessful(move *Move, expectedFEN string) {
	this.game.Reset()
	err := this.game.Move(move)
	this.So(err, should.BeNil)
	this.So(this.game.FEN(), should.Equal, expectedFEN)
}

const positionAfter1A3 = "rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1A4 = "rnbqkbnr/pppppppp/8/8/P7/8/1PPPPPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1B3 = "rnbqkbnr/pppppppp/8/8/8/1P6/P1PPPPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1B4 = "rnbqkbnr/pppppppp/8/8/1P6/8/P1PPPPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1C3 = "rnbqkbnr/pppppppp/8/8/8/2P5/PP1PPPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1C4 = "rnbqkbnr/pppppppp/8/8/2P5/8/PP1PPPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1D3 = "rnbqkbnr/pppppppp/8/8/8/3P4/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1D4 = "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1E3 = "rnbqkbnr/pppppppp/8/8/8/4P3/PPPP1PPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1E4 = "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1"
const positionAfter1F3 = "rnbqkbnr/pppppppp/8/8/8/5P2/PPPPP1PP/RNBQKBNR b KQkq - 0 1"
const positionAfter1F4 = "rnbqkbnr/pppppppp/8/8/5P2/8/PPPPP1PP/RNBQKBNR b KQkq - 0 1"
const positionAfter1G3 = "rnbqkbnr/pppppppp/8/8/8/6P1/PPPPPP1P/RNBQKBNR b KQkq - 0 1"
const positionAfter1G4 = "rnbqkbnr/pppppppp/8/8/6P1/8/PPPPPP1P/RNBQKBNR b KQkq - 0 1"
const positionAfter1H3 = "rnbqkbnr/pppppppp/8/8/8/7P/PPPPPPP1/RNBQKBNR b KQkq - 0 1"
const positionAfter1H4 = "rnbqkbnr/pppppppp/8/8/7P/8/PPPPPPP1/RNBQKBNR b KQkq - 0 1"
