package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

type point struct {
	X int
	Y int
	Z int
}

type cuboid struct {
	low  point
	high point
	size int
}

type instruction struct {
	cube cuboid
	isOn bool
}

const (
	rangeStart = -50
	rangeEnd   = 50
)

func main() {
	//absPath, _ := filepath.Abs("day22/input-example-2.txt")
	absPath, _ := filepath.Abs("day22/input.txt")
	lines := common.ReadLines(absPath)

	instructions, cuboids := parseGrid(lines)

	total := 0
	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]
		if inst.isOn {
			total += determinePoints(inst.cube, cuboids[i+1:])
		}
	}

	fmt.Println("Answer:", total)
}

func determinePoints(c cuboid, blockers []cuboid) int {
	total := c.size

	for i := 0; i < len(blockers); i++ {
		if coverage, ok := c.isIntersected(blockers[i]); ok {
			total -= determinePoints(coverage, blockers[i+1:])
		}
	}

	return total
}

// Calculate intersection and return boolean if any cubes intersect along with intersection cube
func (c cuboid) isIntersected(o cuboid) (cuboid, bool) {
	sectionA := c.high.Sub(o.low)
	sectionB := o.high.Sub(c.low)

	if common.Positive(sectionA.X, sectionA.Y, sectionA.Z) && common.Positive(sectionB.X, sectionB.Y, sectionB.Z) {
		// intersection
		lowCorner := point{common.Max(c.low.X, o.low.X), common.Max(c.low.Y, o.low.Y), common.Max(c.low.Z, o.low.Z)}
		highCorner := point{common.Min(c.high.X, o.high.X), common.Min(c.high.Y, o.high.Y), common.Min(c.high.Z, o.high.Z)}

		newCuboid := cuboid{low: lowCorner, high: highCorner}
		newCuboid.calculateSize()
		return newCuboid, true
	}

	return cuboid{}, false
}

func (p point) Add(p2 point) point {
	return point{p.X + p2.X, p.Y + p2.Y, p.Z + p2.Z}
}

func (p point) Sub(p2 point) point {
	return point{p.X - p2.X, p.Y - p2.Y, p.Z - p2.Z}
}

func (c *cuboid) calculateSize() {
	x := c.high.X - c.low.X + 1
	y := c.high.Y - c.low.Y + 1
	z := c.high.Z - c.low.Z + 1

	c.size = x * y * z
}

func parseGrid(lines []string) (instructions []instruction, cubes []cuboid) {
	instructions = []instruction{}
	cubes = []cuboid{}
	for _, line := range lines {
		var low point
		var high point
		var command string
		_, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &low.X, &high.X, &low.Y, &high.Y, &low.Z, &high.Z)
		if err != nil {
			log.Println("Err:", err)
		}

		cube := cuboid{
			low:  low,
			high: high,
		}
		cube.calculateSize()
		instruct := instruction{
			cube: cube,
			isOn: command == "on",
		}
		instructions = append(instructions, instruct)
		cubes = append(cubes, cube)
	}
	return
}
