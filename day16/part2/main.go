package main

import (
	"encoding/hex"
	"fmt"
	"path/filepath"

	"github.com/emmalp/advent-of-code-2021/pkg/common"
)

func main() {
	// absPath, _ := filepath.Abs("day16/input-example.txt")
	absPath, _ := filepath.Abs("day16/input.txt")
	lines := common.ReadLines(absPath)
	binary, _ := hex.DecodeString(lines[0])

	bitDecorder := &bitDecoder{raw: binary}
	var answer []int

	var parsePacket func()
	parsePacket = func() {
		version := bitDecorder.take(3)
		messageType := bitDecorder.take(3)
		_ = version

		switch messageType {
		case 4:
			num := 0
			for {
				p := bitDecorder.take(5)
				num = (num << 4) | (p & 0xf)
				if p&0x10 == 0 {
					break
				}
			}
			answer = append(answer, num)
		default:
			tlenid := bitDecorder.take(1)
			vcount := 0
			if tlenid == 0 {
				blen := bitDecorder.take(15)
				start := bitDecorder.total
				for bitDecorder.total < start+blen {
					parsePacket()
					vcount++
				}
			} else {
				pcount := bitDecorder.take(11)
				for i := 0; i < pcount; i++ {
					parsePacket()
				}
				vcount = pcount
			}
			tmpStack := answer[len(answer)-vcount:]
			answer = answer[:len(answer)-vcount]

			answer = append(answer, determineValue(messageType, tmpStack))
		}
	}

	parsePacket()
	fmt.Println("Answer: ", answer[0])
}

func determineValue(messageType int, stack []int) (val int) {
	switch messageType {
	case 0: // sum
		for _, v := range stack {
			val += v
		}
	case 1: // product
		val = 1
		for _, v := range stack {
			val *= v
		}
	case 2: // min
		val = stack[0]
		for _, v := range stack[1:] {
			if v < val {
				val = v
			}
		}
	case 3: // max
		val = stack[0]
		for _, v := range stack[1:] {
			if v > val {
				val = v
			}
		}
	case 5: // greater
		if stack[0] > stack[1] {
			val = 1
		}
	case 6: // less
		if stack[0] < stack[1] {
			val = 1
		}
	case 7: // equal
		if stack[0] == stack[1] {
			val = 1
		}
	}
	return val
}

type bitDecoder struct {
	raw   []byte
	bit   int
	total int
}

func (d *bitDecoder) take(bits int) int {
	v := 0
	for i := 0; i < bits; i++ {
		v = (v << 1) | d.takeBit()
	}
	return v
}

func (d *bitDecoder) takeBit() int {
	if d.bit == 8 {
		d.bit = 0
		d.raw = d.raw[1:]
	}
	shift := 7 - d.bit
	v := d.raw[0] & (1 << shift)
	d.bit++
	d.total++
	return int(v >> shift)
}
