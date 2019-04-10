package graph

import "fmt"

// CC represents an implementation of Connected Components.
type CC struct {
	marked []bool
	ids    []int
	count  int
}

// NewCC constructs a new connected component from undirected graph using dfs.
func NewCC(g UndirectedGraph) ConnectedComponent {
	c := &CC{
		marked: make([]bool, g.V()),
		ids:    make([]int, g.V()),
	}

	for i := 0; i < g.V(); i++ {
		if c.marked[i] {
			continue
		}

		c.dfs(g, i)
		c.count++
	}

	return c
}

func (c *CC) dfs(g UndirectedGraph, v int) {
	c.marked[v] = true
	c.ids[v] = c.count

	for _, w := range g.Adj(v) {
		if c.marked[w] {
			continue
		}

		c.dfs(g, w)
	}
}

// Connected checks whether vertix v and w are connected in graph.
func (c *CC) Connected(v, w int) bool {
	return c.ids[v] == c.ids[w]
}

// Count returns number of connected components.
func (c *CC) Count() int {
	return c.count
}

// ID returns connected component identifier of vertix v.
func (c *CC) ID(v int) int {
	return c.ids[v]
}

func (c *CC) String() string {
	s := ""
	for v, id := range c.ids {
		s += fmt.Sprintf("%d - %d\n", v, id)
	}

	return s
}
