package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day13/input-example.txt")
	absPath, _ := filepath.Abs("day13/input.txt")
	lines := common.ReadLines(absPath)
	coords, instructions := populateCoordsAndInstructions(lines)

	for _, instruction := range instructions {
		for i, coord := range coords {
			if instruction[0] == 0 {
				if coord[0] > instruction[1] {
					coord[0] = 2*instruction[1] - coord[0]
				}
			} else {
				if coord[1] > instruction[1] {
					coord[1] = 2*instruction[1] - coord[1]
				}
			}
			coords[i] = coord
		}
	}

	maxX := 0
	maxY := 0
	seen := map[[2]int]bool{}

	for _, coord := range coords {
		seen[coord] = true
		if coord[0] > maxX {
			maxX = coord[0]
		}
		if coord[1] > maxY {
			maxY = coord[1]
		}
	}
	printCode(maxX, maxY, seen)
}

func printCode(maxX, maxY int, populatedValues map[[2]int]bool) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if populatedValues[[2]int{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func populateCoordsAndInstructions(lines []string) (coords, instructions [][2]int) {
	populateInstructions := false
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			populateInstructions = true
			continue
		}

		if !populateInstructions {
			ps := strings.Split(line, ",")
			var p [2]int
			p[0], _ = strconv.Atoi(ps[0])
			p[1], _ = strconv.Atoi(ps[1])
			coords = append(coords, p)
		} else {
			line = line[len("fold along "):]
			instructionSplit := strings.Split(line, "=")

			axis := 0
			if instructionSplit[0] == "y" {
				axis = 1
			}

			axisIndex, _ := strconv.Atoi(instructionSplit[1])
			instructions = append(instructions, [2]int{axis, axisIndex})
		}
	}
	return coords, instructions
}
