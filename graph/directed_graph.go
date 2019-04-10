package graph

import (
	"fmt"
)

// DirectedGraph defines interface for directed graph.
type DirectedGraph interface {
	// Number of vertices.
	V() int

	// Number of edges.
	E() int

	// Add edge from v to w.
	AddEdge(v, w int)

	// Vertices connected to v.
	Adj(v int) []int

	// Reverse of this directed graph.
	Reverse() DirectedGraph

	// Whether target and be reached from source.
	IsReachable(s, t int) bool
}

// Digraph represents an implementation of directed graph.
type Digraph struct {
	v   int
	e   int
	adj [][]int
}

// DepthFirstOrder represents a way to compute directed graph topological order.
type DepthFirstOrder struct {
	marked      []bool
	reversePost []int
	order       []int
}

// NewDigraph constructs an adjacency-list graph representation.
func NewDigraph(v int) DirectedGraph {
	g := &Digraph{
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
func (g *Digraph) V() int {
	return g.v
}

// E returns number of edges.
func (g *Digraph) E() int {
	return g.e
}

// AddEdge adds edge connecting from v to w.
func (g *Digraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

// Adj returns adjacent list of vertix v.
func (g *Digraph) Adj(v int) []int {
	return g.adj[v]
}

// IsReachable checks whether vertix t is reachable from vertix s.
func (g *Digraph) IsReachable(s, t int) bool {
	var stack []int
	marked := make([]bool, g.V())

	stack = append(stack, s)
	marked[s] = true
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if v == t {
			break
		}

		for _, w := range g.Adj(v) {
			if marked[w] {
				continue
			}

			stack = append(stack, w)
			marked[w] = true
		}
	}

	return marked[t]
}

// Reverse returns reversed directed graph.
func (g *Digraph) Reverse() DirectedGraph {
	r := NewDigraph(g.V())
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			r.AddEdge(w, v)
		}
	}
	return r
}

func (g *Digraph) String() string {
	str := ""
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			str += fmt.Sprintf("%d -> %d\n", v, w)
		}
	}
	return str
}

// NewDepthFirstOrder constructs a topological order using dfs.
func NewDepthFirstOrder(g DirectedGraph) *DepthFirstOrder {
	d := &DepthFirstOrder{
		marked:      make([]bool, g.V()),
		reversePost: make([]int, 0),
	}
	for v := 0; v < g.V(); v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}

	d.order = reverse(d.reversePost)

	return d
}

func (d *DepthFirstOrder) dfs(g DirectedGraph, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
	d.reversePost = append(d.reversePost, v)
}

// Order return reverse postorder of the directed graph.
func (d *DepthFirstOrder) Order() []int {
	return d.order
}

func reverse(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
