package day2_test

import (
	"strconv"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"

	"github.com/crystalix007/aoc25/day2"
)

func TestIsRepeatedDigit(t *testing.T) {
	for _, tt := range []struct {
		num      int
		expected bool
	}{
		{num: 11, expected: true},
		{num: 112233, expected: false},
		{num: 123456, expected: false},
		{num: 122122, expected: true},
		{num: 1112111, expected: false},
	} {
		t.Run(strconv.FormatInt(int64(tt.num), 10), func(t *testing.T) {
			result := day2.IsRepeatedDigit(tt.num)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsAnyRepeatedDigit(t *testing.T) {
	for _, tt := range []struct {
		num      int
		expected bool
	}{
		{num: 11, expected: true},
		{num: 1212, expected: true},
		{num: 123123, expected: true},
		{num: 1234, expected: false},
		{num: 123123123, expected: true},
		{num: 1231123, expected: false},
		{num: 123123123123123, expected: true},
		{num: 999999999, expected: true},
	} {
		t.Run(strconv.FormatInt(int64(tt.num), 10), func(t *testing.T) {
			result := day2.IsRepeatedAnyDigit(tt.num)
			assert.Equal[E](t, tt.expected, result)
		})
	}
}
