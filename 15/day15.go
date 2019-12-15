package main

import (
	"fmt"
	"os"
)

type Map map[int]map[int]int

func (m Map) Set(x, y, val int) {
	if m[y] == nil {
		m[y] = map[int]int{}
	}
	m[y][x] = val
}

func (m Map) Get(x, y int) int {
	if m[y] == nil {
		return -1
	}
	val, ok := m[y][x]
	if !ok {
		return -1
	}
	return val
}

func (m Map) Draw(droidX, droidY int) {
	min_y, max_y := 0, 0
	min_x, max_x := 0, 0
	first := true

	for y, row := range m {
		if y < min_y || first {
			min_y = y
		}
		if y > max_y || first {
			max_y = y
		}
		for x := range row {
			if x < min_x || first {
				min_x = x
			}
			if x > max_x || first {
				max_x = x
			}
		}
		first = false
	}
	fmt.Println("\033[2J")
	fmt.Println("============================================")

	for y := min_y; y <= max_y; y++ {
		fmt.Printf("%3d ", y)
		for x := min_x; x <= max_x; x++ {
			if x == droidX && y == droidY {
				fmt.Print("D")
				continue
			}
			v := m.Get(x, y)
			switch v {
			case -1:
				fmt.Print(".")
			case 0:
				fmt.Print("#")
			case 1:
				fmt.Print(" ")
			case 2:
				fmt.Print("o")
			case 5:
				fmt.Print("~")
			}
		}
		fmt.Println()
	}
	fmt.Println("============================================")
}

func WalkMap(x, y int, input, output chan int64, m Map) {
	for dir := 1; dir < 5; dir++ {
		nx, ny := x, y
		backdir := 0
		switch dir {
		case 1:
			ny--
			backdir = 2
		case 2:
			ny++
			backdir = 1
		case 3:
			nx++
			backdir = 4
		case 4:
			nx--
			backdir = 3
		}
		if m.Get(nx, ny) >= 0 {
			continue
		}
		input <- int64(dir)
		val := <-output
		m.Set(nx, ny, int(val))
		// hit the wall
		if val == 0 {
			continue
		}
		// bot arrived at destination
		if val > 0 {
			WalkMap(nx, ny, input, output, m)
		}
		input <- int64(backdir)
		val = <-output
		if val != 1 {
			fmt.Println("Error %d", val)
		}
	}
}

func BFS(x, y int, m Map) int {
	queue := [][3]int{{x, y, 0}}
	visited := map[int]bool{}
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		key := first[1]<<8 + first[0]
		if visited[key] {
			continue
		}
		visited[key] = true
		val := m.Get(first[0], first[1])
		if val == 2 {
			fmt.Println(first)
			return first[2]
		}
		if val <= 0 {
			continue
		}
		queue = append(queue, [3]int{first[0], first[1] - 1, first[2] + 1})
		queue = append(queue, [3]int{first[0], first[1] + 1, first[2] + 1})
		queue = append(queue, [3]int{first[0] - 1, first[1], first[2] + 1})
		queue = append(queue, [3]int{first[0] + 1, first[1], first[2] + 1})
	}
	return 0
}

func BFS2(x, y int, m Map) int {
	queue := [][3]int{{x, y, 0}}
	visited := map[int]bool{}
	distance := 0
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		key := first[1]<<8 + first[0]
		if visited[key] {
			continue
		}
		visited[key] = true
		val := m.Get(first[0], first[1])
		if val <= 0 {
			continue
		}
		//m.Set(first[0], first[1], 5)
		//m.Draw(0, 0)
		//fmt.Println(first[2])
		if first[2] > distance {
			distance = first[2]
		}
		queue = append(queue, [3]int{first[0], first[1] - 1, first[2] + 1})
		queue = append(queue, [3]int{first[0], first[1] + 1, first[2] + 1})
		queue = append(queue, [3]int{first[0] - 1, first[1], first[2] + 1})
		queue = append(queue, [3]int{first[0] + 1, first[1], first[2] + 1})
	}
	return distance
}
func FindOxygen(m Map) (int, int) {
	for y, row := range m {
		for x, v := range row {
			if v == 2 {
				return x, y
			}
		}
	}
	return 0, 0
}

func main() {
	c := ReadProgram(os.Stdin)
	m := Map{0: {0: 1}}
	m.Set(0, 0, 1)
	input, output := make(chan int64, 1), make(chan int64)

	c.input = func() int64 {
		return <-input
	}
	c.output = func(val int64) {
		output <- val
	}
	go func() {
		WalkMap(0, 0, input, output, m)
		m.Draw(0, 0)
		close(input)
	}()
	c.Run()
	close(output)
	fmt.Println(BFS(0, 0, m))
	x, y := FindOxygen(m)
	fmt.Println(BFS2(x, y, m))
}
