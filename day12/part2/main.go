package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day12/input-example.txt")
	absPath, _ := filepath.Abs("day12/input.txt")
	lines := common.ReadLines(absPath)

	paths := map[string][]string{}
	for _, line := range lines {
		ps := strings.Split(line, "-")
		a, b := ps[0], ps[1]
		paths[a] = append(paths[a], b)
		paths[b] = append(paths[b], a)
	}

	res := 0
	seenPath := map[string]bool{}
	var search func(searchStr string, alreadySeen bool)
	search = func(searchStr string, alreadySeen bool) {
		currentSeen := seenPath[searchStr]

		if searchStr == "end" {
			res++
			return
		}

		if searchStr[0] >= 'a' && searchStr[0] <= 'z' && currentSeen {
			if alreadySeen || searchStr == "start" {
				return
			}
			alreadySeen = true
		}
		seenPath[searchStr] = true
		for _, n := range paths[searchStr] {
			search(n, alreadySeen)
		}
		seenPath[searchStr] = currentSeen
	}

	search("start", false)

	fmt.Println("Answer: ", res)
}
