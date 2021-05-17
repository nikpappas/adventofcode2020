package src

import (
	"fmt"
	"strconv"
	"strings"

	"nikpappas.com/adventofcode2021/files"
)

func Day2() {
	fmt.Println("Day 2.")
	lines := files.ReadLines("inputs/day2.txt")

	day2Sol1(lines)
	day2Sol2(lines)
}

func day2Sol1(lines []string) {
	fmt.Println("solution1")
	count := 0
	for _, line := range lines {
		fmt.Println(line)
		pass := getPasword(line)
		limits, c := getPolicy(line)
		occur := getCharCount(pass, c)
		fmt.Println(pass)
		fmt.Println(limits, c)
		fmt.Println(occur)
		if limits[0] <= occur && occur <= limits[1] {
			count += 1
			fmt.Println("valid")
		} else {
			fmt.Println("invalid")

		}

	}
	fmt.Println(count)
}

func day2Sol2(lines []string) {
	fmt.Println("solution2")
	count := 0
	for _, line := range lines {
		fmt.Println(line)
		pass := getPasword(line)
		places, c := getPolicy(line)
		if charAt(pass, places[0]-1) == c && charAt(pass, places[1]-1) != c ||
			charAt(pass, places[1]-1) == c && charAt(pass, places[0]-1) != c {
			count += 1
			fmt.Println("valid")
		} else {
			fmt.Println("invalid")

		}

	}
	fmt.Println(count)
}

func getCharCount(pass string, c string) int {
	count := 0
	for _, character := range pass {
		if c == string(character) {
			count += 1
		}
	}
	return count
}

func charAt(s string, i int) string {
	return string(s[i])
}

func getPolicy(line string) ([2]int, string) {
	policy := strings.Split(line, ":")[0]
	policyToks := strings.Split(strings.TrimSpace(policy), " ")
	rangeToks := strings.Split(policyToks[0], "-")
	min, _ := strconv.Atoi(rangeToks[0])
	max, _ := strconv.Atoi(rangeToks[1])
	charRange := [2]int{min, max}
	return charRange, policyToks[1]
}

func getPasword(line string) string {
	return strings.TrimSpace(strings.Split(line, ":")[1])
}
