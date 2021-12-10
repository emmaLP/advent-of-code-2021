package main

import (
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type NumGrid struct {
	grid [][]int
}

func main() {
	//absPath, _ := filepath.Abs("day4/input-example.txt")
	absPath, _ := filepath.Abs("day4/input.txt")
	lines := common.ReadLines(absPath)
	fmt.Println(lines)
	bingNumbers := strings.Split(lines[0], ",")
	fmt.Println("Bingo numbers: ", bingNumbers)
	grids := getsBoards(lines)
	fmt.Println("number of boards", len(grids))

	for _, num := range bingNumbers {
		number, _ := strconv.Atoi(num)
		fmt.Println("Drawn number:", number)
		for bingoNumIndex, grid := range grids {
			markChecked(number, grid.grid)
			if bingoNumIndex < 4 {
				// No point checking for bingo when you need at least 5 callouts
				continue
			}
			if checkForBingo(grid.grid) {
				fmt.Println(calculateResult(grid.grid, number))
				return
			}
		}
	}

}

func calculateResult(grid [][]int, num int) int {
	fmt.Println("Winning grid is ", grid)
	return sumUnchecked(grid) * num
}

func sumUnchecked(grid [][]int) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != -1 {
				sum += grid[i][j]
			}
		}
	}
	return sum
}

func markChecked(num int, grid [][]int) bool {
	changed := false
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == num {
				grid[i][j] = -1
				changed = true
			}
		}
	}
	return changed
}

func checkForBingo(grid [][]int) bool {
	isBingo := false
	//Row Matching
	for i := 0; i < len(grid); i++ {
		//fmt.Println("Grid row to check", grid[i])
		if isBingo = reflect.DeepEqual(grid[i], []int{-1, -1, -1, -1, -1}); isBingo {
			return isBingo
		}

	}

	//Column matching
	for colIndex := 0; colIndex < len(grid[0]); colIndex++ {
		allMarked := true
		for _, row := range grid {
			if row[colIndex] != -1 {
				allMarked = false
				break
			}
		}
		if allMarked {
			return true
		}

	}
	return false
}

func getsBoards(lines []string) []NumGrid {
	var grids []NumGrid
	gridLines := ""
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" || i == len(lines)-1 {
			if i == len(lines)-1 {
				gridLines += line
			}
			board := parseBingoBoard(gridLines, "\n", " ")
			grids = append(grids, board)
			gridLines = ""
			continue
		}

		gridLines += line
		gridLines += "\n"

	}
	return grids
}

func parseBingoBoard(str string, lineDelimiter string, itemDelimiter string) NumGrid {
	fmt.Println("Line to make as a board", str)
	rows := strings.Split(str, lineDelimiter)
	fmt.Println()
	grid := make([][]int, 5)
	for i := 0; i < len(rows); i++ {
		fmt.Println("Processing row", rows[i])
		if strings.TrimSuffix(strings.TrimSpace(rows[i]), "\n") == "" {
			continue
		}
		columns := strings.Split(rows[i], itemDelimiter)
		cols, _ := sliceAtoi(columns)
		if len(cols) == 5 {
			grid[i] = cols
		}

	}
	fmt.Println("Board structure completed", grid)
	return NumGrid{grid: grid}

}

func sliceAtoi(sa []string) ([]int, error) {
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
