package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day4"
)

func main() {
	grid := day4.Parse(os.Stdin)

	var count int

	for {
		removed := grid.RemoveAccessiblePaper()
		count += removed

		if removed == 0 {
			break
		}
	}

	for y := range grid {
		row := make([]rune, len(grid[y]))

		for x := range grid[y] {
			isPaper := grid[y][x]

			switch {
			case isPaper:
				row[x] = 'x'
			default:
				row[x] = '.'
			}
		}

		fmt.Println(string(row))
	}

	fmt.Printf("\n")
	fmt.Printf("Total accessible paper cells: %d\n", count)
}
