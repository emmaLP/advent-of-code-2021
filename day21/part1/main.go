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
	p1Pos, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	p2Pos, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	fmt.Println("Player starting pos:", p1Pos, p2Pos)

	p1Score, p2Score, diceRollCount := 0, 0, 0
	nextDiceRoll := 1
	for {
		//Player 1
		p1Pos, p1Score, nextDiceRoll, diceRollCount = playTurn(p1Pos, p1Score, nextDiceRoll, diceRollCount)
		fmt.Println("P1 turn output:", p1Pos, p1Score, nextDiceRoll, diceRollCount)
		if p1Score >= 1000 {
			break
		}
		//Player 2
		p2Pos, p2Score, nextDiceRoll, diceRollCount = playTurn(p2Pos, p2Score, nextDiceRoll, diceRollCount)
		fmt.Println("P2 turn output", p2Pos, p2Score, nextDiceRoll, diceRollCount)
		if p2Score >= 1000 {
			break
		}
	}
	fmt.Println("Player scores:", p1Score, p2Score)
	if p1Score > p2Score {
		fmt.Println("Answer:", p2Score*diceRollCount)
	} else {
		fmt.Println("Answer:", p1Score*diceRollCount)
	}
}

func rollDice(currentPlayPos, nextDiceRoll, diceRollCount int) (int, int, int) {
	position := currentPlayPos + nextDiceRoll
	nextDiceRoll++
	if nextDiceRoll > 100 {
		nextDiceRoll = 1
	}
	diceRollCount++
	position = (position-1)%10 + 1
	return position, nextDiceRoll, diceRollCount
}

func playTurn(playerPos, playerScore, nextDiceRoll, diceRollCount int) (int, int, int, int) {
	for i := 0; i < 3; i++ {
		playerPos, nextDiceRoll, diceRollCount = rollDice(playerPos, nextDiceRoll, diceRollCount)
	}
	playerScore += playerPos

	return playerPos, playerScore, nextDiceRoll, diceRollCount
}
