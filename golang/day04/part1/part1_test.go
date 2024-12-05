package part1

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input_small.txt")
	if err != nil {
		t.Errorf("Error reading input file: %s", err)
	}

	want := 18
	got := Solve(string(input))

	if got != want {
		t.Errorf("Want %d, got %d", want, got)
	}
}
