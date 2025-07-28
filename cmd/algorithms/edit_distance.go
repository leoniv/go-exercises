package algorithms

func EditDistance(w10 string, w20 string) [][]int {
	w1 := []rune(w10)
	w2 := []rune(w20)
	calc := make([][]int, len(w1)+1)
	for i := range calc {
		calc[i] = make([]int, len(w2)+1)
	}

	for i := range calc {
		for j := range calc[i] {
			println("i, j", i, j)
			switch {
			case i == 0 && j == 0: //nop
			case j == 0:
				calc[i][j] = calc[i-1][j] + 1
			case i == 0:
				calc[i][j] = calc[i][j-1] + 1
			default:
				calc[i][j] = min(
					calc[i-1][j]+1,
					calc[i][j-1]+1,
					func() int {
						if w1[i-1] == w2[j-1] {
							return 0
						} else {
							return 1
						}
					}()+calc[i-1][j-1],
				)
			}
		}
	}

	return calc
}
