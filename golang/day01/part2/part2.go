package part2

import (
	"strconv"
	"strings"
)

func Solve(input string) int {
	col1 := make([]int, 0)
	col2 := make([]int, 0)

	inputs := strings.Split(input, "\n")
	for _, v := range inputs {
		if v == "" {
			continue
		}
		cols := strings.Fields(v)

		val, _ := strconv.Atoi(cols[0])
		col1 = append(col1, val)

		val, _ = strconv.Atoi(cols[1])
		col2 = append(col2, val)
	}

	freq := make(map[int]int)

	for _, v := range col2 {
		freq[v]++
	}

	sum := 0
	for _, v := range col1 {
		sum += v * freq[v]
	}
	return sum
}
