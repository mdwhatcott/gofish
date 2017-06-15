package rules

import (
	"strconv"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSquareFixture(t *testing.T) {
	gunit.Run(new(SquareFixture), t)
}

type SquareFixture struct {
	*gunit.Fixture
}

func (this *SquareFixture) TestAllSquaresOnTheBoardAreValid() {
	x := 0
	for rank := 1; rank <= 8; rank++ {
		for _, file := range "abcdefgh" {
			square := NewSquare(x)
			this.So(square.IsValid(), should.BeTrue)
			this.So(square.Int(), should.Equal, x)
			this.So(square.String(), should.Equal, string(file)+strconv.Itoa(rank))
			x++
		}
	}
}

func (this *SquareFixture) TestOffsetsProduceInvalidSquaresAsTheyGoOffTheBoard() {
	var (
		leftEdge   = []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8"}
		rightEdge  = []string{"h1", "h2", "h3", "h4", "h5", "h6", "h7", "h8"}
		topEdge    = []string{"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8"}
		bottomEdge = []string{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"}
		inner      = []string{
			"b2", "b3", "b4", "b5", "b6", "b7",
			"c2", "c3", "c4", "c5", "c6", "c7",
			"d2", "d3", "d4", "d5", "d6", "d7",
			"e2", "e3", "e4", "e5", "e6", "e7",
			"f2", "f3", "f4", "f5", "f6", "f7",
			"g2", "g3", "g4", "g5", "g6", "g7",
		}
	)

	for _, square := range leftEdge { // Falling off the left edge:
		this.So(ParseSquare(square).Offset(Square{File: -1}).IsValid(), should.BeFalse)
	}
	for _, square := range rightEdge { // Falling off the right edge:
		this.So(ParseSquare(square).Offset(Square{File: 1}).IsValid(), should.BeFalse)
	}
	for _, square := range bottomEdge { // Falling off the bottom edge:
		this.So(ParseSquare(square).Offset(Square{Rank: -1}).IsValid(), should.BeFalse)
	}
	for _, square := range topEdge { // Falling off the top edge:
		this.So(ParseSquare(square).Offset(Square{Rank: 1}).IsValid(), should.BeFalse)
	}
	for _, square := range inner { // Offsets of 1 are fine for non-edge squares:
		from := ParseSquare(square)
		this.So(from.Offset(Square{Rank: 1}).IsValid(), should.BeTrue)
		this.So(from.Offset(Square{Rank: -1}).IsValid(), should.BeTrue)
		this.So(from.Offset(Square{File: 1}).IsValid(), should.BeTrue)
		this.So(from.Offset(Square{File: -1}).IsValid(), should.BeTrue)
	}
}
