package graph

import "fmt"

func ExampleCC() {
	v, _, pairs := readGraph([]byte(tinyUG))

	g := NewUndigraph(v)

	for _, pair := range pairs {
		g.AddEdge(pair[0], pair[1])
	}

	cc := NewCC(g)

	fmt.Println(cc)

	fmt.Println(cc.Count())
	fmt.Println(cc.ID(1))
	fmt.Println(cc.Connected(0, 7))
	fmt.Println(cc.Connected(0, 2))

	// Output:
	// 0 - 0
	// 1 - 0
	// 2 - 0
	// 3 - 0
	// 4 - 0
	// 5 - 0
	// 6 - 0
	// 7 - 1
	// 8 - 1
	// 9 - 2
	// 10 - 2
	// 11 - 2
	// 12 - 2
	//
	// 3
	// 0
	// false
	// true

}
