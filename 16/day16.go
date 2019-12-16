package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func FFT(input, pattern []int, phase int) (sum int) {
	patLen := len(pattern)
	idx := 0
	for ; pattern[(idx+1)/phase]%patLen == 0; idx++ {
	}
	for ; idx < len(input); idx++ {
		val := input[idx]
		val = val * pattern[(idx+1)/(phase)%patLen]
		sum += val
	}
	sum = sum % 10
	if sum < 0 {
		sum = -sum
	}
	return sum
}

func StrToIntSlice(s string) (out []int) {
	for _, r := range s {
		out = append(out, int(r-'0'))
	}
	return out
}

func Multiply(input []int, n int) (out []int) {
	for n > 0 {
		out = append(out, input...)
		n--
	}
	return out
}

func SingleFFT(input, pattern []int) []int {
	out := []int{}
	inLen := len(input)
	fmt.Println("SingleFFT", inLen)
	for phase := 1; phase < inLen+1; phase++ {
		out = append(out, FFT(input, pattern, phase))
	}
	return out
}

func MultiFFT(input, pattern []int, numphases int) []int {
	for i := 0; i < numphases; i++ {
		input = SingleFFT(input, pattern)
	}
	return input
}

func Part2(input []int, skip int) {
	fmt.Println(len(input), skip)
	input = input[skip:]
	l := len(input)
	for phase := 0; phase < 100; phase++ {
		for i := l - 2; i > -1; i-- {
			input[i] += input[i+1]
			input[i] %= 10
		}
	}
}

func main() {
	fmt.Println("vim-go")
	pattern := []int{0, 1, 0, -1}
	data, _ := ioutil.ReadAll(os.Stdin)
	line := strings.TrimSpace(string(data))
	input := StrToIntSlice(line)
	out := MultiFFT(input, pattern, 100)
	fmt.Println(out[:8])

	input = Multiply(input, 10000)
	offset, _ := strconv.ParseInt(line[:7], 10, 32)
	Part2(input, int(offset))
	fmt.Println(input[offset : offset+8])
}
