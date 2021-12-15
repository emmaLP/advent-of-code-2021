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
	lanternFish, _ := common.SliceAtoi(strings.Split(lines[0], ","))
	lanternFishInternals := parseToMap(lanternFish)
	fmt.Println(lanternFishInternals)
	for i := 1; i <= 256; i++ {
		fmt.Println("Day: ", i)
		lanternFishInternals = simulate(lanternFishInternals)
		fmt.Println("Fishes:", lanternFishInternals)
	}

	fishCount := 0
	for _, value := range lanternFishInternals {
		fishCount += value
	}

	fmt.Println("Answer:", fishCount)
}

func parseToMap(sa []int) map[int]int {
	mapOutput := map[int]int{}
	for _, a := range sa {

		value, okay := mapOutput[a]
		if okay {
			mapOutput[a] = value + 1
		} else {
			mapOutput[a] = 1
		}
	}

	return mapOutput
}

func simulate(internalTimers map[int]int) map[int]int {
	newAgeMap := map[int]int{}
	for key, value := range internalTimers {
		// timer := internalTimers[i]

		if key == 0 {
			if value6, ok := newAgeMap[6]; ok {
				newAgeMap[6] = value6 + value
			} else {
				newAgeMap[6] = value
			}
			if value8, ok := newAgeMap[8]; ok {
				newAgeMap[8] = value8 + value
			} else {
				newAgeMap[8] = value
			}

		} else {
			if tmpValue, ok := newAgeMap[key-1]; ok {
				newAgeMap[key-1] = tmpValue + value
			} else {
				newAgeMap[key-1] = value
			}
		}
	}
	//for i := 0; i < numOfNewLanternFishToAdd; i++ {
	//	internalTimers = append(internalTimers, 8)
	//}
	return newAgeMap
}
