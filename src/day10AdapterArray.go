package main

import (
	"fmt"
	"sort"
	"time"
)

var sec = time.Now().Unix()
func day10(){
	lines := readLines("inputs/day10test.txt")
	
	ints := mapLinesToInts(lines)
	sort.Ints(ints)
	ints = append([]int{0}, ints...)
	ints = append(ints, ints[len(ints)-1]+3)

	prod := day10sol1(ints)
	fmt.Println(prod)
	count := day10sol2(ints)
	fmt.Println(count)
}

func day10sol1(ints []int) int{
	histo:=make(map[int](int))
	for i:=1;i<len(ints);i++{
		histo[ints[i-1]-ints[i]]+=1
		fmt.Println(ints[i-1],ints[i], ints[i-1]-ints[i], histo[ints[i-1]-ints[i]])
	}
	fmt.Println(histo[-1],histo[-3])
	return histo[-1]*histo[-3]

}


func day10sol2(ints []int) int{
	perms :=countPermutations(ints, 0)
	return perms+1
}

func countPermutations(ints []int, count int ) int{
	if len(ints) < 2{
		return count
	}

	for i:=0;i<len(ints)-1;i++{
		diffNext:=ints[i+1] - ints[i-1]
		// fmt.Println("'i'",i, "diff", diffNext)

		if diffNext <=3{
			var newInts []int
			for j:=i+1;j<len(ints);j++{
				if i!=j{
					newInts=append(newInts,ints[j])
				}
			}
			// fmt.Println("newints",newInts)
			// fmt.Println(permHash)
			count = countPermutations(newInts, count+1)
			// fmt.Println("up a level")
		}
	}
	return count
}