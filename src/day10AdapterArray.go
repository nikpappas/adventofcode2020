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
	perms, _ :=countPermutations(ints, [][]int{}, make(map[string](bool)))
	return len(perms)+1
}

func countPermutations(ints []int, perms [][]int, hashes map[string](bool), ) ([][]int, map[string](bool)){
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
			permHash:= hash(newInts)
			if !hashes[permHash]{
				perms  = append(perms, newInts)
				hashes[permHash]=true
				perms, hashes = countPermutations(newInts, perms, hashes)
			}
			// fmt.Println("up a level")
		}
	}
	return perms, hashes
}

func hash(perm[]int) string{
	toRet:= string(perm[0])
	for _,n := range perm{
		toRet+=string(n)+","
	}
	return toRet
}