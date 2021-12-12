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
	fmt.Println("Paths:", paths)

	res := 0
	seenPath := map[string]bool{}
	var search func(searchStr string)
	search = func(searchStr string) {
		if searchStr == "end" {
			res++
			return
		}
		if seenPath[searchStr] && searchStr[0] >= 'a' && searchStr[0] <= 'z' {
			return
		}
		seenPath[searchStr] = true
		for _, n := range paths[searchStr] {
			search(n)
		}
		seenPath[searchStr] = false
	}

	search("start")

	fmt.Println("Answer: ", res)
}
