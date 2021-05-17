package src

import (
	"fmt"
	"strconv"
	"strings"

	"nikpappas.com/adventofcode2021/files"
)

func Day14() {
	fmt.Println("Day 14.")

	lines := files.ReadLines("inputs/day14.txt")
	// day14sol1(lines)
	day14sol2(lines)

}
func day14sol1(lines []string) {
	mem := make(map[int]uint64)
	mask := ""
	for _, line := range lines {
		instruction := strings.Split(line, " = ")
		op := instruction[0]
		fmt.Println(instruction)
		if op == "mask" {
			mask = instruction[1]
		} else {
			n, _ := strconv.ParseUint(instruction[1], 10, 64)
			add := parseAddress(instruction[0])
			bit := strconv.FormatUint(n, 2)
			masked := applyMask(bit, mask)
			val, _ := strconv.ParseUint(masked, 2, 64)

			mem[add] = val
		}
	}
	fmt.Println(mem)
	sum := uint64(0)
	for _, m := range mem {
		sum += m
	}
	fmt.Println(sum)

}

func day14sol2(lines []string) {
	mem := make(map[uint64]uint64)
	mask := ""
	for _, line := range lines {
		instruction := strings.Split(line, " = ")
		op := instruction[0]
		fmt.Println(instruction)
		if op == "mask" {
			mask = instruction[1]
		} else {
			n, _ := strconv.ParseUint(instruction[1], 10, 64)
			add := uint64(parseAddress(instruction[0]))
			byteVal := strconv.FormatUint(add, 2)
			// bit := strconv.FormatUint(n, 2)
			masked := applyMaskAddresses(byteVal, mask)
			for _, addByte := range masked {
				a, _ := strconv.ParseUint(addByte, 2, 64)
				mem[a] = n
			}

		}
	}
	// fmt.Println(mem)
	sum := uint64(0)
	for _, m := range mem {
		sum += m
	}
	fmt.Println(sum)

}

func applyMaskAddresses(byteVal string, mask string) []string {
	adds := []string{""}
	for i, _ := range mask {
		bit := mask[len(mask)-1-i]
		switch bit {
		case '0':
			if i < len(byteVal) {
				for x, _ := range adds {
					adds[x] = string(byteVal[len(byteVal)-1-i]) + adds[x]
				}
			} else {
				for x, _ := range adds {
					adds[x] = "0" + adds[x]
				}

			}
		case '1':
			for x, _ := range adds {
				adds[x] = "1" + adds[x]
			}
			break
		case 'X':
			initLen := len(adds)
			for _, add := range adds {
				adds = append(adds, add)
			}
			for x, a := range adds {
				if x < initLen {
					adds[x] = "0" + a
				} else {
					adds[x] = "1" + a
				}
			}
			break
		}

	}

	return adds
}

func applyMask(bit string, mask string) string {
	newBit := ""
	fmt.Println("mask  ", mask)
	for i, _ := range mask {
		maskbit := mask[len(mask)-1-i]
		switch maskbit {
		case 'X':
			if i >= len(bit) {
				newBit = "0" + newBit
			} else {
				newBit = string(bit[len(bit)-1-i]) + newBit
			}
			break
		case '0':
			newBit = "0" + newBit
			break
		case '1':
			newBit = string(maskbit) + newBit
			break
		}

	}
	return newBit
}

func parseAddress(op string) int {
	fmt.Println(op)
	add, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(op, "]", ""), "mem[", ""))
	return add
}
