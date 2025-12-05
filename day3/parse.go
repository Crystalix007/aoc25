package day3

import (
	"bufio"
	"io"
)

type Bank []int

type Batteries []Bank

func ParseBatteries(r io.Reader) Batteries {
	var batteries Batteries

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		var bank Bank
		for _, ch := range line {
			bank = append(bank, int(ch-'0'))
		}

		batteries = append(batteries, bank)
	}

	return batteries
}
