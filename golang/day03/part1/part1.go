package part1

import (
	"regexp"
	"strconv"
)

// https://adventofcode.com/2024/day/3#part1
func Solve(input string) int {
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	instructions := re.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, v := range instructions {
		if len(v) != 3 {
			continue
		}
		x, _ := strconv.Atoi(v[1])
		y, _ := strconv.Atoi(v[2])

		sum += x * y
	}
	return sum
}
