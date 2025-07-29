package main

import (
	"fmt"
	"local/go-exercises/cmd/algorithms"
	"math"
)

const (
	Inf = math.MaxInt
)

func main() {
	fmt.Println(algorithms.FloydWarshal([][3]int{
		// A - 0, B - 1, C - 2, D - 3, S - 4, T - 5
		{0, 1, 2},
		{0, 3, 5},
		{0, 4, 1},
	}))
}
