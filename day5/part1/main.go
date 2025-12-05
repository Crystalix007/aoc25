package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day5"
)

func main() {
	db := day5.Parse(os.Stdin)

	freshIngredientCount := db.CountFreshIngredients()

	fmt.Printf("\n")
	fmt.Printf("Fresh ingredients: %d\n", freshIngredientCount)
}
