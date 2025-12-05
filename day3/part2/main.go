package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day3"
)

func main() {
	batteries := day3.ParseBatteries(os.Stdin)

	total := day3.TotalMaxJoltageInN(batteries, 12)

	fmt.Printf("Total of largest joltage pairs: %d\n", total)
}
