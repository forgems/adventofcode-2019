package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFFT(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	pattern := []int{0, 1, 0, -1}
	out := []int{}
	for phase := 1; phase < len(input)+1; phase++ {
		out = append(out, FFT(input, pattern, phase))
	}
	fmt.Println(out)
}

func TestMultiFFT(t *testing.T) {
	pattern := []int{0, 1, 0, -1}
	data := []struct {
		input  string
		output string
	}{
		{
			"80871224585914546619083218645595",
			"24176176",
		},
		{
			"19617804207202209144916044189917",
			"73745418",
		},
		{
			"69317163492948606335995924319873",
			"52432133",
		},
	}

	for i := range data {
		i := i
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			in := StrToIntSlice(data[i].input)
			expected := StrToIntSlice(data[i].output)
			out := MultiFFT(in, pattern, 100)
			same := true
			for i, v := range expected {
				if v != out[i] {
					same = false
					break
				}
			}
			if !same {
				t.Errorf("%v is not equal to %v", expected, out[:len(expected)])
			}
		})
	}
}

func TestPart2(t *testing.T) {
	data := []struct {
		input  string
		output string
	}{
		{
			"03036732577212944063491565474664",
			"84462026",
		},
		{
			"02935109699940807407585447034323",
			"78725270",
		},
		{
			"03081770884921959731165446850517",
			"53553731",
		},
	}

	for i := range data {
		i := i
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			in := StrToIntSlice(data[i].input)
			offset, _ := strconv.ParseInt(data[i].input[:7], 10, 32)
			fmt.Println(offset)
			in = Multiply(in, 10000)
			expected := StrToIntSlice(data[i].output)
			Part2(in, int(offset))
			out := in[offset : offset+8]
			same := true
			for i, v := range expected {
				if v != out[i] {
					same = false
					break
				}
			}
			if !same {
				t.Errorf("%v is not equal to %v", expected, out[:len(expected)])
			}
		})
	}
}
