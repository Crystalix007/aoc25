package day1

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var rotationRegex = regexp.MustCompile(`([LR])(\d+)`)

func ParseRotations(input io.Reader) ([]Rotation, error) {
	var rotations []Rotation

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		components := rotationRegex.FindStringSubmatch(line)
		if len(components) != 3 {
			break
		}

		rotation := Rotation{
			Direction: DecodeDirection(components[1]),
			Amount:    parseAmount(components[2]),
		}

		rotations = append(rotations, rotation)
	}

	return rotations, nil
}

func parseAmount(s string) int {
	amount, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid amount: %s", s))
	}

	return int(amount)
}
