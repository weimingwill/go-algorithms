package pq

import "fmt"

func ExampleMinPQ() {
	pqueue := NewMinPQ(1)
	for _, s := range keys {
		if s != "-" {
			pqueue.Insert(&MockKey{value: s})
		} else if !pqueue.IsEmpty() {
			fmt.Print(pqueue.Poll().Value())
		}
	}
	// Output:
	// EAE
}
