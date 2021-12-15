package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day6/input-example.txt")
	absPath, _ := filepath.Abs("day6/input.txt")
	lines := common.ReadLines(absPath)
	fmt.Println(lines)
	lanternFish := strings.Split(lines[0], ",")
	lanternFishInternals, _ := common.SliceAtoi(lanternFish)
	for i := 1; i <= 80; i++ {
		fmt.Println("Day: ", i)
		lanternFishInternals = simulate(lanternFishInternals)
		// fmt.Println("Fishes:", lanternFishInternals)
	}

	// fmt.Println(lanternFishInternals)
	fmt.Println("Answer:", len(lanternFishInternals))
}

func simulate(internalTimers []int) []int {
	numOfNewLanternFishToAdd := 0
	for i := 0; i < len(internalTimers); i++ {
		timer := internalTimers[i]

		if timer == 0 {
			internalTimers[i] = 6
			numOfNewLanternFishToAdd++
		} else {
			internalTimers[i] = timer - 1
		}

	}
	for i := 0; i < numOfNewLanternFishToAdd; i++ {
		internalTimers = append(internalTimers, 8)
	}
	return internalTimers
}
