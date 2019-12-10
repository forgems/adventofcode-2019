package main

import (
	"fmt"
	"io"
	"math/cmplx"
	"os"
	"sort"
)

type AsteroidMap [][]int

type Point struct {
	x, y int
}

func ReadMap(reader io.Reader) AsteroidMap {
	var line string
	field := AsteroidMap{}
	for y := 0; ; y++ {
		n, err := fmt.Fscanln(reader, &line)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Errorf("Error during read %s", err)
			break
		}
		field = append(field, make([]int, len(line)))
		for x, c := range line {
			point := 0
			if c == '#' {
				point = 1
			}
			field[y][x] = point
		}
	}
	return field
}

func (field AsteroidMap) FindBestAsteroid() (int, int, int) {
	max_pos := struct{ x, y int }{}
	max_number := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] == 1 {
				points := field.findVisibleAsteroids(x, y)
				if len(points) > max_number {
					max_number = len(points)
					max_pos.x = x
					max_pos.y = y
				}
			}
		}
	}
	return max_pos.x, max_pos.y, max_number
}

func (field AsteroidMap) findVisibleAsteroids(x, y int) []complex128 {
	points := []complex128{}
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[row]); col++ {
			if row == y && col == x {
				continue
			}
			if field[row][col] == 0 {
				continue
			}
			if field.isVisible(x, y, col, row) {
				//fmt.Println(row, col, "isVisible from", x, y)
				points = append(points, complex(float64(col), float64(row)))
			}
		}
	}
	return points
}

func (field AsteroidMap) isVisible(x1, y1, x2, y2 int) bool {
	//fmt.Println("isVisible", x1, y1, x2, y2)
	// calculate unit vector
	dx, dy := x2-x1, y2-y1
	gcd := GCD(dx, dy)
	//fmt.Println("dx", dx, "dy", dy)
	if gcd < 0 {
		gcd = -gcd
	}
	if gcd > 0 {
		//fmt.Println("GCD", gcd)
		dx /= gcd
		dy /= gcd
	}
	if dy == 0 {
		if dx > 0 {
			dx = 1
		} else {
			dx = -1
		}
	}
	if dx == 0 {
		if dy > 0 {
			dy = 1
		} else {
			dy = -1
		}
	}
	// fmt.Println("dx", dx, "dy", dy)
	// move from origin to destination checking for obstacles
	for {
		x1 += dx
		y1 += dy
		//fmt.Println(x1, y1)
		if x1 == x2 && y1 == y2 {
			break
		}
		if field[y1][x1] == 1 {
			//fmt.Println("Obstacle at ", x1, y1)
			return false
		}
	}

	return true
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	field := ReadMap(os.Stdin)
	x, y, n := field.FindBestAsteroid()
	fmt.Println(x, y, n)
	points := field.findVisibleAsteroids(x, y)
	asteroid := complex(float64(x), float64(y))
	sort.Slice(points, func(i, j int) bool {
		// normalize the point against the asteroid, rotate clockwise and get the phase
		return cmplx.Phase((points[i]-asteroid)*-1i) < cmplx.Phase((points[j]-asteroid)*-1i)
	})
	fmt.Println(points[199])
}
