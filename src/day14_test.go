package src

import (
	"fmt"
	"testing"
)

func TestDay14(t *testing.T) {
	bit := "10"
	mask := "0101"
	masked := applyMaskAddresses(bit, mask)
	fmt.Println(bit)
	fmt.Println(mask)
	fmt.Println(masked)

	mask = "0X1"
	fmt.Println(bit)
	fmt.Println(mask)
	masked = applyMaskAddresses(bit, mask)
	fmt.Println(masked)

	mask = "0X01"
	fmt.Println(bit)
	fmt.Println(mask)
	masked = applyMaskAddresses(bit, mask)
	fmt.Println(masked)

	mask = "XX01"
	fmt.Println(bit)
	fmt.Println(mask)
	masked = applyMaskAddresses(bit, mask)
	fmt.Println(masked)

}
