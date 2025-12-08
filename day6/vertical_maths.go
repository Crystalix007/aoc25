package day6

import "fmt"

func (h Homework) VerticalMaths() []int {
	results := make([]int, len(h))

	for i, column := range h {
		results[i] = column.Values[0]

		for _, value := range column.Values[1:] {
			// Perform the operation based on column.Operator
			results[i] = doOperator(results[i], column.Operator, value)
		}
	}

	return results
}

func doOperator(value1 int, operator Operator, value2 int) int {
	switch operator {
	case ADD:
		return value1 + value2
	case MULT:
		return value1 * value2
	}

	panic(fmt.Sprintf("unexpected operator: %v", operator))
}
