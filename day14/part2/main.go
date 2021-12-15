package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day14/input-example.txt")
	absPath, _ := filepath.Abs("day14/input.txt")
	lines := common.ReadLines(absPath)

	template := lines[0]
	pairIns := getPairInsertions(lines)

	elements := determineElements(pairIns, template, 40)
	fmt.Println("Elements: ", elements)

	max := max(elements)
	min := min(elements, max)
	fmt.Println("Answer:", max-min)
}

func getPairInsertions(lines []string) map[string]string {
	pairInsertions := map[string]string{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		pairs := strings.Split(line, " -> ")
		pairInsertions[pairs[0]] = pairs[1]
	}
	return pairInsertions
}

func determineElements(pairIns map[string]string, template string, steps int) map[string]int {
	elementCount := map[string]int{}
	pairs := map[string]int{}
	// Determine initial pairs from template
	for i := 0; i < len(template)-1; i++ {
		first := string(template[i])
		second := string(template[i+1])
		pairs[first+second] += 1
	}
	// Populate element count with initial template letters
	for i := 0; i < len(template); i++ {
		elementCount[string(template[i])] += 1
	}

	for i := 0; i < steps; i++ {
		fmt.Println("Step:", i+1)
		newPairs := map[string]int{}
		for key, value := range pairs {
			insKey, ok := pairIns[key]

			if ok {
				n1 := string(key[0]) + insKey
				n2 := insKey + string(key[1])
				newPairs[n1] += value
				newPairs[n2] += value
				elementCount[string(insKey[0])] += value
			} else {
				newPairs[key] += value
			}
		}
		pairs = newPairs
	}
	return elementCount
}

func max(numbers map[string]int) (maxNumber int) {
	for _, value := range numbers {
		if value > maxNumber {
			maxNumber = value
		}
	}
	return maxNumber
}

func min(numbers map[string]int, initialNum int) (minNumber int) {
	minNumber = initialNum
	for _, value := range numbers {
		if value < minNumber {
			minNumber = value
		}
	}
	return minNumber
}
