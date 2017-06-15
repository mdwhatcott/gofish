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
	this.So(this.game.ExportFEN(), should.Equal, startingPositionFEN)
}

func (this *GameFixture) TestGameConditionsAfterFirstPawnMove() {
	err := this.game.Move(Move{From: ParseSquare("a2"), To: ParseSquare("a3"), Piece: WhitePawn})

	this.So(err, should.BeNil)
	this.So(this.game.IsOver(), should.BeFalse)
	this.So(this.game.PlayerToMove(), should.Equal, Black)
	this.So(this.game.FullMoveCount(), should.Equal, 1)
	this.So(this.game.HalfMoveCount(), should.Equal, 0) // pawn move
	this.So(this.game.ExportFEN(), should.Equal, positionAfter1A3)
}

func (this *GameFixture) TestLoadFEN() {
	const kingsOnBackRanks = "4k3/8/8/8/8/8/8/4K3 w - - 0 1"
	err := this.game.LoadFEN(kingsOnBackRanks)
	this.So(err, should.BeNil)
	this.So(this.game.ExportFEN(), should.Equal, kingsOnBackRanks)
}

func (this *GameFixture) TestLegalKingMoves() {
	this.game.MustLoadFEN("8/1k6/8/p7/P7/8/1K6/8 w - - 0 1") // white king on b2, surrounding squares empty
	moves := this.game.CalculateAvailableMoves()
	this.So(moves, should.HaveLength, 8)
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("a1")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("a2")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("a3")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("b1")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("b3")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("c1")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("c2")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b2"), To: ParseSquare("c3")})
}

func (this *GameFixture) TestLegalKingMoves_KingOnBottomEdge() {
	this.game.MustLoadFEN("8/1k6/8/p7/P7/8/8/1K6 w - - 0 1") // white king on b1
	moves := this.game.CalculateAvailableMoves()
	this.So(moves, should.HaveLength, 5)
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b1"), To: ParseSquare("a1")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b1"), To: ParseSquare("c1")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b1"), To: ParseSquare("a2")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b1"), To: ParseSquare("b2")})
	this.So(moves, should.Contain, Move{Piece: WhiteKing, From: ParseSquare("b1"), To: ParseSquare("c2")})
}

func (this *GameFixture) TestLegalKingMoves_KingOnTopEdge() {
	this.game.MustLoadFEN("1k6/8/8/p7/P7/8/8/1K6 b - - 0 1") // white king on b1
	moves := this.game.CalculateAvailableMoves()
	this.So(moves, should.HaveLength, 5)
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("b8"), To: ParseSquare("a8")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("b8"), To: ParseSquare("c8")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("b8"), To: ParseSquare("a7")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("b8"), To: ParseSquare("b7")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("b8"), To: ParseSquare("c7")})
}

func (this *GameFixture) TestLegalKingMoves_KingOnLeftEdge() {
	this.game.MustLoadFEN("8/k7/8/p7/P7/8/8/1K6 b - - 0 1") // white king on b1
	moves := this.game.CalculateAvailableMoves()
	this.So(moves, should.HaveLength, 5)
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("a7"), To: ParseSquare("a8")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("a7"), To: ParseSquare("a6")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("a7"), To: ParseSquare("b8")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("a7"), To: ParseSquare("b7")})
	this.So(moves, should.Contain, Move{Piece: BlackKing, From: ParseSquare("a7"), To: ParseSquare("b6")})
}

// TODO: Rook moves (Can land vertically or horizontally as far as first obstacle that is enemy (capture), is blocked by ally)
// TODO: Bishop moves: Can land diagonally as far as first obstacle that is enemy (capture), is blocked by ally
// TODO: Knight moves: Can jump over any piece to land on empty or enemy (capture)
// TODO: Queen moves: Combined movements of Bishop and Rook
// TODO: Legal pawn moves: advance 1 rank or optionally 2 ranks if on starting square as long as ending square is empty,
// TODO: Legal pawn captures: capture diagonally or en-passant if on its "5th" rank and eligible opposing pawn target exists
// TODO: enforce castling limitations
// TODO: detect discovered check
// TODO: prevent illegal move into check
// TODO: detect checkmate
// TODO: detect stalemate
// TODO: detect draw by insufficient material
// TODO: detect three-fold repetition
// TODO: detect 50-move rule violation
// TODO: Load/Export PGN

func (this *GameFixture) TestLegalFirstMoves() {
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("a2"), To: ParseSquare("a3"), Piece: WhitePawn}, positionAfter1A3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("a2"), To: ParseSquare("a4"), Piece: WhitePawn}, positionAfter1A4)

	this.assertFirstMoveSuccessful(Move{From: ParseSquare("b2"), To: ParseSquare("b3"), Piece: WhitePawn}, positionAfter1B3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("b2"), To: ParseSquare("b4"), Piece: WhitePawn}, positionAfter1B4)

	this.assertFirstMoveSuccessful(Move{From: ParseSquare("c2"), To: ParseSquare("c3"), Piece: WhitePawn}, positionAfter1C3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("c2"), To: ParseSquare("c4"), Piece: WhitePawn}, positionAfter1C4)

	this.assertFirstMoveSuccessful(Move{From: ParseSquare("d2"), To: ParseSquare("d3"), Piece: WhitePawn}, positionAfter1D3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("d2"), To: ParseSquare("d4"), Piece: WhitePawn}, positionAfter1D4)

	this.assertFirstMoveSuccessful(Move{From: ParseSquare("e2"), To: ParseSquare("e3"), Piece: WhitePawn}, positionAfter1E3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("e2"), To: ParseSquare("e4"), Piece: WhitePawn}, positionAfter1E4)

	this.assertFirstMoveSuccessful(Move{From: ParseSquare("f2"), To: ParseSquare("f3"), Piece: WhitePawn}, positionAfter1F3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("f2"), To: ParseSquare("f4"), Piece: WhitePawn}, positionAfter1F4)

	this.assertFirstMoveSuccessful(Move{From: ParseSquare("g2"), To: ParseSquare("g3"), Piece: WhitePawn}, positionAfter1G3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("g2"), To: ParseSquare("g4"), Piece: WhitePawn}, positionAfter1G4)

	this.assertFirstMoveSuccessful(Move{From: ParseSquare("h2"), To: ParseSquare("h3"), Piece: WhitePawn}, positionAfter1H3)
	this.assertFirstMoveSuccessful(Move{From: ParseSquare("h2"), To: ParseSquare("h4"), Piece: WhitePawn}, positionAfter1H4)
}
func (this *GameFixture) assertFirstMoveSuccessful(move Move, expectedFEN string) {
	this.game.Reset()
	err := this.game.Move(move)
	this.So(err, should.BeNil)
	this.So(this.game.ExportFEN(), should.Equal, expectedFEN)
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
