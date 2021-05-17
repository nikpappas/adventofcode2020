package src

import (
	"fmt"

	"nikpappas.com/adventofcode2021/files"
	"nikpappas.com/adventofcode2021/maps"
)

func Day9() {
	fmt.Println("Day 9.")

	lines := files.ReadLines("inputs/day9.txt")
	ints := maps.MapLinesToInts(lines)
	firstInvalidPos := day9sol1(ints, 25)
	fmt.Println("the first Invalid number is", firstInvalidPos, ints[firstInvalidPos])
	minRange, maxRange := day9sol2(ints, firstInvalidPos)
	fmt.Println("The sum is", minRange, "+", maxRange, "=", minRange+maxRange)

}

func day9sol2(ints []int, firstInvalidPos int) (int, int) {
	for i := 0; i < len(ints); i++ {
		start, end := contingentSetThatSumsToN(ints, ints[firstInvalidPos], i)
		if start > 0 && end > 0 {
			min := int(^uint(0) >> 1)
			max := -min
			fmt.Println(min, max)
			for x := start; x <= end; x++ {
				if ints[x] > max {
					max = ints[x]
				}
				if ints[x] < min {
					min = ints[x]
				}
			}
			return min, max
		}
	}

	return -1, -1

}

func day9sol1(ints []int, preamble int) int {
	for i := preamble; i < len(ints); i++ {
		if !isSumOfPrevs(ints, i, preamble) {
			return i
		}

	}
	return -1
}

func contingentSetThatSumsToN(ints []int, n int, start int) (int, int) {
	sum := 0
	for i := start; i < len(ints); i++ {
		sum += ints[i]
		if sum > n {
			return -1, -1
		} else if sum == n {
			return start, i
		}
	}
	return -1, -1

}

func isSumOfPrevs(ints []int, i int, preamble int) bool {
	for x := 1; x <= preamble; x++ {
		for y := 1; y <= preamble; y++ {
			if ints[i] == (ints[i-x] + ints[i-y]) {
				return true
			}
		}
	}
	return false
}
