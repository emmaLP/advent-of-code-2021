package common

import (
	"io"
	"log"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	log.Println("Err:", err)
	defer file.Close()

	data, _ := io.ReadAll(file)
	return strings.Split(string(data), "\n")
}
