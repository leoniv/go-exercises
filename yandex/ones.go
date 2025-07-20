package yandex

func SolveOnes(in []int32) int32 {
	var result, accum int32
	for _, i := range in {
		if i == 1 {
			accum += 1
		} else {
			result = max(result, accum)
			accum = 0
		}
	}

	return max(result, accum)
}
