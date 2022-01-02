package main

import (
	"fmt"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

var down = common.Point{X: 0, Y: 1}
var right = common.Point{X: 1, Y: 0}

func main() {
	// absPath, _ := filepath.Abs("day25/input-example.txt")
	absPath, _ := filepath.Abs("day25/input.txt")
	lines := common.ReadLines(absPath)

	height, width, grid := parseInput(lines)

	step := 0
	var newGrid map[common.Point]rune

	for {
		motion := false

		// Iterate
		newGrid = make(map[common.Point]rune)

		for p, c := range grid {
			if c == '>' {
				destination := p.Add(right)
				if destination.X >= width {
					destination.X = 0
				}
				if grid[destination] == 0 {
					newGrid[destination] = c
					motion = true
				} else {
					newGrid[p] = c
				}
			} else {
				newGrid[p] = c
			}
		}
		grid = newGrid

		newGrid = make(map[common.Point]rune)

		for p, c := range grid {
			if c == 'v' {
				destination := p.Add(down)
				if destination.Y >= height {
					destination.Y = 0
				}
				if grid[destination] == 0 {
					newGrid[destination] = c
					motion = true
				} else {
					newGrid[p] = c
				}
			} else {
				newGrid[p] = c
			}
		}
		grid = newGrid
		step++

		if motion == false {
			break
		}
	}
	fmt.Println("Answer: ", step)
}

func parseInput(lines []string) (height, width int, res map[common.Point]rune) {
	res = make(map[common.Point]rune)
	height = len(lines)
	width = len(lines[0])

	for j, ln := range lines {
		for i, c := range ln {
			if c != '.' {
				res[common.Point{X: i, Y: j}] = c
			}
		}
	}

	return
}
