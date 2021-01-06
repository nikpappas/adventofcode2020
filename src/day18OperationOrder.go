package main

import (
	"fmt"
	"strconv"
	"strings"

	"./files"
)

func day18() {
	lines := files.ReadLines("inputs/day18.txt")
	res := day18sol1(lines)
	fmt.Println(res)

}

func day18sol2(lines []string) int {
	sum := 0
	for _, line := range lines {
		fmt.Println("======================", line)
		line = strings.ReplaceAll(line, " ", "")
		fmt.Println("cleanline", line)

		// line := addBrackets(line)
		res := calcBrackets(line)
		sum += res
		fmt.Println("======================", res)
		fmt.Println(sum)
	}
	return sum

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

// func addBrackets(line string) string {
// 	firstPlus := indexOfFirst(line, '+')
// 	left, right, isBracketed := getOperands(line, firstPlus)

// 	return line

// }

// func getOperands(line string, index int) (string, string, bool) {

// }
func calcBrackets(line string) int {
	latestOpeningBracket := indexOfLast(line, '(')
	for latestOpeningBracket != -1 {
		fmt.Println(line)
		offset := indexOfFirst(line[latestOpeningBracket:], ')')
		firstClosing := latestOpeningBracket + offset

		fmt.Println("->", line[latestOpeningBracket+1:firstClosing])
		res := calcPrecedenceLine(line[latestOpeningBracket+1 : firstClosing])
		line = line[:latestOpeningBracket] + fmt.Sprint(res) + line[firstClosing+1:]
		fmt.Println("newline", line)
		latestOpeningBracket = indexOfLast(line, '(')
	}
	return calcPrecedenceLine(line)

}

func indexOfLast(line string, cQuery rune) int {
	latestOpeningBracket := -1
	for i, c := range line {
		if c == cQuery {
			latestOpeningBracket = i
		}
	}
	return latestOpeningBracket
}
func indexOfFirst(line string, cQuery rune) int {
	for i, c := range line {
		if c == cQuery {
			return i
		}
	}
	return -1
}
func calcPrecedenceLine(line string) int {
	fmt.Println(line)
	indexPlus := strings.IndexRune(line, '+')
	indexMult := strings.IndexRune(line, '*')
	fmt.Println("indexPlus", indexPlus)
	fmt.Println("indexMult", indexMult)
	if indexPlus == -1 && indexMult == -1 {
		val, _ := strconv.Atoi(line)
		return val
	} else if indexPlus >= 0 {
		left := line[:indexPlus]
		right := line[indexPlus+1:]
		multL := indexOfLast(left, '*')
		multR := indexOfFirst(right, '*')
		if multL == -1 && multR == -1 {
			fmt.Println(left, "^+", right)
			return calcPrecedenceLine(left) + calcPrecedenceLine(right)
		} else if multL >= 0 {
			left = left[multL+1:]
			remainder := line[:multL]
			fmt.Println(remainder, "^^*", left, "^+", right)

			return (calcPrecedenceLine(left) + calcPrecedenceLine(right)) * calcPrecedenceLine(remainder)
		} else {
			remainder := right[multR+1:]
			right = right[:multR]
			fmt.Println(left, "^+", right, "^^*", remainder)

			return (calcPrecedenceLine(left) + calcPrecedenceLine(right)) * calcPrecedenceLine(remainder)

		}
	} else {
		left := line[:indexMult]
		right := line[indexMult+1:]
		fmt.Println(left, "^*", right)
		return calcPrecedenceLine(left) * calcPrecedenceLine(right)
	}

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
	// 2 * 3 + (4 * 5) becomes 26.
	// 5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
	// 5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
	// ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.
	fmt.Println("Testign solution 2")
	day18sol2([]string{"1 + (2 * 3) + (4 * (5 + 6))"})
	day18sol2([]string{"2 * 3 + (4 * 5)"})
	day18sol2([]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"})
	day18sol2([]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"})
	day18sol2([]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"})

	// 1 + (2 * 3) + (4 * (5 + 6)) still becomes 51.
	// 2 * 3 + (4 * 5) becomes 46.
	// 5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 1445.
	// 5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 669060.
	// ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 23340.
	day18sol2([]string{"1 + 2 + 3"})
	day18sol2([]string{"2 * 2 + 3"})
	day18sol2([]string{"3+2 * 2"})
	day18sol2([]string{"3+2 * 2+1"})

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
