package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bagLimitation struct {
	colour string
	num    int
}

func day7() {
	lines := readLines("inputs/day7.txt")
	sum := day7sol1(lines)
	fmt.Println(sum)
	sum2 := day7sol2(lines)
	fmt.Println(sum2)

}

func day7sol1(lines []string) int {

	resColours := make(map[string](bool))
	initColour := []string{"shiny gold"}
	resColours = canContainColour(lines, initColour, resColours)

	return len(resColours) - 1

}
func day7sol2(lines []string) int {
	fmt.Println("Solution 2")

	limitations := compileLimitations(lines)

	initColour := []bagLimitation{{"shiny gold", 1}}
	resColours := mustContainColour(limitations, initColour, 0)

	return resColours - 1

}

func canContainColour(lines []string, colours []string, resColours map[string](bool)) map[string](bool) {
	if len(colours) == 0 {
		return resColours
	}
	var outerColours []string
	for _, colour := range colours {
		resColours[colour] = true
		for _, l := range lines {
			limitations := strings.Split(l, "contain")
			if strings.Contains(limitations[1], colour) {
				outerColours = append(outerColours, strings.TrimSpace(strings.ReplaceAll(limitations[0], "bags", "")))
			}
		}

	}

	return canContainColour(lines, outerColours, resColours)

}

func mustContainColour(limitations map[string](string), colours []bagLimitation, res int) int {
	if len(colours) == 0 {
		return res
	}
	var innerColours []bagLimitation

	fmt.Println(colours)
	for _, colour := range colours {
		fmt.Println("colour", colour)
		fmt.Println("limitations", limitations[colour.colour])
		var toks []string
		if strings.Contains(limitations[colour.colour], ",") {
			toks = strings.Split(limitations[colour.colour], ",")
		} else {
			toks = []string{limitations[colour.colour]}
		}

		for _, tok := range toks {
			tok = strings.TrimSpace(tok)
			innerColour := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tok, "bags", ""), "bag", ""), ".", ""))

			if !strings.Contains(innerColour, "no other") {
				num, _ := strconv.Atoi(innerColour[:1])
				bagLim := bagLimitation{innerColour[2:], num * colour.num}
				innerColours = append(innerColours, bagLim)
			}

		}
		res += colour.num
		fmt.Println(colour.num, res)

	}
	return mustContainColour(limitations, innerColours, res)
}

func compileLimitations(lines []string) map[string](string) {
	limitations := make(map[string](string))
	for _, line := range lines {
		toks := strings.Split(line, "contain")
		key := strings.TrimSpace(strings.ReplaceAll(toks[0], "bags", ""))
		colours := strings.TrimSpace(toks[1])
		limitations[key] = colours

	}
	return limitations

}
