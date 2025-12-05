package day2_test

import (
	"strings"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"

	"github.com/crystalix007/aoc25/day2"
)

func TestParse(t *testing.T) {
	input := "1-3,5-7,10-12"
	expected := []day2.Range{
		{Begin: 1, End: 3},
		{Begin: 5, End: 7},
		{Begin: 10, End: 12},
	}

	ranges := day2.Parse(strings.NewReader(input))

	assert.Equal[E](t, expected, ranges)
}
