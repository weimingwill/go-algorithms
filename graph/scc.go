package graph

import (
	"fmt"
)

// SCC implements Kosaraju Sharir algorithm for getting strongly
// connected components.
type SCC struct {
	marked []bool
	ids    []int
	count  int
}

// NewSCC contructs a new strongly connected component using
// Kosaraju Sharir algorithm.
// 1. It computes reverse post order of its reversed graph.
// 2. Runs dfs through the order from step 1 to count connected components.
func NewSCC(g DirectedGraph) ConnectedComponent {
	s := &CC{
		marked: make([]bool, g.V()),
		ids:    make([]int, g.V()),
	}

	d := NewDepthFirstOrder(g.Reverse())

	for _, v := range d.Order() {
		if !s.marked[v] {
			s.dfs(g, v)
			s.count++
		}
	}

	return s
}

func (s *SCC) dfs(g DirectedGraph, v int) {
	s.marked[v] = true
	s.ids[v] = s.count

	for _, w := range g.Adj(v) {
		if !s.marked[w] {
			s.dfs(g, w)
		}
	}
}

// Connected checks whether vertix v and w are connected in graph.
func (s *SCC) Connected(v, w int) bool {
	return s.ids[v] == s.ids[w]
}

// Count returns number of strongly connected components.
func (s *SCC) Count() int {
	return s.count
}

// ID returns strongly connected component identifier of vertix v.
func (s *SCC) ID(v int) int {
	return s.ids[v]
}

func (s *SCC) String() string {
	str := ""
	for v, id := range s.ids {
		str += fmt.Sprintf("%d - %d\n", v, id)
	}

	return str
}
