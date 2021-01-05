package main

import (
	"fmt"

	"./files"
)

func day18() {
	lines := files.ReadLines("inputs/day18.txt")
	res := day18sol1(lines)
	fmt.Println(res)

}

func day18sol1(lines []string) int {
	sum := 0
	for _, line := range lines {

		fmt.Println("======================", line)
		res, _ := calculateLine(line)
		sum += res
		fmt.Println("======================", res)
		fmt.Println(sum)
	}
	return sum

}

func calculateLine(line string) (int, int) {
	fmt.Println(line)
	res := 0
	op := ' '
	first := true
	for i := 0; i < len(line); i++ {
		c := rune(line[i])
		if c == '(' {
			lRes, li := calculateLine(line[i+1:])
			res, first = calc(res, lRes, op, first)
			i += li
			continue
		}
		if c == ')' {
			fmt.Println(res, "i:", i+1)
			return res, i + 1
		}
		if c == ' ' {
			continue
		}
		if c == '*' || c == '+' {
			op = c
			continue
		}
		res, first = calc(res, int(c-'0'), op, first)
		fmt.Println("=", res)
	}
	return res, len(line)
}

func day18test() {
	day18sol1([]string{"1 + 2 + 3"})
	day18sol1([]string{"2 * 2 + 3"})
	day18sol1([]string{"1 + 2 * 3"})
	day18sol1([]string{"1 + 2 * (3 + 1)"})
	day18sol1([]string{"1 + 2 * ((3 + 1))"})
	day18sol1([]string{"9 + (2 * ((3 + 1)))"})

	day18sol1([]string{"2 * 3 + (4 * 5)"})
	day18sol1([]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"})
	day18sol1([]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"})
	day18sol1([]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"})
}

func calc(res int, operand int, op rune, first bool) (int, bool) {
	if first {
		first = false
		return operand, false
	}
	if op == '+' {
		fmt.Println(res, "+", operand)
		res += operand
	}
	if op == '*' {
		fmt.Println(res, "*", operand)
		res *= operand
	}
	return res, false

}

// 2 * 3 + (4 * 5) becomes 26.
// 5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
// 5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
// ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.
