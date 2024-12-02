package main

import (
	"aoc-2024/day01/part2"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("day01/part2/input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	result := part2.Solve(string(input))
	fmt.Printf("Day 01 - Part 2: %d\n", result)
}
