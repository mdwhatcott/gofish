package console

import (
	"bytes"
	"unicode"
)

type Board struct {
	squares []unit
}

func NewBoard() *Board {
	squares := make([]unit, rankCount*fileCount)
	for x := 0; x < len(squares); x++ {
		squares[x] = Void
	}
	return &Board{squares: squares}
}

func (this *Board) Setup(fen string) {
	i := 0
	for _, c := range fen {
		if c == ' ' {
			break
		} else if c == '/' {
			continue
		} else if unicode.IsDigit(c) {
			offset := int(c - '0')
			for x := 0; x < offset; x++ {
				this.place(Void, i)
				i++
			}
		} else {
			this.place(FENUnits[c], i)
			i++
		}
	}
}

func (this *Board) place(unit unit, square int) {
	this.squares[square] = unit
}

func (this *Board) String() string {
	var buffer bytes.Buffer
	for i, square := range this.squares {
		if i%rankCount == 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteRune(rune(square))
	}
	buffer.WriteString("\n")
	return buffer.String()
}

const rankCount = 8
const fileCount = 8
