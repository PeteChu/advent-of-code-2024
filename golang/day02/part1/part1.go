package part1

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) int {
	reports := strings.Split(input, "\n")

	count := 0
	for _, v := range reports {
		if v == "" {
			continue
		}

		safe := true

		inc := false
		des := false

		levels := strings.Fields(v)
		fmt.Println(levels)
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

		if safe {
			count++
			fmt.Println("Safe")
		} else {
			fmt.Println("Unsafe")
		}
	}
	return count
}
