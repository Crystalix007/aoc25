package day4

import (
	"bufio"
	"io"
)

func Parse(r io.Reader) Grid {
	grid := Grid{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := make([]bool, len(scanner.Text()))

		for i, char := range scanner.Text() {
			row[i] = (char == '@')
		}

		grid = append(grid, row)
	}

	return grid
}
