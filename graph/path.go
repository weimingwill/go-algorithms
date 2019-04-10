package graph

// Pather defines interface for an undirected graph path.
type Pather interface {
	// find paths in G from source s
	HasPathTo(int) bool

	// path from s to v
	PathTo(int) []int
}

// Path represents paths from vertix s for a graph.
type Path struct {
	source int
	edgeTo []int
	marked []bool
}

// HasPathTo finds whether there is a path from source to vertix v.
func (p *Path) HasPathTo(v int) bool {
	return p.marked[v]
}

// PathTo returns the path from source to vertix v.
// If no, return empty array.
func (p *Path) PathTo(v int) (path []int) {
	if !p.HasPathTo(v) {
		return []int{}
	}

	for w := v; p.source != w; w = p.edgeTo[w] {
		defer func(w int) {
			path = append(path, w)
		}(w)
	}

	return path
}

// NewDFSPath constructs path from undirected graph using dfs.
func NewDFSPath(g UndirectedGraph, s int) Pather {
	p := &Path{
		source: s,
		edgeTo: make([]int, g.V()),
		marked: make([]bool, g.V()),
	}
	p.dfs(g, s)
	return p
}

func (p *Path) dfs(g UndirectedGraph, v int) {
	p.marked[v] = true
	for _, w := range g.Adj(v) {
		if !p.marked[w] {
			p.dfs(g, w)
			p.edgeTo[w] = v
		}
	}
}

// NewBFSPath constructs path from undirected graph using bfs.
func NewBFSPath(g UndirectedGraph, s int) Pather {
	p := &Path{
		source: s,
		edgeTo: make([]int, g.V()),
		marked: make([]bool, g.V()),
	}

	// maybe to replace with a real queue implementation
	queue := make(chan int, g.V())
	queue <- s
	p.marked[s] = true
	for len(queue) > 0 {
		v := <-queue
		for _, w := range g.Adj(v) {
			if !p.marked[w] {
				queue <- w
				// for bfs, enqueue first before marking as true.
				p.marked[w] = true
				p.edgeTo[w] = v
			}
		}
	}

	return p
}
