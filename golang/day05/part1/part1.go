package part1

import (
	"strconv"
	"strings"
)

// --- Day 5: Print Queue ---
// https://adventofcode.com/2024/day/5
func Solve(input string) int {
	data := strings.Split(input, "\n\n")
	rules := strings.Split(data[0], "\n")
	updates := strings.Split(data[1], "\n")

	hashedRules := make(map[string]bool)
	for _, rule := range rules {
		hashedRules[rule] = true
	}

	sum := 0
	for _, update := range updates {
		x := strings.Split(update, ",")
		valid := true
		for i := 0; i < len(x)-1; i++ {
			key := x[i] + "|" + x[i+1]
			if _, ok := hashedRules[key]; !ok {
				valid = false
				break
			}
		}
		if valid {
			n, _ := strconv.Atoi(x[len(x)/2])
			sum += n
		}
	}

	return sum
}
