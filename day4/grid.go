package day4

type Grid [][]bool

func (g Grid) CanAccess(x, y int) bool {
	// Bounds-check.
	if y < 0 || y >= len(g) ||
		x < 0 || x >= len(g[y]) {
		return false
	}

	var nearCount int

	for dy := -1; dy <= 1; dy++ {
		if y+dy < 0 ||
			y+dy >= len(g) {
			continue
		}

		for dx := -1; dx <= 1; dx++ {
			if x+dx < 0 ||
				x+dx >= len(g[y+dy]) {
				continue
			}

			// Skip self.
			if dx == 0 && dy == 0 {
				continue
			}

			if g[y+dy][x+dx] {
				nearCount++
			}
		}
	}

	return nearCount < 4
}

func (g Grid) IsAccessiblePaper(x, y int) bool {
	if y < 0 || y >= len(g) ||
		x < 0 || x >= len(g[y]) {
		return false
	}

	return g[y][x] && g.CanAccess(x, y)
}

func (g Grid) RemoveAccessiblePaper() int {
	var removed int

	for y := range g {
		for x := range g[y] {
			if g.IsAccessiblePaper(x, y) {
				g[y][x] = false
				removed++
			}
		}
	}

	return removed
}
