package day6_test

import (
	_ "embed"
	"strings"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"

	"github.com/crystalix007/aoc25/day6"
)

//go:embed testdata/sample.txt
var str string

func TestParseLTR(t *testing.T) {
	t.Parallel()

	expected := day6.Homework{
		day6.Column{
			Values: []int{
				123,
				45,
				6,
			},
			Operator: day6.MULT,
		},
		day6.Column{
			Values: []int{
				328,
				64,
				98,
			},
			Operator: day6.ADD,
		},
		day6.Column{
			Values: []int{
				51,
				387,
				215,
			},
			Operator: day6.MULT,
		},
		day6.Column{
			Values: []int{
				64,
				23,
				314,
			},
			Operator: day6.ADD,
		},
	}

	r := strings.NewReader(str)

	assert.Equal[E](t, day6.ParseLTR(r), expected)
}

func TestParseRTL(t *testing.T) {
	t.Parallel()

	expected := day6.Homework{
		day6.Column{
			Values: []int{
				4,
				431,
				623,
			},
			Operator: day6.ADD,
		},
		day6.Column{
			Values: []int{
				175,
				581,
				32,
			},
			Operator: day6.MULT,
		},
		day6.Column{
			Values: []int{
				8,
				248,
				369,
			},
			Operator: day6.ADD,
		},
		day6.Column{
			Values: []int{
				356,
				24,
				1,
			},
			Operator: day6.MULT,
		},
	}

	r := strings.NewReader(str)

	assert.Equal[E](t, day6.ParseRTL(r), expected)
}
