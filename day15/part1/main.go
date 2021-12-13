package main

import (
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
)

func main() {
	absPath, _ := filepath.Abs("day15/input-example.txt")
	// absPath, _ := filepath.Abs("day15/input.txt")
	lines := common.ReadLines(absPath)
}
