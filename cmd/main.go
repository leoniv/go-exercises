package main

import (
	"fmt"
	. "local/go-exercises/data/list"
	"local/go-exercises/yandex"
	"math"
)

const (
	Inf = math.MaxInt
)

func main() {
	list := func() *List[int] { return MakeList([]int{1, 2, 3, 4, 5, 6, 7}) }
	fmt.Println(list().ToSlice())
	fmt.Println(yandex.IsItPolindrom(list()))
	ll := MakeList([]int{1, 2, 3})
	lr := MakeList([]int{4, 3, 2, 6})
	fmt.Println(ll, lr)
	fmt.Println(Zip(ll, lr).ToSlice(), ll, lr)
	zero := 1
	fmt.Println(Fold(list(), zero, func(a int, b int) int { return a * b }), zero)
}
