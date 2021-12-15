package common

import (
	"fmt"
	"strconv"
	"strings"
)

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		if strings.TrimSpace(a) == "" {
			continue
		}
		i, err := strconv.Atoi(strings.TrimSpace(a))
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SortCharsInStrings(strs []string) []string {
	for pos, str := range strs {
		strs[pos] = SortString(str)
	}
	return strs
}

func ParsIntGrid(rows []string, itemDelimiter string) [][]int {
	grid := make([][]int, len(rows))
	for i := 0; i < len(rows); i++ {
		if strings.TrimSuffix(strings.TrimSpace(rows[i]), "\n") == "" {
			continue
		}
		columns := strings.Split(rows[i], itemDelimiter)
		cols, _ := SliceAtoi(columns)
		grid[i] = cols

	}
	fmt.Println("Grid completed", grid)
	return grid
}
