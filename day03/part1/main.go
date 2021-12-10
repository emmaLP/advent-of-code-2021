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
	gamma, epsilon := "", ""
	for i := 0; i < len(lines[0]); i++ {
		most, least := mostAndLeastCommon(lines, i)
		gamma += string(most)
		epsilon += string(least)
		fmt.Println(gamma, epsilon)
	}
	fmt.Println(convertToBinary(gamma) * convertToBinary(epsilon))
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

func convertToBinary(s string) (r int) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
