package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFindAlignment(t *testing.T) {
	picture := [][]byte{
		[]byte("..#.........."),
		[]byte("..#.........."),
		[]byte("#######...###"),
		[]byte("#.#...#...#.#"),
		[]byte("#############"),
		[]byte("..#...#...#.."),
		[]byte("..#####...^.."),
	}

	if FindAlignment(picture) != 76 {
		t.Error("Invalid alignment value")
	}
}

func TestFindRotation(t *testing.T) {
	picture := [][]byte{
		[]byte("#######...#####"),
		[]byte("#.....#...#...#"),
		[]byte("#.....#...#...#"),
		[]byte("......#...#...#"),
		[]byte("......#...###.#"),
		[]byte("......#.....#.#"),
		[]byte("^########...#.#"),
		[]byte("......#.#...#.#"),
		[]byte("......#########"),
		[]byte("........#...#.."),
		[]byte("....#########.."),
		[]byte("....#...#......"),
		[]byte("....#...#......"),
		[]byte("....#...#......"),
		[]byte("....#####......"),
	}
	expected := strings.Split("R,8,R,8,R,4,R,4,R,8,L,6,L,2,R,4,R,4,R,8,R,8,R,8,L,6,L,2", ",")
	path := FindPath(picture)
	for idx := range expected {
		if expected[idx] != path[idx] {
			t.Errorf("Invalid step %s at %d. Expected %s", path[idx], idx, expected[idx])
		}
	}
}

func TestCompress(t *testing.T) {
	path := "R8R8R4R4R8L6L2R4R4R8R8R8L6L2"
	//path = "PPQQPRSQQPPPRS"
	//expected := "A,B,C,A,B,C"
	//a := "R,8,R,8"
	//b := "R,4,R,4,R,8"
	//c := "L,6,L,2"
	//funcPath, funcDict := CompressPath(path, 3)
	//fmt.Println(funcPath, funcDict)
	fmt.Println(CompressPath(path, 0))
}
