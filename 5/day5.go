package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	Add int = iota + 1
	Multiply
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThen
	Equals
	Exit = 99
)

type Program []int

func (p Program) get(address, mode int) int {
	value := p[address]
	if mode%10 == 1 {
		return value
	} else {
		return p[value]
	}
}
func (p Program) set(address, val int) {
	p[address] = val
}

func readProgram(reader io.Reader) Program {
	program := Program{}
	for {
		var opcode int
		_, err := fmt.Fscanf(reader, "%d", &opcode)
		if err != nil {
			fmt.Println(err)
			break
		}
		program = append(program, opcode)
	}
	return program
}

func executeProgram(program Program) {
	pc := 0
	length := len(program)
	for pc < length {
		instruction := program[pc]
		// fmt.Println(program[pc:])
		opcode := instruction % 100
		fmt.Println(opcode)
		switch opcode {
		case Add:
			//fmt.Println("add", pc)
			program.set(program[pc+3], program.get(pc+1, instruction/100)+program.get(pc+2, instruction/1000))
			pc += 4
		case Multiply:
			// fmt.Println("multiply", pc)
			program.set(program[pc+3], program.get(pc+1, instruction/100)*program.get(pc+2, instruction/1000))
			pc += 4
		case Input:
			fmt.Printf("Input: ")
			var val int
			fmt.Scanf("%d\n", &val)
			program.set(program[pc+1], val)
			pc += 2
		case Output:
			fmt.Printf("Output: %d\n", program.get(pc+1, instruction/100))
			pc += 2
		case JumpIfTrue:
			if program.get(pc+1, instruction/100) != 0 {
				pc = program.get(pc+2, instruction/1000)
			} else {
				pc += 3
			}
		case JumpIfFalse:
			if program.get(pc+1, instruction/100) == 0 {
				pc = program.get(pc+2, instruction/1000)
			} else {
				pc += 3
			}
		case LessThen:
			val := 0
			if program.get(pc+1, instruction/100) < program.get(pc+2, instruction/1000) {
				val = 1
			}
			program.set(program[pc+3], val)
			pc += 4
		case Equals:
			val := 0
			if program.get(pc+1, instruction/100) == program.get(pc+2, instruction/1000) {
				val = 1
			}
			program.set(program[pc+3], val)
			pc += 4
		default:
			log.Fatal("Unknown instruction ", instruction)
			break
		case Exit:
			//fmt.Println("exit", pc)
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing command input")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	program := readProgram(f)
	executeProgram(program)
}
