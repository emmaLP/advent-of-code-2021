package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

var pairs = map[byte]byte{
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

var scores = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	// absPath, _ := filepath.Abs("day10/input-example.txt")
	absPath, _ := filepath.Abs("day10/input.txt")
	lines := common.ReadLines(absPath)

	allscores := []int{}
	for _, line := range lines {

		score, wrongChar := determineClosingSequenceScores(line)
		if wrongChar != 0 {
			continue
		}
		allscores = append(allscores, score)
	}
	sort.Ints(allscores)
	fmt.Println(allscores[len(allscores)/2])
}

func determineClosingSequenceScores(str string) (int, byte) {
	var legalChars []byte
	var wrongChar byte
	for _, c := range []byte(str) {
		c2, ok := pairs[c]
		if ok {
			legalChars = append(legalChars, c2)
		} else if len(legalChars) > 0 && legalChars[len(legalChars)-1] == c {
			legalChars = legalChars[:len(legalChars)-1]
		} else {
			wrongChar = c
			break
		}
	}

	if wrongChar != 0 {
		return 0, wrongChar
	}
	score := 0
	for i := len(legalChars) - 1; i >= 0; i-- {
		score = score*5 + scores[legalChars[i]]
	}
	return score, wrongChar
}
