package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
	"strings"
)

func main() {
	//absPath, _ := filepath.Abs("day7/input-example.txt")
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
	fmt.Println("Guess:", guess)
	for _, position := range positions {
		move := common.Max(guess, position) - common.Min(guess, position)
		moveCost := (move * (move + 1)) / 2
		sumOfMoves += moveCost
	}
	return sumOfMoves
}
