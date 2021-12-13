package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
)

func main() {
	// absPath, _ := filepath.Abs("day11/input-example.txt")
	absPath, _ := filepath.Abs("day11/input.txt")
	lines := common.ReadLines(absPath)
	grid := common.ParsIntGrid(lines, "")

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for step := 0; step < 500; step++ {
		fcount := 0
		stepIncrement(grid)
		
		var checkFlash func(y, x int, add int)
		checkFlash = func(y, x int, add int) {
			if y < 0 || x < 0 || y >= len(grid) || x >= len(grid[y]) {
				return
			}
			if grid[y][x] == -1 {
				return
			}
			grid[y][x] += add
			if grid[y][x] <= 9 {
				return
			}
			fcount++
			grid[y][x] = -1

			for _, d := range directions {
				checkFlash(y+d[0], x+d[1], 1)
			}
		}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				checkFlash(i, j, 0)
			}
		}

		if fcount == len(grid)*len(grid[0]) {
			fmt.Println("Ans:", step + 1)
			return
		}
	}
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
