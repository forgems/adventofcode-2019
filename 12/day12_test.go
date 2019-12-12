package main

import (
	"fmt"
	"testing"
)

func TestGravity(t *testing.T) {
	moons := []Moon{
		Moon{
			Position: Vec3D{-1, 0, 2},
		},
		Moon{
			Position: Vec3D{2, -10, -7},
		},
		Moon{
			Position: Vec3D{4, -8, 8},
		},
		Moon{
			Position: Vec3D{3, 5, -1},
		},
	}
	for s := 0; s < 10; s++ {
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
	final_positions := []struct {
		v      Vec3D
		energy float64
	}{
		{Vec3D{2, 1, -3}, 36},
		{Vec3D{1, -8, 0}, 45},
		{Vec3D{3, -6, 1}, 80},
		{Vec3D{2, 0, 4}, 18},
	}

	for m := range moons {
		if !moons[m].Position.Equal(final_positions[m].v) ||
			moons[m].Energy() != final_positions[m].energy {
			t.Errorf(
				"Invalid position for moon(%f) %d = %s. Should be %v",
				moons[m].Energy(), m, moons[m], final_positions[m],
			)
		}
	}

}

func TestFindCycles(t *testing.T) {
	test_data := []struct {
		moons []Moon
		cycle int64
	}{
		{[]Moon{
			Moon{
				Position: Vec3D{-1, 0, 2},
			},
			Moon{
				Position: Vec3D{2, -10, -7},
			},
			Moon{
				Position: Vec3D{4, -8, 8},
			},
			Moon{
				Position: Vec3D{3, 5, -1},
			},
		}, 2772},
		{[]Moon{
			Moon{
				Position: Vec3D{-8, -10, 0},
			},
			Moon{
				Position: Vec3D{5, 5, 10},
			},
			Moon{
				Position: Vec3D{2, -7, 3},
			},
			Moon{
				Position: Vec3D{9, -8, -3},
			},
		}, 4686774924},
	}
	for d := range test_data {
		i := d
		t.Run(fmt.Sprintf("TestCycle %d", i), func(t *testing.T) {
			v := FindCycle(test_data[i].moons)
			if v != test_data[i].cycle {
				t.Errorf("Invalid value %d. Should be %d", v, test_data[i].cycle)
			}
		})
	}

}
