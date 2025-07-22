package algorithms

import "fmt"

func Selection(in []int, k int) int {
	fmt.Println("in: ", in, "k: ", k)
	sl, sv, sr := split(in, k)
	fmt.Println("sl: ", sl, "sv: ", sv, "sr: ", sr)
	if k <= len(sl) {
		return Selection(sl, k)
	} else if k <= len(sl)+len(sv) {
		return sv[0]
	} else {
		return Selection(sr, k-len(sl)-len(sv))
	}
}

func split(in []int, k int) ([]int, []int, []int) {
	sl := filter(in, func(i int) bool { return i < in[k-1] })
	sv := filter(in, func(i int) bool { return i == in[k-1] })
	sr := filter(in, func(i int) bool { return i > in[k-1] })

	return sl, sv, sr
}

func filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range input {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
