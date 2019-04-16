package mst

// MST defines minimun spanning tree interface.
type MST interface {
	// Edges in MST
	Edges() []Edge

	// Weight of MST
	Weight() float64
}
