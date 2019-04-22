package pq

var (
	keys          = []string{"P", "Q", "E", "-", "X", "A", "M", "-", "P", "L", "E", "-"}
	indexedPQKeys = []string{"P", "Q", "E", "X", "A", "M", "L"}
)

type MockKey struct {
	value string
}

func (m *MockKey) CompareTo(k Key) int {
	if m.value < k.Value().(string) {
		return -1
	} else if m.value > k.Value().(string) {
		return 1
	} else {
		return 0
	}
}

func (m *MockKey) Value() interface{} {
	if m == nil {
		return ""
	}

	return m.value
}
