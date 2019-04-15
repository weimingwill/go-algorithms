package pq

// Key defines interface for the key in priority queue.
type Key interface {
	// CompareTo returns -1, 0, 1,
	// if the element is smaller, equal, bigger than the input.
	CompareTo(Key) int

	// Value return the value of the key.
	Value() interface{}
}
