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
	res64 := day18sol2(lines)
	fmt.Println("day2", res64)

}

// Day 1 __vv
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

// Day 1 --^^
// Day 2 __vv

func day18sol2(lines []string) int {
	sum := 0
	for _, line := range lines {
		fmt.Println("======================", line)
		line = strings.ReplaceAll(line, " ", "")
		fmt.Println("cleanline", line)
		res := calcBrackets(line)
		sum += res
		fmt.Println("======================", res)
		fmt.Println(sum)
	}
	return sum

}

func calcBrackets(line string) int {
	latestOpeningBracket := indexOfLast(line, '(')
	for latestOpeningBracket != -1 {
		fmt.Println("hasBrackets", line)
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
	fmt.Println("line", line)
	operations := splitOps(line)
	hasPlus := hasOp(operations, "+")
	for hasPlus > 0 {
		left, _ := strconv.Atoi(operations[hasPlus-1])
		right, _ := strconv.Atoi(operations[hasPlus+1])
		res := fmt.Sprint(left + right)
		operations[hasPlus-1] = res
		operations = append(operations[:hasPlus], operations[hasPlus+2:]...)
		fmt.Println("operations", operations)
		hasPlus = hasOp(operations, "+")
	}
	hasMult := hasOp(operations, "*")
	for hasMult > 0 {
		left, _ := strconv.Atoi(operations[hasMult-1])
		right, _ := strconv.Atoi(operations[hasMult+1])
		res := fmt.Sprint(left * right)
		operations[hasMult-1] = res
		operations = append(operations[:hasMult], operations[hasMult+2:]...)
		fmt.Println("operations", operations)
		hasMult = hasOp(operations, "*")
	}
	fmt.Println(operations)
	if len(operations) > 1 {
		fmt.Println("Error errore reerorro")
		return -1
	}

	res, _ := strconv.Atoi(operations[0])
	// res, _ := strconv.Atoi(operations[0])
	return res
}
func splitOps(line string) []string {
	var mult []string
	var operations []string
	toks := strings.Split(line, "*")
	for i, tok := range toks {
		mult = append(mult, tok)
		if i < (len(toks) - 1) {
			mult = append(mult, "*")
		}
	}
	for _, m := range mult {
		if strings.Contains(m, "+") {
			plus := strings.Split(m, "+")
			for i, p := range plus {
				operations = append(operations, p)
				if i < (len(plus) - 1) {
					operations = append(operations, "+")
				}
			}
		} else {
			operations = append(operations, m)
		}
	}
	return operations
}

func hasOp(operations []string, op string) int {
	for i, o := range operations {
		if o == op {
			return i
		}
	}
	return -1
}

// Day 2 --^^

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
