package main

import (
	"strings"
	"strconv"
	"fmt"
)


type Instruction struct{
	op string
	operand int
}

type State struct{
	acc int
	add int
}

func day8(){
	lines := readLines("inputs/day8.txt")
	var instructions []Instruction
	for _,line := range lines{
		toks := strings.Split(line, " ")
		operand,_:= strconv.Atoi(toks[1])
		instruction := Instruction{toks[0], operand}
		instructions = append(instructions, instruction)
	}

	// day8sol1(instructions)
	badInstruction := day8sol2(instructions)
	fmt.Println("The instruction that goes bang is ", badInstruction)

}

func day8sol2(instructions []Instruction) int{
	for i,_ := range instructions{
		var tweagedInst []Instruction
		for j,inst:= range instructions{
			if j==i{
				switch inst.op{
					case "nop":
						tweagedInst = append(tweagedInst, Instruction{"jmp", inst.operand})
						break;
					case "jmp":
						tweagedInst = append(tweagedInst, Instruction{"nop", inst.operand})
						break;
					case "acc":
						break;
				}
			}else {
				tweagedInst = append(tweagedInst, inst)
			}
		}
		
		fmt.Println (i)
		state,isFixed := day8sol1(tweagedInst)
		fmt.Println (state, isFixed)
		if isFixed{
			return i
		} 
	}
	return -1
}

func day8sol1(instructions []Instruction) (State, bool){
	state := State{0,0}
	addressesVisited:=make(map[int](bool))

	for  0 <= state.add && state.add < len(instructions) {
		_,visited := addressesVisited[state.add]
		if visited{
			fmt.Println("Loop detected")
			return state, false
		}
		addressesVisited[state.add] =true
		state = operate(instructions[state.add], state)

	}
	if state.add == len(instructions){
		return state, true
	}
	return state, false
}

func operate(instruction Instruction, state State) State{
	switch instruction.op {
		case "acc":
			state.acc +=  instruction.operand;
		case "jmp":
			state.add += instruction.operand;
			return state
			
		}
		state.add +=1
		return state
}