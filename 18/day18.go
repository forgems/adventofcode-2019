package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Map []string

type Pos [2]int

type Visited [256]bool

func (v *Visited) String() string {
	b := strings.Builder{}
	for i, v := range v {
		if v {
			fmt.Fprintf(&b, "%c", i)
		}
	}
	return b.String()

}

func (m Map) Width() int {
	if len(m) > 0 {
		return len(m[0])
	}
	return 0
}

func (m Map) Height() int {
	return len(m)
}

func (m Map) Find(what rune) Pos {
	for y := range m {
		for x, r := range m[y] {
			if r == what {
				return Pos{x, y}
			}
		}
	}
	return Pos{-1, -1}
}

func (m Map) CanWalk(p Pos) bool {
	if p[0] < 0 || p[0] >= m.Width() {
		return false
	}
	if p[1] < 0 || p[1] >= m.Height() {
		return false
	}
	return m[p[1]][p[0]] != '#'
}

func (m Map) FindKeys() map[rune]Pos {
	out := map[rune]Pos{}
	for y := range m {
		for x, r := range m[y] {
			if r >= 'a' && r <= 'z' || r == '@' {
				out[r] = Pos{x, y}
			}
		}
	}
	return out
}

func (m Map) FindAll(needle rune) (out []Pos) {
	for y := range m {
		for x, r := range m[y] {
			if r == needle {
				out = append(out, Pos{x, y})
			}
		}
	}
	return out
}

func (m Map) BFSKeys(s Pos, keys Visited) (out map[rune]int) {
	out = map[rune]int{}
	visited := map[Pos]bool{}
	q := [][3]int{
		{s[0], s[1], 0},
	}
	for len(q) > 0 {
		pt := q[0]
		q = q[1:]
		if visited[Pos{pt[0], pt[1]}] {
			continue
		}
		visited[Pos{pt[0], pt[1]}] = true
		v := rune(m[pt[1]][pt[0]])
		if v >= 'A' && v <= 'Z' {
			if !keys[unicode.ToLower(v)] {
				continue
			}
		}
		if v != '.' && v >= 'a' && v <= 'z' && !keys[v] {
			out[v] = pt[2]
		}
		dist := pt[2] + 1
		if m.CanWalk(Pos{pt[0] - 1, pt[1]}) {
			q = append(q, [3]int{pt[0] - 1, pt[1], dist})
		}
		if m.CanWalk(Pos{pt[0] + 1, pt[1]}) {
			q = append(q, [3]int{pt[0] + 1, pt[1], dist})
		}
		if m.CanWalk(Pos{pt[0], pt[1] - 1}) {
			q = append(q, [3]int{pt[0], pt[1] - 1, dist})
		}
		if m.CanWalk(Pos{pt[0], pt[1] + 1}) {
			q = append(q, [3]int{pt[0], pt[1] + 1, dist})
		}
	}
	return out
}

func (m Map) BFS(s, e Pos) int {
	q := [][3]int{
		{s[0], s[1], 0},
	}
	for len(q) > 0 {
		pt := q[0]
		q = q[1:]
		if pt[0] == e[0] && pt[1] == e[1] {
			return pt[2]
		}
		if m.CanWalk(Pos{pt[0] - 1, pt[1]}) {
			q = append(q, [3]int{pt[0] - 1, pt[1], pt[2] + 1})
		}
		if m.CanWalk(Pos{pt[0] + 1, pt[1]}) {
			q = append(q, [3]int{pt[0] + 1, pt[1], pt[2] + 1})
		}
		if m.CanWalk(Pos{pt[0], pt[1] - 1}) {
			q = append(q, [3]int{pt[0], pt[1] - 1, pt[2] + 1})
		}
		if m.CanWalk(Pos{pt[0], pt[1] + 1}) {
			q = append(q, [3]int{pt[0], pt[1] + 1, pt[2] + 1})
		}
	}
	return -1
}

func (m Map) String() string {
	b := strings.Builder{}
	for i := range m {
		fmt.Fprintf(&b, "%s\n", m[i])
	}
	return b.String()
}

func ReadMap(r io.Reader) Map {
	m := Map{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}
	return m
}

func WalkMap(m Map, k rune, visited Visited, cache map[string]int) int {
	if cache == nil {
		cache = map[string]int{}
	}
	key := fmt.Sprintf("%c,%s", k, visited.String())
	if val, ok := cache[key]; ok {
		return val
	}
	//fmt.Printf("WalkMap: %c %s \n", k, visited.String())
	pos := m.Find(k)
	visited[k] = true
	result := m.BFSKeys(pos, visited)
	//fmt.Println("BFSKEYS=", result)
	if len(result) == 0 {
		//cache[key] = distance
		return 0
	}
	min_dist := 1000000
	for next, dst := range result {
		d := WalkMap(m, next, visited, cache)
		if dst+d < min_dist {
			min_dist = d + dst
		}
	}
	cache[key] = min_dist
	return min_dist
}

func Part2(m Map) int {
	starts := m.FindAll('@')
	fmt.Println(starts)
	keys := Visited{}
	allKeys := m.FindKeys()
	collected := 0
	distance := 0
	for collected < len(allKeys)-1 {
		for i := range starts {
			s := starts[i]
			// get one key
			for k, dist := range m.BFSKeys(s, keys) {
				distance += dist
				fmt.Printf("Collecting %c, %d\n", k, dist)
				keys[k] = true
				starts[i] = m.Find(k)
				collected++
				break
			}
		}
	}
	return distance
}

func main() {
	m := ReadMap(os.Stdin)
	fmt.Println(WalkMap(m, '@', Visited{}, nil))
}
