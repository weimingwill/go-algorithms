package pq

import (
	"fmt"
	"strconv"
)

// IndexedMinPQ keeps the min element in the lower index of the array.
type IndexedMinPQ struct {
	keys []Key

	// pq starts from index 1, i.e. index 0 remain empty.
	pq []int

	// qp use index to get position
	qp []int

	// size is the length of priority queue.
	size int

	maxSize int
}

// NewIndexedMinPQ constructs a min priority queue with capacity + 1 for
// inner array length. It is to keep index 0 empty.
func NewIndexedMinPQ(capacity int) IndexedPQ {
	ipq := &IndexedMinPQ{
		pq:      make([]int, capacity+1),
		qp:      make([]int, capacity+1),
		keys:    make([]Key, capacity+1),
		size:    0,
		maxSize: capacity,
	}

	for i := 0; i <= capacity; i++ {
		ipq.qp[i] = -1
		ipq.pq[i] = -1
	}
	return ipq
}

// KeyOf returns key in index i.
func (m *IndexedMinPQ) KeyOf(i int) Key {
	if i < 0 || i >= m.maxSize {
		fmt.Println("Index out of range")
		return nil
	}

	return m.keys[i]
}

// Insert adds a new element to the priority queue.
func (m *IndexedMinPQ) Insert(i int, key Key) {
	if i < 0 || i >= m.maxSize {
		fmt.Println("Index out of range")
		return
	}

	if m.Contains(i) {
		fmt.Println("Key already exist")
		return
	}

	m.size++
	m.keys[i] = key
	m.pq[m.size] = i
	m.qp[i] = m.size
	m.swim(m.size)
}

// Poll deletes the min element from the priority queue
// returns its associated index.
func (m *IndexedMinPQ) Poll() int {
	if m.IsEmpty() {
		return -1
	}

	min := m.pq[1]
	m.swap(1, m.size)
	m.size--
	m.sink(1)

	m.qp[min] = -1
	m.keys[min] = nil
	m.pq[m.size+1] = -1

	return min
}

// UpdateKey updates key in index i.
func (m *IndexedMinPQ) UpdateKey(i int, key Key) {
	if i < 0 || i >= m.maxSize {
		fmt.Println("Index out of range")
		return
	}

	if !m.Contains(i) {
		fmt.Println("Key does not exist")
		return
	}

	m.keys[i] = key
	m.swim(m.qp[i])
	m.sink(m.qp[i])
}

// IncreaseKey increase the min element from the priority queue.
func (m *IndexedMinPQ) IncreaseKey(i int, key Key) {
	if i < 0 || i >= m.maxSize {
		fmt.Println("Index out of range")
		return
	}

	if !m.Contains(i) {
		fmt.Println("Key does not exist")
		return
	}

	if m.keys[i].CompareTo(key) > 0 {
		fmt.Println("Given key cannot be increased")
		return
	}

	m.keys[i] = key
	m.sink(m.qp[i])
}

// DecreaseKey deletes the min element from the priority queue.
func (m *IndexedMinPQ) DecreaseKey(i int, key Key) {
	if i < 0 || i >= m.maxSize {
		fmt.Println("Index out of range")
		return
	}

	if !m.Contains(i) {
		fmt.Println("Key does not exist")
		return
	}

	if m.keys[i].CompareTo(key) < 0 {
		fmt.Println("Given key cannot be decreased")
		return
	}

	m.keys[i] = key
	m.swim(m.qp[i])
}

// IsEmpty checks whether the priority queue is empty.
func (m *IndexedMinPQ) IsEmpty() bool {
	return m.size == 0
}

// Size returns the size of the priority queue.
func (m *IndexedMinPQ) Size() int {
	return m.size
}

// Contains checks whether i is index of the pq.
func (m *IndexedMinPQ) Contains(i int) bool {
	if i < 0 || i >= m.maxSize {
		fmt.Println("Index out of range")
		return false
	}

	return m.qp[i] != -1
}

func (m *IndexedMinPQ) swim(k int) {
	for k > 1 && m.less(k, k/2) {
		m.swap(k/2, k)
		k = k / 2
	}
}

func (m *IndexedMinPQ) sink(k int) {
	for 2*k <= m.size {
		j := 2 * k
		if j < m.size && m.less(j+1, j) {
			j++
		}
		if m.less(k, j) {
			break
		}
		m.swap(k, j)
		k = j
	}
}

// less compares two values at position a and b.
func (m *IndexedMinPQ) less(a, b int) bool {
	return m.keys[m.pq[a]].CompareTo(m.keys[m.pq[b]]) < 0
}

// swap exchanges values at position a and b.
func (m *IndexedMinPQ) swap(a, b int) {
	m.pq[a], m.pq[b] = m.pq[b], m.pq[a]

	m.qp[m.pq[a]] = a
	m.qp[m.pq[b]] = b
}

func (m *IndexedMinPQ) String() string {
	keys := ""
	pq := ""
	qp := ""
	for i := 0; i <= m.size; i++ {
		if m.keys[i] != nil {
			keys += fmt.Sprintf("%s", m.keys[i].Value())
		} else {
			keys += "-"
		}
		pq += strconv.Itoa(m.pq[i])
		qp += strconv.Itoa(m.qp[i])
	}

	return keys + "\n" + pq + "\n" + qp
}
