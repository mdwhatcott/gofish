package rules

type square struct {
	rank int
	file int
}

func IntSquare(i int) square {
	return square{
		rank: i / 8,
		file: i % 8,
	}
}

func Square(algebraic string) square {
	return square{
		file: int(algebraic[0] - 'a'),
		rank: int(algebraic[1] - '1'),
	}
}

func (this square) Offset(delta square) square {
	return square{
		rank: this.rank + delta.rank,
		file: this.file + delta.file,
	}
}

func (this square) IsValidSquare() bool {
	return this.rank >= 0 && this.rank < 8 &&
		this.file >= 0 && this.file < 8
}

func (this square) String() string {
	return string('a'+this.file) + string('1'+this.rank)
}

func (this square) Int() int {
	return this.rank*8 + this.file
}

func (this square) Rank() string {
	return string(this.rank + 1 + '0')
}
func (this square) File() string {
	return string('a' + this.file)
}

/**************************************************************************/

type squares []square

func (this squares) ranksAreUnique() bool {
	unique := make(map[string]struct{})
	for _, square := range this {
		unique[square.Rank()] = struct{}{}
	}
	return len(unique) == len(this)
}

func (this squares) filesAreUnique() bool {
	unique := make(map[string]struct{})
	for _, square := range this {
		unique[square.File()] = struct{}{}
	}
	return len(unique) == len(this)
}
