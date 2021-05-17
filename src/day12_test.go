package src

import (
	"fmt"
	"testing"

	"nikpappas.com/adventofcode2021/files"
)

func TestDay12(t *testing.T) {
	var ship ShipWithWayPoint
	ship.wp.x = 10
	ship.wp.y = -1
	var err error

	fmt.Println("1. test")
	ship, err = turnLeft(ship)
	assert(-1, ship.wp.x, t)
	assert(-10, ship.wp.y, t)
	fmt.Println()

	fmt.Println("2. test")
	ship, err = turnLeft(ship)
	assert(-10, ship.wp.x, t)
	assert(1, ship.wp.y, t)
	fmt.Println()

	fmt.Println("3. test")
	ship, err = turnLeft(ship)
	assert(1, ship.wp.x, t)
	assert(10, ship.wp.y, t)
	fmt.Println()

	fmt.Println("4. test")
	ship, err = turnLeft(ship)
	assert(10, ship.wp.x, t)
	assert(-1, ship.wp.y, t)
	fmt.Println()

	fmt.Println("1. test")
	ship, err = turnRight(ship)
	assert(1, ship.wp.x, t)
	assert(10, ship.wp.y, t)
	fmt.Println()

	fmt.Println("2. test")
	ship, err = turnRight(ship)
	assert(-10, ship.wp.x, t)
	assert(1, ship.wp.y, t)
	fmt.Println()

	fmt.Println("3. test")
	ship, err = turnRight(ship)
	assert(-1, ship.wp.x, t)
	assert(-10, ship.wp.y, t)
	fmt.Println()

	fmt.Println("4. test")
	ship, err = turnRight(ship)
	assert(10, ship.wp.x, t)
	assert(-1, ship.wp.y, t)
	fmt.Println()

	fmt.Println("1a. test")
	ship.wp.x = 0
	ship, err = turnLeft(ship)
	assert(-1, ship.wp.x, t)
	assert(0, ship.wp.y, t)
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
