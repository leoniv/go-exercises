package main

import (
	"fmt"
	"math"
)

const (
	Infinity uint = math.MaxUint
)

func main() {
	var n int
	fmt.Scan(&n)
	coordinates := make([][2]int, n)
	visited := make([]bool, n)
	for i, _ := range visited {
		visited[i] = false
	}
	minDistance := make([]uint, n)
	for i, _ := range minDistance {
		minDistance[i] = Infinity
	}

	for i := 0; i < n; i++ {
		coord := [2]int{0, 0}
		fmt.Scan(&coord[0])
		fmt.Scan(&coord[1])
		coordinates[i] = coord
	}
	var maxRange, src, dest int
	fmt.Scan(&maxRange)
	fmt.Scan(&src)
	fmt.Scan(&dest)

	graph := MakeGraph(coordinates, maxRange)
	minDistance[src-1] = 0

	FindRoutes(graph, visited, minDistance, src-1)

	if minDistance[dest-1] == Infinity {
		fmt.Println(-1)
	} else {
		fmt.Println(minDistance[dest-1])
	}
}

func FindRoutes(graph map[int][]int, visited []bool, distance []uint, src int) {
	if visited[src] {
		return
	}
	visited[src] = true
	forVisit := graph[src]

	for _, city := range forVisit {
		distance[city] = min(distance[src]+1, distance[city])
	}
	for _, city := range forVisit {
		FindRoutes(graph, visited, distance, city)
	}
}

func MakeGraph(coordinates [][2]int, maxRange int) map[int][]int {
	result := make(map[int][]int)
	for i, point := range coordinates {
		for j, _ := range coordinates {
			if i != j && CalcDistanse(point, coordinates[j]) <= maxRange {
				result[i] = append(result[i], j)
			}
		}
	}
	return result
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func CalcDistanse(point1 [2]int, point2 [2]int) int {
	return Abs(point1[0]-point2[0]) + Abs(point1[1]-point2[1])
}
