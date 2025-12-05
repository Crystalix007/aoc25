package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day5"
)

func main() {
	db := day5.Parse(os.Stdin)

	freshRange := db.FreshRange()

	fmt.Printf("Fresh ingredient count: %d\n", freshRange)
}
