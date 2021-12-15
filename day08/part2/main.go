package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

var chars = map[byte]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
}

func toMask(str string) int {
	m := 0
	for _, c := range []byte(str) {
		m |= 1 << chars[c]
	}
	return m
}

func main() {
	// absPath, _ := filepath.Abs("day8/input-example.txt")
	absPath, _ := filepath.Abs("day8/input.txt")
	lines := common.ReadLines(absPath)

	result := 0
	for _, line := range lines {
		inputOutputSplit := strings.Split(line, " | ")
		output := common.SortCharsInStrings(strings.Fields(inputOutputSplit[1]))
		solved, remainingPatterns := solveUnique(common.SortCharsInStrings(strings.Fields(inputOutputSplit[0])))
		fmt.Println("Easy determineNum:", solved, remainingPatterns)
		nonUnique := solveNonUnique(solved, remainingPatterns)
		fmt.Println("Non unique:", nonUnique)

		resString := ""
		for _, digit := range output {
			for pos, pattern := range solved {
				if digit == pattern {
					resString += strconv.Itoa(pos)
				}
			}
		}
		i, _ := strconv.Atoi(resString)
		result += i

	}

	fmt.Println("Answer: ", result)
}

func determineNum(n int, solved, patterns []string, fn func(pattern string) bool) ([]string, []string) {
	for pos, pattern := range patterns {
		if fn(pattern) {
			solved[n] = pattern
			patterns = append(patterns[:pos], patterns[pos+1:]...)
		}
	}

	return solved, patterns
}

func solveNonUnique(solved, patterns []string) []string {
	// 3 == 5 segments, contains 1
	solved, patterns = determineNum(3, solved, patterns, func(pattern string) bool { return len(pattern) == 5 && common.ContainsChars(pattern, solved[1]) })
	// 6 == 6 segments, does not contain 1
	solved, patterns = determineNum(6, solved, patterns, func(pattern string) bool { return len(pattern) == 6 && !common.ContainsChars(pattern, solved[1]) })

	// separate 2 and 5 based on top left segment
	topLeft := setermineTopLeftSegment(solved)
	solved, patterns = determineNum(2, solved, patterns, func(pattern string) bool { return len(pattern) == 5 && strings.Contains(pattern, topLeft) })
	solved, patterns = determineNum(5, solved, patterns, func(pattern string) bool { return len(pattern) == 5 && !strings.Contains(pattern, topLeft) })

	// separate 0 and 9 based on bottom right segment
	bottomRight := determineBottomRightSegment(solved)
	solved, patterns = determineNum(9, solved, patterns, func(pattern string) bool { return len(pattern) == 6 && !strings.Contains(pattern, bottomRight) })
	solved, _ = determineNum(0, solved, patterns, func(pattern string) bool { return len(pattern) == 6 && strings.Contains(pattern, bottomRight) })

	return solved
}

func solveUnique(patterns []string) ([]string, []string) {
	digits := make([]string, 10)
	var remaining []string

	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			digits[1] = pattern
		case 3:
			digits[7] = pattern
		case 4:
			digits[4] = pattern
		case 7:
			digits[8] = pattern
		default:
			remaining = append(remaining, pattern)
		}
	}
	return digits, remaining
}

func determineBottomRightSegment(solved []string) string {
	bottomRight := solved[6]
	for _, c := range solved[5] {
		bottomRight = strings.Replace(bottomRight, string(c), "", 1)
	}
	return bottomRight
}

func setermineTopLeftSegment(solved []string) string {
	for _, c := range solved[1] {
		if !strings.ContainsRune(solved[6], c) {
			return string(c)
		}
	}
	return ""
}
