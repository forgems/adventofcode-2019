package main

import (
	"fmt"
	"os"
)

func main() {
	c := ReadProgram(os.Stdin)
	tmp := make([]int64, len(c.memory))
	copy(tmp, c.memory)
	c.input = func() int64 {
		return 1
	}
	c.output = func(val int64) {
		fmt.Println("Output=", val)
	}
	c.Run()
	// part2
	fmt.Println("=======================\nPart2")
	copy(c.memory, tmp)
	c.additional_memory = map[int64]int64{}
	c.pc = 0
	c.rb = 0
	c.input = func() int64 {
		return 2
	}
	c.Run()
}
