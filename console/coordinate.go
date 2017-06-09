package console

import (
	"fmt"
	"strconv"
)

// TODO: this will probably be moved elsewhere as it's useful, but not here (anymore)

func squareIndex(algebraic string) int {
	fileLetter := algebraic[0]
	rankNumber, _ := strconv.Atoi(algebraic[1:])
	if fileLetter < 'a' || fileLetter > 'h' || rankNumber < 1 || rankNumber > 8 {
		panic(fmt.Sprintf("Invalid algebraic coordinate: %s", algebraic))
	}
	file := int(fileLetter - 'a')
	startOfRank := (rankCount - rankNumber) * rankCount
	return startOfRank + file
}
