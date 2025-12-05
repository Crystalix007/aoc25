package day4_test

import (
	"os"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"

	"github.com/crystalix007/aoc25/day4"
)

func TestGrid_CanAccess(t *testing.T) {
	const (
		x = true
		o = false
	)

	grid := day4.Grid{
		{x, x, x, x},
		{x, x, x, o},
		{x, x, o, o},
		{x, o, o, o},
	}

	expectedAccess := [][]bool{
		{true, false, false, true},
		{false, false, false, true},
		{false, false, true, true},
		{true, true, true, true},
	}

	gotAccess := make([][]bool, len(grid))

	for y := range grid {
		gotAccess[y] = make([]bool, len(grid[y]))

		for x := range grid[y] {
			gotAccess[y][x] = grid.CanAccess(x, y)
		}
	}

	assert.Equal[E](t, expectedAccess, gotAccess)
}

func TestGrid_CanAccess_sample(t *testing.T) {
	sampleFile, err := os.Open("testdata/sample.txt")
	assert.NoErr[F](t, err)

	sample := day4.Parse(sampleFile)

	expectedFile, err := os.Open("testdata/sample_expected.txt")
	assert.NoErr[F](t, err)

	expected := day4.Parse(expectedFile)

	for y := range sample {
		for x := range sample[y] {
			// Only test cells that are marked as having rolls of paper.
			if !sample[y][x] {
				continue
			}

			gotAccess := sample.CanAccess(x, y)

			// Expected will be true if the cell is _not_ marked accessible (but
			// if it has rolls of paper).
			expected := !expected[y][x]

			assert.Equal[E](t, expected, gotAccess)
		}
	}
}
