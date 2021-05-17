package src

import (
	"fmt"
	"testing"
)

func TestDay18(t *testing.T) {

	day18sol1([]string{"1 + 2 + 3"})
	day18sol1([]string{"2 * 2 + 3"})
	day18sol1([]string{"1 + 2 * 3"})
	day18sol1([]string{"1 + 2 * (3 + 1)"})
	day18sol1([]string{"1 + 2 * ((3 + 1))"})
	day18sol1([]string{"9 + (2 * ((3 + 1)))"})

	day18sol1([]string{"2 * 3 + (4 * 5)"})
	day18sol1([]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"})
	day18sol1([]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"})
	day18sol1([]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"})
	// 2 * 3 + (4 * 5) becomes 26.
	// 5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
	// 5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
	// ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.
	fmt.Println("Testign solution 2")
	day18sol2([]string{"1 + (2 * 3) + (4 * (5 + 6))"})
	day18sol2([]string{"2 * 3 + (4 * 5)"})
	day18sol2([]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"})
	day18sol2([]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"})
	day18sol2([]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"})

	// 1 + (2 * 3) + (4 * (5 + 6)) still becomes 51.
	// 2 * 3 + (4 * 5) becomes 46.
	// 5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 1445.
	// 5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 669060.
	// ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 23340.
	day18sol2([]string{"1 + 2 + 3"})
	day18sol2([]string{"2 * 2 + 3"})
	day18sol2([]string{"3+2 * 2"})
	day18sol2([]string{"3+2 * 2+1"})

}
