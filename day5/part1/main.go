package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	//absPath, _ := filepath.Abs("day5/input-example.txt")
	absPath, _ := filepath.Abs("day5/input.txt")
	lines := common.ReadLines(absPath)
	fmt.Println(lines)

	all := map[point]bool{}
	ventsIntersecting := map[point]bool{}
	for _, line := range lines {
		fmt.Println("Line:", line)
		coords := strings.Split(line, " -> ")
		fmt.Println(coords)
		plotCoordsOnDiagram(coords, all, ventsIntersecting)
	}
	fmt.Println("Answer: ", len(ventsIntersecting))
}

func plotCoordsOnDiagram(coords []string, all, ventsIntersectings map[point]bool) {
	var x1, y1 int
	var x2, y2 int
	coord1 := strings.Split(coords[0], ",")
	coord2 := strings.Split(coords[1], ",")
	x1, _ = strconv.Atoi(coord1[0])
	y1, _ = strconv.Atoi(coord1[1])
	x2, _ = strconv.Atoi(coord2[0])
	y2, _ = strconv.Atoi(coord2[1])
	if x1 == x2 {
		for y := min(y1, y2); y <= max(y1, y2); y++ {
			p := point{
				x: x1,
				y: y,
			}
			if _, ok := all[p]; ok {
				ventsIntersectings[p] = true
			} else {
				all[p] = true
			}
		}
	} else if y1 == y2 {
		for x := min(x1, x2); x <= max(x1, x2); x++ {
			p := point{
				x: x,
				y: y1,
			}
			if _, ok := all[p]; ok {
				ventsIntersectings[p] = true
			} else {
				all[p] = true
			}
		}
	}

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
