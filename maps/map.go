package maps

import (
	"fmt"
	"strconv"
)

func MapLinesToInts(lines []string) []int {
	var ints []int
	for _, line := range lines {
		if string(line[len(line)-1:]) == "\n" {
			line = line[:len(line)-1]
		}

		var n, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			break
		}
		ints = append(ints, n)
	}
	return ints

}

func MapLinesToBoard(lines []string) [][]string {
	var board [][]string

	for _, line := range lines {
		if string(line[len(line)-1:]) == "\n" {
			line = line[:len(line)-1]
		}
		var toAppend []string
		for _, c := range line {
			toAppend = append(toAppend, string(c))
		}
		board = append(board, toAppend)

	}
	return board
}
