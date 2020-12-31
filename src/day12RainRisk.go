package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"./files"
)

const (
	E = iota // == 0
	S = iota // == 1
	W = iota // == 2
	N = iota // == 3
)

var directionLookup = [...]string{"E", "S", "W", "N"}

type NavInstruction struct {
	op   string
	dist int
}

type Ship struct {
	x   int
	y   int
	dir int
}

type WayPoint struct {
	x int
	y int
}

type ShipWithWayPoint struct {
	x  int
	y  int
	wp WayPoint
}

func day12() {
	fmt.Println("Day 12.")

	lines := files.ReadLines("inputs/day12.txt")
	instructions := parseNavInstructions(lines)

	fmt.Println("1. =================")
	x, y := day12sol1(instructions)
	fmt.Println(x, y)
	fmt.Println(getManhatan(x, y))
	fmt.Println("2. =================")
	x, y = day12sol2(instructions)
	fmt.Println(x, y)
	fmt.Println(getManhatan(x, y))
}
func day12sol1(instructions []NavInstruction) (int, int) {
	var ship Ship
	for _, inst := range instructions {
		switch inst.op {
		case "L":
			dirJumps := inst.dist / 90
			ship.dir = (4 + ship.dir - dirJumps) % 4
			continue
		case "R":
			dirJumps := inst.dist / 90
			ship.dir = (ship.dir + dirJumps) % 4
			continue
		case "F":
			inst.op = directionLookup[ship.dir]

		}
		switch inst.op {
		case "E":
			ship.x += inst.dist
			break
		case "S":
			ship.y += inst.dist
			break
		case "W":
			ship.x -= inst.dist
			break
		case "N":
			ship.y -= inst.dist
			break
		}
	}
	return ship.x, ship.y

}

func day12sol2(instructions []NavInstruction) (int, int) {
	var ship ShipWithWayPoint
	ship.wp.x = 10
	ship.wp.y = -1
	var err error
	for _, inst := range instructions {
		switch inst.op {
		case "L":
			dirJumps := inst.dist / 90
			for i := 0; i < dirJumps; i++ {
				ship, err = turnLeft(ship)
			}
			break
		case "R":
			dirJumps := inst.dist / 90
			for i := 0; i < dirJumps; i++ {
				ship, err = turnRight(ship)
			}
			break
		case "F":
			ship.x += inst.dist * ship.wp.x
			ship.y += inst.dist * ship.wp.y
			break

		case "E":
			ship.wp.x += inst.dist
			break
		case "S":
			ship.wp.y += inst.dist
			break
		case "W":
			ship.wp.x -= inst.dist
			break
		case "N":
			ship.wp.y -= inst.dist
			break
		}
		fmt.Println(inst)
		fmt.Println(ship.x, ship.y, "->", ship.wp.x, ship.wp.y)
		if err != nil {
			return -1, -1
		}
	}
	return ship.x, ship.y

}

func getManhatan(x int, y int) int {
	sum := 0
	if x > 0 {
		sum += x
	} else {
		sum -= x
	}
	if y > 0 {
		sum += y
	} else {
		sum -= y
	}
	return sum
}

func turnLeft(ship ShipWithWayPoint) (ShipWithWayPoint, error) {
	return turn90(ship, -1)
}
func turnRight(ship ShipWithWayPoint) (ShipWithWayPoint, error) {
	return turn90(ship, 1)
}
func turn90(ship ShipWithWayPoint, sig int) (ShipWithWayPoint, error) {
	var newWP WayPoint
	x := ship.wp.x
	y := ship.wp.y

	quarter, err := findQuarter(x, y)
	fmt.Println(x, y, quarter)
	xSig, ySig := sigsForQuarters((4 + quarter + sig) % 4)

	newWP.x = xSig * int(math.Abs(float64(y)))
	newWP.y = ySig * int(math.Abs(float64(x)))

	ship.wp = newWP
	return ship, err
}

func day12test() {
	var ship ShipWithWayPoint
	ship.wp.x = 10
	ship.wp.y = -1
	var err error

	fmt.Println("1. test")
	ship, err = turnLeft(ship)
	assert(-1, ship.wp.x)
	assert(-10, ship.wp.y)
	fmt.Println()

	fmt.Println("2. test")
	ship, err = turnLeft(ship)
	assert(-10, ship.wp.x)
	assert(1, ship.wp.y)
	fmt.Println()

	fmt.Println("3. test")
	ship, err = turnLeft(ship)
	assert(1, ship.wp.x)
	assert(10, ship.wp.y)
	fmt.Println()

	fmt.Println("4. test")
	ship, err = turnLeft(ship)
	assert(10, ship.wp.x)
	assert(-1, ship.wp.y)
	fmt.Println()

	fmt.Println("1. test")
	ship, err = turnRight(ship)
	assert(1, ship.wp.x)
	assert(10, ship.wp.y)
	fmt.Println()

	fmt.Println("2. test")
	ship, err = turnRight(ship)
	assert(-10, ship.wp.x)
	assert(1, ship.wp.y)
	fmt.Println()

	fmt.Println("3. test")
	ship, err = turnRight(ship)
	assert(-1, ship.wp.x)
	assert(-10, ship.wp.y)
	fmt.Println()

	fmt.Println("4. test")
	ship, err = turnRight(ship)
	assert(10, ship.wp.x)
	assert(-1, ship.wp.y)
	fmt.Println()

	fmt.Println("1a. test")
	ship.wp.x = 0
	ship, err = turnLeft(ship)
	assert(-1, ship.wp.x)
	assert(0, ship.wp.y)
	fmt.Println()

	fmt.Println(err)
	lines := files.ReadLines("inputs/day12test.txt")
	instructions := parseNavInstructions(lines)
	fmt.Println("1. =================")
	x, y := day12sol1(instructions)
	fmt.Println(x, y)
	fmt.Println(x + y)
	fmt.Println("2. =================")
	x, y = day12sol2(instructions)
	fmt.Println(x, y)
	fmt.Println(x + y)

}

func findQuarter(x int, y int) (int, error) {
	if x >= 0 && y <= 0 {
		return 0, nil
	}
	if x >= 0 && y >= 0 {
		return 1, nil
	}
	if x <= 0 && y >= 0 {
		return 2, nil
	}
	if x <= 0 && y <= 0 {
		return 3, nil
	}

	return -1, errors.New("cant find quarter")
}

func sigsForQuarters(quarter int) (int, int) {
	switch quarter {
	case 0:
		return 1, -1
	case 1:
		return 1, 1
	case 2:
		return -1, 1
	case 3:
		return -1, -1

	}
	return -100, -100

}

func parseNavInstructions(lines []string) []NavInstruction {
	var instructions []NavInstruction
	for _, line := range lines {
		n, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, NavInstruction{line[:1], n})
	}

	return instructions

}
