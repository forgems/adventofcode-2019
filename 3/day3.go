package main

import (
	"fmt"
	"log"
	"strings"
)

type point struct {
	x, y int32
}
type pointval struct {
	wire     int
	distance int
}

var min_point *point = nil
var min_distance int = -1

func abs(val int32) int32 {
	if val < 0 {
		return -val
	}
	return val
}

func trace1(wire int, segment string, field map[int64]int, pos point) point {
	var direction rune
	var amount int
	_, err := fmt.Sscanf(segment, "%c%d", &direction, &amount)
	if err != nil {
		log.Fatal(err)
		return pos
	}
	for i := 0; i < amount; i++ {
		switch direction {
		case 'D':
			pos.y--
		case 'U':
			pos.y++
		case 'L':
			pos.x--
		case 'R':
			pos.x++
		}
		fmt.Println(wire, segment, pos)
		key := int64(pos.x)<<32 + int64(pos.y)
		if field[key] != 0 && field[key] != wire {
			fmt.Println("Collision at", pos)
			if min_point == nil {
				min_point = &point{pos.x, pos.y}
			} else if abs(min_point.x)+abs(min_point.y) > abs(pos.x)+abs(pos.y) {
				min_point = &point{pos.x, pos.y}
			}
		}
		field[key] = wire
	}
	return pos
}

func trace2(wire int, segment string, field map[int64]pointval, pos point, distance int) (point, int) {
	var direction rune
	var amount int
	_, err := fmt.Sscanf(segment, "%c%d", &direction, &amount)
	if err != nil {
		log.Fatal(err)
		return pos, distance
	}
	for i := 0; i < amount; i++ {
		switch direction {
		case 'D':
			pos.y--
		case 'U':
			pos.y++
		case 'L':
			pos.x--
		case 'R':
			pos.x++
		}
		distance++
		fmt.Println(wire, segment, pos)
		key := int64(pos.x)<<32 + int64(pos.y)
		if field[key].wire != 0 && field[key].wire != wire {
			fmt.Println("Collision at", pos)
			if min_distance < 0 || field[key].distance+distance < min_distance {
				min_distance = field[key].distance + distance
			}
		}
		field[key] = pointval{wire, distance}
	}
	return pos, distance
}

func main() {
	var line string
	field := map[int64]int{}
	distance_field := map[int64]pointval{}
	for wire := 1; ; wire++ {
		_, err := fmt.Scanf("%s", &line)
		if err != nil {
			break
		}
		fmt.Println("line")
		segments := strings.Split(line, ",")
		pt := point{}
		distance := 0
		for i := range segments {
			trace1(wire, segments[i], field, pt)
			pt, distance = trace2(wire, segments[i], distance_field, pt, distance)
		}
	}
	fmt.Println(min_point)
	fmt.Println(abs(min_point.x) + abs(min_point.y))
	fmt.Println(min_distance)
}
