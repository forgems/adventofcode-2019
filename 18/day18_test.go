package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`,
			86},
		{`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`,
			132},
		{`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`, 136},
		{`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`, 81},
	}
	for i := range data {
		fmt.Printf("Test %d\n", i)
		m := ReadMap(strings.NewReader(data[i].s))
		distance := WalkMap(m, '@', Visited{}, nil)
		if distance != data[i].expected {
			t.Errorf("Invalid distance %d. Expected %d for %d", distance, data[i].expected, i)
		}
	}
}

func TestFindKeys(t *testing.T) {
	data := `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`
	m := ReadMap(strings.NewReader(data))
	fmt.Println(m.FindKeys())
}

func TestPart2(t *testing.T) {
	data := []struct {
		input    string
		expected int
	}{
		{`#######
#a.#Cd#
##@#@##
#######
##@#@##
#cB#Ab#
#######
`, 8},
		{`###############
#d.ABC.#.....a#
######@#@######
###############
######@#@######
#b.....#.....c#
###############`, 24},
		{`#############
#DcBa.#.GhKl#
#.###@#@#I###
#e#d#####j#k#
###C#@#@###J#
#fEbA.#.FgHi#
#############`, 32},
		{`#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba@#@BcIJ#
#############
#nK.L@#@G...#
#M###N#H###.#
#o#m..#i#jk.#
#############`, 72},
	}
	for i := range data {
		m := ReadMap(strings.NewReader(data[i].input))
		d := Part2(m)
		if d != data[i].expected {
			t.Errorf("%d. Invalid value %d. Should be %d", i, d, data[i].expected)
		}
	}
}
