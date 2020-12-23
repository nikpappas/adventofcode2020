package main

import "fmt"

func day1() {
	var lines = readLines("inputs/day1.txt")
	var ints = mapLinesToInts(lines)

	day1Sol1(ints)
	day1Sol2(ints)
}

func day1Sol1(ints []int) int {
	fmt.Println("Solution 1")
	for i, a := range ints {
		for j, b := range ints {
			if a+b == 2020 {
				fmt.Println(i, j)
				var res = a * b
				fmt.Println(res)
				return res
			}
		}
	}
	return -1

}
func day1Sol2(ints []int) int {
	fmt.Println("Solution 2")
	for i, a := range ints {
		for j, b := range ints {
			for k, c := range ints {

				if a+b+c == 2020 {
					fmt.Println(i, j, k)
					var res = a * b * c
					fmt.Println(res)
					return res

				}
			}
		}
	}
	return -1
}