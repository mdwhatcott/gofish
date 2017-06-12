package console

import (
	"bytes"
	"strconv"
	"strings"
)

const fullWidthFiles = "ＡＢＣＤＥＦＧＨ"

type CoordinateBoard struct {
	buffer *bytes.Buffer
	board  []string
}

func NewCoordinateBoard(board string) *CoordinateBoard {
	return &CoordinateBoard{
		buffer: bytes.NewBufferString("\n"),
		board:  strings.Split(strings.TrimSpace(board), "\n"),
	}
}

func (this *CoordinateBoard) String() string {
	this.writeFiles()
	this.writeRanks()
	this.writeFiles()
	return this.buffer.String()
}

func (this *CoordinateBoard) writeFiles() {
	this.buffer.WriteString("  ")
	this.buffer.WriteString(fullWidthFiles)
	this.buffer.WriteString("\n")
}
func (this *CoordinateBoard) writeRanks() {
	for x, line := range this.board{
		this.writeRank(x, line)
	}

}
func (this *CoordinateBoard) writeRank(index int, line string) {
	rank := strconv.Itoa(8 - index)
	this.buffer.WriteString(rank)
	this.buffer.WriteString(" ")
	this.buffer.WriteString(line)
	this.buffer.WriteString(" ")
	this.buffer.WriteString(rank)
	this.buffer.WriteString("\n")
}
