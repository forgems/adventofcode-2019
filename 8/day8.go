package main

import (
	"fmt"
	"io"
	"os"
)

func readImage(width, height int, reader io.Reader) [][]byte {
	n := width * height
	out := [][]byte{}
	for {
		tmp := make([]byte, n)
		read, err := reader.Read(tmp)
		if read < n {
			fmt.Println("aaaa")
			return out
		}
		if err != nil {
			fmt.Errorf("Error during read %s", err)
			return out
		}
		for i := range tmp {
			tmp[i] = tmp[i] - '0'
		}
		out = append(out, tmp)
	}
	return out
}

func renderImage(width, height int, layers [][]byte) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
		L:
			for z := range layers {
				val := layers[z][y*width+x]
				switch val {
				case 1:
					fmt.Printf("X")
					break L
				case 0:
					fmt.Printf(" ")
					break L
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	layers := readImage(25, 6, os.Stdin)
	min_idx := -1
	min := 0
	for i := range layers {
		zeroes := 0
		for j := range layers[i] {
			if layers[i][j] == 0 {
				zeroes += 1
			}
		}
		if min_idx < 0 || zeroes < min {
			min = zeroes
			min_idx = i
		}
	}
	ones, twos := 0, 0
	for i := range layers[min_idx] {
		switch layers[min_idx][i] {
		case 1:
			ones += 1
		case 2:
			twos += 1
		}
	}
	fmt.Println(ones * twos)
	renderImage(25, 6, layers)
}
