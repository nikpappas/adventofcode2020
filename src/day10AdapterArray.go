package main

import (
	"fmt"
	"sort"
	"strings"
)
var prevDepth = 0
var hashes = make(map[string](int))
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
	return  countPermutations(ints, 0)
}

func countPermutations(ints []int, depth int) int{
	count:=1
	if len(ints) < 3{
		return count
	}
	intsHash:=hash(ints)
	preCalced, has:=hashes[intsHash]
	if has{
		return preCalced
	}


	for i:=0;i<len(ints)-2;i++{
		diffNext:=ints[i+2] - ints[i]

		if diffNext <=3{
			var newInts []int
			for j:=i;j<len(ints);j++{
				if j!=(i+1){
					newInts=append(newInts,ints[j])
				}
			}

			count +=  countPermutations(newInts, depth+1)
			if depth!=prevDepth{
				fmt.Println(strings.Repeat(" ", depth),"*")
				prevDepth = depth
			}

		}
	}
	hashes[intsHash] = count
	return count
}


func hash(ints []int) string{
	toRet:=""
	for _,i := range ints{
		toRet+= fmt.Sprint(i)+","
	}	
	return toRet
}

func day10test(){
	expected:=0
	res:=0
	startTest(1)
	expected=2
	res=countPermutations([]int{0,1,3}, 0)
	assert(expected, res)

	startTest(2)
	expected=3
	res= countPermutations([]int{0,1,3, 4}, 0)
	assert(expected, res)

	startTest(3)
	expected=8
	res= countPermutations([]int{0,1,4,5,6,7,10,11,12,15,16,19, 22}, 0)
	assert(expected, res)

	startTest(4)
	expected=19208
	toCalc:=[]int{0,28,33,18,42,31,14,46,20,48,47,24,23,49,45,19,38,39,11,1,32,25,35,8,17,7,9,4,2,34,10,3, 52}
	sort.Ints(toCalc)
	res= countPermutations(toCalc, 0)
	assert(expected, res)
}

func assert(expected int, res int){
	if res!=expected{
		fmt.Println("Error","!!!!","!!!!","!!!!","!!!!")
		fmt.Println("Error",res, "!=", expected)
		fmt.Println("Error","!!!!","!!!!","!!!!","!!!!")
	}
} 

func startTest(testNum int){
	fmt.Println("============",testNum,"============")
	fmt.Println()
}