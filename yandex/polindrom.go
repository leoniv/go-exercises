package yandex

import (
	"fmt"
	. "local/go-exercises/data/list"
)

func IsItPolindrom[T comparable](in *List[T]) bool {
	secondHalf := TakeSecondHalf(in).Reverse()
	fmt.Println(in.ToSlice())
	fmt.Println(secondHalf.ToSlice())

	loop := true
	for loop {
		if in.Value != secondHalf.Value {
			return false
		}
		loop = secondHalf.Next() != nil
		in = in.Next()
		secondHalf = secondHalf.Next()
	}
	return true
}

// It finds beginning second half of list
func TakeSecondHalf[T any](l *List[T]) *List[T] {
	if l == nil {
		return l
	}
	if l.Next() == nil {
		return l
	}
	// slow: 1, 2, 3, 4, 5, 6, 7
	// fast: 1,    3,    5,    7,  9
	slow := l
	fast := l
	even := false
	for fast.Next() != nil {
		slow = slow.Next()
		fast = fast.Next()
		even = !even
		if fast.Next() != nil {
			even = !even
			fast = fast.Next()
		}
	}
	if even {
		return slow
	} else {
		return slow.Next()
	}
}
