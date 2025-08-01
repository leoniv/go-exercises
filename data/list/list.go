package list

import (
	. "local/go-exercises/data"
)

type List[T any] struct {
	Value T
	next  *List[T]
}

// [1,2,3] -> (1, (2, (3)))
func MakeList[T any](in []T) *List[T] {
	var head *List[T]
	for i := len(in) - 1; i > -1; i-- {
		head = &List[T]{Value: in[i], next: head}
	}
	return head
}

func Fold[A any, B any](ls *List[A], zero B, f func(A, B) B) B {
	for ls != nil {
		zero = f(ls.Value, zero)
		ls = ls.next
	}
	return zero
}

func (l *List[T]) ToSlice() []T {
	current := l
	var result []T
	for current != nil {
		result = append(result, current.Value)
		current = current.next
	}

	return result
}

/*
* Here is chain of flips from right to left
* (nil)   1 -> 2 -> 3 -> nil // flipRight 1
* (nil <- 1)   2 -> 3 -> nil // flipRight 2
* (nil <- 1 <- 2)   3 -> nil // flipRight 3
* (nil <- 1 <- 2 <- 3)
 */
func (rest *List[T]) Reverse() *List[T] {
	var head *List[T]
	head, rest = flipRight(head, rest)
	for rest != nil {
		head, rest = flipRight(head, rest)
	}

	return head
}

// (. <- left) (rigth -> .) flip to (. <- left <- right) (right.Next -> .)
func flipRight[T any](left *List[T], right *List[T]) (*List[T], *List[T]) {
	next := right.next // (. <- l)   r  -> nx -> .
	right.next = left  // (. <- l <- r)    nx -> .
	return right, next // (. <- l <- r)    nx -> .
}

func Map[A, B any](ls *List[A], f func(A) B) *List[B] {

	panic("FIXME")
}

/*
	  nil
		(1, 1) -> nil
		(1, 1) -> (2, 2) -> nil
		(1, 1) -> (2, 2) -> (3, 3) -> nil
*/
func Zip[A any, B any](ls *List[A], rs *List[B]) *List[Pair[A, B]] {
	var head *List[Pair[A, B]]
	var tail *List[Pair[A, B]]
	init := true

	makeNode := func(a A, b B) *List[Pair[A, B]] {
		return &List[Pair[A, B]]{
			Value: Pair[A, B]{F1: a, F2: b},
			next:  nil,
		}
	}

	for ls != nil && rs != nil {
		if init {
			init = false
			head = makeNode(ls.Value, rs.Value)
			tail = head
		} else {
			tail.next = makeNode(ls.Value, rs.Value)
			tail = tail.next
		}
		ls = ls.next
		rs = rs.next
	}
	return head
}

func (ls *List[A]) Next() *List[A] {
	return ls.next
}
