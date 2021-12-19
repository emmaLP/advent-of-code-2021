package main

import (
	"fmt"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

var numRegex, _ = regexp.Compile(`[0-9]+`)
var lineRegex, _ = regexp.Compile(`([[][0-9]*,[0-9]*[]])`)

func main() {
	// absPath, _ := filepath.Abs("day18/input-example.txt")
	absPath, _ := filepath.Abs("day18/input.txt")
	lines := common.ReadLines(absPath)

	answer := 0
	for i := range lines {
		for j := range lines {
			if j == i {
				continue
			}
			r := add(lines[i], lines[j])
			answer = common.Max(r, answer)

		}
	}

	fmt.Println("Answer: ", answer)
}

func add(x, y string) int {
	return magnitude(reduce("[" + x + "," + y + "]"))
}

func magnitude(str string) int {
	for {
		strIndices := lineRegex.FindAllStringIndex(str, -1)
		if len(strIndices) == 0 {
			break
		}
		numstrIndexers := strings.Split(str[strIndices[0][0]+1:strIndices[0][1]-1], ",")
		l, _ := strconv.Atoi(numstrIndexers[0])
		r, _ := strconv.Atoi(numstrIndexers[1])
		str = str[:strIndices[0][0]] + strconv.Itoa(3*l+2*r) + str[strIndices[0][1]:]
	}
	// output, _ := strconv.ParseUint(str, 10, 64)
	output, _ := strconv.Atoi(str)
	return output
}

func reduce(str string) string {
	for {
		for explode(&str) {
		}
		if !split(&str) {
			break
		}
	}
	return str
}

func explode(str *string) (found bool) {
	process := true
	for process {
		process = false
		strIndices := lineRegex.FindAllStringIndex(*str, -1)
		if len(strIndices) == 0 {
			break
		}
		for _, strIndex := range strIndices {
			if isExpectedDepth(*str, strIndex[0]) {
				numstrIndexers := strings.Split((*str)[strIndex[0]+1:strIndex[1]-1], ",")
				l, _ := strconv.Atoi(numstrIndexers[0])
				r, _ := strconv.Atoi(numstrIndexers[1])
				*str = (*str)[:strIndex[0]] + "0" + (*str)[strIndex[1]:]
				sumLeft(str, strIndex[0], l)
				sumRight(str, strIndex[0]+2, r)

				process = true
				break
			}
		}
	}
	return
}

func sumLeft(str *string, i int, nstrIndex int) {
	strIndices := numRegex.FindAllStringIndex((*str)[:i], -1)
	if len(strIndices) == 0 {
		return
	}
	strIndexaseNstrIndex, _ := strconv.Atoi((*str)[strIndices[len(strIndices)-1][0]:strIndices[len(strIndices)-1][1]])
	*str = (*str)[:strIndices[len(strIndices)-1][0]] + strconv.Itoa(strIndexaseNstrIndex+nstrIndex) + (*str)[strIndices[len(strIndices)-1][1]:]
}

func sumRight(str *string, i int, nstrIndex int) {
	strIndices := numRegex.FindAllStringIndex((*str)[i:], -1)
	if len(strIndices) == 0 {
		return
	}
	strIndexaseNstrIndex, _ := strconv.Atoi((*str)[i+strIndices[0][0] : i+strIndices[0][1]])
	*str = (*str)[:i+strIndices[0][0]] + strconv.Itoa(strIndexaseNstrIndex+nstrIndex) + (*str)[i+strIndices[0][1]:]
}

func split(str *string) bool {
	compile, _ := regexp.Compile(`[1-9][0-9]+`)
	strIndices := compile.FindAllStringIndex(*str, -1)
	if len(strIndices) == 0 {
		return false
	}
	strIndex := strIndices[0]
	strIndexig, _ := strconv.Atoi((*str)[strIndex[0]:strIndex[1]])
	newPair := fmt.Sprintf("[%d,%d]", strIndexig/2, int(math.Round(float64(strIndexig)/2.0)))
	*str = (*str)[:strIndex[0]] + newPair + (*str)[strIndex[1]:]
	return true
}

func isExpectedDepth(str string, i int) bool {
	return strings.Count(str[:i], "[")-strings.Count(str[:i], "]") >= 4
}
