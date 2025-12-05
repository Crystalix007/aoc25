package day2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Range struct {
	Begin int
	End   int
}

func Parse(r io.Reader) []Range {
	var ranges []Range

	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF || len(data) == 0 {
			return 0, nil, nil
		}

		next := bytes.Index(data, []byte{','})
		if next == -1 {
			return len(data), data, nil
		}

		token := data[:next]

		return next + 1, token, nil
	})

	for scanner.Scan() {
		var begin, end int
		line := scanner.Text()

		fmt.Sscanf(line, "%d-%d", &begin, &end)
		ranges = append(ranges, Range{Begin: begin, End: end})
	}

	return ranges
}
