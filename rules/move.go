package rules

type move struct {
	Piece piece
	From  square
	To    square

	Promotion piece

	Capture   piece
	CaptureOn square

	Castles bool // When true, the Piece will represent the Rook involved.
}
