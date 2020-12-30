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
	lcmInt := 1
	buffer := 0
	lcm := make(map[int](int))
	for _, d := range details {
		lcm[d.id] = d.delay
		if lcm[d.id] > buffer {
			buffer = lcm[d.id]
		}
		lcmInt *= d.id
	}
	it := 0
	for solved := false; !solved; solved = allEqual(lcm) {
		for _, det := range details {
			if lcm[det.id] < buffer {
				if (buffer-lcm[det.id])%det.id == 0 {
					lcm[det.id] = buffer
				} else {
					lcm[det.id] = buffer + det.id - (buffer-lcm[det.id])%det.id
					buffer = lcm[det.id]
				}
				it++
				if it%10000000 == 0 {
					fmt.Println(buffer)
				}
			}
		}
	}
	return lcmInt - buffer

	// prod := 1
	// for _, id := range idsTahtMatter {
	// 	prod *= id
	// }
	// return prod

	// positions := make(map[int](int))
	// for _, det := range details {
	// 	positions[det.id] = det.delay
	// }
	// fmt.Println(positions)
	// for solved := false; !solved; {
	// 	for _, det := range details {
	// 		fmt.Println(det)
	// 		solved = true

	// 	}
	// }
	// return 0

}

func allEqual(times map[int]int) bool {
	var any int
	for _, time := range times {
		any = time
		break
	}

	for _, time := range times {
		if any != time {
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

}
