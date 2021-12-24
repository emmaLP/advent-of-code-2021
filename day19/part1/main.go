package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day19/input-example.txt")
	absPath, _ := filepath.Abs("day19/input.txt")
	parts := strings.Split(common.ReadFile(absPath), "\n\n")

	locs := make([][][3]int, len(parts))
	for i, part := range parts {
		lines := strings.Split(part, "\n")[1:]
		locs[i] = make([][3]int, len(lines))
		for j, line := range lines {
			ps := strings.Split(line, ",")
			for k, p := range ps {
				locs[i][j][k], _ = strconv.Atoi(p)
			}
		}
	}

	full := locs[0]
	fullMap := map[[3]int]bool{}
	for _, pt := range full {
		fullMap[pt] = true
	}
	done := map[int]bool{}
	done[0] = true

	for len(done) < len(locs) {
		for i := 1; i < len(locs); i++ {
			if done[i] {
				continue
			}

			offset, orient, ok := matchPair(full, locs[i])
			if !ok {
				continue
			}

			done[i] = true

			for _, pt := range locs[i] {
				opt := add(rotPoint(pt, orient), offset)
				if _, ok := fullMap[opt]; !ok {
					full = append(full, opt)
					fullMap[opt] = true
				}
			}
		}
	}

	fmt.Println(len(fullMap))
}

var orientations = [][3]int{}

func init() {
	os := map[[3]int]bool{}
	llen := 0
	os[[3]int{1, 2, 3}] = true
	for {
		for k := range os {
			os[[3]int{k[0], k[2], -k[1]}] = true
			os[[3]int{k[0], -k[2], k[1]}] = true

			os[[3]int{k[2], k[1], -k[0]}] = true
			os[[3]int{-k[2], k[1], k[0]}] = true

			os[[3]int{k[1], -k[0], k[2]}] = true
			os[[3]int{-k[1], k[0], k[2]}] = true
		}
		if len(os) == llen {
			break
		}
		llen = len(os)
	}
	for k := range os {
		orientations = append(orientations, k)
	}
}

func rotPoint(pt [3]int, orient [3]int) [3]int {
	var res [3]int
	for i := 0; i < 3; i++ {
		p := orient[i]
		if p < 0 {
			res[-p-1] = -pt[i]
		} else {
			res[p-1] = pt[i]
		}
	}
	return res
}

func sub(pi [3]int, pj [3]int) [3]int {
	return [3]int{pi[0] - pj[0], pi[1] - pj[1], pi[2] - pj[2]}
}

func add(pi [3]int, pj [3]int) [3]int {
	return [3]int{pi[0] + pj[0], pi[1] + pj[1], pi[2] + pj[2]}
}

func inRange(p [3]int) bool {
	return p[0] >= -1000 && p[0] <= 1000 && p[1] >= -1000 && p[1] <= 1000 && p[2] >= -1000 && p[2] <= 1000
}

func matchPair(li [][3]int, lj [][3]int) (offset [3]int, dir [3]int, ok bool) {
	pis := map[[3]int]bool{}
	for _, pi := range li {
		pis[pi] = true
	}
	for _, pi := range li {
		for _, pj := range lj {
			// match up each pair of points
			for _, orient := range orientations {
				opj := rotPoint(pj, orient)
				offset = sub(pi, opj)

				count := 0
				for _, pk := range lj {
					opk := add(rotPoint(pk, orient), offset)
					if pis[opk] {
						count++
					}
				}
				if count >= 12 {
					dir = orient
					ok = true
					return
				}
			}
		}
	}

	ok = false
	return
}

