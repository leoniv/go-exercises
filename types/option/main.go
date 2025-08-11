package main

func main() {
}

type some[A any] struct {
	value A
}
type none struct{}

type option[A any] interface {
	some[A] | none
}

type Option[A any] func(func(A), func())

func None[T any]() Option[T] {
	return func(_ func(T), zero func()) {
		zero()
	}
}

func Some[T any](val T) Option[T] {
	return func(some func(T), _ func()) {
		some(val)
	}
}

func Fold[A, B any](opt Option[A], succ func(A) B, zero func() B) B {
	var b B
	opt(
		func(val A) {
			b = succ(val)
		},
		func() {
			b = zero()
		},
	)
	return b
}

func Map[A, B any](opt Option[A], f func(A) B) Option[B] {
	return Fold[A, B](
		opt,
		func(val A) Option[B] { return Some(f(value)) },
		func() Option[B] { return None[B]() },
	)
}

func Flatten[A any](opt Option[Option[A]]) Option[A] {
	Fold(
		opt,
		func(value Option[A]) Option[A] { return value },
		func() Option[A] { return None[A]() },
	)
}
