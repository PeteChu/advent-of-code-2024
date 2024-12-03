package part2

import (
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input_small.txt")
	if err != nil {
		t.Errorf(`Error while reading input file %v`, err)
	}

	want := 4
	got := Solve(string(input))

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
