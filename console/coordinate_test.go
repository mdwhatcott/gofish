package console

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSquareIndexFixture(t *testing.T) { gunit.Run(new(SquareIndexFixture), t) }

type SquareIndexFixture struct{ *gunit.Fixture }

func (this *SquareIndexFixture) TestInvalidSquaresShouldCausePanic() {
	this.So(func() { squareIndex("") }, should.Panic)
	this.So(func() { squareIndex(" ") }, should.Panic)
	this.So(func() { squareIndex("a") }, should.Panic)
	this.So(func() { squareIndex("1") }, should.Panic)
	this.So(func() { squareIndex(" 1") }, should.Panic)
	this.So(func() { squareIndex("a ") }, should.Panic)
	this.So(func() { squareIndex("a9") }, should.Panic)
	this.So(func() { squareIndex("a0") }, should.Panic)
	this.So(func() { squareIndex("q5") }, should.Panic)
	this.So(func() { squareIndex("A5") }, should.Panic)
}

func (this *SquareIndexFixture) TestAllSquaresCanBeIndexedCorrectly() {
	this.AssertEqual(0, squareIndex("a8"))
	this.AssertEqual(1, squareIndex("b8"))
	this.AssertEqual(2, squareIndex("c8"))
	this.AssertEqual(3, squareIndex("d8"))
	this.AssertEqual(4, squareIndex("e8"))
	this.AssertEqual(5, squareIndex("f8"))
	this.AssertEqual(6, squareIndex("g8"))
	this.AssertEqual(7, squareIndex("h8"))

	this.AssertEqual(8, squareIndex("a7"))
	this.AssertEqual(9, squareIndex("b7"))
	this.AssertEqual(10, squareIndex("c7"))
	this.AssertEqual(11, squareIndex("d7"))
	this.AssertEqual(12, squareIndex("e7"))
	this.AssertEqual(13, squareIndex("f7"))
	this.AssertEqual(14, squareIndex("g7"))
	this.AssertEqual(15, squareIndex("h7"))

	this.AssertEqual(16, squareIndex("a6"))
	this.AssertEqual(17, squareIndex("b6"))
	this.AssertEqual(18, squareIndex("c6"))
	this.AssertEqual(19, squareIndex("d6"))
	this.AssertEqual(20, squareIndex("e6"))
	this.AssertEqual(21, squareIndex("f6"))
	this.AssertEqual(22, squareIndex("g6"))
	this.AssertEqual(23, squareIndex("h6"))

	this.AssertEqual(24, squareIndex("a5"))
	this.AssertEqual(25, squareIndex("b5"))
	this.AssertEqual(26, squareIndex("c5"))
	this.AssertEqual(27, squareIndex("d5"))
	this.AssertEqual(28, squareIndex("e5"))
	this.AssertEqual(29, squareIndex("f5"))
	this.AssertEqual(30, squareIndex("g5"))
	this.AssertEqual(31, squareIndex("h5"))

	this.AssertEqual(32, squareIndex("a4"))
	this.AssertEqual(33, squareIndex("b4"))
	this.AssertEqual(34, squareIndex("c4"))
	this.AssertEqual(35, squareIndex("d4"))
	this.AssertEqual(36, squareIndex("e4"))
	this.AssertEqual(37, squareIndex("f4"))
	this.AssertEqual(38, squareIndex("g4"))
	this.AssertEqual(39, squareIndex("h4"))

	this.AssertEqual(40, squareIndex("a3"))
	this.AssertEqual(41, squareIndex("b3"))
	this.AssertEqual(42, squareIndex("c3"))
	this.AssertEqual(43, squareIndex("d3"))
	this.AssertEqual(44, squareIndex("e3"))
	this.AssertEqual(45, squareIndex("f3"))
	this.AssertEqual(46, squareIndex("g3"))
	this.AssertEqual(47, squareIndex("h3"))

	this.AssertEqual(48, squareIndex("a2"))
	this.AssertEqual(49, squareIndex("b2"))
	this.AssertEqual(50, squareIndex("c2"))
	this.AssertEqual(51, squareIndex("d2"))
	this.AssertEqual(52, squareIndex("e2"))
	this.AssertEqual(53, squareIndex("f2"))
	this.AssertEqual(54, squareIndex("g2"))
	this.AssertEqual(55, squareIndex("h2"))

	this.AssertEqual(56, squareIndex("a1"))
	this.AssertEqual(57, squareIndex("b1"))
	this.AssertEqual(58, squareIndex("c1"))
	this.AssertEqual(59, squareIndex("d1"))
	this.AssertEqual(60, squareIndex("e1"))
	this.AssertEqual(61, squareIndex("f1"))
	this.AssertEqual(62, squareIndex("g1"))
	this.AssertEqual(63, squareIndex("h1"))
}
