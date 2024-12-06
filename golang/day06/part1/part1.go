package part1

import (
	"fmt"
	"regexp"
	"strings"
)

type StartPoint struct {
	row int
	col int
}

func BuildGrid(input string) (StartPoint, [][]string) {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	re := regexp.MustCompile(`\^|v|<|>`)
	var startPoint StartPoint
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
		index := re.FindStringIndex(line)
		if len(index) > 0 {
			startPoint = StartPoint{i, index[0]}
		}
	}
	return startPoint, grid[:len(grid)-1]
}

func Move(grid [][]string, point *StartPoint, visited map[string]bool, direction string) int {
	if point.row < 0 || point.row > len(grid)-1 || point.col < 0 || point.col > len(grid[0])-1 {
		return len(visited)
	}

	switch direction {
	case "^":
		if point.row-1 < 0 {
			return len(visited)
		}
		point.row--
		if point.row < 0 || grid[point.row][point.col] == "#" {
			point.row++
			direction = ">"
		}
	case "<":
		if point.col-1 < 0 {
			return len(visited)
		}
		point.col--
		if point.col < 0 || grid[point.row][point.col] == "#" {
			point.col++
			direction = "^"
		}
	case ">":
		if point.col+1 > len(grid[0])-1 {
			return len(visited)
		}
		point.col++
		if point.col > len(grid[0])-1 || grid[point.row][point.col] == "#" {
			point.col--
			direction = "v"
		}
	case "v":
		if point.row+1 > len(grid)-1 {
			return len(visited)
		}
		point.row++
		if point.row > len(grid)-1 || grid[point.row][point.col] == "#" {
			point.row--
			direction = "<"
		}
	}

	key := fmt.Sprintf("%d-%d", point.row, point.col)
	visited[key] = true
	return Move(grid, point, visited, direction)
}

func Solve(input string) int {
	point, grid := BuildGrid(input)
	visited := make(map[string]bool)
	key := fmt.Sprintf("%d-%d", point.row, point.col)
	visited[key] = true
	return Move(grid, &point, visited, grid[point.row][point.col])
}
