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
