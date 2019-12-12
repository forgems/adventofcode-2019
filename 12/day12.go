package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Vec3D [3]float64

func (v Vec3D) String() string {
	return fmt.Sprintf("x=%3.0f y=%3.0f z=%3.0f", v[0], v[1], v[2])
}

func (v Vec3D) Energy() float64 {
	total := 0.0
	for c := range v {
		total += math.Abs(v[c])
	}
	return total
}

func (v Vec3D) Equal(o Vec3D) bool {
	for c := range v {
		if v[c] != o[c] {
			return false
		}
	}
	return true
}

type Moon struct {
	Position Vec3D
	Velocity Vec3D
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<%s>, vel=<%s>", m.Position, m.Velocity)
}

func (m *Moon) Move() {
	for d := range m.Position {
		m.Position[d] += m.Velocity[d]
	}
}

func (m Moon) Energy() float64 {
	return m.Position.Energy() * m.Velocity.Energy()
}

func applyGravity(m1, m2 *Moon) {
	for d := range m1.Position {
		if m1.Position[d] > m2.Position[d] {
			m1.Velocity[d] -= 1
			m2.Velocity[d] += 1
		} else if m1.Position[d] < m2.Position[d] {
			m1.Velocity[d] += 1
			m2.Velocity[d] -= 1
		}
	}
}

func ReadMoons(r io.Reader) []Moon {
	moons := []Moon{}
	for {
		m := Moon{}
		_, err := fmt.Fscanf(r, "<x=%f, y=%f, z=%f>\n", &m.Position[0], &m.Position[1], &m.Position[2])
		if err != nil {
			break
		}
		moons = append(moons, m)
	}
	return moons
}

func FindCycle(moons []Moon) (total int64) {
	/*
		Dimenions are independant from each other
		so we find a system cycle for each of the dimensions
		and we find a Least common multiple of those cycles.
	*/
	x := FindCyclesInDimension(moons, 0)
	y := FindCyclesInDimension(moons, 1)
	z := FindCyclesInDimension(moons, 2)
	return LCM(x, y, z)
}

func LCM(vals ...int64) int64 {
	/*
		Least common multiplier
	*/
	for len(vals) > 1 {
		v := vals[0] * vals[1] / GCD(vals[0], vals[1])
		vals = vals[1:]
		vals[0] = v
	}
	return vals[0]
}

func GCD(a, b int64) int64 {
	/*
		Greatest Commond Divisor
	*/
	for b != 0 {
		a, b = b, a%b
	}
	return int64(a)
}

func FindCyclesInDimension(moons []Moon, d int) (cycle int64) {
	/*
		Return the number of cycles until the system goes back to the same state on given
		dimension
	*/
	// copy initial state
	initial := make([]Moon, len(moons))
	copy(initial, moons)

	stop := false
	for cycle = 0; !stop; cycle++ {
		// apply Gravity
		for j := range moons {
			for k := j + 1; k < len(moons); k++ {
				applyGravity(&moons[j], &moons[k])
			}
		}
		// move Moons
		for j := range moons {
			moons[j].Move()
		}
		// for each moon check id it's position and velocity in given dimension
		// is the same as initial state
		stop = true
		for m := range moons {
			stop = stop && (moons[m].Position[d] == initial[m].Position[d] && moons[m].Velocity[d] == initial[m].Velocity[d])
		}
	}
	return cycle
}

func Part1(moons []Moon, n int) (total_energy float64) {
	for ; n > 0; n-- {
		for idx := range moons {
			for idx2 := range moons {
				if idx2 < idx {
					continue
				}
				applyGravity(&moons[idx], &moons[idx2])
			}
		}
		for idx := range moons {
			(&moons[idx]).Move()
		}
	}
	for idx := range moons {
		total_energy += moons[idx].Energy()
	}
	return total_energy
}

func main() {
	moons := ReadMoons(os.Stdin)
	fmt.Println(Part1(moons, 1000))
	fmt.Println(FindCycle(moons))
}
