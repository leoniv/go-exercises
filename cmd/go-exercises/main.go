package main

import (
	"fmt"
	"local/go-exercises/cmd/go-exercises/algorithms"
)

func main() {
	fmt.Println(algorithms.RadixSort([]uint32{329, 457, 657, 839, 436, 720, 355}))
}
