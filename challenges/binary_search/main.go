package main

import (
	"errors"
	"fmt"
)

func main() {

	i, error := exact([]int{1, 2, 4, 5}, 6)
	fmt.Println("Found exact", i, "error", error)

	i, error = closestBigger([]int{3, 5, 9, 13}, 5)
	fmt.Println("Found bigger", i, "error", error)
}

// 3, 5, 7, 1, 2, 4, 6

// Finds
func closestBigger(sorted []int, val int) (int, error) {
	i := 0
	j := len(sorted) - 1

	for i <= j {
		m := (i + j) / 2
		if val <= sorted[i] {
			return i, nil
		} else if val < sorted[m] {
			j = m
		} else if sorted[m] <= val {
			i = m + 1
		}
	}
	return -1, errors.New("Bigger")
}

// Finds exact value in sorted array
func exact(sorted []int, val int) (int, error) {

	i := 0
	j := len(sorted) - 1

	for i <= j {
		m := (i + j) / 2
		if sorted[m] == val {
			return m, nil
		} else if sorted[m] < m {
			j = m - 1
		} else {
			i = m + 1
		}
	}

	return 0, errors.New("Not found")
}
