package main

import (
	"fmt"
	"math/cmplx"
	"strings"
	"testing"
)

func TestReadMap(t *testing.T) {
	// t.Fatal("not implemented")
	t.Run("large field", func(t *testing.T) {
		reader := strings.NewReader(`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`)
		field := ReadMap(reader)
		if len(field) != 20 {
			t.Errorf("Invalid field height: %d", len(field))
		}
		if len(field[0]) != 20 {
			t.Errorf("Invalid field width: %d", len(field))
		}
	})
}

func TestFindBestAsteroid(t *testing.T) {
	var test_data = []struct {
		data    string
		x, y, n int
	}{
		{`.#..#
.....
#####
....#
...##`, 3, 4, 8},
		{`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`, 5, 8, 33},
		{`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`, 1, 2, 35},
		{`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`, 6, 3, 41},
		{`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`, 11, 13, 210},
	}

	for idx := range test_data {
		i := idx
		t.Run(fmt.Sprintf("example %d", i), func(t *testing.T) {
			field := ReadMap(strings.NewReader(test_data[i].data))
			x, y, number := field.FindBestAsteroid()
			if x != test_data[i].x || y != test_data[i].y || number != test_data[i].n {
				t.Errorf("Invalid output for %v, %d, %d, %d", test_data[i], x, y, number)
			}
		})
	}
}

func TestIsVisible(t *testing.T) {
	field := ReadMap(
		strings.NewReader(
			`.#..#
.....
#####
....#
...##`))
	t.Run("0,1 -> 1,1", func(t *testing.T) {
		if !field.isVisible(0, 2, 1, 2) {
			t.Error("Not visible")
		}
	})
	t.Run("0,2 -> 2,2", func(t *testing.T) {
		if field.isVisible(0, 2, 2, 2) {
			t.Error("visible")
		}
	})
	t.Run("1,2 -> 0,2", func(t *testing.T) {
		if !field.isVisible(1, 2, 0, 2) {
			t.Error("not visible")
		}
	})
	t.Run("0,2 -> 4,2", func(t *testing.T) {
		if field.isVisible(0, 2, 4, 2) {
			t.Error("visible")
		}
	})
	t.Run("0,2 -> 3,4", func(t *testing.T) {
		if !field.isVisible(0, 2, 3, 4) {
			t.Error("not visible")
		}
	})
	t.Run("Find visible asteroids", func(t *testing.T) {
		n := field.findVisibleAsteroids(3, 4)
		if len(n) != 8 {
			t.Errorf("Invalid number of visible asteroids %d", len(n))
		}
	})
	t.Run("Find visible asteroids 1, 2", func(t *testing.T) {
		n := field.findVisibleAsteroids(1, 2)
		if len(n) != 7 {
			t.Errorf("Invalid number of visible asteroids %d", len(n))
		}
	})
}

func TestPoint(t *testing.T) {
	fmt.Println(cmplx.Abs(1 + 1i))
	fmt.Println(cmplx.Phase((0 + 1i) * 1i))
	fmt.Println(cmplx.Phase((1) * 1i))
	fmt.Println(cmplx.Phase(-1i))
}
