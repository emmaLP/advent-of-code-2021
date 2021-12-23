package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

type Point struct {
	X int
	Y int
	Z int
}

const (
	rangeStart = -50
	rangeEnd   = 50
)

func main() {
	//absPath, _ := filepath.Abs("day22/input-example.txt")
	absPath, _ := filepath.Abs("day22/input.txt")
	lines := common.ReadLines(absPath)

	grid := parseGrid(lines)

	total := 0
	for _, value := range grid {
		if value {
			total++
		}
	}

	fmt.Println("Answer:", total)
}

func parseGrid(lines []string) (grid map[Point]bool) {
	grid = map[Point]bool{}
	for _, line := range lines {
		var low Point
		var high Point
		var command string
		_, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &low.X, &high.X, &low.Y, &high.Y, &low.Z, &high.Z)
		if err != nil {
			log.Println("Err:", err)
		}

		if low.X >= rangeStart && high.X <= rangeEnd && low.Y >= rangeStart && high.Y <= rangeEnd && low.Z >= rangeStart && high.Z <= rangeEnd {
			for k := low.Z; k <= high.Z; k++ {
				for j := low.Y; j <= high.Y; j++ {
					for i := low.X; i <= high.X; i++ {
						if command == "on" {
							grid[Point{i, j, k}] = true
						} else {
							grid[Point{i, j, k}] = false
						}
					}
				}
			}
		}
	}
	return
}
