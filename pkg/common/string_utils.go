package common

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func ContainsChars(s1, s2 string) bool {
	for _, c := range s2 {
		if !strings.ContainsRune(s1, c) {
			return false
		}
	}
	return true
}
