package main

import (
	"fmt"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day11/input-example.txt")
	absPath, _ := filepath.Abs("day11/input.txt")
	lines := common.ReadLines(absPath)
	grid := common.ParsIntGrid(lines, "")

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	fcount := 0
	for step := 0; step < 100; step++ {
		stepIncrement(grid)
		var checkForFlash func(row, col, add int)
		checkForFlash = func(row, col, add int) {
			if (row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row])) || grid[row][col] == -1 {
				return
			}
			grid[row][col] += add
			if grid[row][col] <= 9 {
				return
			}
			fcount++
			grid[row][col] = -1

			for _, d := range directions {
				checkForFlash(row+d[0], col+d[1], 1)
			}
		}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				checkForFlash(i, j, 0)
			}
		}
	}
	fmt.Println(fcount)
}

func stepIncrement(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == -1 {
				grid[i][j] = 0
			}
			grid[i][j]++
		}
	}
}
