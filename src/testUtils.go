package src

import (
	"fmt"
	"testing"
)

func assert(expected int, res int, t *testing.T) {
	if res != expected {
		fmt.Println("Error", "!!!!", "!!!!", "!!!!", "!!!!")
		fmt.Println("Error", res, "!=", expected)
		fmt.Println("Error", "!!!!", "!!!!", "!!!!", "!!!!")
	}
	t.Error("Error", res, "!=", expected)
}

func startTest(testNum int) {
	fmt.Println("============", testNum, "============")
	fmt.Println()
}
