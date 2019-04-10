package graph

import (
	"fmt"
)

// UndirectedGraph defines interface for undirected graph.
type UndirectedGraph interface {
	// Number of vertices.
	V() int

	// Number of edges.
	E() int

	// Add edge from v to w.
	AddEdge(v, w int)

	// Vertices connected to v.
	Adj(v int) []int
}

// Undigraph represents an implementation of Undirected Graph.
type Undigraph struct {
	v   int
	e   int
	adj [][]int
}

// NewUndigraph constructs an adjacency-list undirected graph representation.
func NewUndigraph(v int) UndirectedGraph {
	g := &Undigraph{
		v:   v,
		e:   0,
		adj: make([][]int, v),
	}

	for i := 0; i < v; i++ {
		g.adj[i] = []int{}
	}

	return g
}

// V returns number of vertices.
func (g *Undigraph) V() int {
	return g.v
}

// E returns number of edges.
func (g *Undigraph) E() int {
	return g.e
}

// AddEdge adds edge connecting two vertices v and w.
func (g *Undigraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

// Adj returns adjacent list of vertix v.
func (g *Undigraph) Adj(v int) []int {
	return g.adj[v]
}

func (g *Undigraph) String() string {
	str := ""
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			str += fmt.Sprintf("%d - %d\n", v, w)
		}
	}
	return str
}
