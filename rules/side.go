package rules

type Player int

const (
	White Player = 0
	Black Player = 1
)

func (this Player) Alternate() Player {
	if this == White {
		return Black
	}
	return White
}

func (this Player) String() string {
	if this == White {
		return "White"
	}
	return "Black"
}
