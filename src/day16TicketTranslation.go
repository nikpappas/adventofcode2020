package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"./files"
)

type Rule struct {
	name string
	r1   []int
	r2   []int
}

func day16() {
	lines := files.ReadLines("inputs/day16.txt")
	rules, ticket, tickets := parseInput(lines)

	sort.Slice(rules, func(i int, j int) bool {
		return rules[i].r1[0] < rules[j].r1[0]
	})

	errorRate, potentiallyValidTickets := day16sol1(rules, tickets)
	fmt.Println(errorRate)
	fmt.Println("valid tickets:", len(potentiallyValidTickets))
	prod := day16sol2(rules, potentiallyValidTickets, ticket)
	fmt.Println(prod)

}

func day16sol2(rules []Rule, validTickets [][]int, ticket []int) int {
	finalTicketMappings := getTicketValMappings(rules, validTickets)
	prod := 1
	for i, val := range ticket {
		if strings.HasPrefix(finalTicketMappings[i], "departure") {
			prod *= val
		}
	}

	return prod
}
func day16sol1(rules []Rule, tickets [][]int) (int, [][]int) {
	sum := 0
	var validTickets [][]int
	for _, ticket := range tickets {
		isValid := true
		var validTicket []int
		for _, val := range ticket {
			validTicket = append(validTicket, val)
			if !isWithinAValidRange(rules, val) {
				sum += val
				isValid = false
			}
		}

		if isValid {
			validTickets = append(validTickets, validTicket)
		}

	}
	return sum, validTickets

}

func getTicketValMappings(rules []Rule, validTickets [][]int) map[int](string) {
	histo := getRulesValidHisto(rules, validTickets)

	finalTicketMappings := make(map[int](string))
	prevCount := 10000000
	for len(histo) > 0 {
		min := len(validTickets) + 1
		minRuleSeq := -1
		ticketSeq := -0
		for k, v := range histo {
			for ruleSeq, vla := range v {
				if min > vla {
					min = vla
					minRuleSeq = ruleSeq
					ticketSeq = k
				}
			}
		}

		if min < len(validTickets) {
			delete(histo[ticketSeq], minRuleSeq)
		}
		count := countElements(histo)
		if count == prevCount {
			break
		} else {
			prevCount = count
		}
	}
	for len(histo) > 0 {
		key := -1
		toDel := -1
		for i, v := range histo {
			if len(v) == 1 {
				toDel = i
				key = getKey(v)
				finalTicketMappings[i] = rules[key].name
			}
		}
		for _, v := range histo {
			delete(v, key)
		}
		delete(histo, toDel)
	}

	fmt.Println(finalTicketMappings)
	return finalTicketMappings

}

func isWithinAValidRange(ranges []Rule, val int) bool {
	for _, r := range ranges {
		if conformsToRule(r, val) {
			return true
		}
	}
	return false
}

func conformsToRule(rule Rule, val int) bool {
	return (val >= rule.r1[0] && val <= rule.r1[1]) || (val >= rule.r2[0] && val <= rule.r2[1])
}
func parseInput(lines []string) ([]Rule, []int, [][]int) {
	var rules []Rule
	var tickets [][]int
	i := 0
	for lines[i] != "" {
		ruleTokens := strings.Split(lines[i], ":")
		name := strings.TrimSpace(ruleTokens[0])
		rangeStr := strings.TrimSpace(ruleTokens[1])
		limits := strings.Split(rangeStr, " or ")
		toks1 := strings.Split(limits[0], "-")
		min1, _ := strconv.Atoi(toks1[0])
		max1, _ := strconv.Atoi(toks1[1])
		toks2 := strings.Split(limits[1], "-")
		min2, _ := strconv.Atoi(toks2[0])
		max2, _ := strconv.Atoi(toks2[1])

		rules = append(rules, Rule{name, []int{min1, max1}, []int{min2, max2}})
		i++
	}

	i += 2
	ticket := parseTicket(lines[i])

	fmt.Println(i, "myTicket", lines[i])
	i += 3
	for i < len(lines) {
		tickets = append(tickets, parseTicket(lines[i]))
		i++
	}

	fmt.Println("tickets:", len(tickets))
	return rules, ticket, tickets

}

func parseTicket(ticketStr string) []int {
	var ticket []int
	for _, v := range strings.Split(ticketStr, ",") {
		val, _ := strconv.Atoi(v)
		ticket = append(ticket, val)
	}
	return ticket
}

func inncrementHisto(histo map[int](map[int](int)), i int, j int) {
	_, ok := histo[i]
	if ok {
		histo[i][j] = histo[i][j] + 1
	} else {
		histo[i] = make(map[int](int))
		histo[i][j] = 1
	}

}
func countElements(histo map[int](map[int](int))) int {
	count := 0
	for _, v := range histo {
		count += len(v)
	}
	return count

}

func getKey(mapVar map[int](int)) int {
	for k, _ := range mapVar {
		return k
	}
	return -1
}

func getRulesValidHisto(rules []Rule, validTickets [][]int) map[int](map[int](int)) {
	histo := make(map[int](map[int](int)))
	for _, vt := range validTickets {
		for i, tval := range vt {
			for j, rule := range rules {
				if conformsToRule(rule, tval) {
					inncrementHisto(histo, i, j)
				}
			}
		}
	}
	return histo
}
