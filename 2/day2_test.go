package main

import "testing"

func TestExecuteProgram(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		p := []int{1, 0, 0, 0, 99}
		executeProgram(p)
		if p[0] != 2 {
			t.Error(p)
		}
	})
	t.Run("multiply", func(t *testing.T) {
		p := []int{2, 3, 0, 3, 99}
		executeProgram(p)
		if p[3] != 6 {
			t.Error(p)
		}
	})
	t.Run("multiply 2", func(t *testing.T) {
		p := []int{2, 4, 4, 5, 99, 0}
		executeProgram(p)
		if p[5] != 9801 {
			t.Error(p)
		}
	})
	t.Run("3", func(t *testing.T) {
		p := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
		executeProgram(p)
		if p[0] != 30 || p[4] != 2 {
			t.Error(p)
		}
	})
}
