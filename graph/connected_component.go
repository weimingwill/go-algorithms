package graph

// ConnectedComponent represents interface for connected components in graph.
// For undirected graph, vertices connected is in same connected compoent.
// For directed graph, if there is a path from v to w and from w to v,
// then they are strongly connected.
type ConnectedComponent interface {
	// whether are two vertices connected
	Connected(int, int) bool

	// number of connected components
	Count() int

	// component identifier
	ID(int) int
}
