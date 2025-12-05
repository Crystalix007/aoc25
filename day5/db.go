package day5

import (
	"fmt"
	"slices"
)

type IngredientRange struct {
	Min int
	Max int
}

type DB struct {
	FreshIngredients     []IngredientRange
	AvailableIngredients []int
}

func (i IngredientRange) Contains(ingredient int) bool {
	return ingredient >= i.Min && ingredient <= i.Max
}

func (db DB) IsIngredientFresh(ingredient int) bool {
	for _, ir := range db.FreshIngredients {
		if ir.Contains(ingredient) {
			return true
		}
	}

	return false
}

func (db DB) CountFreshIngredients() int {
	var count int

	for _, ingredient := range db.AvailableIngredients {
		fresh := db.IsIngredientFresh(ingredient)

		if fresh {
			fmt.Printf("Ingredient %d is fresh\n", ingredient)

			count++
		} else {
			fmt.Printf("Ingredient %d is spoiled\n", ingredient)
		}
	}

	return count
}

func (db DB) FreshRange() int {
	remainingRanges := make([]IngredientRange, len(db.FreshIngredients))
	copy(remainingRanges, db.FreshIngredients)

	var (
		minIndex   int
		freshRange int
	)

	for len(remainingRanges) > 0 {
		minRange := slices.MinFunc(
			remainingRanges,
			func(r1 IngredientRange, r2 IngredientRange) int {
				return r1.Min - r2.Min
			},
		)
		remainingRanges = slices.DeleteFunc(
			remainingRanges,
			func(r IngredientRange) bool {
				return r == minRange
			},
		)

		if minIndex > minRange.Max {
			continue
		}

		if minIndex < minRange.Min {
			minIndex = minRange.Min
		}

		freshRange += (minRange.Max - minIndex + 1)

		minIndex = minRange.Max + 1
	}

	return freshRange
}
