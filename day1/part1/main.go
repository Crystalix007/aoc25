package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day1"
)

func main() {
	rotations, err := day1.ParseRotations(os.Stdin)
	if err != nil {
		panic(err)
	}

	positions := day1.Positions(50, rotations)

	var zeroPositions int
	for i, pos := range positions {
		if pos == 0 {
			zeroPositions++
		}

		fmt.Fprintf(os.Stderr, "%03d: at %02d\n", i, pos)
	}

	fmt.Printf("Total times at 0: %d\n", zeroPositions)
}
