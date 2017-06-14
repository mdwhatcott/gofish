package rules

type Piece string

const (
	Void        Piece = ""
	WhitePawn   Piece = "P"
	WhiteKnight Piece = "N"
	WhiteBishop Piece = "B"
	WhiteRook   Piece = "R"
	WhiteQueen  Piece = "Q"
	WhiteKing   Piece = "K"
	BlackPawn   Piece = "p"
	BlackKnight Piece = "n"
	BlackBishop Piece = "b"
	BlackRook   Piece = "r"
	BlackQueen  Piece = "q"
	BlackKing   Piece = "k"
)

func startingPosition() map[int]Piece {
	return map[int]Piece{
		0*8 + 0: WhiteRook,
		0*8 + 1: WhiteKnight,
		0*8 + 2: WhiteBishop,
		0*8 + 3: WhiteQueen,
		0*8 + 4: WhiteKing,
		0*8 + 5: WhiteBishop,
		0*8 + 6: WhiteKnight,
		0*8 + 7: WhiteRook,

		1*8 + 0: WhitePawn,
		1*8 + 1: WhitePawn,
		1*8 + 2: WhitePawn,
		1*8 + 3: WhitePawn,
		1*8 + 4: WhitePawn,
		1*8 + 5: WhitePawn,
		1*8 + 6: WhitePawn,
		1*8 + 7: WhitePawn,

		2*8 + 0: Void,
		2*8 + 1: Void,
		2*8 + 2: Void,
		2*8 + 3: Void,
		2*8 + 4: Void,
		2*8 + 5: Void,
		2*8 + 6: Void,
		2*8 + 7: Void,

		3*8 + 0: Void,
		3*8 + 1: Void,
		3*8 + 2: Void,
		3*8 + 3: Void,
		3*8 + 4: Void,
		3*8 + 5: Void,
		3*8 + 6: Void,
		3*8 + 7: Void,

		4*8 + 0: Void,
		4*8 + 1: Void,
		4*8 + 2: Void,
		4*8 + 3: Void,
		4*8 + 4: Void,
		4*8 + 5: Void,
		4*8 + 6: Void,
		4*8 + 7: Void,

		5*8 + 0: Void,
		5*8 + 1: Void,
		5*8 + 2: Void,
		5*8 + 3: Void,
		5*8 + 4: Void,
		5*8 + 5: Void,
		5*8 + 6: Void,
		5*8 + 7: Void,

		6*8 + 0: BlackPawn,
		6*8 + 1: BlackPawn,
		6*8 + 2: BlackPawn,
		6*8 + 3: BlackPawn,
		6*8 + 4: BlackPawn,
		6*8 + 5: BlackPawn,
		6*8 + 6: BlackPawn,
		6*8 + 7: BlackPawn,

		7*8 + 0: BlackRook,
		7*8 + 1: BlackKnight,
		7*8 + 2: BlackBishop,
		7*8 + 3: BlackQueen,
		7*8 + 4: BlackKing,
		7*8 + 5: BlackBishop,
		7*8 + 6: BlackKnight,
		7*8 + 7: BlackRook,
	}
}
