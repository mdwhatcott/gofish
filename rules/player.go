package rules

type Player int

const (
	Neither Player = iota
	White
	Black
)

func (this Player) Other() Player {
	if this == White {
		return Black
	}
	return White
}

func (this Player) String() string {
	if this == White {
		return "White"
	} else if this == Black {
		return "Black"
	}
	return ""
}
