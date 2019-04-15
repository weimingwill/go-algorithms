package pq

// MinPQ keeps the min element in the lower index of the array.
type MinPQ struct {
	// pq stores list of keys.
	// pq starts from index 1, i.e. index 0 remain empty.
	pq []Key

	// n is the length of priority queue.
	n int
}

// NewMinPQ constructs a min priority queue with capacity + 1 for
// inner array length. It is to keep index 0 empty.
func NewMinPQ(capacity int) PQ {
	return &MinPQ{
		pq: make([]Key, capacity+1),
		n:  0,
	}
}

// Insert adds a new element to the priority queue.
func (m *MinPQ) Insert(k Key) {
	if m.n == len(m.pq)-1 {
		m.resize(2 * len(m.pq))
	}

	m.n++
	m.pq[m.n] = k
	m.swim(m.n)
}

// Poll deletes the min element from the priority queue.
func (m *MinPQ) Poll() Key {
	if m.IsEmpty() {
		return nil
	}

	max := m.pq[1]
	m.exech(1, m.n)
	m.pq[m.n] = nil
	m.n--
	m.sink(1)

	if m.n > 0 && m.n == (len(m.pq)-1)/4 {
		m.resize(len(m.pq) / 2)
	}

	return max
}

// Peek returns the max element from the priority queue.
func (m *MinPQ) Peek() Key {
	return m.pq[1]
}

// IsEmpty checks whether the priority queue is empty.
func (m *MinPQ) IsEmpty() bool {
	return m.n == 0
}

// Size returns the size of the priority queue.
func (m *MinPQ) Size() int {
	return m.n
}

func (m *MinPQ) swim(k int) {
	for k > 1 && m.less(k, k/2) {
		m.exech(k/2, k)
		k = k / 2
	}
}

func (m *MinPQ) sink(k int) {
	for 2*k <= m.n {
		j := 2 * k
		if j < m.n && m.less(j+1, j) {
			j++
		}
		if m.less(k, j) {
			break
		}
		m.exech(k, j)
		k = j
	}
}

func (m *MinPQ) less(a, b int) bool {
	return m.pq[a].CompareTo(m.pq[b]) < 0
}

func (m *MinPQ) exech(a, b int) {
	m.pq[a], m.pq[b] = m.pq[b], m.pq[a]
}

func (m *MinPQ) resize(capacity int) {
	keys := make([]Key, capacity)

	for i := 1; i <= m.n; i++ {
		keys[i] = m.pq[i]
	}
	m.pq = keys
}

func (m *MinPQ) String() string {
	str := ""
	for i := 1; i <= m.n; i++ {
		if m.pq[i] == nil {
			str += "-"
		} else {
			str += m.pq[i].Value().(string)
		}
	}
	return str
}
