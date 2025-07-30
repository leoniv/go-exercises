package algorithms

import (
	"fmt"
	"math"
)

// It finds minimal amount of matrix multiplications
// A X (B X C) not the same (A X B) X C
func MatrixMinMult(matrixs [][2]int) int {
	result := make([][]int, len(matrixs))
	for i := range result {
		result[i] = make([]int, len(matrixs))
		for j := range result {
			result[i][j] = func() int {
				if i == j {
					return 0
				} else {
					return math.MaxInt
				}
			}()
		}
	}

	m := getDimensions(matrixs)
	n := len(matrixs)
	fmt.Println(matrixs)
	fmt.Println(m)

	for s := 2; s <= n; s++ {
		for i := 0; i <= n-s; i++ {
			j := i + s - 1
			fmt.Println("s, (i, j)", s, i, j)
			for k := i; k < j; k++ {
				q := result[i][k] + result[k+1][j] + m[i]*m[k+1]*m[j+1]
				if result[i][j] > q {
					result[i][j] = q
				}
			}
		}
	}

	return result[0][n-1]
}

func getDimensions(matrixs [][2]int) []int {
	result := make([]int, len(matrixs)+1)
	for i, xs := range matrixs {
		result[i] = xs[0]
		result[i+1] = xs[1]
	}
	return result
}
