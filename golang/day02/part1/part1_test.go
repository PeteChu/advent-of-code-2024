package part1

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input_small.txt")
	if err != nil {
		t.Errorf("Error reading input file: %v", err)
	}

	want := 2
	got := Solve(string(input))

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
