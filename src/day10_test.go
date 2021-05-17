package src

import (
	"sort"
	"testing"
)

func TestDay10(t *testing.T) {
	expected := 0
	res := 0
	startTest(1)
	expected = 2
	res = countPermutations([]int{0, 1, 3}, 0)
	assert(expected, res, t)

	startTest(2)
	expected = 3
	res = countPermutations([]int{0, 1, 3, 4}, 0)
	assert(expected, res, t)

	startTest(3)
	expected = 8
	res = countPermutations([]int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}, 0)
	assert(expected, res, t)

	startTest(4)
	expected = 19208
	toCalc := []int{0, 28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3, 52}
	sort.Ints(toCalc)
	res = countPermutations(toCalc, 0)
	assert(expected, res, t)

}
