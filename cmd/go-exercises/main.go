package main

import (
	"fmt"
	"local/go-exercises/cmd/algorithms"
)

func main() {
	fmt.Println(algorithms.MatrixMinMult([][2]int{
		{50, 20},
		{20, 1},
		{1, 10},
		{10, 100},
	}))
}
