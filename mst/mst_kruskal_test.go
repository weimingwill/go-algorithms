package mst

import (
	"bytes"
	"fmt"
)

func ExampleKruskalMST() {
	graph := NewEWgraph(bytes.NewReader([]byte(tinyEWG)))
	m := NewKruskalMST(graph)
	fmt.Print(m)
	fmt.Printf("%.2f", m.Weight())

	// Output:
	// 0-7 0.16
	// 2-3 0.17
	// 1-7 0.19
	// 0-2 0.26
	// 5-7 0.28
	// 4-5 0.35
	// 6-2 0.40
	// 1.81
}
