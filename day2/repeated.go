package day2

import (
	"math"
)

// IsRepeatedDigit checks if the number contains only a once-repeated sequence
// of digits.
//
// I.e. 123123 is true.
// But  1234321 is false.
func IsRepeatedDigit(num int) bool {
	digits := int(math.Log10(float64(num))) + 1

	if digits%2 != 0 {
		return false
	}

	halfDigits := digits / 2
	powTen := int(math.Pow10(halfDigits))

	firstHalf := num / powTen
	secondHalf := num % powTen

	return firstHalf == secondHalf
}

// IsRepeatedAnyDigit checks if the number contains any repeated sequence
// of digits.
//
// Note that this means sequences can repeat more than once.
//
// I.e. 1212 is true.
//
//	121212 is also true.
//
// But  1234 is false.
func IsRepeatedAnyDigit(num int) bool {
	digits := int(math.Log10(float64(num))) + 1

SIZE_LOOP:
	for size := 1; size <= digits/2; size++ {
		if digits%size != 0 {
			continue
		}

		powTen := int(math.Pow10(size))
		num := num
		pattern := num % powTen

		for i := size; i < digits; i += size {
			num /= powTen
			if num%powTen != pattern {
				continue SIZE_LOOP
			}
		}

		return true
	}

	return false
}
