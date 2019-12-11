package main

import (
	"fmt"
	"os"
)

const (
	BLACK = 0
	WHITE = 1
)

func main() {
	c := ReadProgram(os.Stdin)
	input, output := make(chan int64, 1), make(chan int64)
	hull := map[string]int64{}

	c.input = func() int64 {
		val := <-input
		return val
	}

	c.output = func(val int64) {
		output <- val
	}
	go func() {
		c.Run()
		close(output)
	}()

	direction := 0 - 1i
	pos := 0 + 0i

	hull[fmt.Sprintf("%v", pos)] = WHITE // part2
	for {
		input <- hull[fmt.Sprintf("%v", pos)]
		color, ok := <-output
		if !ok {
			break
		}
		hull[fmt.Sprintf("%v", pos)] = color
		turn, ok := <-output
		if !ok {
			break
		}
		if turn == 0 {
			direction *= 1i
		} else {
			direction *= -1i
		}
		pos += direction
	}
	close(input)
	fmt.Println(len(hull))
	drawHull(hull)
}

func drawHull(hull map[string]int64) {
	var min_x, max_x, min_y, max_y int
	force := true
	for k := range hull {
		var c complex128
		fmt.Sscanf(k, "%g", &c)
		x := int(real(c))
		y := int(imag(c))
		if force || x <= min_x {
			min_x = x
		}
		if force || x > max_x {
			max_x = x
		}
		if force || y < min_y {
			min_y = y
		}
		if force || y > max_y {
			max_y = y
		}
		force = false
	}
	fmt.Println(min_x, max_x, min_y, max_y)
	for y := min_y; y <= max_y; y++ {
		fmt.Printf("%04d", y)
		for x := max_x; x >= min_x; x-- {
			c := complex(float64(x), float64(y))
			switch hull[fmt.Sprintf("%v", c)] {
			case BLACK:
				fmt.Print(" ")
			case WHITE:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
