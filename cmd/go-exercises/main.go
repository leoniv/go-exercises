package main

import (
	"fmt"
	"local/go-exercises/cmd/go-exercises/algorithms"
)

func main() {
	var k int
	fmt.Scan(&k)
	fmt.Println(algorithms.Selection([]int{2, 36, 5, 21, 8, 13, 11, 20, 5, 4, 1}, k))
}
