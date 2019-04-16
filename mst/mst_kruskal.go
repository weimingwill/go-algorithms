package mst

import (
	"fmt"
	"go-algorithms/pq"
)

// KruskalMST is Kruskal's algorithm implementation of finding mst in a graph.
type KruskalMST struct {
	mst    []Edge
	weight float64
}

// NewKruskalMST constructs a MST using union find.
func NewKruskalMST(g EdgeWeightedGraph) MST {
	k := &KruskalMST{
		mst:    make([]Edge, 0),
		weight: 0,
	}

	queue := pq.NewMinPQ(1)
	for _, edge := range g.Edges() {
		e := edge.(*WeightedEdge)
		queue.Insert(e)
	}

	uf := NewUnionFind(g.V())
	for !queue.IsEmpty() && len(k.mst) < g.V()-1 {
		key := queue.Poll()
		edge := key.(*WeightedEdge)
		v := edge.Either()
		w := edge.Other(v)

		if !uf.Connected(v, w) {
			uf.Union(v, w)
			k.mst = append(k.mst, edge)
			k.weight += edge.weight
		}
	}
	return k
}

// Edges in the mst.
func (k *KruskalMST) Edges() []Edge {
	return k.mst
}

// Weight of the mst.
func (k *KruskalMST) Weight() float64 {
	return k.weight
}

func (k *KruskalMST) String() string {
	str := ""
	for _, e := range k.mst {
		v := e.Either()
		w := e.Other(v)
		str += fmt.Sprintf("%d-%d %.2f\n", v, w, e.Weight())
	}
	return str
}
