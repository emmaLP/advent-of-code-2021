package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day7/input-example.txt")
	absPath, _ := filepath.Abs("day7/input.txt")
	lines := common.ReadLines(absPath)
	fmt.Println(lines)
	pos := strings.Split(lines[0], ",")
	crabPositions, _ := common.SliceAtoi(pos)
	moves := 0
	for _, position := range crabPositions {
		sumOfMoves := getMoves(crabPositions, position)
		if moves == 0 {
			moves = sumOfMoves
		} else {
			moves = common.Min(moves, sumOfMoves)
		}
	}

	fmt.Println("Answer: ", moves)
}

func getMoves(positions []int, guess int) int {
	sumOfMoves := 0
	for _, position := range positions {
		sumOfMoves += common.Max(guess, position) - common.Min(guess, position)
	}
	return sumOfMoves
}
