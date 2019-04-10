package graph

import "fmt"

func ExamplePath() {
	v, _, pairs := readGraph([]byte(tinyUG))

	g := NewUndigraph(v)

	for _, pair := range pairs {
		g.AddEdge(pair[0], pair[1])
	}

	dfsPath := NewDFSPath(g, 0)
	fmt.Println(dfsPath.PathTo(1))
	fmt.Println(dfsPath.PathTo(2))
	fmt.Println(dfsPath.PathTo(3))
	fmt.Println(dfsPath.PathTo(4))
	fmt.Println(dfsPath.PathTo(5))
	fmt.Println(dfsPath.PathTo(8))

	bfsPath := NewBFSPath(g, 0)
	fmt.Println(bfsPath.PathTo(1))
	fmt.Println(bfsPath.PathTo(2))
	fmt.Println(bfsPath.PathTo(3))
	fmt.Println(bfsPath.PathTo(4))
	fmt.Println(bfsPath.PathTo(5))
	fmt.Println(bfsPath.PathTo(8))

	// Output:
	// [1]
	// [2]
	// [5 4 3]
	// [5 4]
	// [5]
	// []

	// [1]
	// [2]
	// [5 3]
	// [5 4]
	// [5]
	// []
}
