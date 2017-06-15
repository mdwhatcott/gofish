package rules

type Square struct {
	Rank int
	File int
}

func NewSquare(i int) Square {
	return Square{
		Rank: i / 8,
		File: i % 8,
	}
}

func ParseSquare(algebraic string) Square {
	return Square{
		File: int(algebraic[0] - 'a'),
		Rank: int(algebraic[1] - '1'),
	}
}

func (this Square) Offset(delta Square) Square {
	return Square{
		Rank: this.Rank + delta.Rank,
		File: this.File + delta.File,
	}
}

func (this Square) IsValid() bool {
	return this.Rank >= 0 && this.Rank < 8 &&
		this.File >= 0 && this.File < 8
}

func (this Square) String() string {
	return string('a'+this.File) + string('1'+this.Rank)
}

func (this Square) Int() int {
	return this.Rank*8 + this.File
}
