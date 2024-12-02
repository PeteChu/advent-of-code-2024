package part1

import (
	"sort"
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

	sort.Ints(col1)
	sort.Ints(col2)

	sum := 0
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}
