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

type input_func func() int
type output_func func(int)

func (program Program) execute(in input_func, out output_func) {
	pc := 0
	length := len(program)
	for pc < length {
		instruction := program[pc]
		//fmt.Println("ins=", instruction)
		opcode := instruction % 100
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
			val := in()
			program.set(program[pc+1], val)
			pc += 2
		case Output:
			out(program.get(pc+1, instruction/100))
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

func amplify(program Program, phases []int) int {
	tmp := make(Program, len(program))
	signal := 0
	for _, phase := range phases {
		inputs := []int{phase, signal}
		input := func() int {
			val := inputs[0]
			inputs = inputs[1:]
			return val
		}
		output := func(i int) {
			signal = i
		}
		copy(tmp, program)
		tmp.execute(input, output)
	}
	return signal
}

func amplify2(program Program, phases []int) int {
	n := len(phases)
	amplifiers := make([]Program, n)
	channels := make([]chan int, n+1)
	signal := 0
	channels[n] = make(chan int, 2)
	for i, phase := range phases {
		amplifiers[i] = make(Program, len(program))
		copy(amplifiers[i], program)
		channels[i] = make(chan int, 2)
		channels[i] <- phase
	}
	channels[0] <- signal

	for i := 0; i < n; i++ {
		go func(in <-chan int, out chan<- int, idx int) {
			input := func() int {
				fmt.Println("Getting input for ", idx)
				val := <-in
				fmt.Printf("input[%d]: %d\n", idx, val)
				return val
			}
			output := func(val int) {
				fmt.Printf("Output for %d is %d\n", idx, val)
				out <- val
			}
			amplifiers[idx].execute(input, output)
			close(out)
		}(channels[i], channels[i+1], i)
	}

	for signal = range channels[n] {
		fmt.Println("Final output", signal)
		channels[0] <- signal
	}
	close(channels[0])
	return signal
}

func findMaxTrust(program Program, phases []int, amplify func(p Program, phases []int) int) int {
	n := len(phases)
	max := amplify(program, phases)
	c := make([]int, n)
	for i := 0; i < n; {
		if c[i] < i {
			if i%2 == 0 {
				phases[0], phases[i] = phases[i], phases[0]
			} else {
				phases[c[i]], phases[i] = phases[i], phases[c[i]]
			}
			signal := amplify(program, phases)
			fmt.Printf("signal is %d for phases %v\n", signal, phases)
			if signal > max {
				max = signal
			}
			c[i] += 1
			i = 0
		} else {
			c[i] = 0
			i += 1
		}
	}
	return max
}

func main() {
	phases := []int{0, 1, 2, 3, 4}
	program := readProgram(os.Stdin)
	fmt.Println(findMaxTrust(program, phases, amplify))
	fmt.Println(findMaxTrust(program, []int{5, 6, 7, 8, 9}, amplify2))
}
