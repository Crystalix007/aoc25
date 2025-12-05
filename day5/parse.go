package day5

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Parse(r io.Reader) DB {
	scanner := bufio.NewScanner(r)

	type State int

	const (
		ParsingFresh State = iota
		ParsingAvailable
	)

	var (
		state State = ParsingFresh
		db    DB
	)

	for scanner.Scan() {
		switch state {
		case ParsingFresh:
			if scanner.Text() == "" {
				state = ParsingAvailable
				continue
			}

			minStr, maxStr, _ := strings.Cut(scanner.Text(), "-")

			minStr, maxStr = strings.TrimSpace(minStr), strings.TrimSpace(maxStr)

			min, err := strconv.ParseInt(minStr, 10, 64)
			if err != nil {
				panic(err)
			}

			max, err := strconv.ParseInt(maxStr, 10, 64)
			if err != nil {
				panic(err)
			}

			db.FreshIngredients = append(
				db.FreshIngredients,
				IngredientRange{
					Min: int(min),
					Max: int(max),
				},
			)

		case ParsingAvailable:
			ingredient, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}

			db.AvailableIngredients = append(
				db.AvailableIngredients,
				int(ingredient),
			)
		}
	}

	return db
}
