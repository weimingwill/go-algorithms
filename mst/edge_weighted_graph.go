package mst

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// EdgeWeightedGraph defines interface for edge weighted graph.
type EdgeWeightedGraph interface {
	// Add weighted edge to this graph
	AddEdge(Edge)

	// Edges adjacent to the vertix
	Adj(int) []Edge

	// All edges in this graph
	Edges() []Edge

	// Number of vertices
	V() int

	// Number of edges
	E() int
}

// EWgraph represents edge weighted graph.
type EWgraph struct {
	v     int
	e     int
	edges []Edge
	adj   [][]Edge
}

// NewEWgraph constructs a new edge weighted graph from input stream.
func NewEWgraph(in io.Reader) EdgeWeightedGraph {
	graph := &EWgraph{
		edges: make([]Edge, 0),
	}

	scanner := bufio.NewScanner(in)
	line := 0
	for scanner.Scan() {
		if line == 0 {
			graph.v, _ = strconv.Atoi(scanner.Text())
			graph.adj = make([][]Edge, graph.v)
			line++
			continue
		} else if line == 1 {
			graph.e, _ = strconv.Atoi(scanner.Text())
			line++
			continue
		}

		parts := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		weight, _ := strconv.ParseFloat(parts[2], 32)

		e := NewWeightedEdge(v, w, weight)
		graph.AddEdge(e)

		line++
	}

	return graph
}

// AddEdge adds a new edge to the graph.
func (g *EWgraph) AddEdge(e Edge) {
	g.edges = append(g.edges, e)
	v := e.Either()
	w := e.Other(v)
	g.adj[v] = append(g.adj[v], e)
	g.adj[w] = append(g.adj[w], e)
}

// Edges return all edges of the graph.
func (g *EWgraph) Edges() []Edge {
	return g.edges
}

// Adj returns edges adjacent to input vertix.
func (g *EWgraph) Adj(v int) []Edge {
	return g.adj[v]
}

// V returns number of vertices.
func (g *EWgraph) V() int {
	return g.v
}

// E returns number of edges.
func (g *EWgraph) E() int {
	return g.e
}

func (g *EWgraph) String() string {
	str := ""

	for _, edge := range g.Edges() {
		v := edge.Either()
		w := edge.Other(v)
		str += fmt.Sprintf("%d-%d %.2f\n", v, w, edge.Weight())
	}

	return str
}
