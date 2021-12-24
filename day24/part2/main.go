package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
	"strconv"
	"strings"
)

type monad struct {
	start, end int
}

func main() {
	//absPath, _ := filepath.Abs("day24/input-example.txt")
	absPath, _ := filepath.Abs("day24/input.txt")
	data := strings.Split(common.ReadFile(absPath), "inp w\n")[1:]
	modelNumber := make([]int, 14)
	z := []monad{}
	for i, chunk := range data {
		fmt.Println("Index:", i)
		lines := strings.Split(chunk, "\n")
		popValue, _ := strconv.Atoi(strings.Split(lines[3], " ")[2])
		pop := popValue == 26
		xAdd, _ := strconv.Atoi(strings.Split(lines[4], " ")[2])
		yAdd, _ := strconv.Atoi(strings.Split(lines[14], " ")[2])

		if !pop {
			z = append(z, monad{i, yAdd})
			fmt.Println("Z:", z)
		} else {
			var zVal monad
			zVal, z = popArr(z)
			j, yAdd := zVal.start, zVal.end
			difference := xAdd + yAdd
			if difference < 0 {
				modelNumber[i] = 1
				modelNumber[j] = 1 - difference
			} else if difference > 0 {
				modelNumber[i] = 1 + difference
				modelNumber[j] = 1
			} else {
				modelNumber[i] = 1
				modelNumber[j] = 1
			}
		}
	}

	fmt.Println("Answer: ", common.IntSliceToString(modelNumber, ""))
}

func popArr(arr []monad) (monad, []monad) {
	val := arr[len(arr)-1]
	return val, arr[:len(arr)-1]
}
