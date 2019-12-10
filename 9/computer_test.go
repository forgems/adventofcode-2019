package main

import (
	"fmt"
	"testing"
)

func TestComputer(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		program := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
		c := NewComputer(program)
		output := []int64{}
		c.output = func(val int64) {
			output = append(output, val)
		}
		c.Run()
		if len(output) != len(program) {
			t.Errorf("invalid output length %d", len(output))
		}
	})
	t.Run("Test 2", func(t *testing.T) {
		program := []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
		c := NewComputer(program)
		output := []int64{}
		c.output = func(val int64) {
			fmt.Println(val)
			output = append(output, val)
		}
		c.Run()
		if len(output) != 1 {
			t.Errorf("Invalid output %v", output)
		}
		if len(fmt.Sprintf("%d", output[0])) != 16 {
			t.Errorf("Invalid number of digits: %d", len(fmt.Sprintf("%d", output[0])))
		}
	})
	t.Run("Test 3", func(t *testing.T) {
		c := NewComputer([]int64{104, 1125899906842624, 99})
		out := int64(0)
		c.output = func(val int64) {
			out = val
		}
		c.Run()
		if out != c.memory[1] {
			t.Errorf("Invalid output %d", out)
		}
	})
}
