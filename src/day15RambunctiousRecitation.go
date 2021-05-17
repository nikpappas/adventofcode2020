package src

import (
	"fmt"
	"strings"

	"nikpappas.com/adventofcode2021/files"
	"nikpappas.com/adventofcode2021/maps"
)

func Day15() {
	lines := files.ReadLines("inputs/day15.txt")
	toks := strings.Split(lines[0], ",")
	ints := maps.MapLinesToInts(toks)
	res := day15sol1(ints, 2020)
	fmt.Println(res)
	res = day15sol2(ints, 2020)
	fmt.Println(res)
	res = day15sol2(ints, 30000000)
	fmt.Println(res)

}

func day15sol1(ints []int, limit int) int {
	fmt.Println(ints)
	for len(ints) < limit {
		i := findLast(ints, getLast(ints))
		if i >= 0 {
			ints = append(ints, (len(ints)-1)-i)
		} else {
			ints = append(ints, 0)
		}
	}
	return getLast(ints)
}
func day15sol2(ints []int, limit int) int {
	mem := make(map[int]([]int))
	for i, n := range ints {
		mem[n] = []int{i}
	}
	fmt.Println(mem)
	fmt.Println(ints)
	for len(ints) < limit {
		last := getLast(ints)
		l, ok := mem[last]
		i := -1
		if ok {
			i = l[0]
		}
		if i == len(ints)-1 {
			if len(l) > 1 {
				i = l[1]
			}
		}
		next := 0
		if i >= 0 {
			next = (len(ints) - 1) - i
		}
		ints = append(ints, next)
		l, ok = mem[next]
		if ok {
			mem[next] = append([]int{len(ints) - 1}, l[0])
		} else {
			mem[next] = []int{len(ints) - 1}

		}
	}
	return getLast(ints)
}

func findLast(ints []int, x int) int {
	for i := len(ints) - 2; i >= 0; i-- {
		if ints[i] == x {
			return i
		}
	}
	return -1

}

func getLast(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return arr[len(arr)-1]
}
