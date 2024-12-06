package part1

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	want := 41
	got := Solve(string(input))

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
