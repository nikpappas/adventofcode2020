package src

import (
	"fmt"
	"strconv"
	"strings"

	"nikpappas.com/adventofcode2021/files"
)

type BusDetails struct {
	id    int
	delay int
}

func Day13() {
	fmt.Println("Day 13.")

	lines := files.ReadLines("inputs/day13.txt")
	arrival, _ := strconv.Atoi(lines[0])
	var ids []int
	for _, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			n, _ := strconv.Atoi(id)
			ids = append(ids, n)
		}
	}

	minHash := day13sol1(ids, arrival)
	fmt.Println(minHash)

	fmt.Println("Solution 2")

	var details []BusDetails
	for i, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			n, _ := strconv.Atoi(id)
			details = append(details, BusDetails{n, i})
		}
	}

	timestamp := day13sol2(details)
	fmt.Println(timestamp)

}

func day13sol1(ids []int, arrival int) int {
	minWait := 100000000000
	minId := -1
	for _, id := range ids {
		wait := id - arrival%id
		if minWait > wait {
			minWait = wait
			minId = id
		}
	}
	return minWait * minId

}

func day13sol2(details []BusDetails) int {
	fmt.Println(details)
	timestamp := 0
	maxDelay := 0
	maxId := 0
	prevTimestamp := timestamp
	for _, d := range details {
		if maxId < d.id {
			maxId = d.id
			maxDelay = d.delay
		}

	}
	fmt.Println(timestamp, maxDelay, maxId)
	for timestamp = -maxDelay; !isValidTimestamp(details, timestamp); timestamp += maxId {
		if timestamp-prevTimestamp > 100000000000 {
			fmt.Println(timestamp)
			prevTimestamp = timestamp
		}
	}

	return timestamp

}

func isValidTimestamp(details []BusDetails, timestamp int) bool {
	for _, d := range details {
		if (timestamp+d.delay)%d.id != 0 {
			return false
		}
	}
	return true

}
