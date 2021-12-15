package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("../input.txt")
	fileReader := bufio.NewReader(file)
	ints, err := ReadInts(fileReader)
	if err != nil {
		log.Fatal("An error occurred.", err)
	}
	log.Print(ints)

	slidingWindow := sumSlidingWindow(ints)
	log.Printf("Sliding Window Values:, %d", slidingWindow)

	count := countLargerMeasurements(slidingWindow)
	log.Printf("Count of Measurements:, %d", count)
}

func countLargerMeasurements(values []int) int {
	count := 0
	for i := 0; i < len(values); i++ {
		if i == 0 {
			continue
		}
		if values[i] > values[i-1] {
			log.Println(values[i])
			count++
		}
	}
	return count
}

func sumSlidingWindow(values []int) []int {
	var count []int
	for i := 0; i < len(values); i++ {
		var sumValue int
		if i == len(values)-2 {
			break
		} else {
			sumValue = values[i] + values[i+1] + values[i+2]
		}
		log.Println(sumValue)
		count = append(count, sumValue)

	}
	return count
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
