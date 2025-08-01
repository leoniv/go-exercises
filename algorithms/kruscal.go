package algorithms

import "sort"

type SetEntry struct {
	Parent *SetEntry
	Num    int
	Rank   int
}

type Edg struct {
	V1     *SetEntry
	V2     *SetEntry
	Weight int
}

type Egs = []*Edg

func MakeGraph(in [][2]int, weights []int) Egs {
	vrtx := make(map[int]*SetEntry)
	var egs Egs
	if len(in) != len(weights) {
		panic("Len must be equal")
	}
	for i, e := range in {
		if _, ok := vrtx[e[0]]; ok {
			vrtx[e[0]] = MakeSet(e[0])
		}
		if _, ok := vrtx[e[1]]; ok {
			vrtx[e[1]] = MakeSet(e[1])
		}
		edge := Edg{vrtx[e[0]], vrtx[e[1]], weights[i]}
		egs = append(egs, &edge)
	}
	return egs
}

func Kruscal(egs Egs) Egs {
	var resul Egs
	for _, edge := range Sort(egs) {
		r1 := Find(edge.V1)
		r2 := Find(edge.V2)
		if r1 != r2 {
			resul = append(resul, edge)
			Union(r1, r2)
		}
	}
	return resul
}

func Sort(edges Egs) Egs {
	sort.Slice(edges, func(i int, j int) bool { return edges[i].Weight > edges[j].Weight })
	return edges
}

func MakeSet(num int) *SetEntry {
	result := SetEntry{
		nil,
		num,
		0,
	}
	result.Parent = &result
	return &result
}

func Find(v *SetEntry) *SetEntry {
	resul := v
	for resul.Parent != resul {
		resul = resul.Parent
	}
	return resul
}

func Union(v1 *SetEntry, v2 *SetEntry) {
	r1 := Find(v1)
	r2 := Find(v2)
	if r1 == r2 {
		return
	}
	if r1.Rank > r2.Rank {
		r2.Parent = r1
		return
	}

	r1.Parent = r2
	if r1.Rank == r2.Rank {
		r2.Rank += 1
	}
}
