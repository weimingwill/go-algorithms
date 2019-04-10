package graph

import (
	"fmt"
)

func ExampleDigraph() {
	v, _, pairs := readGraph([]byte(tinyDG))

	g := NewDigraph(v)

	for _, pair := range pairs {
		g.AddEdge(pair[0], pair[1])
	}

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g)

	fmt.Println(g.IsReachable(0, 3))
	fmt.Println(g.IsReachable(0, 7))

	// Output:
	// 13
	// 22
	// 0 -> 1
	// 0 -> 5
	// 2 -> 3
	// 2 -> 0
	// 3 -> 2
	// 3 -> 5
	// 4 -> 2
	// 4 -> 3
	// 5 -> 4
	// 6 -> 0
	// 6 -> 8
	// 6 -> 4
	// 6 -> 9
	// 7 -> 9
	// 7 -> 6
	// 8 -> 6
	// 9 -> 10
	// 9 -> 11
	// 10 -> 12
	// 11 -> 12
	// 11 -> 4
	// 12 -> 9
	//
	// true
	// false
}

func ExampleDepthFirstOrder() {
	v, _, pairs := readGraph([]byte(tinyDAG))

	g := NewDigraph(v)
	for _, pair := range pairs {
		g.AddEdge(pair[0], pair[1])
	}

	d := NewDepthFirstOrder(g)

	fmt.Println(g)

	fmt.Println(d.Order())

	// Output:
	// 0 -> 6
	// 0 -> 1
	// 0 -> 5
	// 2 -> 3
	// 2 -> 0
	// 3 -> 5
	// 5 -> 4
	// 6 -> 4
	// 6 -> 9
	// 7 -> 6
	// 8 -> 7
	// 9 -> 12
	// 9 -> 10
	// 9 -> 11
	// 11 -> 12
	//
	// [8 7 2 3 0 5 1 6 9 11 10 12 4]

}
