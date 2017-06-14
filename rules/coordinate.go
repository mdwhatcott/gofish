package rules

import (
	"fmt"
	"strconv"
)

const rankCount = 8

func squareIndex(algebraic string) int {
	fileLetter := algebraic[0]
	rankNumber, _ := strconv.Atoi(algebraic[1:])
	if fileLetter < 'a' || fileLetter > 'h' || rankNumber < 1 || rankNumber > 8 {
		panic(fmt.Sprintf("Invalid algebraic coordinate: %s", algebraic))
	}
	file := int(fileLetter - 'a')
	startOfRank := (rankNumber -1) * rankCount
	return startOfRank + file
}
