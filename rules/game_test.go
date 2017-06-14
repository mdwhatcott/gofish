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

func (this *GameFixture) TestLoadFEN() {
}

const startingPositionFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
