package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day17/input-example.txt")
	absPath, _ := filepath.Abs("day17/input.txt")
	lines := common.ReadLines(absPath)
	fmt.Println(lines)
	parts := strings.Split(lines[0][len("target area:"):], ",")
	x1, x2 := parseRange(parts[0])
	y2, y1 := parseRange(parts[1])

	hits := 0.

	for xv := 0.; xv <= x2; xv++ {
		for yv := y2; yv < 999; yv++ {
			probe := launchProbe(Probe{xv: xv, yv: yv}, x1, x2, y1, y2)
			if probe.isHit(x1, x2, y1, y2) {
				hits++
			}
		}
	}
	fmt.Println("Answer:", hits)
}

type Probe struct {
	x, y, xv, yv, maxHeight float64
}

func launchProbe(probe Probe, x1, x2, y1, y2 float64) Probe {
	for !probe.isMiss(x2, y2) && !probe.isHit(x1, x2, y1, y2) {
		probe = probe.step()
	}
	return probe
}

func (p Probe) step() Probe {
	return Probe{p.x + p.xv, p.y + p.yv, math.Max(0, p.xv-1), p.yv - 1, math.Max(p.maxHeight, p.y+p.yv)}
}

func (p Probe) isHit(x1, x2, y1, y2 float64) bool {
	return p.x >= x1 && p.x <= x2 && p.y <= y1 && p.y >= y2
}

func (p Probe) isMiss(x2, y2 float64) bool {
	return p.y < y2 || p.x > x2
}

func parseRange(part string) (float64, float64) {
	parts := strings.Split(part[3:], "..")
	start, _ := strconv.ParseFloat(parts[0], 32)
	end, _ := strconv.ParseFloat(parts[1], 32)
	return start, end
}
