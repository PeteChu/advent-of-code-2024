package part1

import (
	"strings"
)

func makeGrid(n int) [][]string {
	grid := make([][]string, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]string, n)
	}
	return grid
}

func fillGrid(grid [][]string, data []string) [][]string {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			grid[i][j] = string(data[i][j])
		}
	}
	return grid
}

func extract(grid [][]string) (lines []string) {
	for _, row := range grid {
		lines = append(lines, strings.Join(row, ""))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if j+3 < len(grid) && i+3 < len(grid) {
				x := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2] + grid[i+3][j+3]
				lines = append(lines, x)
			}
			if j-3 >= 0 && i+3 < len(grid) {
				y := grid[i][j] + grid[i+1][j-1] + grid[i+2][j-2] + grid[i+3][j-3]
				lines = append(lines, y)
			}
			if j+3 < len(grid) {
				z := grid[j][i] + grid[j+1][i] + grid[j+2][i] + grid[j+3][i]
				lines = append(lines, z)
			}
		}
	}

	return lines
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines) - 1
	grid := fillGrid(makeGrid(n), lines)
	sum := 0
	for _, v := range extract(grid) {
		sum += strings.Count(v, "XMAS")
		sum += strings.Count(v, "SAMX")
	}
	return sum
}
