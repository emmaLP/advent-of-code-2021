package common

import (
	"io"
	"log"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	return strings.Split(ReadFile(path), "\n")
}

func ReadFile(path string) string {
	file, err := os.Open(path)
	log.Println("Err:", err)
	defer file.Close()

	data, _ := io.ReadAll(file)
	return string(data)
}
