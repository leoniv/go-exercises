package yandex

func FindGems(gems string, stones string) int32 {
	gemsMap := make(map[rune]int32)
	for _, gem := range gems {
		gemsMap[gem] = 0
	}

	for _, stone := range stones {
		if _, ok := gemsMap[stone]; ok {
			gemsMap[stone] += 1
		}
	}

	var result int32
	for _, b := range gemsMap {
		result += b
	}
	return result
}
