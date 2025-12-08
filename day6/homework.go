package day6

import (
	"fmt"
	"io"
	"strings"
)

func (h Homework) Print(o io.Writer) {
	for _, column := range h {
		var buffer strings.Builder

		for _, value := range column.Values {
			fmt.Fprintf(&buffer, "%d ", value)
		}

		fmt.Fprintf(o, "%s %c\n", buffer.String(), column.Operator)
	}
}
