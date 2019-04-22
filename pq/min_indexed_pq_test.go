package pq

import "fmt"

func ExampleIndexedMinPQ() {
	pqueue := NewIndexedMinPQ(15)
	for i, s := range indexedPQKeys {
		pqueue.Insert(i, &MockKey{value: s})
	}

	fmt.Println("origin:")
	fmt.Println(pqueue)
	fmt.Println()

	pqueue.Poll()
	fmt.Println("poll:")
	fmt.Println(pqueue)
	fmt.Println()

	pqueue.DecreaseKey(1, &MockKey{value: "A"})
	fmt.Println("decrease key:")
	fmt.Println(pqueue)
	fmt.Println()

	// Output:
	// origin:
	// PQEXAML-
	// -14263105
	// 6524173-1
	//
	// poll:
	// PQEX-ML
	// -1256310
	// 6514-123
	//
	// decrease key:
	// PAEX-ML
	// -1126350
	// 6124-153
	//
}
