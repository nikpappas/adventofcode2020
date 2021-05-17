package src

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDay13(t *testing.T) {
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
	assert(1068781, minHash, t)

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
	assert(3417, minHash, t)

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
	assert(754018, minHash, t)

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
	assert(779210, minHash, t)

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
	assert(1261476, minHash, t)

}
