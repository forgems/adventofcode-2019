package main

import (
	"strings"
	"testing"
)

func TestChecksum(t *testing.T) {
	// t.Fatal("not implemented")
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`
	checksum := orbitChecksum(readMap(strings.NewReader(input)))
	if checksum != 42 {
		t.Error("Invalid checksum", checksum)
	}
}

func TestOrbitTransfer(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
`
	m := readMap(strings.NewReader(input))
	transfers := orbitTransfer(m, "YOU", "SAN")
	if transfers != 4 {
		t.Error("Invalid checksum", transfers)
	}
}
