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

	seen := map[[2]int]bool{}
	for _, p := range coords {
		firstInstruction := instructions[0][1]
		if p[0] > firstInstruction {
			p[0] = 2*firstInstruction - p[0]
		}
		seen[p] = true
	}
	fmt.Println("Answer:", len(seen))
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
