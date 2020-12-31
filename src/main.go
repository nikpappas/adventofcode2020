package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	day, _ := strconv.Atoi(os.Args[1])

	switch day {
	case 0:
		test()
		break
	case 1:
		day1()
		break
	case 2:
		day2()
		break
	case 3:
		day3()
		break
	case 4:
		day4()
		break
	case 5:
		day5()
		break
	case 6:
		day6()
		break
	case 7:
		day7()
		break
	case 8:
		day8()
		break
	case 9:
		day9()
		break
	case 10:
		day10()
		break
	case 11:
		day11()
		break
	case 12:
		day12()
		break
	case 13:
		day13()
		break
	case 14:
		day14()
		break
	}
}

func test() {
	fmt.Println("Testing")
	day10test()
	day12test()
	day13test()
	day14test()
}
