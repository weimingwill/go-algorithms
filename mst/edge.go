package mst

import "go-algorithms/pq"

// Edge defines interface for edge in graph.
type Edge interface {
	// Either endpoint of edge
	Either() int

	// The endpoint that is not v
	Other(v int) int

	// Weight of the edge
	Weight() float64

	// Cmp compare two edges
	Cmp(Edge) int
}

// WeightedEdge is edge with two vertics and has weight on the edge.
type WeightedEdge struct {
	v      int
	w      int
	weight float64
}

// NewWeightedEdge constructs a new weighted edge.
func NewWeightedEdge(v, w int, weight float64) Edge {
	return &WeightedEdge{
		v:      v,
		w:      w,
		weight: weight,
	}
}

// Either returns either vertix of the edge.
func (e *WeightedEdge) Either() int {
	return e.v
}

// Other returns the other vertix of the edge other than the input one.
func (e *WeightedEdge) Other(v int) int {
	if e.v == v {
		return e.w
	}

	return e.v
}

// Weight returns weight of the edge.
func (e *WeightedEdge) Weight() float64 {
	return e.weight
}

// Cmp compares weight of two edges.
func (e *WeightedEdge) Cmp(edge Edge) int {
	return e.compare(edge)
}

func (e *WeightedEdge) compare(edge Edge) int {
	if e.weight < edge.Weight() {
		return -1
	} else if e.weight > edge.Weight() {
		return 1
	} else {
		return 0
	}
}

// CompareTo implements priority queue Key.
// This is for the purpose of using WeightedEdge in priority queue.
// Better implementation is with golang generic.
func (e *WeightedEdge) CompareTo(k pq.Key) int {
	edge := k.(Edge)
	return e.compare(edge)
}

// Value returns weight of the weighted edge.
// It implements priority queue Key.
func (e *WeightedEdge) Value() interface{} {
	return e.weight
}
