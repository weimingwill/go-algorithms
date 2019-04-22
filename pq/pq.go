package pq

// PQ defines interface for priority queue.
// Assumes Key is immutable.
type PQ interface {
	// Insert adds an element to the priority queue.
	Insert(Key)

	// Peek retrieves the head of the priority queue.
	Peek() Key

	// Poll retrives and removes the head of the priority queue.
	// It returns nil if the queue is empty.
	Poll() Key

	// IsEmpty checks whether the priority queue is empty.
	IsEmpty() bool

	// Size returns the size of the priority queue.
	Size() int
}

// IndexedPQ defines interface for indexed priority queue.
type IndexedPQ interface {
	// gets key in index i.
	KeyOf(i int) Key

	// associate key with index i.
	Insert(i int, key Key)

	// increase the key associated with index i.
	IncreaseKey(i int, key Key)

	// decrease the key associated with index i.
	DecreaseKey(i int, key Key)

	// UpdateKey updates key in index i
	UpdateKey(i int, key Key)

	// is i an index on the priority queue?
	Contains(i int) bool

	// remove a minimal key and return its index
	Poll() int

	// is the priority queue empty?
	IsEmpty() bool

	// number of entries in the priority queue
	Size() int
}
