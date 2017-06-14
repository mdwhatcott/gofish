package rules

import (
	"strings"
	"unicode"
)

type Game struct {
	squares  map[int]Piece
	player   Player
	castling map[rune]struct{}
}

func NewGame() *Game {
	game := &Game{
		squares:  make(map[int]Piece, 64),
		castling: make(map[rune]struct{}),
	}
	game.Reset()
	return game
}

func (this *Game) Reset() {
	this.LoadFEN(startingPositionFEN)
}

func (this *Game) LoadFEN(fen string) error {
	for key := range this.castling {
		delete(this.castling, key)
	}
	for key := range this.squares {
		delete(this.squares, key)
	}

	fields := strings.Split(fen, " ")
	ranks := strings.Split(fields[0], "/")
	squares := make(map[int]Piece, 64)
	for r, rank := range ranks {
		square := 64 - ((r + 1) * 8)
		for _, c := range rank {
			if unicode.IsDigit(c) {
				square += int(c - '0')
			} else {
				squares[square] = Piece(string(c))
				square++
			}
		}
	}
	this.squares = squares
	if fields[1] == "w" {
		this.player = White
	} else {
		this.player = Black
	}

	if fields[2] != "-" {
		for _, c := range fields[2] {
			this.castling[c] = struct{}{}
		}
	}

	// TODO: fields[3], fields[4], fields[5]
	return nil
}

func (this *Game) PlayerToMove() Player {
	return this.player
}

func (this *Game) IsOver() bool {
	return false
}

func (this *Game) FullMoveCount() int {
	return 1
}

func (this *Game) HalfMoveCount() int {
	return 0
}

func (this *Game) CanCastleKingside(player Player) bool {
	if player == White {
		_, valid := this.castling['K']
		return valid
	} else {
		_, valid := this.castling['k']
		return valid
	}
}

func (this *Game) CanCastleQueenside(player Player) bool {
	if player == White {
		_, valid := this.castling['Q']
		return valid
	} else {
		_, valid := this.castling['q']
		return valid
	}
}

func (this *Game) FEN() string {
	return PrepareFEN(this.squares, this).String()
}

func (this *Game) Move(move *Move) error {
	from := squareIndex(move.From)
	to := squareIndex(move.To)
	this.squares[to], this.squares[from] = this.squares[from], Void
	this.player = this.player.Alternate()
	return nil
}

func (this *Game) CalculateAvailableMoves() (moves []Move) {
	moves = []Move{}
	return moves
}
