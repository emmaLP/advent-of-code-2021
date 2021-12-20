package main

import (
	"fmt"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	//absPath, _ := filepath.Abs("day20/input-example.txt")
	absPath, _ := filepath.Abs("day20/input.txt")
	lines := common.ReadLines(absPath)
	imageEnhancement := lines[0]
	inputImageLines := lines[2:] //caters for blank line between image enhancement and input

	imageMap := populateMap(inputImageLines)

	answer := 0
	loopStartSize := -99
	loopSize := len(inputImageLines) + 100
	for y := loopStartSize; y <= loopSize; y++ {
		for x := loopStartSize; x <= len(inputImageLines[0])+100; x++ {
			if containsPixel(x, y, 50, imageEnhancement, imageMap) {
				answer++
			}
		}
	}

	fmt.Println("Answer:", answer)
}

func populateMap(input []string) map[[3]int]bool {
	imageMap := map[[3]int]bool{}
	for y, line := range input {
		for x, b := range []byte(line) {
			imageMap[[3]int{x, y, 0}] = b == '#'
		}
	}

	return imageMap
}

func containsPixel(x, y, stepCount int, imageEnhancement string, imageMap map[[3]int]bool) bool {
	initPoint := [3]int{x, y, stepCount}
	if v, ok := imageMap[initPoint]; ok {
		return v
	} else if stepCount == 0 {
		return false
	}

	idx := 0
	bit := 8
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			x2, y2 := x+dx, y+dy
			if containsPixel(x2, y2, stepCount-1, imageEnhancement, imageMap) {
				idx += 1 << bit
			}
			bit--
		}
	}

	isPixel := imageEnhancement[idx] == '#'
	imageMap[initPoint] = isPixel
	return isPixel
}
