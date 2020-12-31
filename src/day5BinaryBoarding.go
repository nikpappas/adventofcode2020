package main

import (
	"fmt"
	"sort"

	"./files"
)

func day5() {
	fmt.Println("Day 5.")

	lines := files.ReadLines("inputs/day5.txt")
	max, ids := day5Sol1(lines)
	fmt.Println("amx", max)

	mySeatId := day5Sol2(ids)
	fmt.Println(mySeatId)

}

func day5Sol2(ids []int) int {
	fmt.Println("Solution 2")
	sort.Ints(ids)
	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1] != ids[i]+1 {
			fmt.Println("i:", i, "-->", ids[i])
			return ids[i] + 1
		}
	}
	return -1
}
func day5Sol1(lines []string) (int, []int) {
	var ids []int
	for _, line := range lines {
		rowCode := line[:7]
		colCode := line[7:]
		fmt.Println(rowCode, colCode)
		rowU := 127
		rowL := 0
		for _, c := range rowCode {
			rowL, rowU = disect(rowL, rowU, c == 'B')
		}
		fmt.Println(rowL, rowU)
		colU := 7
		colL := 0
		for _, c := range colCode {
			colL, colU = disect(colL, colU, c == 'R')
		}
		fmt.Println(colL, colU)
		id := (rowU)*8 + (colU)
		fmt.Println(id)
		ids = append(ids, id)

		fmt.Println("_____")
	}
	max := 0
	for _, id := range ids {
		if id > max {
			max = id
		}

	}

	return max, ids

}

func disect(lower int, upper int, higher bool) (int, int) {
	diff := upper - lower + 1
	offset := diff / 2

	if higher {
		lower += offset
	} else {
		upper -= offset
	}
	return lower, upper
}
