package main

import (
	"fmt"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day9/input-example.txt")
	absPath, _ := filepath.Abs("day9/input.txt")
	lines := common.ReadLines(absPath)
	grid := common.ParsIntGrid(lines, "")
	fmt.Println("Answer:", determineBasins(grid))
}

func determineBasins(grid [][]int) int {
	visited := populateVisitedArray(len(grid), len(grid[0]))
	var areas []int
	for i, row := range grid {
		for j, col := range row {
			visit := visited[i][j]
			if !visit {
				visited[i][j] = true
				if col != 9 {
					area := calculateArea(i, j, visited, len(grid), len(row), grid)
					areas = append(areas, area)
				}
			}
		}
	}

	max1, max2, max3 := 0, 0, 0
	for _, area := range areas {
		if area > max1 {
			max3 = max2
			max2 = max1
			max1 = area
		} else if area > max2 {
			max3 = max2
			max2 = area
		} else if area > max3 {
			max3 = area
		}
	}

	return max1 * max2 * max3
}

func populateVisitedArray(rowSize, colSize int) [][]bool {
	visited := make([][]bool, rowSize)
	for i := 0; i < rowSize; i++ {
		colVisited := make([]bool, colSize)
		for j := 0; j < colSize; j++ {
			colVisited[j] = false
		}
		visited[i] = colVisited
	}
	return visited
}

func calculateArea(i, j int, visited [][]bool, rowSize, colSize int, grid [][]int) int {
	dxArr := []int{1, -1, 0, 0}
	dyArr := []int{0, 0, 1, -1}
	area := 1
	for k := 0; k < 4; k++ {
		dx := dxArr[k]
		dy := dyArr[k]
		if dx+i >= 0 && dx+i < rowSize && dy+j >= 0 && dy+j < colSize {
			ii := dx + i
			jj := dy + j
			if !visited[ii][jj] && grid[ii][jj] < 9 {
				visited[ii][jj] = true
				area += calculateArea(ii, jj, visited, rowSize, colSize, grid)
			}
		}
	}
	return area
}
