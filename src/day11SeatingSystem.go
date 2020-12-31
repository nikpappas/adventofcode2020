package main

import (
	"fmt"

	"./files"
	"./maps"
)

func day11() {
	fmt.Println("Day 11.")

	lines := files.ReadLines("inputs/day11.txt")
	board := maps.MapLinesToBoard(lines)
	occupied := 0
	occupied = day11sol1(clone(board))
	fmt.Println("Seats occupied:", occupied)
	occupied = day11sol2(clone(board))
	fmt.Println("Seats occupied:", occupied)

}
func day11sol1(board [][]string) int {
	var buffer [][]string
	var occupied int
	for ok := true; ok; ok = !equals2D(buffer, board) {
		buffer = clone(board)
		for i, line := range buffer {
			for j, _ := range line {
				if buffer[i][j] == "." {
					continue
				}
				occup := countAdjasent(i, j, buffer)

				if occup >= 4 {
					board[i][j] = "L"
				} else if occup == 0 {
					board[i][j] = "#"
				}
			}
		}
		occupied = countTotOccupied(board)
	}
	return occupied

}

func day11sol2(board [][]string) int {
	var buffer [][]string
	var occupied int
	for ok := true; ok; ok = !equals2D(buffer, board) {
		buffer = clone(board)
		for i, line := range buffer {
			for j, _ := range line {
				if buffer[i][j] == "." {
					continue
				}
				occup := countVisible(i, j, buffer)

				if occup >= 5 {
					board[i][j] = "L"
				} else if occup == 0 {
					board[i][j] = "#"
				}
			}
		}
		occupied = countTotOccupied(board)
	}
	return occupied

}

func countTotOccupied(board [][]string) int {
	occupied := 0
	for _, line := range board {
		for _, seat := range line {
			if seat == "#" {
				occupied += 1
			}
		}
	}
	return occupied
}

func countVisible(i int, j int, board [][]string) int {
	sum := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if !(x == 0 && y == 0) {
				if hasVisible(i, j, board, x, y) {
					sum += 1
				}
			}
		}
	}
	return sum

}

func hasVisible(i int, j int, board [][]string, offsetX int, offsetY int) bool {
	x := i + offsetX
	y := j + offsetY
	if !isInBoard(x, y, board) {
		return false
	}
	if board[x][y] == "L" {
		return false
	}

	if board[x][y] == "#" {
		return true
	}
	return hasVisible(x, y, board, offsetX, offsetY)

}

func countAdjasent(i int, j int, board [][]string) int {
	sum := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if !(x == 0 && y == 0) {
				if isOccupied(i+x, j+y, board) {
					sum += 1
				}
			}
		}
	}
	return sum

}
func isInBoard(i int, j int, board [][]string) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= len(board) || j >= len(board[i]) {
		return false
	}
	return true
}
func isOccupied(i int, j int, board [][]string) bool {
	if !isInBoard(i, j, board) {
		return false
	}

	if board[i][j] != "#" {
		return false
	} else {
		return true
	}

}

func clone(a [][]string) [][]string {
	var toRet [][]string
	for _, ai := range a {
		var line []string
		for _, i := range ai {
			line = append(line, i)
		}
		toRet = append(toRet, line)
	}
	return toRet
}

func equals2D(a [][]string, b [][]string) bool {
	for i, ai := range a {
		if !equals1D(ai, b[i]) {
			return false
		}
	}
	return true
}
func equals1D(a []string, b []string) bool {
	for i, ai := range a {
		if b[i] != ai {
			return false
		}
	}
	return true
}
