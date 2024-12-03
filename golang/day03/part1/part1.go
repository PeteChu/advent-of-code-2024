package part1

import (
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/3#part1
func Solve(input string) int {
	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)
	instructions := re.FindAllString(input, -1)

	re = regexp.MustCompile(`[mul\(|\)]`)

	sum := 0
	for _, v := range instructions {
		o := strings.Split(
			re.ReplaceAllString(v, ""),
			",",
		)
		x, _ := strconv.Atoi(o[0])
		y, _ := strconv.Atoi(o[1])

		sum += x * y
	}
	return sum
}
