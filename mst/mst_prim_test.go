package mst

import (
	"bytes"
	"fmt"
)

func ExampleLazyPrimMST() {
	graph := NewEWgraph(bytes.NewReader([]byte(tinyEWG)))
	m := NewLazyPrimMST(graph)
	fmt.Print(m)
	fmt.Printf("%.2f", m.Weight())

	// Output:
	// 0-7 0.16
	// 1-7 0.19
	// 0-2 0.26
	// 2-3 0.17
	// 5-7 0.28
	// 4-5 0.35
	// 6-2 0.40
	// 1.81
}
