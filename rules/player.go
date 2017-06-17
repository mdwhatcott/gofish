package rules

type player int

const (
	Neither player = iota
	White
	Black
)

func (this player) Other() player {
	if this == White {
		return Black
	}
	return White
}

func (this player) String() string {
	if this == White {
		return "White"
	} else if this == Black {
		return "Black"
	}
	return ""
}
