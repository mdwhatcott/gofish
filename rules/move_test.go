package rules

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestMoveRepresentationFixture(t *testing.T) {
	gunit.Run(new(MoveRepresentationFixture), t)
}

type MoveRepresentationFixture struct {
	*gunit.Fixture
}

func (this *MoveRepresentationFixture) TestPawnMoves() {
	this.assertMove("a4", move{
		Piece: WhitePawn,
		From:  Square("a2"),
		To:    Square("a4"),
	})
	this.assertMove("a3+", move{
		Piece: WhitePawn, From: Square("a2"), To: Square("a3"),
		Check: true,
	})
	this.assertMove("a3#", move{
		Piece: WhitePawn, From: Square("a2"), To: Square("a3"),
		Checkmate: true,
	})
}
func (this *MoveRepresentationFixture) TestPawnCaptures() {
	this.assertMove("axb3", move{
		Piece: WhitePawn, From: Square("a2"), To: Square("b3"),
		Captured: BlackPawn, CapturedOn: Square("b3"),
	})
	this.assertMove("axb3+", move{
		Piece: WhitePawn, From: Square("a2"), To: Square("b3"),
		Captured: BlackPawn, CapturedOn: Square("b3"),
		Check: true,
	})
	this.assertMove("axb3#", move{
		Piece: WhitePawn, From: Square("a2"), To: Square("b3"),
		Captured: BlackPawn, CapturedOn: Square("b3"),
		Checkmate: true,
	})
}
func (this *MoveRepresentationFixture) TestPawnPromotions() {
	this.assertMove("a8=Q", move{
		Piece: WhitePawn, From: Square("a7"), To: Square("a8"),
		Promotion: WhiteQueen,
	})
	this.assertMove("a8=R", move{
		Piece: WhitePawn, From: Square("a7"), To: Square("a8"),
		Promotion: WhiteRook,
	})
	this.assertMove("a8=B", move{
		Piece: WhitePawn, From: Square("a7"), To: Square("a8"),
		Promotion: WhiteBishop,
	})
	this.assertMove("a8=N", move{
		Piece: WhitePawn, From: Square("a7"), To: Square("a8"),
		Promotion: WhiteKnight,
	})
	this.assertMove("a8=Q+", move{
		Piece: WhitePawn, From: Square("a7"), To: Square("a8"),
		Promotion: WhiteQueen,
		Check:     true,
	})
	this.assertMove("a8=N#", move{
		Piece: WhitePawn, From: Square("a7"), To: Square("a8"),
		Promotion: WhiteKnight,
		Checkmate: true,
	})
}
func (this *MoveRepresentationFixture) TestPawnCapturePromotions() {
	this.assertMove("bxa8=Q", move{
		Piece: WhitePawn, From: Square("b7"), To: Square("a8"),
		Captured: BlackKing, CapturedOn: Square("a8"),
		Promotion: WhiteQueen,
	})
	this.assertMove("bxa8=Q+", move{
		Piece: WhitePawn, From: Square("b7"), To: Square("a8"),
		Captured: BlackKing, CapturedOn: Square("a8"),
		Promotion: WhiteQueen,
		Check:     true,
	})
	this.assertMove("bxa8=Q#", move{
		Piece: WhitePawn, From: Square("b7"), To: Square("a8"),
		Captured: BlackKing, CapturedOn: Square("a8"),
		Promotion: WhiteQueen,
		Checkmate: true,
	})
}

func (this *MoveRepresentationFixture) TestPieceMoves() {
	this.assertMove("Nc3", move{Piece: WhiteKnight, From: Square("b1"), To: Square("c3")})
	this.assertMove("Nc6", move{Piece: BlackKnight, From: Square("b8"), To: Square("c6")})

	this.assertMove("Bg2", move{Piece: WhiteBishop, From: Square("f1"), To: Square("g2")})
	this.assertMove("Bg7", move{Piece: BlackBishop, From: Square("f8"), To: Square("g7")})

	this.assertMove("Re1", move{Piece: WhiteRook, From: Square("h1"), To: Square("e1")})
	this.assertMove("Re8", move{Piece: BlackRook, From: Square("h8"), To: Square("e8")})

	this.assertMove("Qd2", move{Piece: WhiteQueen, From: Square("d1"), To: Square("d2")})
	this.assertMove("Qd7", move{Piece: BlackQueen, From: Square("d8"), To: Square("d7")})

	this.assertMove("Kd1", move{Piece: WhiteKing, From: Square("e1"), To: Square("d1")})
	this.assertMove("Kd8", move{Piece: BlackKing, From: Square("e8"), To: Square("d8")})
}

func (this *MoveRepresentationFixture) TestPieceMoves_Check() {
	this.assertMove("Nc3+", move{Piece: WhiteKnight, From: Square("b1"), To: Square("c3"), Check: true})
	this.assertMove("Nc6+", move{Piece: BlackKnight, From: Square("b8"), To: Square("c6"), Check: true})

	this.assertMove("Bg2+", move{Piece: WhiteBishop, From: Square("f1"), To: Square("g2"), Check: true})
	this.assertMove("Bg7+", move{Piece: BlackBishop, From: Square("f8"), To: Square("g7"), Check: true})

	this.assertMove("Re1+", move{Piece: WhiteRook, From: Square("h1"), To: Square("e1"), Check: true})
	this.assertMove("Re8+", move{Piece: BlackRook, From: Square("h8"), To: Square("e8"), Check: true})

	this.assertMove("Qd2+", move{Piece: WhiteQueen, From: Square("d1"), To: Square("d2"), Check: true})
	this.assertMove("Qd7+", move{Piece: BlackQueen, From: Square("d8"), To: Square("d7"), Check: true})

	this.assertMove("Kd1+", move{Piece: WhiteKing, From: Square("e1"), To: Square("d1"), Check: true})
	this.assertMove("Kd8+", move{Piece: BlackKing, From: Square("e8"), To: Square("d8"), Check: true})
}

