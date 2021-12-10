package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
)

func main() {
	//absPath, _ := filepath.Abs("day9/input-example.txt")
	absPath, _ := filepath.Abs("day9/input.txt")
	lines := common.ReadLines(absPath)
	grid := common.ParsIntGrid(lines, "")
	fmt.Println("Answer:", determineLowPointsSum(grid))
}

func determineLowPointsSum(grid [][]int) int {
	lowPoints := 0
	for i, row := range grid {
		for j, col := range row {
			var top, bottom, right, left bool
			top = i == 0 || col < grid[i-1][j]
			bottom = i == len(grid)-1 || col < grid[i+1][j]
			right = j == 0 || col < grid[i][j-1]
			left = j == len(row)-1 || col < grid[i][j+1]

			if top && bottom && left && right {
				lowPoints += col + 1
			}
		}
	}
	return lowPoints
}
