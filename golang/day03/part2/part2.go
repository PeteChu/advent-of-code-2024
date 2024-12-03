package part2

import (
	"regexp"
	"strconv"
	"strings"
)

func Solve(input string) int {
	pattern := `(mul\(\d+,\d+\)|don\'t\(\)|do\(\))`
	re := regexp.MustCompile(pattern)
	instructions := re.FindAllString(input, -1)

	re = regexp.MustCompile(`[mul\(|\)]`)
	ignore := false

	sum := 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == "don't()" {
			ignore = true
		}

		if instructions[i] == "do()" {
			ignore = false
			continue
		}

		if ignore {
			continue
		}

		o := strings.Split(
			re.ReplaceAllString(instructions[i], ""),
			",",
		)
		x, _ := strconv.Atoi(o[0])
		y, _ := strconv.Atoi(o[1])

		sum += x * y
	}

	return sum
}
