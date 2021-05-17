package src

import (
	"fmt"
	"strconv"
	"strings"

	"nikpappas.com/adventofcode2021/files"
)

func Day19() {
	lines := files.ReadLines("inputs/day19.txt")
	var rulesStr []string
	var messages []string
	i := 0
	for i < len(lines) && lines[i] != "" {
		rulesStr = append(rulesStr, lines[i])
		i++
	}
	i++
	for i < len(lines) {
		messages = append(messages, lines[i])
		i++
	}
	rules := parseRules(rulesStr)

	validMessages := day19sol1(rules, messages)
	fmt.Println(validMessages)
}

func day19sol1(rules map[int]([]int), messages []string) int {
	fmt.Println(rules)
	fmt.Println(messages)
	return 0
}

func parseRules(lines []string) map[int]([]int) {
	rules := make(map[int]([]int))
	for _, line := range lines {
		toks := strings.Split(line, ":")
		index, _ := strconv.Atoi(toks[0])
		nestedRulesStr := strings.Split(toks[1], " ")
		var nestedRules []int
		for _, i := range nestedRulesStr {
			ruleIndex, _ := strconv.Atoi(i)
			nestedRules = append(nestedRules, ruleIndex)
		}
		rules[index] = nestedRules
	}

	return rules
}
