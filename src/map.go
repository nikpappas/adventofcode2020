package main

import (
	"fmt"
	"strconv"
)

func mapLinesToInts(lines []string) []int {
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
