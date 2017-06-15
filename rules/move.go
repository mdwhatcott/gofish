package rules

type Move struct {
	Piece Piece
	From  Square
	To    Square

	Captures   Piece
	CapturesOn Square

	Castles bool // When true, the Piece will represent the involved Rook.
}
