package rules

import (
	"sort"
	"testing"

	"github.com/mdwhatcott/gofish/console"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestLegalFirstMovesFixture(t *testing.T) {
	gunit.Run(new(LegalFirstMovesFixture), t)
}

type LegalFirstMovesFixture struct {
	*gunit.Fixture
	*LegalMovesFixture
}

func (this *LegalFirstMovesFixture) Setup() {
	this.LegalMovesFixture = NewLegalGameMovesFixture(this.Fixture)
}

func (this *LegalFirstMovesFixture) TestLegalFirstMovesForWhite() {
	this.assertAllLegalMoves(this.game.GetLegalMoves(this.game.PlayerToMove()),
		"a3", "a4",
		"b3", "b4",
		"c3", "c4",
		"d3", "d4",
		"e3", "e4",
		"f3", "f4",
		"g3", "g4",
		"h3", "h4",
		"Na3", "Nc3",
		"Nf3", "Nh3",
	)
}

func (this *LegalFirstMovesFixture) TestLegalFirstMovesForBlack() {
	this.game.Execute(move{From: Square("a2"), To: Square("a3"), Piece: WhitePawn})
	this.assertAllLegalMoves(this.game.GetLegalMoves(this.game.PlayerToMove()),
		"a6", "a5",
		"b6", "b5",
		"c6", "c5",
		"d6", "d5",
		"e6", "e5",
		"f6", "f5",
		"g6", "g5",
		"h6", "h5",
		"Na6", "Nc6",
		"Nf6", "Nh6",
	)
}

/**************************************************************************/

type LegalMovesFixture struct {
	*gunit.Fixture

	game *Game
}

func (this *LegalMovesFixture) Setup() {
	this.game = NewGame()
}

func NewLegalGameMovesFixture(fixture *gunit.Fixture) *LegalMovesFixture {
	this := &LegalMovesFixture{Fixture: fixture}
	this.Setup()
	return this
}

func (this *LegalMovesFixture) assertAllLegalMoves(actualMoves []move, expectedMovesSAN ...string) {
	ok := this.So(actualMoves, should.HaveLength, len(expectedMovesSAN))
	actualMovesSan := []string{}
	for _, move := range actualMoves {
		ok = ok && this.So(move.To.String(), should.NotResemble, move.From.String())
		actualMovesSan = append(actualMovesSan, move.String())
	}
	sort.Strings(actualMovesSan)
	sort.Strings(expectedMovesSAN)
	ok = ok && this.So(actualMovesSan, should.Resemble, expectedMovesSAN)
	if !ok {
		picture := console.NewBoard()
		picture.Setup(this.game.ExportFEN().String())
		this.Println(console.NewCoordinateBoard(picture.String()))
		this.Println("Moves:", actualMoves)
	}
}

func (this *LegalMovesFixture) assertLegalPieceMoves(
	position string, from string, piece piece, expectedPieceMovesSAN ...string) {

	if len(expectedPieceMovesSAN) == 0 {
		expectedPieceMovesSAN = []string{}
	}

	this.game.MustLoadFEN(position)

	moves := this.game.GetLegalMoves(piece.Player())
	actualPieceMoves := filterMovesByPieceOnSquare(moves, piece, from)

	this.assertAllLegalMoves(actualPieceMoves, expectedPieceMovesSAN...)
}

func filterMovesByPieceOnSquare(moves []move, piece piece, from string) (filtered []move) {
	for _, move := range moves {
		if move.Piece == piece && move.From.String() == from {
			filtered = append(filtered, move)
		}
	}
	return filtered
}

func (this *LegalMovesFixture) assertPosition(expected string) {
	actual := this.game.ExportFEN().String()
	if !this.So(actual, should.StartWith, expected) {
		expectedBoard := console.NewBoard()
		expectedBoard.Setup(expected)

		actualBoard := console.NewBoard()
		actualBoard.Setup(actual)

		this.Println("Expected:\n", expectedBoard.String())
		this.Println("Actual:\n", actualBoard.String())
	}
}

/**************************************************************************/

type LegalMovesSetup struct {
	InitialPositionFEN  string
	PreparatoryMovesSAN []string
	ExpectedMovesSAN    []string
	ExpectedPositionFEN string
	FocusOnPiece        piece
	FromSquare          string
}

func (this *LegalMovesFixture) Play(setup LegalMovesSetup) {
	if setup.InitialPositionFEN != "" {
		this.game.MustLoadFEN(setup.InitialPositionFEN)
	} else {
		this.game.Reset()
	}

	for _, move := range setup.PreparatoryMovesSAN {
		this.game.Attempt(move)
	}
}
func (this *LegalMovesFixture) PlayAndValidate(setup LegalMovesSetup) {
	this.Play(setup)

	resultPosition := this.game.ExportFEN().String()

	if setup.ExpectedPositionFEN != "" {
		this.assertPosition(setup.ExpectedPositionFEN)
	}

	if setup.FocusOnPiece != Void {
		this.assertLegalPieceMoves(resultPosition, setup.FromSquare, setup.FocusOnPiece, setup.ExpectedMovesSAN...)
	} else {
		this.assertAllLegalMoves(this.game.GetLegalMoves(this.game.PlayerToMove()), setup.ExpectedMovesSAN...)
	}

}
