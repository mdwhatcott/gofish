package rules

import (
	"testing"

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

const ()

func (this *KnightMovesFixture) SkipTestAwayFromAnyEdge() {
}
