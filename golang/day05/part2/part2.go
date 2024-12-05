package part2

import (
	"strconv"
	"strings"
)

// --- Day 5: Print Queue ---
// https://adventofcode.com/2024/day/5

func isOrdered(str []string, rules map[string]bool) bool {
	for i := 0; i < len(str)-1; i++ {
		key := str[i] + "|" + str[i+1]
		if _, ok := rules[key]; !ok {
			return false
		}
	}
	return true
}

func makeOrdered(str []string, rules map[string]bool) []string {
	for i := 0; i < len(str)-1; i++ {
		key := str[i] + "|" + str[i+1]
		if _, ok := rules[key]; !ok {
			str[i], str[i+1] = str[i+1], str[i]
		}
	}

	if !isOrdered(str, rules) {
		return makeOrdered(str, rules)
	}
	return str
}

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
		isValid := isOrdered(x, hashedRules)
		if isValid {
			continue
		}

		n := makeOrdered(x, hashedRules)
		y, _ := strconv.Atoi(n[len(n)/2])
		sum += y
	}

	return sum
}
