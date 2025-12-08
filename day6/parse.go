package day6

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Operator rune

const (
	MULT    Operator = '*'
	ADD     Operator = '+'
	UNKNOWN Operator = '?'
)

type Column struct {
	Values   []int
	Operator Operator
}

type Homework []Column

func ParseLTR(r io.Reader) Homework {
	var homework Homework

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		if len(homework) == 0 {
			homework = make(Homework, len(fields))
		}

		// Last line of operators.
		if operatorFromString(fields[0]) != UNKNOWN {
			for i, field := range fields {
				homework[i].Operator = operatorFromString(field)
			}

			break
		}

		for i, field := range fields {
			value, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				panic(err)
			}

			homework[i].Values = append(homework[i].Values, int(value))
		}
	}

	return homework
}

func ParseRTL(r io.Reader) Homework {
	var (
		homework Homework
		digits   [][]int
	)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		// If this is the last line of operators.
		if operatorFromString(line[0:1]) != UNKNOWN {
			operators := strings.Fields(line)
			homework = make(Homework, len(operators))

			for i, operator := range operators {
				reversedIndex := len(operators) - 1 - i
				homework[reversedIndex].Operator = operatorFromString(operator)
			}

			break
		}

		var digitRow []int

		for _, char := range line {
			if char == ' ' {
				digitRow = append(digitRow, 0)
			} else {
				digitRow = append(digitRow, int(char-'0'))
			}
		}

		fmt.Printf("Read digit row: %v\n", digitRow)

		digits = append(digits, digitRow)
	}

	var column int

	for x := len(digits[0]) - 1; x >= 0; x-- {
		var value int

		for y := 0; y < len(digits); y++ {
			if len(digits[y]) <= x {
				panic(fmt.Sprintf(
					"unexpected column index %d for row %d, only have %d",
					x,
					y,
					len(digits[y]),
				))
			}

			if digits[y][x] == 0 {
				continue
			}

			value = value*10 + digits[y][x]
		}

		// Skip empty columns.
		if value == 0 {
			column++
			continue
		}

		homework[column].Values = append(homework[column].Values, value)
	}

	return homework
}

func operatorFromString(s string) Operator {
	if len(s) != 1 {
		return UNKNOWN
	}

	switch o := Operator(s[0]); o {
	case ADD, MULT:
		return o
	default:
		return UNKNOWN
	}
}
