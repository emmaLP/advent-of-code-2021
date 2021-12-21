package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	//absPath, _ := filepath.Abs("day21/input-example.txt")
	absPath, _ := filepath.Abs("day21/input.txt")
	lines := common.ReadLines(absPath)
	p1Position, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	p2Position, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	fmt.Println("Player starting pos:", p1Position, p2Position)

	maxScore := 21
	p1UniverseWins, p2UniverseWins := 0, 0
	playerTurns := [22][22][11][11][2]int{} //[p1Score][p2Score][p1Pos][p2Pos][playerNumber]
	playerTurns[0][0][p1Position][p2Position][0] = 1

	diceRolls := diceRollVariations()

	for p1Score := 0; p1Score < maxScore; p1Score++ {
		for p2Score := 0; p2Score < maxScore; p2Score++ {
			for p1Pos := 1; p1Pos <= 10; p1Pos++ {
				for p2Pos := 1; p2Pos <= 10; p2Pos++ {
					for roll := 1; roll < len(diceRolls); roll++ {
						if playerTurns[p1Score][p2Score][p1Pos][p2Pos][0] > 0 {
							newP1Pos := (p1Pos+roll-1)%10 + 1
							newScore := p1Score + newP1Pos
							newCount := playerTurns[p1Score][p2Score][p1Pos][p2Pos][0] * diceRolls[roll]
							if newScore > 20 {
								p1UniverseWins += newCount
							} else {
								playerTurns[newScore][p2Score][newP1Pos][p2Pos][1] += newCount
							}
						}
						if playerTurns[p1Score][p2Score][p1Pos][p2Pos][1] > 0 {
							newP2Pos := (p2Pos+roll-1)%10 + 1
							newScore := p2Score + newP2Pos
							newCount := playerTurns[p1Score][p2Score][p1Pos][p2Pos][1] * diceRolls[roll]
							if newScore > 20 {
								p2UniverseWins += newCount
							} else {
								playerTurns[p1Score][newScore][p1Pos][newP2Pos][0] += newCount
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Answer:", common.Max(p1UniverseWins, p2UniverseWins))
}

func diceRollVariations() []int {
	rolls := make([]int, 10)
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				rolls[r1+r2+r3]++
			}
		}
	}
	return rolls
}
