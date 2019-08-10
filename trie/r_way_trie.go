package trie

import "fmt"

// TrieST using trie as data structure for symbol table.
type TrieST struct {
	r    int // r is the number of trie branches -> r-way trie
	root Node
	n    int
}

type Node struct {
	value *int // it can also be an object
	next  []Node
}

func NewTrieST() *TrieST {

}

func (t *TrieST) Get(s string) *int {

}

func (t *TrieST) Put(s string, val *int) {

}

func (t *TrieST) Delete(s string) {

}

func (t *TrieST) Contains(s string) bool {
}

func (t *TrieST) KeysWithPrefix(prefix string) []string {

}

func (t *TrieST) LongestPrefixOf(query string) string {

}

func (t *TrieST) KeysThatMatch(pattern string) []string {

}

func (t *TrieST) String() string {
	str := ""
	for _, key := range t.KeysWithPrefix("") {
		str += fmt.Sprintf("%s %d", key, *t.Get(key))
	}
	return str
}
