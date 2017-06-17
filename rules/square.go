package rules

type square struct {
	Rank int
	File int
}

func IntSquare(i int) square {
	return square{
		Rank: i / 8,
		File: i % 8,
	}
}

func Square(algebraic string) square {
	return square{
		File: int(algebraic[0] - 'a'),
		Rank: int(algebraic[1] - '1'),
	}
}

func (this square) Offset(delta square) square {
	return square{
		Rank: this.Rank + delta.Rank,
		File: this.File + delta.File,
	}
}

func (this square) IsValidSquare() bool {
	return this.Rank >= 0 && this.Rank < 8 &&
		this.File >= 0 && this.File < 8
}

func (this square) String() string {
	return string('a'+this.File) + string('1'+this.Rank)
}

func (this square) Int() int {
	return this.Rank*8 + this.File
}
