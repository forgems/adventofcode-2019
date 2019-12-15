package main

import (
	"fmt"
	"os"
)

const (
	W = 43
	H = 23
)

type Vec2D [2]int64

func (v Vec2D) Add(o Vec2D) Vec2D {
	for i := range v {
		v[i] += o[i]
	}
	return v
}

func (v Vec2D) Sub(o Vec2D) Vec2D {
	for i := range v {
		v[i] -= o[i]
	}
	return v
}

func Part1(c *Computer) {
	instructions := make(chan int64)
	screen := make([]int, W*H)

	c.output = func(v int64) {
		instructions <- v
	}
	go func() {
		c.Run()
		close(instructions)
	}()

	ins := [3]int64{}
	i := 0
	for out := range instructions {
		ins[i] = out
		i++
		if i >= len(ins) {
			i = 0
			screen[ins[0]+W*ins[1]] = int(ins[2])
		}
	}
	count := 0
	for _, v := range screen {
		if v == 2 {
			count++
		}
	}
	fmt.Println(count)

}

func Part2(c *Computer) {
	c.memory[0] = 2
	screen := make([]int, W*H)
	prev_input_ball_pos := Vec2D{}
	var score int64 = 0
	ins := [3]int64{}
	i := 0

	c.output = func(v int64) {
		ins[i] = v
		i++
		if i >= len(ins) {
			i = 0
			if ins[0] == -1 && ins[1] == 0 {
				score = ins[2]
				fmt.Println("Set score", score)
				return
			}
			screen[ins[0]+W*ins[1]] = int(ins[2])
		}
	}

	c.input = func() int64 {
		ball_pos := Find(screen, W, H, 4)
		paddle_pos := Find(screen, W, H, 3)
		//DrawScreen(screen, W, H)
		ball_direction := ball_pos.Sub(prev_input_ball_pos)
		if prev_input_ball_pos[0] == 0 && prev_input_ball_pos[1] == 0 {
			ball_direction = paddle_pos.Sub(ball_pos)
		}
		prev_input_ball_pos = ball_pos
		predicted_pos := ball_pos
		for predicted_pos[1] > 0 && predicted_pos[1] < paddle_pos[1] {
			predicted_pos = predicted_pos.Add(ball_direction)
		}
		if ball_direction[0] < 0 {
			paddle_pos[0] -= 1
		} else {
			paddle_pos[0] += 1
		}
		if predicted_pos[0] < paddle_pos[0] {
			return -1
		} else if predicted_pos[0] > paddle_pos[0] {
			return 1
		}
		return 0
	}
	c.Run()
	fmt.Println(score)
}

func Find(screen []int, w, h, what int) Vec2D {
	for i, v := range screen {
		if v == what {
			return Vec2D{int64(i % w), int64(i / w)}
		}
	}
	return Vec2D{0, 0}
}

func DrawScreen(screen []int, w, h int) {
	//fmt.Print("\033[H\033[2J")
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := ' '
			switch screen[y*w+x] {
			case 0:
				c = ' '
			case 1:
				c = 'W'
			case 2:
				c = 'B'
			case 3:
				c = '-'
			case 4:
				c = 'o'
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func main() {
	c := ReadProgram(os.Stdin)
	Part2(c)
}
