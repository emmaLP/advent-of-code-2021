package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
)

func main() {
	absPath, _ := filepath.Abs("day10/input-example.txt")
	//absPath, _ := filepath.Abs("day10/input.txt")
	lines := common.ReadLines(absPath)
	fmt.Println(lines)

}
