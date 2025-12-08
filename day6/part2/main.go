package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day6"
)

func main() {
	homework := day6.ParseRTL(os.Stdin)

	homework.Print(os.Stderr)

	vm := homework.VerticalMaths()
	var total int

	for _, result := range vm {
		fmt.Printf("%d\n", result)
		total += result
	}

	fmt.Printf("\n")
	fmt.Printf("Grand total: %d\n", total)
}
