package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BusDetails struct {
	id    int
	delay int
}

func day13() {
	lines := readLines("inputs/day13.txt")
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

func day13test() {
	minHash := day13sol1([]int{7, 13, 59, 31, 19}, 939)
	fmt.Println(minHash)

	deets := []string{"7", "13", "x", "x", "59", "x", "31", "19"}
	var deetss []BusDetails
	for i, d := range deets {
		if d != "x" {
			toi, _ := strconv.Atoi(d)
			deetss = append(deetss, BusDetails{toi, i})
		}

	}
	minHash = day13sol2(deetss)
	fmt.Println(minHash)
	assert(1068781, minHash)

	deets = []string{"17", "x", "13", "19"}
	deetss = []BusDetails{}
	for i, d := range deets {
		if d != "x" {
			toi, _ := strconv.Atoi(d)
			deetss = append(deetss, BusDetails{toi, i})
		}

	}
	minHash = day13sol2(deetss)
	fmt.Println(minHash)
	assert(3417, minHash)

	deets = []string{"67", "7", "59", "61"}
	deetss = []BusDetails{}
	for i, d := range deets {
		if d != "x" {
			toi, _ := strconv.Atoi(d)
			deetss = append(deetss, BusDetails{toi, i})
		}

	}
	minHash = day13sol2(deetss)
	fmt.Println(minHash)
	assert(754018, minHash)

	deets = []string{"67", "x", "7", "59", "61"}
	deetss = []BusDetails{}
	for i, d := range deets {
		if d != "x" {
			toi, _ := strconv.Atoi(d)
			deetss = append(deetss, BusDetails{toi, i})
		}
	}
	minHash = day13sol2(deetss)
	fmt.Println(minHash)
	assert(779210, minHash)

	deets = []string{"67", "7", "x", "59", "61"}
	deetss = []BusDetails{}
	for i, d := range deets {
		if d != "x" {
			toi, _ := strconv.Atoi(d)
			deetss = append(deetss, BusDetails{toi, i})
		}
	}
	minHash = day13sol2(deetss)
	fmt.Println(minHash)
	assert(1261476, minHash)

}
