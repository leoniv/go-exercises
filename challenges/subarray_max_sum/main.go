package main

import (
	"fmt"
)

func main() {

	cases := []struct {
		Ex uint64
		In []uint32
	}{
		{0, []uint32{}},
		{4, []uint32{1, 1, 1, 1}},
		{7, []uint32{1, 2, 3, 4}},
		{7, []uint32{1, 2, 2, 1, 1, 3, 1, 1}},
		{8, []uint32{1, 2, 2, 1, 1, 3, 3}},
		{9, []uint32{1, 2, 2, 1, 1, 1, 3, 6}},
		{8, []uint32{1, 2, 2, 1, 1, 1, 3, 2}},
		{120, []uint32{10, 5, 10, 10, 100}},
	}

	for _, case_ := range cases {
		result := solution(case_.In)
		fmt.Println("Passed:", case_.Ex == result, case_.In, "Expexted:", case_.Ex, "Actual:", result)
	}
}

/*
	  Дан массив неотрицательных интов.
		  Нужно найти непрерывный подмассив с наибольшей суммой элементов такой,
			в котором не более двух разных типов чисел. И вывести сумму.
			Например можно [5,10,5,10,5,10] тут всего два типа чисел: десятка и пятерка.
*/
func solution(in []uint32) uint64 {
	var maximum uint64
	state := state{stat{0, 0}, stat{0, 0}, stat{0, 0}}

	for _, val := range in {
		maximum = max(maximum, sum(state))
		state = inc(state, val)
	}

	return max(maximum, sum(state))
}

type state struct {
	a1 stat
	a2 stat
	a3 stat
}

type stat struct {
	val   uint32
	count uint32
}

func (s stat) inc() stat {
	s.count += 1
	return s
}

func (s stat) calc() uint64 {
	return uint64(s.val * s.count)
}

func sum(in state) uint64 {
	return in.a1.calc() + in.a2.calc()
}

func inc(in state, newVal uint32) state {
	if in.a1.val == newVal {
		return state{in.a2, in.a1.inc(), stat{in.a1.val, 1}}
	}

	if in.a2.val == newVal {
		return state{in.a1, in.a2.inc(), in.a3.inc()}
	}

	return state{in.a3, stat{newVal, 1}, stat{newVal, 1}}
}
