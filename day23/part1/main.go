package main

import (
	"container/heap"
	"fmt"
	"github.com/emmalp/advent-of-code-2021/pkg/common"
	"math"
	"path/filepath"
)

var (
	stepsFromHallwayToPodDoor = [][]int{
		{2, 4, 6, 8},
		{1, 3, 5, 7},
		{1, 1, 3, 5},
		{3, 1, 1, 3},
		{5, 3, 1, 1},
		{7, 5, 3, 1},
		{8, 6, 4, 2},
	}
	energyPerStep = map[rune]int{
		'A': 1,
		'B': 10,
		'C': 100,
		'D': 1000,
	}
)

type stateHeap []StateEnergy
type State struct {
	Pods    [4][2]rune
	Hallway [7]rune
}
type StateEnergy struct {
	Value  State
	Energy int
}

func main() {
	//absPath, _ := filepath.Abs("day23/input-example.txt")
	absPath, _ := filepath.Abs("day23/input.txt")
	lines := common.ReadLines(absPath)

	initState := parseInput(lines)

	fmt.Println("Answer:", calculateMinEnergy(initState))

}

func parseInput(lines []string) State {
	s := State{}

	for r := 0; r < 2; r++ {
		values := []rune(lines[r+2])
		for p := 0; p < 4; p++ {
			c := values[2*p+3]
			s.Pods[p][r] = c
		}
	}

	return s
}

func calculateMinEnergy(initState State) (minScore int) {
	minScore = math.MaxInt
	finalState := State{[4][2]rune{{'A', 'A'}, {'B', 'B'}, {'C', 'C'}, {'D', 'D'}}, [7]rune{}}

	processed := map[State]int{}
	h := make(stateHeap, 1)
	h[0] = StateEnergy{initState, 0}

	heap.Init(&h)
	fmt.Println("Heap Length", h.Len())
	for h.Len() > 0 {
		current := heap.Pop(&h).(StateEnergy)
		fmt.Println("Current State:", current)
		for _, se := range nextStates(current.Value, current.Energy) {
			s := se.Value
			e := se.Energy

			foundEnergy, ok := processed[s]
			if !ok || e < foundEnergy {
				processed[s] = e
				heap.Push(&h, se)
			}
		}
	}

	return processed[finalState]
}

func nextStates(s State, startEnergy int) []StateEnergy {
	ret := make([]StateEnergy, 0)
	var energy int

	// Check everything in hallway to see if it can go home
outerHallway:
	for pos := 0; pos < 7; pos++ {
		c := s.Hallway[pos]
		if c == 0 {
			continue
		}

		// Letter found, check if the corresponding pod is ready
		podDesired := int(c - 'A')
		for row := 0; row < 2; row++ {
			if s.Pods[podDesired][row] != c && s.Pods[podDesired][row] != 0 {
				// Pod is not ready
				continue outerHallway
			}
		}

		// Check if the path to pod is open
		stopping := podDesired + 2
		if pos < stopping {
			for i := pos + 1; i < stopping; i++ {
				if s.Hallway[i] != 0 {
					continue outerHallway
				}
			}
		} else {
			for i := pos - 1; i >= stopping; i-- {
				if s.Hallway[i] != 0 {
					continue outerHallway
				}
			}
		}

		// It is possible to move piece home
		var endRow int
		if s.Pods[podDesired][1] == 0 {
			endRow = 1
		} else {
			endRow = 0
		}

		energy = startEnergy + (stepsFromHallwayToPodDoor[pos][podDesired]+endRow+1)*energyPerStep[c]
		nState := s
		nState.Hallway[pos] = 0
		nState.Pods[podDesired][endRow] = c
		ret = append(ret, StateEnergy{nState, energy})
	}

	if len(ret) > 0 {
		return ret
	}

	// Run through all pods to see if they can evolve
outerPods:
	for pod := 0; pod < 4; pod++ {
		endState := 'A' + rune(pod)
		for row := 0; row < 2; row++ {
			c := s.Pods[pod][row]
			if c == endState {
				skip := true
				// Check if all descending runes are correct
				for d := row + 1; d < 2; d++ {
					if s.Pods[pod][d] != endState {
						skip = false
					}
				}
				if skip {
					continue outerPods
				}
			}

			if c != 0 {
				// Will have to try moving available piece then end the pod

				// 2, 3, 4, 5
				stopping := pod + 2 // This is the split point in the hallway for this pod

				// Iterate from hallway position down and up until you hit another letter
				for i := stopping - 1; i >= 0; i-- {
					if s.Hallway[i] == 0 {
						nState := s
						nState.Pods[pod][row] = 0
						nState.Hallway[i] = c
						energy = startEnergy + (stepsFromHallwayToPodDoor[i][pod]+row+1)*energyPerStep[c]
						ret = append(ret, StateEnergy{nState, energy})
					} else {
						break
					}
				}

				for i := stopping; i < 7; i++ {
					if s.Hallway[i] == 0 {
						nState := s
						nState.Pods[pod][row] = 0
						nState.Hallway[i] = c
						energy = startEnergy + (stepsFromHallwayToPodDoor[i][pod]+row+1)*energyPerStep[c]
						ret = append(ret, StateEnergy{nState, energy})
					} else {
						break
					}
				}
				// Can't move the letter under this one
				continue outerPods
			}
		}
	}

	return ret
}

func (h stateHeap) Len() int { return len(h) }
func (h stateHeap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
func (h *stateHeap) Push(x interface{}) {
	se := x.(StateEnergy)
	*h = append(*h, se)
}
func (h *stateHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func (h stateHeap) Less(i, j int) bool {
	return h[i].Energy < h[j].Energy
}
