package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file, _ := os.Open("input-example.txt")
	file, _ := os.Open("../input.txt")
	defer file.Close()
	fileReader := bufio.NewReader(file)
	horizontal, depth := calcPosition(fileReader)
	log.Print(horizontal, depth)

	count := horizontal * depth
	log.Printf("Overall of position:, %d", count)
}

func calcPosition(r io.Reader) (int, int) {
	scanner := bufio.NewScanner(r)
	horizontal := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("Line:", line)
		input := strings.Split(line, " ")
		direction := input[0]
		value, _ := strconv.Atoi(input[1])
		switch {
		case direction == "forward":
			horizontal = horizontal + value
			if aim > 0 {
				depth = depth + (value * aim)
			}
		case direction == "down":
			aim = aim + value
		case direction == "up":
			aim = aim - value
		}
		log.Println("Horizontal:", horizontal)
		log.Println("Depth:", depth)
	}
	return horizontal, depth
}
