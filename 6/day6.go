package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func path(m map[string]string, start string) []string {
	out := []string{}
	for m[start] != "" {
		out = append(out, start)
		fmt.Printf("%s ", start)
		start = m[start]
	}
	fmt.Println()
	return out
}
func readMap(input io.Reader) map[string]string {
	m := map[string]string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		var a, b string
		strings := strings.Split(line, ")")
		a = strings[0]
		b = strings[1]
		fmt.Println(a, b)
		m[b] = a
	}
	return m
}

func orbitChecksum(m map[string]string) int {
	checksum := 0
	for k := range m {
		checksum += len(path(m, k))
	}
	return checksum
}

func orbitTransfer(m map[string]string, start, end string) int {
	path1 := path(m, start)
	path2 := path(m, end)
	p1 := len(path1) - 1
	p2 := len(path2) - 1
	for path1[p1] == path2[p2] {
		p1--
		p2--
	}
	return p1 + p2
}

func main() {
	m := readMap(os.Stdin)
	fmt.Println(orbitChecksum(m))
	fmt.Println(orbitTransfer(m, "YOU", "SAN"))
}
