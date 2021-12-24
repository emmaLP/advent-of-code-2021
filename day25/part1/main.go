package main

import (
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day25/input-example.txt")
	// absPath, _ := filepath.Abs("day25/input.txt")
	lines := common.ReadLines(absPath)
}
