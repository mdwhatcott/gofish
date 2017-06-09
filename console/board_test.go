package console

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestBoardFixture(t *testing.T) {
	gunit.Run(new(BoardFixture), t)
}

type BoardFixture struct {
	*gunit.Fixture
	board *Board
}

func (this *BoardFixture) Setup() {
	this.board = NewBoard()
}

func (this *BoardFixture) assertPosition(expected string) {
	this.AssertEqual(expected, this.board.String())
}

func (this *BoardFixture) TestInitialStateOfBoard_Blank() {
	this.assertPosition(blankPosition)
}

const blankPosition = `
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
`

func (this *BoardFixture) TestDeriveInitialPositionFromFEN() {
	this.board.Setup(startingPositionFEN)
	this.assertPosition(startingPosition)
}

const startingPositionFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
const startingPosition = `
♜♞♝♛♚♝♞♜
♟♟♟♟♟♟♟♟
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
＿＿＿＿＿＿＿＿
♙♙♙♙♙♙♙♙
♖♘♗♕♔♗♘♖
`

func (this *BoardFixture) TestDeriveRuyLopezPositionFromFEN() {
	this.board.Setup(ruyLopezFEN)
	this.assertPosition(ruyLopezPosition)
}

const ruyLopezFEN = "r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R"
const ruyLopezPosition = `
♜＿♝♛♚♝♞♜
♟♟♟♟＿♟♟♟
＿＿♞＿＿＿＿＿
＿♗＿＿♟＿＿＿
＿＿＿＿♙＿＿＿
＿＿＿＿＿♘＿＿
♙♙♙♙＿♙♙♙
♖♘♗♕♔＿＿♖
`

func (this *BoardFixture) TestDeriveMiddleGamePositionFromFEN() {
	this.board.Setup(middleGameFEN)
	this.assertPosition(middleGamePosition)
}

const middleGameFEN = "r1k4r/p2nb1p1/2b4p/1p1n1p2/2PP4/3Q1NB1/1P3PPP/R5K1"
const middleGamePosition = `
♜＿♚＿＿＿＿♜
♟＿＿♞♝＿♟＿
＿＿♝＿＿＿＿♟
＿♟＿♞＿♟＿＿
＿＿♙♙＿＿＿＿
＿＿＿♕＿♘♗＿
＿♙＿＿＿♙♙♙
♖＿＿＿＿＿♔＿
`

func (this *BoardFixture) TestSetupFENShouldResetEverySquare() {
	this.board.Setup(middleGameFEN)
	this.board.Setup(blankBoardFEN)
	this.assertPosition(blankPosition)
}

const blankBoardFEN = "8/8/8/8/8/8/8/8"
