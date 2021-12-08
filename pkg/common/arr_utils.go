package common

import (
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
