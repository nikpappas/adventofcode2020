package main

import (
	"fmt"
	"sort"
	"time"
)

var sec = time.Now().Unix()
func day10(){
	lines := readLines("inputs/day10.txt")
	
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
	return len(countPermutations(ints, [][]int{}))+1
}

func countPermutations(ints []int, perms [][]int) [][]int{
	// fmt.Println("len(ints)",len(ints))
	// fmt.Println("ints",ints)
	lenPerms:=len(perms)
	if lenPerms%1000==0{
		now := time.Now().Unix()
		fmt.Println(lenPerms, now-sec,"s")
		sec =now
	}
	for i:=1;i<len(ints)-1;i++{
		diffNext:=ints[i+1] - ints[i-1]
		// fmt.Println("'i'",i, "diff", diffNext)

		if diffNext <=3{
			var newInts []int
			for j:=0;j<len(ints);j++{
				if i!=j{
					newInts=append(newInts,ints[j])
				}
			}
			// fmt.Println("newints",newInts)
			if !contains(perms, newInts){
				perms  = append(perms, newInts)
				perms = countPermutations(newInts, perms)
			}
			// fmt.Println("up a level")
		}
	}
	return perms
}
func equals(a []int, b[]int)bool{
	if len(a)!= len(b) {return false}
	for i,ai :=range a{
		if ai != b[i] {return false}
	}
	return true
}
func contains(a[][]int, b[]int) bool{
	for _,ai:=range a{
		if equals(ai,b) {return true}
	}
	return false
}