package src

import (
	"fmt"
	"strings"

	"nikpappas.com/adventofcode2021/files"
)

func Day6() {
	fmt.Println("Day 6.")

	lines := files.ReadLines("inputs/day6.txt")
	groups := splitGroups(lines)
	sum := day6Sol1(groups)
	fmt.Println(sum)
	sum2 := day6Sol2(groups)
	fmt.Println(sum2)
}
func day6Sol1(groups [][]string) int {
	var yesSets []map[rune](bool)
	for _, group := range groups {
		yesSet := anyoneAnsweredYes(group)
		yesSets = append(yesSets, yesSet)
	}
	sum := sumOfYeses(yesSets)
	return sum
}
func day6Sol2(groups [][]string) int {
	var yesSets []map[rune](bool)
	for _, group := range groups {
		yesSet := allAnsweredYes(group)
		yesSets = append(yesSets, yesSet)
	}
	sum := sumOfYeses(yesSets)
	return sum
}

func anyoneAnsweredYes(group []string) map[rune](bool) {
	yes := make(map[rune](bool))
	for _, l := range group {
		for _, c := range l {
			yes[c] = true
		}
	}
	return yes
}

func allAnsweredYes(group []string) map[rune](bool) {
	yes := make(map[rune](bool))
	for _, l := range group {
		for _, c := range l {
			yes[c] = true
		}
	}
	for k, _ := range yes {
		yes[k] = allInGroupAnseredYes(group, k)
	}
	return yes
}

func sumOfYeses(yesSets []map[rune](bool)) int {
	sum := 0
	for _, yesSet := range yesSets {
		for _, yes := range yesSet {
			if yes {
				sum += 1
			}
		}
	}
	return sum

}
func splitGroups(lines []string) [][]string {
	var groups [][]string
	var group []string
	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = []string{}
		} else {
			group = append(group, line)
		}
	}
	groups = append(groups, group)
	return groups
}

func allInGroupAnseredYes(group []string, c rune) bool {
	for _, l := range group {

		if strings.IndexRune(l, c) == -1 {
			return false
		}
	}
	return true
}
