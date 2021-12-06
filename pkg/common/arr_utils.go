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
