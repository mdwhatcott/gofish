package console

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestBoardDecorationsFixture(t *testing.T) {
	gunit.Run(new(BoardDecorationsFixture), t)
}

type BoardDecorationsFixture struct {
	*gunit.Fixture
}

func (this *BoardDecorationsFixture) TestCoordinateBorders() {
	decorated := NewCoordinateBoard(startingPosition).String()
	this.AssertEqual(`
  ＡＢＣＤＥＦＧＨ
8 ♜♞♝♛♚♝♞♜ 8
7 ♟♟♟♟♟♟♟♟ 7
6 ＿＿＿＿＿＿＿＿ 6
5 ＿＿＿＿＿＿＿＿ 5
4 ＿＿＿＿＿＿＿＿ 4
3 ＿＿＿＿＿＿＿＿ 3
2 ♙♙♙♙♙♙♙♙ 2
1 ♖♘♗♕♔♗♘♖ 1
  ＡＢＣＤＥＦＧＨ
`, decorated)
}
