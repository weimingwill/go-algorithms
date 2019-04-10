package graph

import "fmt"

func ExampleSCC() {
	v, _, pairs := readGraph([]byte(tinyDG))

	g := NewDigraph(v)

	for _, pair := range pairs {
		g.AddEdge(pair[0], pair[1])
	}

	scc := NewSCC(g)

	fmt.Println(scc)

	fmt.Println(scc.Count())
	fmt.Println(scc.ID(0))
	fmt.Println(scc.Connected(0, 1))
	fmt.Println(scc.Connected(0, 2))

	// Output:
	// 0 - 1
	// 1 - 0
	// 2 - 1
	// 3 - 1
	// 4 - 1
	// 5 - 1
	// 6 - 3
	// 7 - 4
	// 8 - 3
	// 9 - 2
	// 10 - 2
	// 11 - 2
	// 12 - 2
	//
	// 5
	// 1
	// false
	// true
}
