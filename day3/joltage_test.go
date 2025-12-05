package day3_test

import (
	"fmt"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"

	"github.com/crystalix007/aoc25/day3"
)

var sampleBatteries = day3.Batteries{
	day3.Bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
	day3.Bank{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
	day3.Bank{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
	day3.Bank{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
}

func TestLargestJoltagePair(t *testing.T) {
	for i, tt := range []struct {
		expected int
	}{
		{
			expected: 98,
		},
		{
			expected: 89,
		},
		{
			expected: 89,
		},
		{
			expected: 92,
		},
	} {
		t.Run(fmt.Sprintf("%v", sampleBatteries[i]), func(t *testing.T) {
			result := day3.LargestJoltagePair(sampleBatteries[i])

			assert.Equal[E](t, tt.expected, result)
		})
	}
}

func TestLargestJoltageInN(t *testing.T) {
	const testN = 12

	for _, tt := range []struct {
		expected int
	}{
		{
			expected: 987654321111,
		},
		{
			expected: 811111111119,
		},
		{
			expected: 434234234278,
		},
		{
			expected: 888911112111,
		},
	} {
		t.Run(fmt.Sprintf("%v", sampleBatteries[0]), func(t *testing.T) {
			result := day3.LargestJoltageInN(sampleBatteries[0], testN)

			assert.Equal[E](t, tt.expected, result)
		})
	}
}
