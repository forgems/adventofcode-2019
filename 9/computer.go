package main

import (
	"fmt"
	"io"
	"log"
)

const (
	Add int64 = iota + 1
	Multiply
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThen
	Equals
	AdjustRelativeBase
	Exit = 99
)

type input_func func() int64
type output_func func(int64)

type Computer struct {
	memory            []int64
	pc                int64
	rb                int64
	additional_memory map[int64]int64
	input             input_func
	output            output_func
}

func NewComputer(program []int64) *Computer {
	return &Computer{
		memory:            program,
		additional_memory: make(map[int64]int64),
	}
}

func ReadProgram(reader io.Reader) *Computer {
	program := []int64{}
	for {
		var opcode int64
		_, err := fmt.Fscanf(reader, "%d", &opcode)
		if err != nil {
			fmt.Println(err)
			break
		}
		program = append(program, opcode)
	}
	return NewComputer(program)
}

func (c *Computer) ReadMemory(address int64) int64 {
	if address < 0 {
		log.Fatalf("Invalid address %d", address)
	}
	if address > int64(len(c.memory)) {
		return c.additional_memory[address]
	}
	return c.memory[address]
}

func (c *Computer) Get(address, mode int64) int64 {
	value := c.memory[address]
	switch mode % 10 {
	case 1:
		return value
	case 2:
		return c.ReadMemory(c.rb + value)
	default:
		return c.ReadMemory(value)
	}
}

func (c *Computer) Set(address, val int64) {
	//fmt.Println("Set", address, val)
	if address > int64(len(c.memory)) {
		c.additional_memory[address] = val
	} else {
		c.memory[address] = val
	}
}

func (c *Computer) Param(offset int64) int64 {
	//fmt.Printf("Param %d", offset)
	exp := offset
	mode := c.memory[c.pc] / 10
	for ; exp > 0; exp-- {
		mode /= 10
	}
	//fmt.Printf("[%d]=", mode%10)
	out := int64(0)
	switch mode % 10 {
	case 1:
		out = c.ReadMemory(c.pc + offset)
	case 2:
		out = c.ReadMemory(c.rb + c.ReadMemory(c.pc+offset))
	case 0:
		out = c.ReadMemory(c.ReadMemory(c.pc + offset))
	default:
		log.Fatalf("Invalid mode for param %d for instruction %d", offset, c.memory[c.pc])
	}
	//fmt.Println(out)
	return out
}

func (c *Computer) ParamAddress(offset int64) int64 {
	//fmt.Printf("ParamAddress %d", offset)
	exp := offset
	mode := c.memory[c.pc] / 10
	for ; exp > 0; exp-- {
		mode /= 10
	}
	//fmt.Printf("[%d]=", mode)
	out := int64(0)
	switch mode % 10 {
	case 2:
		out = c.rb + c.ReadMemory(c.pc+offset)
	case 0:
		out = c.ReadMemory(c.pc + offset)
	default:
		log.Fatalf("Invalid mode %d for param %d for instruction %d", mode, offset, c.memory[c.pc])
	}
	//fmt.Println(out)
	return out
}

func (c *Computer) Run() {
	length := int64(len(c.memory))
	for c.pc < length {
		//fmt.Println("pc", c.pc, "rb", c.rb)
		instruction := c.memory[c.pc]
		//fmt.Println("ins", instruction)
		opcode := instruction % 100
		switch opcode {
		case Add:
			//fmt.Println("add", pc)
			c.Set(c.ParamAddress(3), c.Param(1)+c.Param(2))
			c.pc += 4
		case Multiply:
			// fmt.Println("multiply", pc)
			c.Set(c.ParamAddress(3), c.Param(1)*c.Param(2))
			c.pc += 4
		case Input:
			val := c.input()
			c.Set(c.ParamAddress(1), val)
			c.pc += 2
		case Output:
			c.output(c.Param(1))
			c.pc += 2
		case JumpIfTrue:
			if c.Param(1) != 0 {
				c.pc = c.Param(2)
			} else {
				c.pc += 3
			}
		case JumpIfFalse:
			if c.Param(1) == 0 {
				c.pc = c.Param(2)
			} else {
				c.pc += 3
			}
		case LessThen:
			val := int64(0)
			if c.Param(1) < c.Param(2) {
				val = 1
			}
			c.Set(c.ParamAddress(3), val)
			c.pc += 4
		case Equals:
			val := int64(0)
			if c.Param(1) == c.Param(2) {
				val = 1
			}
			c.Set(c.ParamAddress(3), val)
			c.pc += 4
		case AdjustRelativeBase:
			val := c.Param(1)
			c.rb += val
			//fmt.Println("rb", c.rb)
			c.pc += 2
		default:
			log.Fatal("Unknown instruction ", instruction)
			break
		case Exit:
			//fmt.Println("exit", pc)
			return
		}
	}
}
