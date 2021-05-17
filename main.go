package main

import (
	"fmt"
	"os"
	"strconv"

	"nikpappas.com/adventofcode2021/src"
)

func main() {
	day, _ := strconv.Atoi(os.Args[1])

	switch day {
	case 0:
		fmt.Println("For tests run ./test.sh")
		break
	case 1:
		src.Day1()
		break
	case 2:
		src.Day2()
		break
	case 3:
		src.Day3()
		break
	case 4:
		src.Day4()
		break
	case 5:
		src.Day5()
		break
	case 6:
		src.Day6()
		break
	case 7:
		src.Day7()
		break
	case 8:
		src.Day8()
		break
	case 9:
		src.Day9()
		break
	case 10:
		src.Day10()
		break
	case 11:
		src.Day11()
		break
	case 12:
		src.Day12()
		break
	case 13:
		src.Day13()
		break
	case 14:
		src.Day14()
		break
	case 15:
		src.Day15()
		break
	case 16:
		src.Day16()
		break
	case 17:
		src.Day17()
		break
	case 18:
		src.Day18()
		break
	case 19:
		src.Day19()
		break
	default:
		fmt.Println("No day ", day)
	}

}
