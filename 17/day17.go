package main

import (
	"fmt"
	"os"
	"strings"
)

func FindPath(picture [][]byte) (out []string) {
	pos, dir := FindRobotPos(picture)
	//fmt.Println(pos, dir)
	visited := map[complex64]bool{pos: true}
	rotation := FindRotation(pos, dir, picture)
	for rotation != 0 {
		if rotation == 1i {
			out = append(out, "R")
		} else {
			out = append(out, "L")
		}
		dir *= rotation
		//fmt.Println("Rotation", rotation, dir)
		steps := 0
		pos += dir
		for CanGo(int(real(pos)), int(imag(pos)), picture) {
			//fmt.Println("Pos", pos)
			steps++
			visited[pos] = true
			pos += dir
		}
		pos -= dir
		out = append(out, fmt.Sprintf("%d", steps))
		rotation = FindRotation(pos, dir, picture)
		//fmt.Println("Rotation", rotation)
	}
	return out
}
func Equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func CompressPath(path string, idx int) (funcPath string, dict map[string]string) {
	alphabet := map[rune]bool{}
	for _, r := range path {
		alphabet[r] = true
	}
	if idx > 3 {
		return
	}
	fmt.Println("CompressPath", path, len(alphabet))
	for i := 0; i < len(path); i += 1 {
		if path[i] >= 'A' && path[i] <= 'C' {
			continue
		}
		for j := len(path); j > i; j-- {
			if path[j-1] >= 'A' && path[j-1] <= 'C' {
				continue
			}
			subpath := path[i:j]
			if strings.Index(path[j:], subpath) >= 0 {
				CompressPath(strings.Replace(path, subpath, string('A'+idx), -1), idx+1)
			}
		}
	}
	return path, map[string]string{}
}

func ReplaceWithDefinition(path []string, definition []string, name string) []string {
	out := []string{}
	defLen := len(definition)
	pathLen := len(path)
	for i := 0; i < pathLen; {
		if i+defLen < pathLen && Equal(path[i:i+defLen], definition) {
			out = append(out, name)
			i += defLen
			continue
		}
		out = append(out, path[i])
		i++
	}
	return out
}

func CanGo(x, y int, pic [][]byte) bool {
	//fmt.Println("CanGo", x, y)
	if y < 0 || y >= len(pic) {
		return false
	}
	if x < 0 || x >= len(pic[y]) {
		return false
	}
	//fmt.Printf("%c\n", pic[y][x])
	return pic[y][x] == '#'
}

func FindRotation(pos, dir complex64, picture [][]byte) complex64 {
	if pos := pos + dir*1i; CanGo(int(real(pos)), int(imag(pos)), picture) {
		return 1i
	}
	if pos := pos + dir*-1i; CanGo(int(real(pos)), int(imag(pos)), picture) {
		return -1i
	}
	return 0
}

func FindRobotPos(picture [][]byte) (complex64, complex64) {
	for y := range picture {
		for x := range picture[y] {
			val := picture[y][x]
			if val == '^' {
				return complex(float32(x), float32(y)), complex(0, -1)
			}
			if val == 'v' {
				return complex(float32(x), float32(y)), complex(0, 1)
			}
			if val == '<' {
				return complex(float32(x), float32(y)), complex(-1, 0)
			}
			if val == '>' {
				return complex(float32(x), float32(y)), complex(1, 0)

			}
		}
	}
	return -1, -1
}

func FindAlignment(picture [][]byte) (sum int) {
	for y := 1; y < len(picture)-1; y++ {
		for x := 1; x < len(picture[y])-1; x++ {
			if picture[y][x] != '#' {
				continue
			}
			if picture[y-1][x] != '#' {
				continue
			}
			if len(picture[y+1]) < x || picture[y+1][x] != '#' {
				continue
			}
			if picture[y][x-1] != '#' {
				continue
			}
			if picture[y][x+1] != '#' {
				continue
			}
			sum += y * x
		}
	}
	return sum
}

func Draw(picture [][]byte) {
	for i, row := range picture {
		if i == 0 {
			fmt.Print("   ")
			for j := range row {
				fmt.Printf("%d", j%10)
			}
			fmt.Println()
		}
		fmt.Printf("%2d ", i)
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}

func main() {
	c := ReadProgram(os.Stdin)
	part2 := make([]int64, len(c.memory))
	copy(part2, c.memory)
	out := [][]byte{[]byte{}}
	y := 0
	x := 0
	c.output = func(val int64) {
		switch val {
		case 10:
			if x == 0 {
				break
			}
			out = append(out, []byte{})
			y++
			x = 0
		default:
			out[y] = append(out[y], byte(val))
			x++
		}
	}
	c.Run()
	Draw(out)
	fmt.Println(FindAlignment(out))
	fmt.Println(len(out))
	fmt.Println(strings.Join(FindPath(out), ","))
	c = NewComputer(part2)
	c.memory[0] = 2
	i := 0
	program := `A,C,C,B,A,C,B,A,C,B
L,6,R,12,L,4,L,6
L,6,L,10,L,10,R,6
R,6,L,6,R,12
n
`
	//fmt.Println([]byte(program))
	c.input = func() int64 {
		fmt.Println("input")
		val := int64(program[i])
		i++
		return val
	}
	c.output = func(val int64) {
		fmt.Println(val)
	}
	fmt.Println("Part2")
	c.Run()
}
