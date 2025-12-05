package day5_test

import (
	"os"
	"testing"

	"github.com/crystalix007/aoc25/day5"
	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"
)

func TestDB_FreshRange(t *testing.T) {
	sampleFile, err := os.Open("testdata/sample.txt")
	assert.NoErr[F](t, err)

	db := day5.Parse(sampleFile)

	freshRange := db.FreshRange()

	assert.Equal[E](t, 19, freshRange)
}
