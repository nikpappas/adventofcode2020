package src

import (
	"fmt"
	"strconv"
	"strings"

	"nikpappas.com/adventofcode2021/files"
)

func Day18() {
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
