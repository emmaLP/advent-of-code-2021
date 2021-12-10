package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
	"strconv"
)

func main() {
	//absPath, _ := filepath.Abs("day3/part1/input-example.txt")
	absPath, _ := filepath.Abs("day3/input.txt")
	lines := common.ReadLines(absPath)
	oxygen := filter(lines, 0, true)
	scrubber := filter(lines, 0, false)
	fmt.Println(convertToInt(oxygen) * convertToInt(scrubber))
}

func mostAndLeastCommon(lines []string, i int) (uint8, uint8) {
	count0, count1 := 0, 0
	for _, line := range lines {
		if line[i] == '0' {
			count0++
		} else {
			count1++
		}
	}
	if count0 > count1 {
		return '0', '1'
	}
	return '1', '0'
}

func filter(lines []string, i int, mostCommon bool) string {
	if len(lines) == 1 {
		return lines[0]
	}
	most, least := mostAndLeastCommon(lines, i)
	comparator := least
	if mostCommon {
		comparator = most
	}
	filtered := make([]string, 0)
	for _, l := range lines {
		if l[i] == comparator {
			filtered = append(filtered, l)
		}
	}
	return filter(filtered, i+1, mostCommon)
}

func convertToInt(s string) (r int) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
