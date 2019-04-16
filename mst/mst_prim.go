package mst

import (
	"fmt"
	"go-algorithms/pq"
)

// LazyPrimMST is Prim's algorithm implementation of mst.
type LazyPrimMST struct {
	mst    []Edge
	weight float64

	marked []bool
	queue  pq.PQ
}

// NewLazyPrimMST constructs a MST by visiting edges with smaller weights first.
func NewLazyPrimMST(g EdgeWeightedGraph) MST {
	l := &LazyPrimMST{
		mst:    make([]Edge, 0),
		weight: 0,
		queue:  pq.NewMinPQ(1),
		marked: make([]bool, g.V()),
	}

	l.visit(g, 0)

	for !l.queue.IsEmpty() && len(l.mst) < g.V()-1 {
		key := l.queue.Poll()
		edge := key.(*WeightedEdge)

		v := edge.Either()
		w := edge.Other(v)

		if l.marked[v] && l.marked[w] {
			continue
		}

		l.mst = append(l.mst, edge)
		l.weight += edge.Weight()

		if !l.marked[v] {
			l.visit(g, v)
		}

		if !l.marked[w] {
			l.visit(g, w)
		}
	}
	return l
}

func (l *LazyPrimMST) visit(g EdgeWeightedGraph, v int) {
	l.marked[v] = true
	for _, e := range g.Adj(v) {
		if !l.marked[e.Other(v)] {
			edge := e.(*WeightedEdge)
			l.queue.Insert(edge)
		}
	}
}

// Edges in the mst.
func (l *LazyPrimMST) Edges() []Edge {
	return l.mst
}

// Weight of the mst.
func (l *LazyPrimMST) Weight() float64 {
	return l.weight
}

func (l *LazyPrimMST) String() string {
	str := ""
	for _, e := range l.mst {
		v := e.Either()
		w := e.Other(v)
		str += fmt.Sprintf("%d-%d %.2f\n", v, w, e.Weight())
	}
	return str
}
