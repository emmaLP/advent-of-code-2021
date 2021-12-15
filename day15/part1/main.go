package main

import (
	"container/heap"
	"fmt"
	"math"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

type pos struct {
	i, j int
	val  uint16
}
type minHeap []pos

func main() {
	//absPath, _ := filepath.Abs("day15/input-example.txt")
	absPath, _ := filepath.Abs("day15/input.txt")
	lines := common.ReadLines(absPath)

	rowLength := len(lines)
	colLength := len(lines[0])

	var grid [100][100]byte
	var dist [100][100]uint16
	for i, row := range lines {
		for j := range row {
			grid[i][j] = lines[i][j] - '0'
			dist[i][j] = math.MaxUint16
		}
	}

	fmt.Println("Answer:", determine(rowLength, colLength, grid, dist))
}

func determine(rowLength, colLength int, grid [100][100]byte, dist [100][100]uint16) int {
	ok := func(i, j int) bool {
		return i >= 0 && i < rowLength && j >= 0 && j < colLength
	} // Keep a min heap of distances
	h := make(minHeap, 1, 100)
	h[0] = pos{0, 0, 0}

	// While there are entries in the min heap (always true for this)
	for {
		x := heap.Pop(&h).(pos)
		if x.i == rowLength-1 && x.j == colLength-1 {
			return int(x.val)
		}
		for _, nei := range [][2]int{
			{x.i + 1, x.j}, {x.i - 1, x.j}, {x.i, x.j - 1}, {x.i, x.j + 1},
		} {
			ii, jj := nei[0], nei[1]
			if !ok(ii, jj) {
				continue
			}
			risk := x.val + uint16(grid[ii][jj])
			if risk >= dist[ii][jj] {
				continue
			}
			dist[ii][jj] = risk
			heap.Push(&h, pos{ii, jj, risk})
		}
	}
}

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pos))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
