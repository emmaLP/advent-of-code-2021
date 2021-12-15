package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day8/input-example.txt")
	absPath, _ := filepath.Abs("day8/input.txt")
	lines := common.ReadLines(absPath)
	var outputs []string
	for _, line := range lines {
		inputOutputSplit := strings.Split(line, " | ")
		outputSplits := strings.Split(inputOutputSplit[1], " ")
		outputs = append(outputs, outputSplits...)
	}
	fmt.Println("Outputs: ", outputs, len(outputs))
	result := 0
	lenArry := []int{2, 3, 4, 7}
	for _, output := range outputs {
		if common.Contains(lenArry, len(output)) {
			result += 1
		}
	}
	fmt.Println("Answer: ", result)
}
