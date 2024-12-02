package part2

import (
	"strconv"
	"strings"
)

func isSafe(levels []string) bool {
	safe := true

	inc := false
	des := false

	for i := 0; i < len(levels)-1; i++ {

		r, _ := strconv.Atoi(levels[i+1])
		l, _ := strconv.Atoi(levels[i])

		if r > l {
			inc = true
		} else if r < l {
			des = true
		} else {
			safe = false
			break
		}
		if inc && des {
			safe = false
			break
		}

		diff := l - r
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			safe = false
			break
		}
	}

	return safe
}

func Solve(input string) int {
	reports := strings.Split(input, "\n")

	count := 0
	for _, v := range reports {
		if v == "" {
			continue
		}
		levels := strings.Fields(v)
		if isSafe(levels) {
			count++
		} else {
			for i := 0; i < len(levels); i++ {
				modifiedReport := append([]string{}, levels[:i]...)
				modifiedReport = append(modifiedReport, levels[i+1:]...)
				if isSafe(modifiedReport) {
					count++
					break
				}
			}
		}
	}
	return count
}
