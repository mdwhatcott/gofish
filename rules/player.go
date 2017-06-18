package rules

type player string

const (
	Neither player = ""
	White   player = "w"
	Black   player = "b"
)

func (this player) Other() player {
	if this == White {
		return Black
	} else if this == Black {
		return White
	} else {
		return Neither
	}
}
