package algorithms

type Tuple3 struct {
	P_1, P_2, P_3 int
}

func FloydWarshal(edges [][3]int) map[Tuple3]int {
	dist := make(map[Tuple3]int)
	n := len(edges)
	for _, edge := range edges {
		dist[Tuple3{edge[0], edge[1], -1}] = edge[2]
	}

	for k := range n {
		for i := range n {
			for j := range n {
				dist[Tuple3{i, j, k}] = min(dist[Tuple3{i, k, k - 1}]+dist[Tuple3{k, i, k - 1}], dist[Tuple3{i, j, k - 1}])
			}
		}
	}

	return dist
}
