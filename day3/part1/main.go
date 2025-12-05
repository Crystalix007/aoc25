package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day3"
)

func main() {
	batteries := day3.ParseBatteries(os.Stdin)

	total := day3.TotalMaxJoltagePairs(batteries)

	fmt.Printf("Total of largest joltage pairs: %d\n", total)
}
