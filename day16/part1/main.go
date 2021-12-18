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

	bDec := &bitDecoder{raw: binary}

	answer := determinePacket(bDec, 0)

	fmt.Println("Answer:", answer)

}

func determinePacket(bitDecorder *bitDecoder, answer int) (ans int) {
	version := bitDecorder.takeBits(3)
	typ := bitDecorder.takeBits(3)

	ans += answer + version

	switch typ {
	case 4:
		num := 0
		for {
			p := bitDecorder.takeBits(5)
			num = (num << 4) | (p & 0xf)
			if p&0x10 == 0 {
				break
			}
		}
	default:
		tlenid := bitDecorder.takeBits(1)
		if tlenid == 0 {
			blen := bitDecorder.takeBits(15)
			start := bitDecorder.total
			for bitDecorder.total < start+blen {
				ans = determinePacket(bitDecorder, ans)
			}
		} else {
			pcount := bitDecorder.takeBits(11)
			for i := 0; i < pcount; i++ {
				ans = determinePacket(bitDecorder, ans)
			}
		}
	}
	return ans
}

type bitDecoder struct {
	raw   []byte
	bit   int
	total int
}

func (d *bitDecoder) takeBits(bits int) int {
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
