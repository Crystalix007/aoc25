package main

import (
	"fmt"
	"os"

	"github.com/crystalix007/aoc25/day2"
)

func main() {
	parse := day2.Parse(os.Stdin)

	var total int64

	for _, r := range parse {
		for i := r.Begin; i <= r.End; i++ {
			if day2.IsRepeatedDigit(i) {
				fmt.Println(i)
				total += int64(i)
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
}
