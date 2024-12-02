package main

import (
	"aoc-2024/day02/part2"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("day02/part2/input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	result := part2.Solve(string(input))
	fmt.Printf("Day 02 - Part 2: %d\n", result)
}
