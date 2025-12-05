package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day4"
)

func main() {
	grid := day4.Parse(os.Stdin)

	var count int

	for y := range grid {
		row := make([]rune, len(grid[y]))

		for x := range grid[y] {
			isPaper := grid[y][x]
			isAccessible := grid.IsAccessiblePaper(x, y)

			switch {
			case isPaper && !isAccessible:
				row[x] = '@'
			case isPaper && isAccessible:
				row[x] = 'x'
			default:
				row[x] = '.'
			}

			if isAccessible {
				count++
			}
		}

		fmt.Println(string(row))
	}

	fmt.Printf("\n")
	fmt.Printf("Total accessible paper cells: %d\n", count)
}