func (this *MoveRepresentationFixture) TestPieceMoves_Checkmate() {
	this.assertMove("Nc3#", move{Piece: WhiteKnight, From: Square("b1"), To: Square("c3"), Checkmate: true})
	this.assertMove("Nc6#", move{Piece: BlackKnight, From: Square("b8"), To: Square("c6"), Checkmate: true})

	this.assertMove("Bg2#", move{Piece: WhiteBishop, From: Square("f1"), To: Square("g2"), Checkmate: true})
	this.assertMove("Bg7#", move{Piece: BlackBishop, From: Square("f8"), To: Square("g7"), Checkmate: true})

	this.assertMove("Re1#", move{Piece: WhiteRook, From: Square("h1"), To: Square("e1"), Checkmate: true})
	this.assertMove("Re8#", move{Piece: BlackRook, From: Square("h8"), To: Square("e8"), Checkmate: true})

	this.assertMove("Qd2#", move{Piece: WhiteQueen, From: Square("d1"), To: Square("d2"), Checkmate: true})
	this.assertMove("Qd7#", move{Piece: BlackQueen, From: Square("d8"), To: Square("d7"), Checkmate: true})

	this.assertMove("Kd1#", move{Piece: WhiteKing, From: Square("e1"), To: Square("d1"), Checkmate: true})
	this.assertMove("Kd8#", move{Piece: BlackKing, From: Square("e8"), To: Square("d8"), Checkmate: true})
}

func (this *MoveRepresentationFixture) TestPieceCaptures() {
	this.assertMove("Nxa1", move{Piece: WhiteKnight, From: Square("b3"), To: Square("a1"), Captured: BlackPawn})
	this.assertMove("Nxa1", move{Piece: BlackKnight, From: Square("b3"), To: Square("a1"), Captured: WhitePawn})

	this.assertMove("Bxa2", move{Piece: WhiteBishop, From: Square("b3"), To: Square("a2"), Captured: BlackPawn})
	this.assertMove("Bxa2", move{Piece: BlackBishop, From: Square("b3"), To: Square("a2"), Captured: WhitePawn})

	this.assertMove("Rxa1", move{Piece: WhiteRook, From: Square("b1"), To: Square("a1"), Captured: BlackPawn})
	this.assertMove("Rxa1", move{Piece: BlackRook, From: Square("b1"), To: Square("a1"), Captured: WhitePawn})

	this.assertMove("Qxa1", move{Piece: WhiteQueen, From: Square("b1"), To: Square("a1"), Captured: BlackPawn})
	this.assertMove("Qxa1", move{Piece: BlackQueen, From: Square("b1"), To: Square("a1"), Captured: WhitePawn})

	this.assertMove("Kxa1", move{Piece: WhiteKing, From: Square("b1"), To: Square("a1"), Captured: BlackPawn})
	this.assertMove("Kxa1", move{Piece: BlackKing, From: Square("b1"), To: Square("a1"), Captured: WhitePawn})
}

func (this *MoveRepresentationFixture) TestAmbiguousPieceMoves() {
	this.assertMove("Rhe1", move{
		Piece:   WhiteRook,
		From:    Square("h1"),
		To:      Square("e1"),
		NotFrom: []square{Square("a1")},
	})
	this.assertMove("R1a2", move{
		Piece:   WhiteRook,
		From:    Square("a1"),
		To:      Square("a2"),
		NotFrom: []square{Square("a3")},
	})
	this.assertMove("Rh1e1", move{
		Piece:   WhiteRook,
		From:    Square("h1"),
		To:      Square("e1"),
		NotFrom: []square{Square("a1"), Square("h2")},
	})
	this.assertMove("Ra1e1", move{
		Piece:   WhiteRook,
		From:    Square("a1"),
		To:      Square("e1"),
		NotFrom: []square{Square("h1"), Square("h2")},
	})
}

func (this *MoveRepresentationFixture) TestCastling() {
	this.assertMove("O-O", move{
		Piece:   WhiteRook,
		From:    Square("h1"),
		To:      Square("e1"),
		Castles: true,
	})
	this.assertMove("O-O-O", move{
		Piece:   WhiteRook,
		From:    Square("a1"),
		To:      Square("d1"),
		Castles: true,
	})
	this.assertMove("O-O", move{
		Piece:   BlackRook,
		From:    Square("h8"),
		To:      Square("e8"),
		Castles: true,
	})
	this.assertMove("O-O-O", move{
		Piece:   BlackRook,
		From:    Square("a8"),
		To:      Square("d8"),
		Castles: true,
	})
}

func (this *MoveRepresentationFixture) assertMove(expected string, move move) {
	this.AssertSprintEqual(expected, move)
}
