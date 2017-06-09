package console

type unit rune

const (
	Void        unit = '＿'
	WhitePawn        = '♙'
	WhiteKnight      = '♘'
	WhiteBishop      = '♗'
	WhiteRook        = '♖'
	WhiteQueen       = '♕'
	WhiteKing        = '♔'
	BlackPawn        = '♟'
	BlackKnight      = '♞'
	BlackBishop      = '♝'
	BlackRook        = '♜'
	BlackQueen       = '♛'
	BlackKing        = '♚'
)

var FENUnits = map[rune]unit {
	'p': BlackPawn,
	'r': BlackRook,
	'n': BlackKnight,
	'b': BlackBishop,
	'q': BlackQueen,
	'k': BlackKing,
	'P': WhitePawn,
	'R': WhiteRook,
	'N': WhiteKnight,
	'B': WhiteBishop,
	'Q': WhiteQueen,
	'K': WhiteKing,
}
