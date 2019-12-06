package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

const (
	Add int = iota + 1
	Multiply
	Exit = 99
)

func readProgram() []int {
	program := []int{}
	for {
		var opcode int
		_, err := fmt.Scanf("%d", &opcode)
		if err != nil {
			fmt.Println(err)
			break
		}
		program = append(program, opcode)
	}
	return program
}

func executeProgram(program []int) {
	pc := 0
	length := len(program)
	for pc < length {
		switch program[pc] {
		case Add:
			//fmt.Println("add", pc)
			program[program[pc+3]] = program[program[pc+1]] + program[program[pc+2]]
			pc += 4
		case Multiply:
			// fmt.Println("multiply", pc)
			program[program[pc+3]] = program[program[pc+1]] * program[program[pc+2]]
			pc += 4
		case Exit:
			//fmt.Println("exit", pc)
			return
		}
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	program := readProgram()
	tmp := make([]int, len(program))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(tmp, program)
			//fmt.Println("noun=", noun, "verb=", verb)
			tmp[1] = noun
			tmp[2] = verb
			executeProgram(tmp)
			if tmp[0] == 19690720 {
				fmt.Println("values = ", 100*noun+verb)
				return
			}
		}
	}
}
