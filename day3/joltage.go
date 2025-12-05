package day3

import (
	"fmt"
	"slices"
)

func LargestJoltagePair(bank Bank) int {
	if len(bank) == 0 {
		return 0
	}

	largestDigit := slices.Max(bank[:len(bank)-1])
	largestDigitPos := slices.Index(bank, largestDigit)

	largestNextDigit := slices.Max(bank[largestDigitPos+1:])

	return largestDigit*10 + largestNextDigit
}

func TotalMaxJoltagePairs(batteries Batteries) int {
	total := 0

	for _, bank := range batteries {
		total += LargestJoltagePair(bank)
	}

	return total
}

func LargestJoltageInN(bank Bank, n int) int {
	var (
		currentLeftIndex int
		total            int
	)

	for tailDigits := n - 1; tailDigits >= 0; tailDigits-- {
		usableBank := bank[currentLeftIndex : len(bank)-tailDigits]

		largestDigit := slices.Max(usableBank)
		largestDigitPos := slices.Index(usableBank, largestDigit)

		// Move the left index to the right of the found largest digit.
		//
		// We're using a subset of the bank slice, so we need to add
		// currentLeftIndex to get the correct position in the original bank.
		currentLeftIndex += largestDigitPos + 1

		total = total*10 + largestDigit
	}

	return total
}

func TotalMaxJoltageInN(batteries Batteries, n int) int {
	total := 0

	for _, bank := range batteries {
		joltage := LargestJoltageInN(bank, n)

		fmt.Printf("Selected: %d\n", joltage)

		total += joltage
	}

	return total
}
