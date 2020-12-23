package main

import (
	"fmt"
	"strings"
)

func day3() {
	fmt.Println("day3")
	lines := readLines("inputs/day3.txt")
	board := mapLinesToBoard(lines)

	day3Sol1(board, 3, 1)
	steps := [][2]int{
		[2]int{1, 1},
		[2]int{3, 1},
		[2]int{5, 1},
		[2]int{7, 1},
		[2]int{1, 2}}

	day3Sol2(board, steps)
}

func day3Sol1(board [][]string, hStep int, vStep int) int {
	fmt.Println("Solution 1")
	sum := 0
	cursor := 0
	for i := 0; i < len(board); i += vStep {
		fmt.Println(board[i])
		fmt.Print(strings.Repeat(" ", cursor*2+1))
		if board[i][cursor] == "#" {
			fmt.Println("X")
			sum += 1
		} else {
			fmt.Println("0")
		}
		cursor = (cursor + hStep) % len(board[i])
	}
	fmt.Println(sum)
	return sum

}

func day3Sol2(board [][]string, steps [][2]int) int {
	product := 1
	for i := 0; i < len(steps); i++ {
		product *= day3Sol1(board, steps[i][0], steps[i][1])
	}
	fmt.Println(product)
	return product
}
