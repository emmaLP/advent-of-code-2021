package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
)

var pairs = map[byte]byte{
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

var scores = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	//absPath, _ := filepath.Abs("day10/input-example.txt")
	absPath, _ := filepath.Abs("day10/input.txt")
	lines := common.ReadLines(absPath)

	res := 0
	for _, line := range lines {

		if wrongChar := findIllegalCharacter(line); wrongChar != 0 {
			res += scores[wrongChar]
		}
	}
	fmt.Println("Answer: ", res)
}

func findIllegalCharacter(str string) byte {
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
	return wrongChar
}
