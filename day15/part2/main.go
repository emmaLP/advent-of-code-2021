package main

import (
	"container/heap"
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"math"
	"path/filepath"
)

type path struct {
	x, y int
	risk int
}
type minHeap []path

func main() {
	//absPath, _ := filepath.Abs("day15/input-example.txt")
	absPath, _ := filepath.Abs("day15/input.txt")
	lines := common.ReadLines(absPath)

	grid := common.ParsIntGrid(lines, "")
	dist := make([][]int, len(grid)*5)

	for i := 0; i < len(dist); i++ {
		col := make([]int, len(grid[0])*5)
		for j := 0; j < len(grid[0])*5; j++ {
			col[j] = math.MaxInt
		}
		dist[i] = col
	}

	fmt.Println("Answer:", determine(grid, dist))
}

func determine(grid [][]int, dist [][]int) int {
	seen := make([][]bool, len(grid)*5)
	for i := 0; i < len(seen); i++ {
		colSeen := make([]bool, len(grid[0])*5)
		for j := 0; j < len(grid[0])*5; j++ {
			colSeen[j] = false
		}
		seen[i] = colSeen
	}
	mm := len(dist)
	nn := len(dist[0])
	insideGrid := func(x, y int) bool {
		return x >= 0 && x < mm && y >= 0 && y < nn
	}

	// Keep a min heap of distances
	h := make(minHeap, 1, 100)
	h[0] = path{0, 0, 0}
	for {
		current := heap.Pop(&h).(path)
		fmt.Println("X:", current)
		for _, nei := range [][2]int{
			{current.x + 1, current.y}, {current.x - 1, current.y}, {current.x, current.y - 1}, {current.x, current.y + 1},
		} {
			rows := nei[0]
			cols := nei[1]

			if insideGrid(rows, cols) {
				gridNum := grid[rows%len(grid)][cols%len(grid)] + (rows / len(grid)) + (cols / len(grid[0]))
				fmt.Println("Grid num:", gridNum)
				if gridNum > 9 {
					gridNum -= 9
				}
				fmt.Println("Mutated Grid num:", gridNum)

				if dist[rows][cols] > dist[current.x][current.y]+gridNum {
					if dist[rows][cols] != math.MaxInt {
						fmt.Println("Supported to remove here?? ", dist[rows][cols])
					}
					risk := current.risk + gridNum
					fmt.Println("Risk", risk)
					dist[rows][cols] = risk

					heap.Push(&h, path{rows, cols, risk})

					if rows == mm-1 && cols == nn-1 {
						return risk
					}
				}
			}
		}
	}
}

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].risk < h[j].risk
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(path))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
