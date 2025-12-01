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

	// newRotations := make([]day1.Rotation, 0, len(rotations))

	// const maxRotation = 1

	// for _, r := range rotations {
	// 	for r.Amount > maxRotation {
	// 		newRotations = append(newRotations, day1.Rotation{
	// 			Direction: r.Direction,
	// 			Amount:    maxRotation,
	// 		})
	// 		r.Amount -= maxRotation
	// 	}

	// 	if r.Amount > 0 {
	// 		newRotations = append(newRotations, day1.Rotation{
	// 			Direction: r.Direction,
	// 			Amount:    r.Amount,
	// 		})
	// 	}
	// }

	// clicks := day1.Clicks(50, newRotations)
	clicks := day1.Clicks(50, rotations)

	fmt.Printf("Clicks: %d\n", clicks)
}
