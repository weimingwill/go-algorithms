package graph

import "fmt"

func ExampleUndigraph() {
	v, _, pairs := readGraph([]byte(tinyUG))

	g := NewUndigraph(v)

	for _, pair := range pairs {
		g.AddEdge(pair[0], pair[1])
	}

	fmt.Println(g.V())
	fmt.Println(g.E())
	fmt.Println(g)

	// Output:
	// 13
	// 13
	// 0 - 5
	// 0 - 1
	// 0 - 2
	// 0 - 6
	// 1 - 0
	// 2 - 0
	// 3 - 4
	// 3 - 5
	// 4 - 3
	// 4 - 6
	// 4 - 5
	// 5 - 0
	// 5 - 4
	// 5 - 3
	// 6 - 4
	// 6 - 0
	// 7 - 8
	// 8 - 7
	// 9 - 12
	// 9 - 10
	// 9 - 11
	// 10 - 9
	// 11 - 12
	// 11 - 9
	// 12 - 9
	// 12 - 11
}
