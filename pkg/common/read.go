package common

import (
	"io"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	data, _ := io.ReadAll(file)
	return strings.Split(string(data), "\n")
}
