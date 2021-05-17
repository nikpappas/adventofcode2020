package src

import (
	"fmt"
	"sort"
	"strings"

	"nikpappas.com/adventofcode2021/files"
	"nikpappas.com/adventofcode2021/maps"
)

var prevDepth = 0
var hashes = make(map[string](int))

func Day10() {
	fmt.Println("Day 10.")

	lines := files.ReadLines("inputs/day10.txt")

	ints := maps.MapLinesToInts(lines)
	sort.Ints(ints)
	ints = append([]int{0}, ints...)
	ints = append(ints, ints[len(ints)-1]+3)

	prod := day10sol1(ints)
	fmt.Println(prod)
	count := day10sol2(ints)
	fmt.Println(count)
}

func day10sol1(ints []int) int {
	histo := make(map[int](int))
	for i := 1; i < len(ints); i++ {
		histo[ints[i-1]-ints[i]] += 1
		fmt.Println(ints[i-1], ints[i], ints[i-1]-ints[i], histo[ints[i-1]-ints[i]])
	}
	fmt.Println(histo[-1], histo[-3])
	return histo[-1] * histo[-3]

}

func day10sol2(ints []int) int {
	return countPermutations(ints, 0)
}

func countPermutations(ints []int, depth int) int {
	count := 1
	if len(ints) < 3 {
		return count
	}
	intsHash := hash(ints)
	preCalced, has := hashes[intsHash]
	if has {
		return preCalced
	}

	for i := 0; i < len(ints)-2; i++ {
		diffNext := ints[i+2] - ints[i]

		if diffNext <= 3 {
			var newInts []int
			for j := i; j < len(ints); j++ {
				if j != (i + 1) {
					newInts = append(newInts, ints[j])
				}
			}

			count += countPermutations(newInts, depth+1)
			if depth != prevDepth {
				fmt.Println(strings.Repeat(" ", depth), "*")
				prevDepth = depth
			}

		}
	}
	hashes[intsHash] = count
	return count
}

func hash(ints []int) string {
	toRet := ""
	for _, i := range ints {
		toRet += fmt.Sprint(i) + ","
	}
	return toRet
}
