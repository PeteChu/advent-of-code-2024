package part2

import (
	"regexp"
	"strings"
)

func makeGrid(n int) [][]string {
	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, n)
	}
	return grid
}

func fillGrid(grid [][]string, lines []string) [][]string {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = string(lines[i][j])
		}
	}
	return grid
}

func findXMAS(grid [][]string) int {
	n := len(grid)

	pattern := "MMASS|SSAMM|SMASM|MSAMS"
	re := regexp.MustCompile(pattern)

	builder := strings.Builder{}
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] != "A" {
				continue
			}
			tl, tr, mid, bl, br := grid[i-1][j-1], grid[i-1][j+1], grid[i][j], grid[i+1][j-1], grid[i+1][j+1]
			builder.WriteString(tl + tr + mid + bl + br)
		}
	}
	return len(re.FindAllString(builder.String(), -1))
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	size := len(lines) - 1
	grid := fillGrid(makeGrid(size), lines)
	return findXMAS(grid)
}
